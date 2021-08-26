package requests

import (
	"openscrm/app/constants"
	"openscrm/common/app"
)

type QueryMainStaffInfoReq struct {
	ExtStaffID      string `form:"ext_staff_id" json:"ext_staff_id" validate:"omitempty"`
	ExtDepartmentID string `json:"ext_department_id" form:"ext_department_id" validate:"omitempty"`
	app.Pager
}

type StaffCustomerCount struct {
	DecreaseUserCount int `json:"decrease_user_count"`
	IncreaseUserCount int `json:"increase_user_count"`
	TotalUserCount    int `json:"total_user_count"`
}

type QueryStaffReq struct {
	// 企业微信部门id, 0-所有部门, 非0-指定部门
	ExtDepartmentIDs constants.Int64ArrayField `json:"ext_department_ids" form:"ext_department_ids" validate:"omitempty,gte=0"`
	// 员工名字
	Name string `json:"name" form:"name" validate:"omitempty,gt=0"`
	// RoleID 角色ID
	RoleID string `form:"role_id" json:"role_id"  validate:"omitempty,int64"`
	// RoleType 角色类型 admin departmentAdmin staff superAdmin
	RoleType string `form:"role_type" json:"role_type" validate:"omitempty,oneof=admin departmentAdmin staff superAdmin"`
	// 开启会话存档 1-是 2-否
	EnableMsgArch constants.Boolean `json:"enable_msg_arch" form:"enable_msg_arch"   validate:"omitempty"`
	app.Pager
	app.Sorter
}

type UpdateCustomerInternalTagsReq struct {
	ExtStaffID    string   `json:"ext_staff_id" form:"ext_staff_id" validate:"required"`
	ExtCustomerID string   `json:"ext_customer_id" form:"ext_customer_id" validate:"required"`
	AddTags       []string `json:"add_tags" form:"add_tags" validate:"omitempty,gt=0"`
	RemoveTags    []string `json:"remove_tags" form:"remove_tags"  validate:"omitempty,gt=0"`
}

// UpdateCustomerTagsReq 更新标签和批量打标签
type UpdateCustomerTagsReq struct {
	ExtStaffID      string   `json:"ext_staff_id" form:"ext_staff_id" validate:"omitempty"`
	ExtCustomerIDs  []string `json:"ext_customer_ids" form:"ext_customer_ids" validate:"required"`
	AddExtTagIDs    []string `json:"add_ext_tag_ids" form:"add_ext_tag_ids" validate:"omitempty"`
	RemoveExtTagIDs []string `json:"remove_ext_tag_ids" form:"remove_ext_tag_ids"  validate:"omitempty"`
}

// EnableStaffs 批量启用、禁用员工
type EnableStaffs struct {
	// 启用员工外部id
	ExtStaffIDs []string `json:"ext_staff_ids" form:"ext_staff_ids"`
	// 禁用员工外部id
	ExcludeExtStaffIDs []string `json:"exclude_ext_staff_ids" form:"exclude_ext_staff_ids"`
}
