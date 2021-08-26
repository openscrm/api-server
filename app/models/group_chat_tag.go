package models

import "gorm.io/gorm/clause"

type GroupChatTag struct {
	ExtCorpModel
	GroupChatTagGroupID string `json:"group_chat_tag_group_id" gorm:"uniqueIndex:idx_group_id_name"`
	Name                string `json:"name" gorm:"type:char(32);uniqueIndex:idx_group_id_name"`
	Timestamp
}

func (o GroupChatTag) Create(tags []GroupChatTag) error {
	return DB.CreateInBatches(&tags, 100).Error
}

func (o GroupChatTag) Delete(ids []string) (int64, error) {
	res := DB.Model(&GroupChatTag{}).Where("id in ?", ids).Delete(&GroupChatTag{})
	return res.RowsAffected, res.Error
}

func (o GroupChatTag) Get(id string) (tag GroupChatTag, err error) {
	err = DB.Model(&GroupChatTag{}).Where("id =?", id).First(&tag).Error
	return
}

func (o GroupChatTag) Updates(tag GroupChatTag) error {
	return DB.Model(&GroupChatTag{}).Where("id = ?", tag.ID).Update("name", tag.Name).Error
}

func (o GroupChatTag) Upsert(tags []GroupChatTag) error {
	err := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "group_chat_tag_group_id"}),
	}).Create(&tags).Error

	return err
}
