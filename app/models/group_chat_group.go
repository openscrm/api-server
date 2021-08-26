package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/app"
)

type GroupChatGroup struct {
	ExtCorpModel
	Name string `json:"name" gorm:"type:text"`
	// IsDefault 是否为默认分组，1：是；2：否
	IsDefault constants.Boolean `json:"is_default" gorm:"default:2;comment:'是否为默认分组，1：是；2：否'"`
	Timestamp
}

func (o GroupChatGroup) Get(groupID string) (GroupChatGroup, error) {
	var g GroupChatGroup
	err := DB.Model(GroupChatGroup{}).Where("id = ?", groupID).First(&g).Error

	return g, err
}

func (o GroupChatGroup) Create(group GroupChatGroup) error {
	return DB.Create(&group).Error
}

func (o GroupChatGroup) Update(group GroupChatGroup) error {
	return DB.Model(&GroupChatGroup{}).
		Where("id = ?", group.ID).Updates(map[string]interface{}{"name": group.Name}).Error
}

func (o GroupChatGroup) Delete(ids []string, extCorpID string) (int64, error) {
	result := DB.Where("ext_corp_id = ?", extCorpID).Where("id in (?)", ids).Delete(&GroupChatGroup{})
	err := result.Error
	if err != nil {
		err = errors.Wrap(err, "Delete ContactWayGroup failed")
		return 0, err
	}

	return result.RowsAffected, nil
}

func (o GroupChatGroup) Query(group GroupChatGroup, extCorpID string, sorter *app.Sorter, pager *app.Pager) ([]GroupChatGroup, int64, error) {
	items := make([]GroupChatGroup, 0)
	db := DB.Model(&GroupChatGroup{}).Where("ext_corp_id = ?", extCorpID)
	if group.Name != "" {
		db = db.Where("name like ?", group.Name+"%")
		group.Name = ""
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count GroupChatGroup failed")
		return nil, 0, err
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err = db.Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find GroupChatGroup failed")
		return nil, 0, err
	}

	return items, total, nil
}
