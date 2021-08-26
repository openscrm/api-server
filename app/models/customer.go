package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
	"openscrm/app/requests"
	app "openscrm/common/app"
	"time"
)

type CustomerExportItem struct {
	ExtID            string          `gorm:"type:char(32);uniqueIndex:idx_ext_customer_id;comment:微信定义的userID" json:"ext_customer_id"`
	CustomerName     string          `json:"customer_name"`
	CustomerCorpName string          `json:"customer_corp_name"`
	StaffName        string          `json:"staff_name"`
	Remark           string          `json:"remark"`
	Description      string          `json:"description"`
	Status           string          `json:"status"`
	Createtime       time.Time       `json:"createtime"`
	AddWay           int64           `json:"add_way"`
	Age              int64           `json:"age"`
	Gender           int64           `json:"gender"`
	Birthday         string          `json:"birthday"`
	PhoneNumber      string          `json:"phone_number"`
	Staffs           []CustomerStaff `gorm:"foreignKey:ExtCustomerID;references:ExtID" json:"staff_relations"`
}

type Customer struct {
	ExtCorpModel
	// 微信定义的客户ID
	ExtID string `gorm:"type:char(32);uniqueIndex:idx_ext_customer_id;comment:微信定义的userID" json:"ext_customer_id"`
	// 微信用户对应微信昵称；企业微信用户，则为联系人或管理员设置的昵称、认证的实名和账号名称
	Name string `gorm:"type:varchar(255);comment:名称，微信用户对应微信昵称；企业微信用户，则为联系人或管理员设置的昵称、认证的实名和账号名称" json:"name"`
	// 职位,客户为企业微信时使用
	Position string `gorm:"varchar(255);comment:职位,客户为企业微信时使用" json:"position"`
	// 客户的公司名称,仅当客户ID为企业微信ID时存在
	CorpName string `gorm:"type:varchar(255);comment:客户的公司名称,仅当客户ID为企业微信ID时存在" json:"corp_name"`
	// 头像
	Avatar string `gorm:"type:varchar(255);comment:头像" json:"avatar"`
	// 客户类型 1-微信用户, 2-企业微信用户
	Type int `gorm:"type:tinyint(1);index;comment:类型,1-微信用户, 2-企业微信用户" json:"type"`
	// 0-未知 1-男性 2-女性
	Gender  int    `gorm:"type:tinyint;comment:性别,0-未知 1-男性 2-女性" json:"gender"`
	Unionid string `gorm:"type:varchar(128);comment:微信开放平台的唯一身份标识(微信unionID)" json:"unionid"`
	// 仅当联系人类型是企业微信用户时有此字段
	ExternalProfile ExternalProfile `gorm:"type:json;comment:仅当联系人类型是企业微信用户时有此字段" json:"external_profile"`
	// 所属员工
	Staffs []CustomerStaff `gorm:"foreignKey:ExtCustomerID;references:ExtID" json:"staff_relations"`
	// 所属员工
	Timestamp
}

func (cs Customer) BatchUpsert(customers []Customer) (err error) {
	err = DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ext_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"position", "name", "corp_name", "avatar", "type", "gender", "unionid", "external_profile"}),
	}).Omit("Staffs").CreateInBatches(&customers, 100).Error
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (cs Customer) Upsert(customer Customer) error {
	updateFields := map[string]interface{}{
		"name":             customer.Name,
		"position":         customer.Position,
		"corp_name":        customer.CorpName,
		"avatar":           customer.Avatar,
		"type":             customer.Type,
		"gender":           customer.Gender,
		"unionid":          customer.Unionid,
		"external_profile": customer.ExternalProfile,
	}
	return DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ext_id"}},
		DoUpdates: clause.Assignments(updateFields),
	}).Omit("Staffs").Create(&customer).Error
}

