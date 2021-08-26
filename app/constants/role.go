package constants

// RoleType 角色类型
type RoleType string

const (
	// RoleTypeSuperAdmin 超级管理员
	RoleTypeSuperAdmin RoleType = "superAdmin"
	// RoleTypeAdmin 企业普通管理员
	RoleTypeAdmin RoleType = "admin"
	// RoleTypeDepartmentAdmin 企业部门管理员
	RoleTypeDepartmentAdmin RoleType = "departmentAdmin"
	// RoleTypeStaff 企业普通员工
	RoleTypeStaff RoleType = "staff"
)
