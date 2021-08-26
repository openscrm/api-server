package models

import (
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
)

type CustomerStaffTag struct {
	ExtCorpModel
	CustomerStaffID string `json:"customer_staff_id" gorm:"type:bigint;uniqueIndex:ext_tag_id_cs_id"`
	// TagID 标签id
	ExtTagID string `json:"ext_tag_id" gorm:"type:char(32);uniqueIndex:ext_tag_id_cs_id"`
	// GroupName 该成员添加此外部联系人所打标签的分组名称（标签功能需要企业微信升级到2.7.5及以上版本）
	GroupName string `json:"group_name"`
	// TagName 该成员添加此外部联系人所打标签名称
	TagName string `json:"tag_name"`
	// Type 该成员添加此外部联系人所打标签类型, 1-企业设置, 2-用户自定义
	Type constants.FollowUserTagType `gorm:"type:tinyint" json:"type"`
	Timestamp
}

// CreateInBatches 批量创建

func (c CustomerStaffTag) CreateInBatches(customerStaffTags []CustomerStaffTag) error {
	return DB.CreateInBatches(customerStaffTags, len(customerStaffTags)).Error
}

func (c CustomerStaffTag) Delete(customerStaffId string, extTagsIDs []string, in bool) error {
	var db = DB
	if extTagsIDs != nil && len(extTagsIDs) > 0 {
		if !in {
			db = DB.Where("ext_tag_id not in (?)", extTagsIDs)
		} else {
			db = DB.Where("ext_tag_id in (?)", extTagsIDs)
		}
	}
	if customerStaffId != "" {
		db = db.Where("customer_staff_id = ?", customerStaffId)
	}
	return db.Delete(&CustomerStaffTag{}).Error
}

// Upsert
// Description: 更新员工给客户的标签
// Detail: 需要将已经删除的标签恢复
func (c CustomerStaffTag) Upsert(tag []CustomerStaffTag) error {
	return DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "customer_staff_id"}, {Name: "ext_tag_id"}},
		DoUpdates: clause.AssignmentColumns(
			[]string{"group_name", "ext_tag_id", "type", "tag_name", "deleted_at"},
		),
	}).CreateInBatches(&tag, len(tag)).Error
}