func (cs Customer) Get(ID string, extCorpID string, withStaffRelation bool) (*Customer, error) {
	customer := Customer{}
	db := DB.Model(&Customer{}).Where("id = ? and ext_corp_id = ?", ID, extCorpID)
	if withStaffRelation {
		db = db.Preload("Staffs").Preload("Staffs.CustomerStaffTags")
	}
	err := db.Find(&customer).Error
	if err != nil {
		err = errors.Wrap(err, "Get customer by id failed")
		return &customer, err
	}
	return &customer, nil
}

//select c.name,
//       c.corp_name,
//       s.name,
//       customer_staff.remark,
//       customer_staff.description,
//       customer_staff.customer_delete_staff_at,
//       customer_staff.createtime,
//       customer_staff.add_way,
//       ci.age,
//       ci.birthday,
//       ci.phone_number
//from customer_staff
//         join customer c on c.ext_customer_id = customer_staff.ext_customer_id
//         join staff s on s.ext_staff_id = customer_staff.ext_staff_id
//         left join customer_info ci on c.id = ci.ext_customer_id;

func (cs Customer) QueryExport(req requests.QueryCustomerReq, extCorpID string, pager *app.Pager) ([]CustomerExportItem, int64, error) {

	db := DB.Table("customer_staff").
		Joins("join customer c on c.ext_id = customer_staff.ext_customer_id").
		Joins("join staff s on s.ext_id = customer_staff.ext_staff_id").
		Joins("left join customer_info ci on c.id = ci.ext_customer_id")

	if req.Name != "" {
		db = db.Where("customer.name like ?", req.Name+"%")
	}
	if req.Gender != 0 {
		db = db.Where("customer.gender = ?", req.Gender)
	}
	if req.Type != 0 {
		db = db.Where("customer.type = ?", req.Type)
	}
	if len(req.ExtStaffIDs) > 0 {
		db = db.Where("cs.ext_staff_ids in (?)", req.ExtStaffIDs)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	pageOffset := app.GetPageOffset(pager.Page, pager.PageSize)
	if pageOffset >= 0 && pager.PageSize > 0 {
		db = db.Offset(pageOffset).Limit(pager.PageSize)
	}
	var res []CustomerExportItem
	if err := db.Preload("Staffs"). /*Preload("Staffs.CustomerStaffTags").*/ Select("customer.*").Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

// ExportQuery
// Description: 查询需要导出的客户
func (cs Customer) ExportQuery(
	req requests.QueryCustomerReq, extCorpID string, pager *app.Pager) ([]*CustomerExportItem, int64, error) {

	var customers []*CustomerExportItem

	db := DB.Table("customer").
		Joins("left join customer_staff cs on customer.ext_id = cs.ext_customer_id").
		Joins("left join customer_staff_tag cst on cst.customer_staff_id = cs.id").
		Joins("join staff s on s.ext_id = cs.ext_staff_id").
		Joins("left join customer_info ci on customer.ext_id = ci.ext_customer_id").
		Where("cs.ext_corp_id = ?", extCorpID)

	if req.Name != "" {
		db = db.Where("customer.name like ?", req.Name+"%")
	}
	if req.Gender != 0 {
		db = db.Where("customer.gender = ?", req.Gender)
	}
	if req.Type != 0 {
		db = db.Where("customer.type = ?", req.Type)
	}
	if len(req.ExtStaffIDs) > 0 {
		db = db.Where("cs.ext_staff_id in (?)", req.ExtStaffIDs)
	}
	if req.StartTime != "" {
		db = db.Where("createtime between ? and ?", req.StartTime, req.EndTime)
	}
	if len(req.ExtTagIDs) > 0 {
		db = db.Where("cst.ext_tag_id in (?)", req.ExtTagIDs)
	}
	if req.ChannelType > 0 {
		db = db.Where("cs.add_way = ?", req.ChannelType)
	}

	var total int64
	if err := db.Distinct("customer.id").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err := db.Preload("Staffs").
		Preload("Staffs.CustomerStaffTags").
		Select("customer.name as customer_name, customer.corp_name as customer_corp_name, " +
			" s.name as staff_name, cs.remark, cs.description, if(cs.deleted_at is null, '未流失', '已流失') as status," +
			" cs.createtime, cs.add_way, ci.age,  customer.gender, ci.birthday ,ci.phone_number").
		Find(&customers).Error
	if err != nil {
		err = errors.WithStack(err)
		return nil, 0, err
	}
	return customers, total, nil
}

// Query
// Description: 查询客户
func (cs Customer) Query(
	req requests.QueryCustomerReq, extCorpID string, pager *app.Pager) ([]*Customer, int64, error) {

	var customers []*Customer

	db := DB.Table("customer").
		Joins("left join customer_staff cs on customer.ext_id = cs.ext_customer_id").
		Joins("left join customer_staff_tag cst on cst.customer_staff_id = cs.id").
		Where("cs.ext_corp_id = ?", extCorpID)

	if req.Name != "" {
		db = db.Where("customer.name like ?", req.Name+"%")
	}
	if req.Gender != 0 {
		db = db.Where("customer.gender = ?", req.Gender)
	}
	if req.Type != 0 {
		db = db.Where("customer.type = ?", req.Type)
	}
	if len(req.ExtStaffIDs) > 0 {
		db = db.Where("cs.ext_staff_id in (?)", req.ExtStaffIDs)
	}
	if req.StartTime != "" {
		db = db.Where("createtime between ? and ?", req.StartTime, req.EndTime)
	}
	if len(req.ExtTagIDs) > 0 {
		//db = db.Where("json_contains(cs.ext_tag_ids, json_array(?))", customerStaff.ExtTagIDs)
		db = db.Where("cst.ext_tag_id in (?)", req.ExtTagIDs)
	}
	if req.ChannelType > 0 {
		db = db.Where("cs.add_way = ?", req.ChannelType)
	}
	if req.OutFlowStatus == 1 {
		db = db.Unscoped().Where("cs.deleted_at is not null")
	} else if req.OutFlowStatus == 2 {
		db = db.Where("cs.deleted_at is null")
	}

	var total int64
	if err := db.Distinct("customer.id").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err := db.Preload("Staffs").
		Preload("Staffs.CustomerStaffTags").
		Select("customer.*").
		Group("customer.ext_id").
		Find(&customers).Error
	if err != nil {
		err = errors.WithStack(err)
		return nil, 0, err
	}
	return customers, total, nil
}

func (cs Customer) GetMassMsg(missionID string) (*MassMsg, error) {
	msg := &MassMsg{}
	err := DB.Model(&MassMsg{}).Where("id = ?", missionID).First(&msg).Error
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (cs Customer) GetByExtID(
	ExtCustomerID string, extStaffIDs []string, withStaffRelation bool) (customer Customer, err error) {

	db := DB.Model(&Customer{})
	if withStaffRelation {
		if extStaffIDs != nil {
			db = db.Preload("Staffs", "ext_staff_id IN (?)", extStaffIDs).Preload("Staffs.CustomerStaffTags")
		} else {
			db = db.Preload("Staffs").Preload("Staffs.CustomerStaffTags")
		}
	}
	err = db.Where("ext_id = ? ", ExtCustomerID).First(&customer).Error
	if err != nil {
		err = errors.Wrap(err, "Get customer by id failed")
		return customer, err
	}
	return customer, nil
}

type CustomerSummary struct {
	CorpName               string `json:"corp_name"`
	TotalStaffsNum         int64  `json:"total_staffs_num"`
	TotalCustomersNum      int64  `json:"total_customers_num"`
	TodayCustomersIncrease int64  `json:"today_customers_increase"`
	TodayCustomersDecrease int64  `json:"today_customers_decrease"`
	TotalGroupsNum         int64  `json:"total_groups_num"`
	TodayGroupsIncrease    int64  `json:"today_groups_increase"`
	TodayGroupsDecrease    int64  `json:"today_groups_decrease"`
}
