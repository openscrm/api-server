package models

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/we_work"
	"openscrm/pkg/easywework"
)

// GroupChatAutoJoinCode 自动拉群码
type GroupChatAutoJoinCode struct {
	ExtCorpModel
	CreateType constants.GroupChatAutoCreateType `gorm:"type:tinyint;comment:拉群方式，1-群二维码，2-企微活码" json:"create_type"`
	// 自动拉群码分组
	GroupID string `gorm:"type:bigint;uniqueIndex:idx_group_id_remark" json:"group_id"`
	// 二维码名称
	Remark string `gorm:"type:varchar(128);index;uniqueIndex:idx_group_id_remark" json:"remark"`
	// 员工被客户扫码添加的自动回复
	AutoReply string `gorm:"type:text" json:"auto_reply"`
	// 每天加群人数是否有上线
	DayAddUserLimitEnable constants.Boolean          `gorm:"type:tinyint unsigned" json:"day_add_user_limit_enable"`
	BackupStaffIDs        constants.StringArrayField `gorm:"type:json" json:"backup_staff_ids" `
	GroupChatQRCode       []GroupChatQRCode          `gorm:"foreignKey:GroupChatAutoJoinID;references:ID" json:"group_chat_qr_code"`
	// ConfigID 自动拉群码配置ID
	ConfigID string `json:"config_id" gorm:"index;comment:自动拉群码配置ID"`
	// QrCode 联系二维码的URL，仅在scene为2时返回
	QrCode string `json:"qr_code" gorm:"comment:联系二维码的URL"`
	// SkipVerify 外部客户添加时是否无需验证，假布尔类型
	SkipVerify constants.Boolean `json:"skip_verify" gorm:"type:tinyint unsigned;default:1;comment:外部客户添加时是否无需验证，假布尔类型"`
	// State 企业自定义的state参数，用于区分不同的添加渠道，在调用“获取外部联系人详情”时会返回该参数值
	State string `json:"state" gorm:"comment:企业自定义的state参数"`
	// AddCustomerCount 扫码添加人次
	AddCustomerCount int `json:"add_customer_count" gorm:"default:0;comment:扫码添加人次"`
	// DailyAddCustomerLimitEnable 是否开启员工每日添加上限
	DailyAddCustomerLimitEnable constants.Boolean `json:"daily_add_customer_limit_enable" gorm:"comment:是否开启员工每日添加上限" `
	// Staffs 绑定员工
	Staffs []GroupChatAutoJoinCodeStaff `json:"staffs" gorm:"foreignKey:GroupChatAutoJoinCodeID;references:ID"`
	// BackupStaffs 绑定备份员工
	BackupStaffs []GroupChatAutoJoinBackupStaff `json:"backup_staffs" gorm:"foreignKey:GroupChatAutoJoinCodeID;references:ID"`
	// AutoTagEnable 是否自动打标签
	AutoTagEnable constants.Boolean `json:"auto_tag_enable" gorm:"comment:'是否自动打标签'"`
	// ExtTagIDs 自动打标签绑定的标签ID数组
	ExtTagIDs constants.StringArrayField `json:"ext_tag_ids" gorm:"type:json;comment:'自动打标签绑定的标签ID数组'"`
	// ExtStaffIDs 关联的外部员工ID
	ExtStaffIDs constants.StringArrayField `json:"ext_staff_ids" gorm:"type:json;comment:'关联的外部员工ID'"`
	Timestamp
}

func (o GroupChatAutoJoinCode) UpdateGroupID(ids []string, newGroupID string) error {
	return DB.Model(&GroupChatAutoJoinCode{}).Where("id in (?)", ids).Update("group_id", newGroupID).Error
}

// GroupChatQRCode 自动拉群码中的群二维码
type GroupChatQRCode struct {
	GroupChatAutoJoinID string `gorm:"type:bigint;comment:自动拉群码id" json:"group_chat_auto_join_id"`
	Order               int    `gorm:"type:int unsigned;" json:"order"`
	QrMediaId           string `gorm:"type:text;comment:群二维码pic media id" json:"qr_media_id"`
	QrUrl               string `gorm:"type:text;comment:群二维码的pic url" json:"qr_url"`
	UserLimit           int    `gorm:"type:int unsigned;comment:群二维码添加好友数上限" json:"user_limit"`
	Status              int    `gorm:"type:tinyint unsigned;comment:群二维码状态,1- 使用中 2-已停用" json:"status"`
}

// GroupChatAutoJoinBackupStaff 自动拉群码绑定的备份员工
type GroupChatAutoJoinBackupStaff struct {
	GroupChatAutoJoinCodeStaff
}

