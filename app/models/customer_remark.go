package models

import (
	"gorm.io/gorm"
)

// CustomerRemark 自定义信息
type CustomerRemark struct {
	ExtCorpModel
	Name         string         `gorm:"type:char(32);uniqueIndex:idx_corp_id_name;" json:"name"`
	FieldType    string         `json:"field_type"` // todo rename to remark_type
	HasStaffUsed bool           `json:"has_staff_used"`
	RankNum      int            `gorm:"type:int unsigned" json:"rank_num"`
	Options      []RemarkOption `gorm:"foreignKey:RemarkID" json:"info_option"`
	Timestamp
}

func (r CustomerRemark) ExchangeOrder(id string, id2 string) error {
	rawSQL := `update custom_remark a, custom_remark b
	set a.rank_num = b.rank_num,
		b.rank_num= a.rank_num
	where a.id =? and b.id = ?;`
	return DB.Exec(rawSQL, id, id2).Error
}

// RemarkOption 对于多选类型信息的选项
type RemarkOption struct {
	Model
	RemarkID string `json:"remark_id" gorm:"type:bigint;uniqueIndex:idx_remark_id_name"`
	Name     string `json:"name" gorm:"type:char(32);uniqueIndex:idx_remark_id_name"`
	Timestamp
}

func (r CustomerRemark) Create(remark CustomerRemark) error {
	tx := DB.Begin()
	defer tx.Rollback()
	err := DB.Create(&remark).Error
	if err != nil {
		return err
	}
	updateRankNumSQL := `UPDATE custom_remark a
	inner join ( select corp_id, max(rank_num) as rank_num from custom_remark where corp_id = ? group by corp_id) b on a.corp_id = b.corp_id
		set a.rank_num=b.rank_num + 1
		where a.id = ? and a.corp_id = ?;`
	return DB.Exec(updateRankNumSQL, remark.ExtCorpID, remark.ID, remark.ExtCorpID).Error
}

func (r CustomerRemark) Delete(ids []string, extCorpID string) error {
	return DB.Model(&CustomerRemark{}).Where("ext_corp_id = ?", extCorpID).Where("id in (?)", ids).Delete(&CustomerRemark{}).Error
}

func (r CustomerRemark) Update(remark CustomerRemark) error {
	return DB.Updates(&remark).Error
}

func (r CustomerRemark) Get(extCorpID string) ([]*CustomerRemark, error) {
	var customerRemarks []*CustomerRemark
	if err := DB.Model(&CustomerRemark{}).Preload("Options").Find(&customerRemarks, "ext_corp_id = ?", extCorpID).Error; err != nil {
		return nil, err
	}
	return customerRemarks, nil
}

//------------------------------

func (o RemarkOption) Create(remark RemarkOption) error {
	return DB.Create(&remark).Error
}

func (o RemarkOption) GetTextOption(db *gorm.DB) (*RemarkOption, error) {
	option := &RemarkOption{}
	err := db.Model(&o).First(option).Error
	if err != nil {
		return nil, err
	}
	return option, nil
}

func (o RemarkOption) Update(option *RemarkOption) error {
	return DB.Model(&RemarkOption{}).
		Where("id = ?", option.ID).
		Updates(option).Error
}

func (o RemarkOption) Delete(ids []string) error {
	return DB.Model(&RemarkOption{}).Where("id in (?)", ids).Delete(&RemarkOption{}).Error
}
