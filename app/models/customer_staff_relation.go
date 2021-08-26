package models

import (
	"database/sql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/app/requests"
	"openscrm/common/app"
	"time"
)

type CustomerTag struct {
	// GroupName 该成员添加此外部联系人所打标签的分组名称（标签功能需要企业微信升级到2.7.5及以上版本）
	GroupName string `json:"group_name"`
	// TagName 该成员添加此外部联系人所打标签名称
	TagName string `json:"tag_name"`
	// Type 该成员添加此外部联系人所打标签类型, 1-企业设置, 2-用户自定义
	Type  constants.FollowUserTagType `json:"type"`
	TagId string                      `json:"tag_id"`
}
type CustomerStaffRelation struct {
	ExtStaffID    string `json:"ext_staff_id"`
	ExtCustomerID string `json:"ext_customer_id"`
}

// CustomerStaff 客户-员工关系
// 员工客户关系的历史数据（流水记录）也在此表中。
// 员工删除客户/客户删除员工时 新增一条数据，写入 customer_delete_staff_at/staff_delete_customer_at, 同时软删除原有记录。
type CustomerStaff struct {
	ExtCorpModel
	// 企微员工ID
	ExtStaffID string `gorm:"type:char(32);index;uniqueIndex:idx_ext_staff_id_ext_customer_id;comment:员工ID" json:"ext_staff_id"`
	// 企微客户ID
	ExtCustomerID string `gorm:"type:char(32);index;uniqueIndex:idx_ext_staff_id_ext_customer_id;comment:客户ID" json:"ext_customer_id"`
	// 员工对客户的备注
	Remark string `gorm:"type:varchar(255);comment:员工对客户的备注" json:"remark"`
	// 员工对客户的描述
	Description string `gorm:"type:varchar(255);comment:员工对此客户的描述" json:"description"`
	// 员工添加客户的时间,与wx返回的一致，以便使用copier
	Createtime time.Time `gorm:"comment:员工添加客户的时间" json:"createtime"`
	// 员工对客户备注的企业名称
	RemarkCorpName string `gorm:"type:varchar(255);comment:员工对客户备注的企业名称" json:"remark_corp_name"`
	// RemarkMobiles 对此客户备注的手机号码，第三方不可获取
	RemarkMobiles constants.StringArrayField `gorm:"type:json;comment:对此客户备注的手机号码" json:"remark_mobiles"`
	// 添加此客户的来源 0-未知来源 1-扫描二维码 2-搜索手机号 3-名片分享 4-群聊 5-手机通讯录 6-微信联系人 7-来自微信的添加好友申请 8-安装第三方应用时自动添加的客服人员 9-搜索邮箱 201-内部成员共享 202-管理员/负责人分配
	AddWay constants.FollowUserAddWay `gorm:"tinyint(8);comment:添加此客户的来源,0-未知来源 1-扫描二维码 2-搜索手机号 3-名片分享 4-群聊 5-手机通讯录 6-微信联系人 7-来自微信的添加好友申请 8-安装第三方应用时自动添加的客服人员 9-搜索邮箱 201-内部成员共享 202-管理员/负责人分配" json:"add_way"`
	// 发起添加的userid
	OperUserID string `gorm:"type:varchar(255);comment:发起添加的userid" json:"oper_user_id"`
	// 区分客户具体是通过哪个「联系我」添加，由企业通过创建「联系我」方式指定
	State string `gorm:"type:varchar(255);comment:区分客户具体是通过哪个「联系我」添加，由企业通过创建「联系我」方式指定" json:"state"`
	// 是否已发送通知，成员删除客户通知管理员，客户删除成员通知成员
	IsNotified constants.IsNotified `gorm:"type:tinyint;comment:是否已发送通知 1-是 2-否" json:"is_notified"`
	// Tags 员工给客户的标签
	CustomerStaffTags []CustomerStaffTag         `gorm:"foreignKey:CustomerStaffID;references:ID" json:"customer_staff_tags"`
	InternalTagIDs    constants.StringArrayField `gorm:"type:json" json:"internal_tag_ids"`
	InternalTags      []InternalTag              `gorm:"-" json:"internal_tags"`
	// 员工给客户的设置的信息
	// CustomerInfo CustomerInfo `gorm:"foreignKey:CustomerInfoID;references:ID" json:"customer_info" `
	// extStaffId-extCustomerID hash,用于批量操作
	Signature string `gorm:"type:char(64);index" json:"signature"`
	Timestamp
}