// GroupChatAutoJoinCodeStaff 自动拉群码绑定的员工
type GroupChatAutoJoinCodeStaff struct {
	ExtCorpModel
	// 自动拉群码id
	GroupChatAutoJoinCodeID string `json:"group_chat_auto_join_code_id" gorm:"type:bigint;index;unique:idx_group_chat_auto_join_code_id_staff_id;comment:'自动拉群码id'"`
	DayAddCustomerCount     int    `json:"day_add_customer_count" gorm:"default:0;comment:'员工每日添加客户计数'"`
	AddCustomerCount        int    `json:"add_customer_count" gorm:"default:0;comment:'员工累计添加客户计数'"`
	DayAddCustomerLimit     int    `json:"day_add_customer_limit" gorm:"comment:'员工每日添加客户上限'"`
	Avatar                  string `json:"avatar" gorm:"comment:'员工头像'"`
	StaffID                 string `json:"staff_id" gorm:"unique:idx_group_chat_auto_join_code_id_staff_id;type:bigint;comment:'员工ID'"`
	ExtStaffID              string `json:"ext_staff_id" gorm:"index;comment:'外部员工ID'"`
	Name                    string `json:"name" gorm:"comment:'员工名称'"`
	Timestamp
}

func (o GroupChatAutoJoinCodeStaff) TableName() string {
	return "group_chat_auto_join_code_staff"
}

func (o GroupChatAutoJoinCode) TableName() string {
	return "group_chat_auto_join_code"
}

func (o GroupChatQRCode) TableName() string {
	return "group_chat_qrcode"
}

func (o GroupChatAutoJoinCode) Create(autoCreateCode GroupChatAutoJoinCode) (err error) {

	err = DB.Create(&autoCreateCode).Error
	if err != nil {
		err = errors.Wrap(err, "Create ContactWay failed")
		return
	}

	return
}

func (o GroupChatAutoJoinCode) Update(autoJoinCodeID string, autoJoinCode GroupChatAutoJoinCode, extCorpID string) (*GroupChatAutoJoinCode, error) {
	tx := DB.Begin()
	defer tx.Rollback()
	var newAutoJoinCode GroupChatAutoJoinCode
	err := tx.Where("id = ?", autoJoinCodeID).Where("ext_corp_id = ?", extCorpID).First(&newAutoJoinCode).Error
	if err != nil {
		err = errors.Wrap(err, "get GroupChatAutoJoinCode failed")
		return nil, err
	}

	err = copier.CopyWithOption(&newAutoJoinCode, autoJoinCode, copier.Option{IgnoreEmpty: true})
	if err != nil {
		err = errors.Wrap(err, "copy param failed")
		return nil, err
	}

	newAutoJoinCode.Staffs = autoJoinCode.Staffs
	newAutoJoinCode.BackupStaffs = autoJoinCode.BackupStaffs

	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		err = errors.Wrap(err, "get Client failed")
		return nil, err
	}

	_, err = client.Customer.UpdateContactWay(workwx.UpdateContactWay{
		ConfigID:   autoJoinCode.ConfigID,
		Remark:     autoJoinCode.Remark,
		SkipVerify: autoJoinCode.SkipVerify == constants.True,
		State:      autoJoinCode.State,
		User:       autoJoinCode.ExtStaffIDs,
	})
	if err != nil {
		err = errors.Wrap(err, "wx UpdateContactWay failed")
		return nil, err
	}

	err = tx.Omit(clause.Associations).Updates(&newAutoJoinCode).Error
	if err != nil {
		err = errors.Wrap(err, "Update ContactWay failed")
		return nil, err
	}

	err = tx.Model(&newAutoJoinCode).Association("Staffs").Replace(&newAutoJoinCode.Staffs)
	if err != nil {
		err = errors.Wrap(err, "Update Staffs failed")
		return nil, err
	}

	// 否则会replace 掉 staffs
	if len(newAutoJoinCode.BackupStaffs) > 0 {
		err = tx.Model(&newAutoJoinCode).Association("BackupStaffs").Replace(&newAutoJoinCode.BackupStaffs)
		if err != nil {
			err = errors.Wrap(err, "Update BackupStaffs failed")
			return nil, err
		}
	}

	err = tx.Commit().Error
	if err != nil {
		err = errors.Wrap(err, "tx.Commit() failed")
		return nil, err
	}

	return &newAutoJoinCode, nil
}

func (o GroupChatAutoJoinCode) Delete(ids []string, extCorpID string) (total int64, err error) {
	result := DB.Where("ext_corp_id = ?", extCorpID).Where("id in (?)", ids).Delete(&GroupChatAutoJoinCode{})
	err = result.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.WithStack(ecode.ItemNotFoundError)
			return
		}
		//err = errors.Wrap(err, "Delete group_chat_auto_join code failed")
		return
	}
	total = result.RowsAffected

	return
}

func (o GroupChatAutoJoinCode) Query(autoJoinCode GroupChatAutoJoinCode, extCorpID string, sorter *app.Sorter, pager *app.Pager) ([]GroupChatAutoJoinCode, int64, error) {
	items := make([]GroupChatAutoJoinCode, 0)
	db := DB.Model(&GroupChatAutoJoinCode{}).Where("ext_corp_id = ?", extCorpID)
	if autoJoinCode.Remark != "" {
		db = db.Where("remark like ?", autoJoinCode.Remark+"%")
	}

	if autoJoinCode.GroupID != "" {
		db = db.Where("group_id = ?", autoJoinCode.GroupID)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count  GroupChatAutoJoinCode failed")
		return nil, 0, err
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err = db.Preload("GroupChatQRCode").Preload("Staffs").Preload("BackupStaffs").Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find GroupChatAutoJoinCode failed")
		return nil, 0, err
	}

	return items, total, nil
}
