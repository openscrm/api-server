package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/app"
)

type MaterialLibTag struct {
	ExtCorpModel
	Name string `json:"name"`
	Timestamp
}

func (o MaterialLibTag) Create(tag []MaterialLibTag) error {
	return DB.CreateInBatches(&tag, 100).Error
}

func (o MaterialLibTag) Delete(ids []string) (int64, error) {
	res := DB.Model(&MaterialLibTag{}).Where("id in (?)", ids).Delete(&MaterialLibTag{})
	return res.RowsAffected, res.Error
}
func (o MaterialLibTag) Query(tag MaterialLibTag, sorter *app.Sorter, pager *app.Pager) ([]MaterialLibTag, int64, error) {
	db := DB.Model(&MaterialLibTag{}).Where("ext_corp_id = ?", tag.ExtCorpID)
	if tag.Name != "" {
		db = db.Where("name like ?", tag.Name)
	}
	var total int64
	tags := make([]MaterialLibTag, 0)
	err := db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count MaterialLibTag failed")
		return tags, 0, err
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err = db.Find(&tags).Error
	if err != nil {
		err = errors.Wrap(err, "Find MaterialLibTag failed")
		return tags, 0, err
	}
	return tags, total, nil
}

func (o MaterialLibTag) GetByIDs(tagIDs []string) (res []MaterialLibTag, err error) {
	err = DB.Model(&MaterialLibTag{}).Where("id in (?)", tagIDs).Find(&res).Error
	return
}