func (o CustomerStaff) GetAllCustomerDeleteStaff(req requests.QueryCustomerLossesReq, extCorpID string, sorter *app.Sorter) ([]CustomerLossInfo, int64, error) {

	db := DB.Table("customer_staff").
		Joins("join customer on customer.ext_id = customer_staff.ext_customer_id").
		Joins("join staff s on customer_staff.ext_staff_id = s.ext_id").
		Select("customer_staff.id as id, customer.ext_id as ext_customer_id, " +
			" customer.avatar as customer_avatar, " +
			" customer.name as ext_customer_name, " +
			" customer_staff.createtime as  relation_create_at, " +
			" customer_staff.customer_delete_staff_at as customer_delete_staff_at, " +
			" s.name as staff_name, " +
			" s.ext_id as ext_staff_id, " +
			" s.id as staff_id, " +
			" s.avatar_url as staff_avatar, " +
			" day(timediff(customer_staff.createtime, customer_delete_staff_at)) as in_connection_time_range, " +
			" customer_staff.ext_tag_ids as ext_tag_ids ").
		Where("customer_delete_staff_at is not null")

	if extCorpID != "" {
		db = db.Where("s.ext_corp_id = ? ", extCorpID)
	}

	if len(req.ExtStaffIDs) > 0 {
		db = db.Where("s.ext_id in(?) ", req.ExtStaffIDs)
	}

	if req.LossStart != "" {
		db = db.Where("customer_staff.customer_delete_staff_at between ? and ?", req.LossStart, req.LossEnd)
	}

	if req.ConnectionCreateStart != "" {
		db = db.Where("customer_staff.createtime between ? and ?", req.ConnectionCreateStart, req.ConnectionCreateEnd)
	}

	if req.TimeSpanLowerLimit > 0 {
		db = db.Where(" day(timediff(staff_delete_customer_at, customer_delete_staff_at))  > ?", req.TimeSpanLowerLimit)
	}

	if req.TimeSpanUpperLimit > 0 {
		db = db.Where(" day(timediff(staff_delete_customer_at, customer_delete_staff_at))  < ?", req.TimeSpanUpperLimit)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count customer_delete_staff failed")
		return nil, 0, err
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: string(sorter.SortField)},
		Desc:   sorter.SortType == constants.SortTypeDesc},
	)

	items := make([]CustomerLossInfo, 0)
	err = db.Preload("Tags").Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find QueryCustomerDeleteStaff failed")
		return nil, 0, err
	}
	return items, total, err
}

func (o CustomerStaff) GetStaffByStaffID(startTime, endTime int64, id string) (*StaffCustomerCount, error) {
	panic("implement me")
}

type CustomerLossInfo struct {
	ID                    string                     `json:"id"`
	ExtCustomerID         string                     `json:"ext_customer_id"`
	CustomerAvatar        string                     `json:"customer_avatar"`
	CustomerCorpName      string                     `json:"customer_corp_name"`
	CustomerType          int                        `json:"customer_type"`
	ExtCustomerName       string                     `json:"ext_customer_name"`
	RelationCreateAt      time.Time                  `json:"relation_create_at"`
	CustomerDeleteStaffAt time.Time                  `json:"customer_delete_staff_at"`
	StaffName             string                     `json:"staff_name"`
	StaffID               int64                      `json:"staff_id"`
	ExtStaffID            string                     `json:"ext_staff_id"`
	StaffAvatar           string                     `json:"staff_avatar"`
	InConnectionTimeRange int64                      `json:"in_connection_time_range"`
	ExtTagIDs             constants.StringArrayField `json:"ext_tag_ids"`
	Tags                  []CustomerStaffTag         `gorm:"foreignKey:CustomerStaffID;references:ID" json:"tags"`
}

