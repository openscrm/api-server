package entities

// CreateQuickReplyGroupReq
// 新建话术分组
type CreateQuickReplyGroupReq struct {
	// 组标题
	Name string `json:"name" validate:"required"`
	// 可用部门，为空表示所有部门可用
	Departments []int64 `json:"departments" form:"departments" validate:"omitempty"`
	// 子分组标题
	SubGroups []SubGroup `json:"sub_groups" form:"sub_groups" validate:"required,dive"`
}

// SubGroup 子分组
type SubGroup struct {
	// 更新子分组时才用
	ID string `json:"id"`
	// 	子分组标题
	Name string `json:"name" form:"name" validate:"required"`
	// 是否删除 0-未删除 1-已删除
	//Deleted int64 `json:"deleted" form:"deleted" validate:"oneof=0 1"`
}

// UpdateQuickReplyGroupReq 更新话术分组
// 非顶级分组只能更新组名
type UpdateQuickReplyGroupReq struct {
	// id
	ID string `json:"id" form:"id" validate:"required,int64"`
	// 组标题
	Name string `json:"name" form:"name" validate:"required"`
	// 可用部门，为空表示所有部门可用
	Departments []int64 `json:"departments" form:"departments" validate:"omitempty"`
	// 删除的子分组
	DeleteGroupIDs []string `json:"delete_group_ids" form:"delete_group_ids"`
	// 子分组，更新子分组时为空
	SubGroups []SubGroup `json:"sub_groups" form:"sub_groups" validate:"omitempty,dive"`
}

// DeleteQuickReplyGroupReq
// 删除话术组
type DeleteQuickReplyGroupReq struct {
	// 话术组ID
	IDs []string `json:"ids" validate:"gt=0,dive,int64"`
}
