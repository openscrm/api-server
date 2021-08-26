package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
)

// CustomerInfo 对客户编辑的用户画像，不同于企业微信员工对客户的备注和描述，后者记录在staff_customer中
type CustomerInfo struct {
	ExtCorpModel
	ExtCustomerID string                        `gorm:"type:char(32);uniqueIndex:idx_ext_staff_id_ext_customer_id;comment:微信客户ID" json:"ext_customer_id"`
	ExtStaffID    string                        `gorm:"type:char(32);uniqueIndex:idx_ext_staff_id_ext_customer_id;comment:微信员工ID" json:"ext_staff_id"`
	Age           int                           `gorm:"type:tinyint(3);comment:年龄" json:"age"`
	Description   string                        `gorm:"type:text;comment:描述" json:"description"`
	Email         string                        `gorm:"type:varchar(32);comment:邮箱" json:"email"`
	PhoneNumber   string                        `gorm:"type:char(32);comment:电话" json:"phone_number"`
	QQ            string                        `gorm:"type:varchar(16);comment:qq" json:"qq"`
	Address       string                        `gorm:"type:varchar(128);comment:地址" json:"address"`
	Birthday      string                        `gorm:"type:char(10);comment:生日" json:"birthday"`
	Weibo         string                        `gorm:"type:varchar(128);comment:微博" json:"weibo" `
	RemarkField   constants.CustomerRemarkField `gorm:"type:json;comment:自定义字段的值"  json:"remark_values"`
	Timestamp
}

func (ci CustomerInfo) Upsert(info CustomerInfo) error {
	return DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ext_customer_id"}, {Name: "ext_staff_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"ext_customer_id", "ext_staff_id"}),
	}).Create(&info).Error
}

func (ci CustomerInfo) Update(info CustomerInfo) error {
	return DB.Model(&CustomerInfo{}).Omit("id", "ext_customer_id", "ext_staff_id").
		Where("ext_customer_id = ?", info.ExtCustomerID).
		Where("ext_staff_id = ?", info.ExtStaffID).
		Updates(&info).Error
}

func (ci CustomerInfo) Get(info CustomerInfo) (res CustomerInfo, err error) {

	db := DB.Model(&CustomerInfo{}).Where("ext_corp_id = ?", info.ExtCorpID)
	if info.ExtStaffID != "" {
		db = db.Where("ext_staff_id = ?", info.ExtStaffID)
	}
	if info.ExtCustomerID != "" {
		db = db.Where("ext_customer_id = ?", info.ExtCustomerID)
	}

	err = db.Find(&res).Error
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}
