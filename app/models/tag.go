package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Tag struct {
	ExtCorpModel
	ExtID      string `gorm:"type:char(32);uniqueIndex;comment:外部标签ID" json:"ext_id"`
	ExtGroupID string `gorm:"type:char(32);index;comment:外部标签组ID" json:"ext_group_id"`
	Name       string `gorm:"type:char(255);uniqueIndex:idx_group_name_tag_name;comment:标签名称" json:"name"`
	GroupName  string `gorm:"type:char(255);uniqueIndex:idx_group_name_tag_name;comment:标签组名称" json:"group_name"`
	CreateTime int    `gorm:"type:int;comment:创建时间" json:"create_time"`
	Order      uint32 `gorm:"type:int;index;comment:标签排序值，值大的在前" json:"order"`
	Type       int    `gorm:"type:tinyint;comment:所打标签类型, 1-企业设置, 2-用户自定义" json:"type"`
	Timestamp
}

func (t Tag) Query(extIDs []string) ([]Tag, error) {
	var tags []Tag
	err := DB.Model(&Tag{}).Where("ext_id in (?)", extIDs).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func (t Tag) GetExistTag(extTagGroupId string, tagNames []string) ([]Tag, error) {
	var tags []Tag
	db := DB.Model(&Tag{}).Where("ext_group_id = ?", extTagGroupId)
	if tagNames != nil && len(tagNames) > 0 {
		db = db.Where("name in (?)", tagNames)
	}
	err := db.Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil

}

func (t Tag) Delete(ids []string) error {
	return DB.Where("ext_id in ?", ids).Delete(&Tag{}).Error
}

func (t Tag) CreateTags(tx *gorm.DB, tagModels []Tag) error {
	return tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ext_tag_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"create_time", "group_name", "order", "ext_group_id", "name"})},
	).CreateInBatches(&tagModels, len(tagModels)).Error
	//return tx.CreateInBatches(tagModels, len(tagModels)).Error
}

func (t Tag) Upsert(tag Tag) error {
	return DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ext_tag_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"create_time", "group_name", "order", "ext_group_id", "name"})},
	).Create(&tag).Error
}
func (t Tag) Update(tag Tag) error {
	return DB.Model(&Tag{}).Where("ext_id = ?", tag.ExtID).Updates(&Tag{Name: tag.Name}).Error
}

func (t Tag) GetCurMaxOrder(tx *gorm.DB, id string) (order int64, err error) {
	err = tx.Model(&Tag{}).Where("ext_group_id = ?", id).Select("max(distinct `order`)").Group("ext_group_id").Scan(&order).Error
	return
}

func (t Tag) Get(extIDs []string) (tags []Tag, err error) {
	err = DB.Model(&Tag{}).Where("ext_id in (?)", extIDs).Find(&tags).Error
	return
}
