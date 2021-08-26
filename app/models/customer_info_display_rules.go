package models

import (
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/id_generator"
	"openscrm/conf"
)

// CustomerInfoDisplayRule 是否显示信息
type CustomerInfoDisplayRule struct {
	Model
	ExtCorpID   string            `gorm:"type:string;size:100;uniqueIndex;comment:企业ID" json:"ext_corp_id"`
	Age         constants.Boolean `gorm:"type:tinyint unsigned;comment:是否展示年龄" json:"age"`
	Description constants.Boolean `gorm:"type:tinyint unsigned;comment:是否展示描述" json:"description"`
	Email       constants.Boolean `gorm:"type:tinyint unsigned;comment:是否展示邮箱" json:"email"`
	PhoneNumber constants.Boolean `gorm:"type:tinyint unsigned;comment:是否展示电话" json:"phone_number"`
	QQ          constants.Boolean `gorm:"type:tinyint unsigned;comment:是否展示mqq" json:"qq"`
	Address     constants.Boolean `gorm:"type:tinyint unsigned;comment:是否展示地址" json:"address"`
	Birthday    constants.Boolean `gorm:"type:tinyint unsigned;comment:是否展示生日" json:"birthday"`
	Weibo       constants.Boolean `gorm:"type:tinyint unsigned;comment:是否展示微博" json:"weibo" `
	Timestamp
}

func (r CustomerInfoDisplayRule) Get(rules *CustomerInfoDisplayRule) (*CustomerInfoDisplayRule, error) {
	infoRules := &CustomerInfoDisplayRule{}
	if err := DB.Model(&CustomerInfoDisplayRule{}).
		Where(" ext_corp_id = ?", rules.ExtCorpID).
		First(&infoRules).Error; err != nil {
		return nil, err
	}
	return infoRules, nil
}

func (r CustomerInfoDisplayRule) Updates(id string, rule map[string]interface{}) error {
	return DB.Model(&CustomerInfoDisplayRule{}).Where("ext_corp_id = ?", id).Updates(&rule).Error
}

func SetupCustomerInfoDisplayRule() {
	rule := CustomerInfoDisplayRule{
		Model:       Model{ID: id_generator.StringID()},
		ExtCorpID:   conf.Settings.WeWork.ExtCorpID,
		Age:         constants.True,
		Description: constants.True,
		Email:       constants.True,
		PhoneNumber: constants.True,
		QQ:          constants.True,
		Address:     constants.True,
		Birthday:    constants.True,
		Weibo:       constants.True,
	}
	err := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ext_corp_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"age", "description", "email", "address", "qq", "weibo", "birthday", "phone_number"}),
	}).Create(&rule).Error
	if err != nil {
		panic(err)
	}
}
