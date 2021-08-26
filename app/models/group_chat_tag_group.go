package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/app"
)

type GroupChatTagGroup struct {
	ExtCorpModel
	Name string         `gorm:"type:char(32);uniqueIndex:name;" json:"name"`
	Tags []GroupChatTag `gorm:"foreignKey:GroupChatTagGroupID;references:ID" json:"tags"`
}

func (o GroupChatTagGroup) Create(group GroupChatTagGroup) error {
	return DB.Model(&GroupChatTagGroup{}).Create(&group).Error
}

func (o GroupChatTagGroup) Query(group GroupChatTagGroup, pager *app.Pager, sorter *app.Sorter) (groups []GroupChatTagGroup, total int64, err error) {
	db := DB.Model(&GroupChatTagGroup{}).Where("ext_corp_id = ?", group.ExtCorpID)
	if group.Name != "" {
		db = db.Where("name like ?", group.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count  GroupChatTagGroup failed")
		return
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err = db.Preload("Tags").Find(&groups).Error
	if err != nil {
		err = errors.Wrap(err, "Find GroupChatTagGroup failed")
		return
	}
	return
}

func (o GroupChatTagGroup) Get(id string) (group GroupChatTagGroup, err error) {
	err = DB.Model(&GroupChatTagGroup{}).Preload("Tags").Where("id = ?", id).First(&group).Error
	return
}

func (o GroupChatTagGroup) Update(group GroupChatTagGroup) error {
	return DB.Model(GroupChatTagGroup{}).Where("id = ?", group.ID).Updates(&group).Error
}

func (o GroupChatTagGroup) Delete(ids []string) (int64, error) {
	res := DB.Model(&GroupChatTagGroup{}).Where("id in ?", ids).Delete(&GroupChatTagGroup{})
	return res.RowsAffected, res.Error
}
