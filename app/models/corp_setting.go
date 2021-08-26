package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
)

type CorpSetting struct {
	// ID
	ID string `json:"id" gorm:"primaryKey;type:bigint;comment:'ID'" validate:"int64"`
	// ExtCorpID 外部企业ID
	ExtCorpID string `json:"ext_corp_id" gorm:"uniqueIndex:ext_corp_id;;type:char(18);comment:外部企业ID" validate:ext_corp_id"`
	// ExtCreatorID 创建者外部员工ID
	ExtCreatorID   string            `json:"ext_creator_id" gorm:"index;type:char(32);comment:创建者外部员工ID" validate:"word"`
	IsMaterialUsed constants.Boolean `json:"is_material_used" gorm:"type:tinyint unsigned"`
	Timestamp
}

func (o CorpSetting) Update(c CorpSetting) error {
	return DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ext_corp_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"is_material_used"}),
	}).Create(&c).Error
}

func (o CorpSetting) Get(extCorpID string) (CorpSetting, error) {
	var c CorpSetting
	err := DB.Model(&CorpSetting{}).Where("ext_corp_id = ?", extCorpID).First(&c).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return CorpSetting{ExtCorpID: extCorpID, IsMaterialUsed: constants.False}, nil
		}
		return c, err
	}
	return c, nil
}
