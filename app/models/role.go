package models

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/log"
	"openscrm/common/redis"
	"openscrm/conf"
	"os"
	"strings"
	"time"
)

// Role 角色
type Role struct {
	ExtCorpModel
	// Name 角色名称
	Name string `json:"name" gorm:"index;comment:'角色名称'"`
	// Description 角色描述
	Description string `json:"description" gorm:"comment:'角色描述'"`
	// Type 角色类型
	Type string `json:"type" gorm:"index;default:Staff;comment:'角色类型'" validate:"oneof=Admin DepartmentAdmin Staff"`
	// SortWeight 角色排序权重
	SortWeight int `json:"sort_weight" gorm:"index;default:1000;comment:'角色排序权重'" validate:"gte=0"`
	// IsDefault 是否为默认角色，1：是；2：否
	IsDefault constants.Boolean `json:"is_default" gorm:"default:2;comment:'是否为默认角色，1：是；2：否'" validate:"oneof=1 2"`
	// PermissionIDs 角色绑定的权限标识数组
	PermissionIDs constants.StringArrayField `json:"permission_ids" gorm:"type:json;comment:'角色绑定的权限标识数组'" validate:"dive,word"`
	Timestamp
}

func (o Role) Query(param Role, extCorpID string, sorter *app.Sorter, pager *app.Pager) (items []Role, total int64, err error) {
	items = make([]Role, 0)
	db := DB.Model(&Role{}).Where("ext_corp_id = ?", extCorpID)
	db = db.Where(param)
	err = db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count Role failed")
		return
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err = db.Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find Role failed")
		return
	}

	return
}

func (o Role) CachedGet(id string) (item Role, err error) {
	err = redis.GetOrSetFunc(fmt.Sprintf(constants.CachedRoleKey, id), func() (interface{}, error) {
		return o.Get(id)
	}, time.Hour*24*7, &item)
	if item.ID == "" {
		err = errors.New("invalid role")
		return
	}
	return
}

func (o Role) Get(id string) (item Role, err error) {
	err = DB.Model(&Role{}).Preload(clause.Associations).Where("id = ?", id).First(&item).Error
	if err == gorm.ErrRecordNotFound {
		err = errors.WithStack(ecode.ItemNotFoundError)
		return
	}
	if err != nil {
		err = errors.Wrap(err, "First Role failed")
		return
	}

	return
}

func (o Role) Create(param Role, extCorpID string) (item Role, err error) {
	err = copier.Copy(&item, param)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}

	item.ExtCorpID = extCorpID

	err = DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&item).Error
	if err != nil {
		err = errors.Wrap(err, "Create Role failed")
		return
	}

	return
}

func (o Role) Update(id string, param Role) (item Role, err error) {
	tx := DB.Begin()
	defer tx.Rollback()
	err = tx.Where("id = ?", id).First(&item).Error
	if err == gorm.ErrRecordNotFound {
		err = errors.WithStack(ecode.ItemNotFoundError)
		return
	}
	if err != nil {
		err = errors.Wrap(err, "First Role failed")
		return
	}
	param.ExtCorpID = ""

	if item.IsDefault == constants.True {
		err = errors.WithStack(ecode.DoNotUpdateDefaultRoleError)
		return
	}

	err = tx.Model(&item).Updates(param).Error
	if err != nil {
		err = errors.Wrap(err, "Update Role failed")
		return
	}

	err = o.CleanCache(id)
	if err != nil {
		err = errors.Wrap(err, "cleanCache failed")
		return
	}

	err = tx.Commit().Error
	if err != nil {
		err = errors.Wrap(err, "tx.Commit() failed")
		return
	}

	return
}

func (o Role) AssignToStaffs(extStaffIDs []string, roleID string) (total int64, err error) {
	role, err := o.Get(roleID)
	if err != nil {
		err = errors.Wrap(err, "Get role failed")
		return
	}

	result := DB.Model(&Staff{}).Where("ext_corp_id = ?", role.ExtCorpID).Where("ext_id in (?)", extStaffIDs).
		Updates(&Staff{RoleID: roleID, RoleType: role.Type})
	if result.Error != nil {
		err = errors.Wrap(err, "update staffs failed")
		return
	}

	total = result.RowsAffected

	return
}

