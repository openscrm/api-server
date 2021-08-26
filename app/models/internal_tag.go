package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/app"
)

type InternalTag struct {
	ExtCorpModel
	ExtStaffID string `gorm:"index;type:char(32)" json:"ext_staff_id"`
	Name       string `gorm:"type:char(32)" json:"name"`
	Timestamp
}

func (o InternalTag) Query(tag InternalTag, sorter *app.Sorter, pager *app.Pager) ([]InternalTag, int64, error) {
	items := make([]InternalTag, 0)
	db := DB.Model(&InternalTag{}).Where("ext_corp_id = ?", tag.ExtCorpID)
	total := int64(0)
	err := db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count InternalTag failed")
		return nil, 0, err
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err = db.Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find InternalTag failed")
		return nil, 0, err
	}

	return items, 0, err
}
func (o InternalTag) CreateInBatches(tag []InternalTag) error {
	return DB.Model(InternalTag{}).CreateInBatches(&tag, 100).Error
}

func (o InternalTag) Create(tag InternalTag) error {
	return DB.Create(&tag).Error
}

func (o InternalTag) Delete(ids []string, extCorpID string) (int64, error) {
	res := DB.Where("ext_corp_id  = ?", extCorpID).Where("id in (?)", ids).Delete(&InternalTag{})
	return res.RowsAffected, res.Error
}

func (o InternalTag) GetByIDs(ids []string) (tags []InternalTag, err error) {
	err = DB.Model(&InternalTag{}).Where("id in (?)", ids).Find(&tags).Error
	return
}
