package entities

import "openscrm/common/app"

type GetSubDepartmentReq struct {
	// 父级部门ID，用于获取该部门下的子集，非零时忽略参数ext_department_ids，非递归
	ExtParentId int64   `form:"ext_parent_id" json:"ext_parent_id" validate:"gte=0"`
	ExtIDs      []int64 `json:"ext_ids" form:"ext_ids" validate:"omitempty,gt=0"`
	app.Pager
	app.Sorter
}
