package requests

import (
	"openscrm/app/constants"
	"openscrm/common/app"
)

// CreateRoleReq 创建角色请求参数
type CreateRoleReq struct {
	// Name 角色名称
	Name string `json:"name" gorm:"index;comment:'角色名称'" validate:"required"`
	// Description 角色描述
	Description string `json:"description" gorm:"comment:'角色描述'"`
	// SortWeight 角色排序权重
	SortWeight int `json:"sort_weight" gorm:"index;default:1000;comment:'角色排序权重'" validate:"gte=0"`
	// PermissionIDs 角色绑定的权限标识数组
	PermissionIDs constants.StringArrayField `json:"permission_ids" gorm:"type:json;comment:'角色绑定的权限标识数组'" validate:"dive,word"`
}

// UpdateRoleReq 更新角色请求参数
type UpdateRoleReq struct {
	// Name 角色名称
	Name string `json:"name" gorm:"index;comment:'角色名称'"`
	// Description 角色描述
	Description string `json:"description" gorm:"comment:'角色描述'"`
	// SortWeight 角色排序权重
	SortWeight int `json:"sort_weight" gorm:"index;default:1000;comment:'角色排序权重'" validate:"gte=0"`
	// PermissionIDs 角色绑定的权限标识数组
	PermissionIDs constants.StringArrayField `json:"permission_ids" gorm:"type:json;comment:'角色绑定的权限标识数组'" validate:"dive,word"`
}

// QueryRoleReq 查询角色列表请求参数
type QueryRoleReq struct {
	app.Pager
	app.Sorter
	// ID 角色ID
	ID string `form:"id" json:"id" gorm:"primaryKey;type:bigint;comment:'ID'" validate:"omitempty,int64" `
	// Name 角色名称
	Name string `form:"name" json:"name" gorm:"type:varchar(255);index;comment:'角色名称'"`
	// Type 角色类型
	Type string `json:"type" gorm:"index;default:Staff;comment:'角色类型'" validate:"omitempty,oneof=Admin DepartmentAdmin Staff"`
	// IsDefault 是否为默认角色
	IsDefault constants.Boolean `form:"is_default" json:"is_default" gorm:"default:2;comment:'是否为默认角色'"  validate:"omitempty,oneof=1 2" `
}

// QueryRoleStaffsReq 查询授权员工列表请求参数
type QueryRoleStaffsReq struct {
	app.Pager
	app.Sorter
	// ID 员工ID
	StaffID string `form:"staff_id" json:"id" validate:"omitempty,int64"`
	// ExtStaffID 员工外部ID
	ExtStaffID string `form:"ext_staff_id"  json:"ext_staff_id" validate:"omitempty"`
	// Name 员工名称
	Name string `form:"name" json:"name" gorm:"type:varchar(255);index;comment:'员工名称'"`
	// RoleID 角色ID
	RoleID string `form:"role_id" json:"role_id" gorm:"type:bigint;comment:角色ID" validate:"omitempty,int64"`
	// RoleType 角色类型
	RoleType string `form:"role_type" json:"type" gorm:"index;default:Staff;comment:'角色类型'" validate:"omitempty,oneof=Admin DepartmentAdmin Staff"`
}

// DeleteRoleReq 删除角色请求参数
type DeleteRoleReq struct {
	// 角色ID
	IDs []string `json:"ids" validate:"gt=0,dive,int64"`
}

// AssignToStaffsReq 授权角色给员工请求
type AssignToStaffsReq struct {
	// ExtStaffIDs 外部员工ID
	ExtStaffIDs []string `json:"ext_staff_ids" validate:"gt=0,dive,alphanum"`
	// RoleID 角色ID
	RoleID string `json:"role_id" validate:"required,int64"`
}
