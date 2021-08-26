package requests

import (
	"openscrm/common/app"
)

type TagListReq struct {
	// 部门ID, 默认0，查询所有
	ExtDepartmentIDs []int64 `json:"ext_department_ids" form:"ext_department_ids" validate:"omitempty,dive,gte=0"`
	// 标签/标签组名
	Name string `json:"name" form:"name" validate:"omitempty"`
	app.Pager
	app.Sorter
}

type CreateTagReq struct {
	// 外部标签组id
	ExtTagGroupId string `json:"ext_tag_group_id" form:"ext_tag_group_id" validate:"required"`
	// 标签名
	Names []string `json:"names" form:"names" validate:"required,gt=0"`
}

type DeleteTagGroupsReq struct {
	// 外部组id列表
	ExtIDs []string `json:"ext_ids" form:"ext_ids" validate:"required,gt=0"`
}

type CreateTagGroupReq struct {
	Name           string  `json:"name" validate:"required"`
	DepartmentList []int64 `json:"department_list" validate:"omitempty"`
	Order          uint32  `json:"order"`
	// 标签列表
	Tags []Tag `json:"tags" form:"tags" validate:"omitempty,gt=0"`
}

type UpdateTagGroupReq struct {
	// 标签组ext_id
	ExtID string `json:"ext_id"`
	// 标签组名
	Name string `json:"name" validate:"required"`
	// 排序权重
	Order uint32 `json:"order"`
	// 删除的标签id
	RemoveExtTagIDs []string `json:"remove_ext_tag_ids" form:"remove_ext_tag_ids" validate:"omitempty,gt=0"`
	// 标签列表
	Tags []Tag `json:"tags" form:"tags" validate:"omitempty,gt=0"`
	// 标签可用部门列表, 缺省所有部门可用
	DepartmentList []int64 `json:"department_list" form:"department_list" validate:"omitempty,dive,gte=0"`
}

type Tag struct {
	// 标签名
	Name string `json:"name" form:"name"`
	// 更新标签时使用，新建标签不用带
	ExtId string `json:"ext_id" form:"ext_id"`
	// 排序权重
	Order uint32 `json:"order"`
}

type ExchangeOrderReq struct {
	ID              string `json:"id" form:"id" validate:"required,int64"`
	ExchangeOrderID string `json:"exchange_order_id" form:"exchange_order_id"  validate:"required,int64"`
}