type StaffDeleteCustomer struct {
	Model
	// 客户id
	ExtCustomerID string `json:"ext_customer_id"`
	// 客户头像 url
	ExtCustomerAvatar string `json:"ext_customer_avatar"`
	// 客户名
	ExtCustomerName string `json:"ext_customer_name"`
	// 客户企业名称
	CustomerCorpName string `json:"customer_corp_name"`
	// 客户账号类型
	CustomerType int `json:"customer_type"`
	// 添加好友时间
	RelationCreateAt time.Time `json:"relation_create_at"`
	// 删除好友时间
	RelationDeleteAt time.Time `json:"relation_delete_at"`
	// 员工头像url
	ExtStaffAvatar string `json:"ext_staff_avatar"`
	// 企微员工id
	ExtStaffId string `json:"ext_staff_id"`
	// 员工id
	StaffId int `json:"staff_id"`
	// 员工名字
	StaffName string `json:"staff_name"`
}

type StaffCustomerCount struct {
	DecreaseUserCount int64 `json:"decrease_user_count"`
	IncreaseUserCount int64 `json:"increase_user_count"`
	TotalUserCount    int64 `json:"total_user_count"`
}

// GetRelationsByExtCustomerID
// Description: 查员工-客户关系
// Param: extCustomerID 外部客户ID
// Param: extStaffID 外部员工ID
// return: 关系数组
func (o CustomerStaff) GetRelationsByExtCustomerID(extCustomerID, extStaffID, extCorpID string) (relations []CustomerStaff, err error) {
	db := DB.Model(&CustomerStaff{}).
		Where("ext_corp_id = ?", extCorpID).
		Where("ext_customer_id = ?", extCustomerID)

	if extStaffID != "" {
		db = db.Where("ext_staff_id = ?", extStaffID)
	}
	err = db.Find(&relations).Error
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (o CustomerStaff) Get(customerStaff CustomerStaff) (CustomerStaff, error) {
	db := DB.Model(&CustomerStaff{}).Where("ext_corp_id = ?", customerStaff.ExtCorpID)

	if customerStaff.ExtStaffID != "" {
		db = db.Where("ext_staff_id = ?", customerStaff.ExtStaffID)
	}
	if customerStaff.ExtCustomerID != "" {
		db = db.Where("ext_customer_id = ?", customerStaff.ExtCustomerID)
	}

	res := CustomerStaff{}
	err := db.First(&res).Error
	if err != nil {
		return CustomerStaff{}, err
	}
	return res, nil
}

func (o CustomerStaff) GetStaffDeleteCustomerHistory(extStaffID, extCorpID string) (cs CustomerStaff, err error) {
	err = DB.Model(&CustomerStaff{}).
		Where(" ext_corp_id = ?", extCorpID).
		Where(" ext_staff_id = ?", extStaffID).
		Where(" staff_delete_customer_at is not null").Find(&cs).Error
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (o CustomerStaff) Upsert(relation *CustomerStaff) error {

	return DB.Transaction(func(tx *gorm.DB) error {
		relation.DeletedAt = gorm.DeletedAt(sql.NullTime{})
		updates := tx.Model(&CustomerStaff{}).
			Where(" ext_customer_id = ? and ext_staff_id = ?", relation.ExtCustomerID, relation.ExtStaffID).
			Where(" customer_delete_staff_at = ? and staff_delete_customer_at = ? ", nil, nil).
			Omit("CustomerStaffTags").Updates(&relation)
		if updates.Error != nil {
			return updates.Error
		}
		if updates.RowsAffected == 0 {
			// DB没有记录
			err := tx.Create(&relation).Error
			if err != nil {
				return err
			}
		} else {
			err := CustomerStaffTag{}.Upsert(relation.CustomerStaffTags)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (o CustomerStaff) Update(staff *CustomerStaff, extStaffID string, extCustomerID string) error {
	if extStaffID == "" || extCustomerID == "" {
		return errors.New(" extCustomerID,extStaffID 不能为空")
	}
	return DB.Model(CustomerStaff{}).
		Where("ext_staff_id = ?", extStaffID).
		Where("ext_customer_id = ?", extCustomerID).
		Omit("ext_staff_id", "ext_customer_id").
		Updates(&staff).Error
}

func (o CustomerStaff) GetAllStaffDeleteCustomer(
	req requests.QueryStaffDeleteCustomerHistoryReq, extCorpID string, sorter *app.Sorter) ([]StaffDeleteCustomer, int64, error) {

	db := DB.Table("customer").
		Joins("join customer_staff cs on customer.ext_id = cs.ext_customer_id").
		Joins("join staff s on cs.ext_staff_id = s.ext_id").
		Select("cs.id as id," +
			" customer.ext_id as ext_customer_id, " +
			" customer.avatar as ext_customer_avatar, " +
			" customer.name as ext_customer_name, " +
			" cs.createtime as relation_create_at, " +
			" cs.staff_delete_customer_at as relation_delete_at,  " +
			" s.name as staff_name, " +
			" s.ext_id as ext_staff_id, " +
			" s.id as staff_id, " +
			" s.avatar_url as ext_staff_avatar ")

	if req.ExtDepartmentID != 0 {
		db = db.Where("  json_contains(s.dept_ids, json_array(?) )", req.ExtDepartmentID)
	}

	if len(req.ExtStaffIDs) > 0 {
		db = db.Where(" s.ext_id in (?) ", req.ExtStaffIDs)
	}

	if req.ConnectionCreateStart != "" && req.ConnectionCreateEnd != "" {
		db = db.Where("cs.createtime between ? and ?", req.ConnectionCreateStart, req.ConnectionCreateEnd)
	}

	if req.DeleteCustomerStart != "" && req.DeleteCustomerEnd != "" {
		db = db.Where("cs.staff_delete_customer_at between ? and ?", req.DeleteCustomerStart, req.DeleteCustomerEnd)
	} else {
		db = db.Where("cs.staff_delete_customer_at is not null")
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count StaffDeleteCustomer failed")
		return nil, 0, err
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: string(sorter.SortField)},
		Desc:   sorter.SortType == constants.SortTypeDesc},
	)

	var items []StaffDeleteCustomer
	err = db.Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find StaffDeleteCustomer failed")
		return nil, 0, err
	}
	return items, total, err
}

// GetCurrentCustomerStaffRelation 员工和客户当前是否是好友
func (o CustomerStaff) GetCurrentCustomerStaffRelation(extCustomerID, extStaffID string) (cs CustomerStaff, err error) {
	err = DB.Model(&CustomerStaff{}).
		Where("ext_customer_id = ? ", extCustomerID).
		Where("ext_staff_id = ?", extStaffID).
		//Where("customer_delete_staff_at is not null and staff_delete_customer_at is not null").
		Find(&cs).Error
	return
}

func (o CustomerStaff) Create(relation *CustomerStaff) error {
	return DB.Create(relation).Error
}

// BatchUpsert
// Description: 批量更新员工客户关系
// Detail:
// 更新的relation,查出relationID
// DB中已经删除的relation,需要将其删除时间置空
// Param: 员工客户关系数组
// return:
func (o CustomerStaff) BatchUpsert(csRelationModels []CustomerStaff) (ids []string, err error) {
	err = DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ext_customer_id"}, {Name: "ext_staff_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"remark", "description", "createtime", "remark_corp_name", "remark_mobiles", "state", "oper_user_id", "add_way", "deleted_at"}),
	}).Omit("CustomerStaffTags").Unscoped().CreateInBatches(&csRelationModels, 100).Error
	if err != nil {
		return
	}
	return
}

type RelationIDSign struct {
	ID        string `json:"id"`
	Signature string `json:"signature"`
}

func (o CustomerStaff) QueryIDs(signatures []string) (ids []RelationIDSign, err error) {
	err = DB.Model(&CustomerStaff{}).
		Select("id, signature").
		Where("signature in (?)", signatures).
		Find(&ids).Error
	return
}
