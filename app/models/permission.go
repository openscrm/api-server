package models

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"os"
)

// Permission 权限
type Permission struct {
	Model
	// Name 权限名称
	Name string `json:"name" gorm:"index;comment:'权限名称'"`
	// Description 权限描述
	Description string `json:"description" gorm:"comment:'权限描述'"`
	// BizName 业务名称
	BizName string `json:"biz_name" gorm:"comment:'业务名称'"`
	// BizIdentity 业务标识
	BizIdentity string `json:"biz_identity" gorm:"index;comment:'业务标识'" validate:"omitempty,alphanum"`
	// Operation 操作
	Operation string `json:"operation" gorm:"comment:'操作'" validate:"omitempty,oneof=read full"`
	// Identity 权限标识
	Identity string `json:"identity" gorm:"unique;size:50;comment:'权限标识'" validate:"omitempty,alphanum,max=50"`
	// SortWeight 权限排序权重
	SortWeight int `json:"sort_weight" gorm:"index;default:1000;comment:'权限排序权重'" validate:"gte=0"`
	Timestamp
}

func (o Permission) Query(param Permission, sorter *app.Sorter, pager *app.Pager) (items []Permission, total int64, err error) {
	items = make([]Permission, 0)
	db := DB.Model(&Permission{})
	if param.Name != "" {
		db = db.Where("name like ?", param.Name+"%")
		param.Name = ""
	}

	db = db.Where(param)
	err = db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count Permission failed")
		return
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err = db.Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find Permission failed")
		return
	}

	return
}

func (o Permission) Get(id string) (item Permission, err error) {
	err = DB.Model(&Permission{}).Where("id = ?", id).First(&item).Error
	if err == gorm.ErrRecordNotFound {
		err = errors.WithStack(ecode.ItemNotFoundError)
		return
	}
	if err != nil {
		err = errors.Wrap(err, "First Permission failed")
		return
	}

	return
}

func (o Permission) BatchGetByIdentities(identities ...string) (items []Permission, err error) {
	err = DB.Model(&Permission{}).Where("identity in (?)", identities).Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find Permissions failed")
		return
	}

	return
}

func (o Permission) Create(param Permission) (item Permission, err error) {
	err = copier.Copy(&item, param)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}

	err = DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&item).Error
	if err != nil {
		err = errors.Wrap(err, "Create Permission failed")
		return
	}

	return
}

func (o Permission) Update(id string, param Permission) (item Permission, err error) {
	err = DB.Where("id = ?", id).First(&item).Error
	if err == gorm.ErrRecordNotFound {
		err = errors.WithStack(ecode.ItemNotFoundError)
		return
	}
	if err != nil {
		err = errors.Wrap(err, "First Permission failed")
		return
	}

	err = DB.Session(&gorm.Session{FullSaveAssociations: true}).Model(&item).Updates(param).Error
	if err != nil {
		err = errors.Wrap(err, "Update Permission failed")
		return
	}

	return
}

func (o Permission) Delete(ids []string) (total int64, err error) {
	result := DB.Where("id in (?)", ids).Delete(&Permission{})
	err = result.Error
	if err != nil {
		err = errors.Wrap(err, "Delete Permission failed")
		return
	}
	total = result.RowsAffected

	return
}

func (o Permission) BatchUpsert(items []*Permission) (err error) {
	err = DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{
			Name: "identify",
		}},
		DoUpdates: clause.AssignmentColumns([]string{`name`, `description`, `biz_name`, `biz_identity`, `operation`, `updated_at`}),
		UpdateAll: false,
	}).CreateInBatches(&items, 500).Error
	if err != nil {
		err = errors.Wrap(err, "BatchUpsert Permission failed")
		return
	}

	return
}

func SetupPermissions() {
	items := make([]*Permission, 0)
	for _, permission := range constants.AdminPermissions {
		item := &Permission{
			Model:       Model{ID: id_generator.StringID()},
			Name:        permission.Name,
			Description: permission.Name,
			BizIdentity: string(permission.BizIdentity),
			Operation:   string(permission.Operation),
			Identity:    fmt.Sprintf("%s_%s", permission.BizIdentity, permission.Operation),
		}
		items = append(items, item)
	}

	for _, permission := range constants.SuperAdminPermissions {
		item := &Permission{
			Model:       Model{ID: id_generator.StringID()},
			Name:        permission.Name,
			Description: permission.Name,
			BizIdentity: string(permission.BizIdentity),
			Operation:   string(permission.Operation),
			Identity:    fmt.Sprintf("%s_%s", permission.BizIdentity, permission.Operation),
		}
		items = append(items, item)
	}

	result := funk.Uniq(items)
	var ok bool
	items, ok = result.([]*Permission)
	if !ok {
		log.Logger.Error("SetupPermissions failed")
		os.Exit(1)
	}

	err := (&Permission{}).BatchUpsert(items)
	if err != nil {
		log.TracedError("SetupPermissions failed", err)
		os.Exit(1)
	}
}
