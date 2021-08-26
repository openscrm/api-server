package models

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/log"
	"openscrm/conf"
	"os"
)

// ContactWayGroup 渠道码分组
type ContactWayGroup struct {
	ExtCorpModel
	// Name 分组名称
	Name string `json:"name" gorm:"index;comment:'分组名称'"`
	// SortWeight 分组排序权重
	SortWeight int `json:"sort_weight" gorm:"index;default:1000;comment:'分组排序权重'" validate:"gte=0"`
	// Count 该分组渠道码数量
	Count int `json:"count" gorm:"comment:'该分组渠道码数量'"`
	// IsDefault 是否为默认分组，1：是；2：否
	IsDefault constants.Boolean `json:"is_default" gorm:"default:2;comment:'是否为默认分组，1：是；2：否'" validate:"oneof=1 2"`
	Timestamp
}

func (o ContactWayGroup) Query(param ContactWayGroup, extCorpID string, sorter *app.Sorter, pager *app.Pager) (items []ContactWayGroup, total int64, err error) {
	items = make([]ContactWayGroup, 0)
	db := DB.Model(&ContactWayGroup{}).Where("ext_corp_id = ?", extCorpID)
	if param.Name != "" {
		db = db.Where("name like ?", param.Name+"%")
		param.Name = ""
	}

	db = db.Where(param)
	err = db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count ContactWayGroup failed")
		return
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err = db.Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find ContactWayGroup failed")
		return
	}

	return
}

func (o ContactWayGroup) Get(id string, extCorpID string) (item ContactWayGroup, err error) {
	err = DB.Model(&ContactWayGroup{}).Where("ext_corp_id = ?", extCorpID).Where("id = ?", id).First(&item).Error
	if err == gorm.ErrRecordNotFound {
		err = errors.WithStack(ecode.ItemNotFoundError)
		return
	}
	if err != nil {
		err = errors.Wrap(err, "First ContactWayGroup failed")
		return
	}

	return
}

func (o ContactWayGroup) Create(param ContactWayGroup, extCorpID string) (item ContactWayGroup, err error) {
	err = copier.Copy(&item, param)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}
	item.ExtCorpID = extCorpID

	err = DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&item).Error
	if err != nil {
		err = errors.Wrap(err, "Create ContactWayGroup failed")
		return
	}

	return
}

func (o ContactWayGroup) Update(id string, param ContactWayGroup, extCorpID string) (item ContactWayGroup, err error) {
	err = DB.Where("ext_corp_id = ?", extCorpID).Where("id = ?", id).First(&item).Error
	if err == gorm.ErrRecordNotFound {
		err = errors.WithStack(ecode.ItemNotFoundError)
		return
	}
	if err != nil {
		err = errors.Wrap(err, "First ContactWayGroup failed")
		return
	}
	param.ExtCorpID = ""

	err = DB.Session(&gorm.Session{FullSaveAssociations: true}).Model(&item).Updates(param).Error
	if err != nil {
		err = errors.Wrap(err, "Update ContactWayGroup failed")
		return
	}

	return
}

func (o ContactWayGroup) Delete(ids []string, extCorpID string) (total int64, err error) {
	result := DB.Where("ext_corp_id = ?", extCorpID).Where("id in (?)", ids).Delete(&ContactWayGroup{})
	err = result.Error
	if err != nil {
		err = errors.Wrap(err, "Delete ContactWayGroup failed")
		return
	}
	total = result.RowsAffected

	return
}

func SetupContactWayGroup() {
	item := &ContactWayGroup{
		ExtCorpModel: ExtCorpModel{
			ID:           string(constants.DefaultContactWayGroupID),
			ExtCorpID:    conf.Settings.WeWork.ExtCorpID,
			ExtCreatorID: "",
		},
		Name:       "默认分组",
		SortWeight: 0,
		Count:      0,
		IsDefault:  constants.True,
	}
	err := DB.FirstOrCreate(&item).Error
	if err != nil {
		log.TracedError("ContactWayGroup failed", errors.WithStack(err))
		os.Exit(1)
	}
}