func (o Role) CleanAllCache() (err error) {
	keys := make([]string, 0)
	keys, err = redis.RedisClient.Keys(context.Background(), strings.ReplaceAll(constants.CachedRoleKey, "%s", "*")).Result()
	if err != nil {
		err = errors.Wrap(err, "redis.RedisClient.Keys failed")
		return
	}
	if len(keys) == 0 {
		return
	}

	err = redis.Delete(keys...)
	if err != nil {
		err = errors.Wrap(err, "delete cached roles failed")
		return
	}
	return
}

func (o Role) CleanCache(ids ...string) (err error) {
	keys := make([]string, 0)
	for _, id := range ids {
		keys = append(keys, fmt.Sprintf(constants.CachedRoleKey, id))
	}

	err = redis.Delete(keys...)
	if err != nil {
		err = errors.Wrap(err, "delete cached roles failed")
		return
	}
	return
}

func (o Role) QueryStaffs(param Staff, extCorpID string, sorter *app.Sorter, pager *app.Pager) (items []Staff, total int64, err error) {
	items = make([]Staff, 0)
	db := DB.Model(&Staff{}).Where("ext_corp_id = ?", extCorpID)
	if param.Name != "" {
		db = db.Where("name like ?", param.Name+"%")
		param.Name = ""
	}

	db = db.Where(param)
	err = db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count Role failed")
		return
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err = db.Preload("Departments").Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find Staff failed")
		return
	}

	return
}

func pluckPermissionIDs(in []constants.Permission) []string {
	ids := make([]string, 0)
	for _, permission := range in {
		ids = append(ids, fmt.Sprintf("%s_%s", permission.BizIdentity, permission.Operation))
	}
	return funk.UniqString(ids)
}

func SetupRoles() {
	tx := DB.Begin()
	defer tx.Rollback()
	roles := []Role{
		{
			ExtCorpModel: ExtCorpModel{
				ID:        string(constants.DefaultCorpStaffRoleID),
				ExtCorpID: conf.Settings.WeWork.ExtCorpID,
			},
			Name:          "员工",
			Description:   "员工",
			Type:          string(constants.RoleTypeStaff),
			IsDefault:     constants.True,
			PermissionIDs: pluckPermissionIDs(constants.StaffPermissions),
			SortWeight:    10000,
		},
		{
			ExtCorpModel: ExtCorpModel{
				ID:        string(constants.DefaultCorpDepartmentAdminRoleID),
				ExtCorpID: conf.Settings.WeWork.ExtCorpID,
			},
			Name:          "部门管理员",
			Description:   "部门管理员",
			Type:          string(constants.RoleTypeDepartmentAdmin),
			IsDefault:     constants.True,
			PermissionIDs: pluckPermissionIDs(constants.DepartmentAdminPermissions),
			SortWeight:    10001,
		},
		{
			ExtCorpModel: ExtCorpModel{
				ID:        string(constants.DefaultCorpAdminRoleID),
				ExtCorpID: conf.Settings.WeWork.ExtCorpID,
			},
			Name:          "管理员",
			Description:   "管理员",
			Type:          string(constants.RoleTypeAdmin),
			IsDefault:     constants.True,
			PermissionIDs: pluckPermissionIDs(constants.AdminPermissions),
			SortWeight:    10002,
		},
		{
			ExtCorpModel: ExtCorpModel{
				ID:        string(constants.DefaultCorpSuperAdminRoleID),
				ExtCorpID: conf.Settings.WeWork.ExtCorpID,
			},
			Name:          "超级管理员",
			Description:   "超级管理员",
			Type:          string(constants.RoleTypeSuperAdmin),
			IsDefault:     constants.True,
			PermissionIDs: pluckPermissionIDs(constants.SuperAdminPermissions),
			SortWeight:    10003,
		},
	}

	err := tx.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{`name`, `sort_weight`, `description`, `type`, `is_default`, `permission_ids`, `updated_at`}),
	}).Create(&roles).Error
	if err != nil {
		log.TracedError("SetupRoles failed", errors.WithStack(err))
		os.Exit(1)
	}

	err = tx.Commit().Error
	if err != nil {
		log.TracedError("Commit failed", errors.WithStack(err))
		os.Exit(1)
	}

	//err = Role{}.CleanAllCache()
	//if err != nil {
	//	log.TracedError("CleanAllCache failed", errors.WithStack(err))
	//	os.Exit(1)
	//}
}
