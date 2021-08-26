package requests

import (
	"openscrm/app/constants"
	"openscrm/common/app"
)

// CreateContactWayGroupReq 创建渠道码分组请求参数
type CreateContactWayGroupReq struct {
	// Name 分组名称
	Name string `json:"name" gorm:"comment:'分组名称'" validate:"required"`
	// SortWeight 分组排序
	SortWeight int `json:"sort_weight" gorm:"comment:'分组排序'" validate:"gte=0"`
}

// UpdateContactWayGroupReq 更新渠道码分组请求参数
type UpdateContactWayGroupReq struct {
	// Name 分组名称
	Name string `json:"name" gorm:"comment:'分组名称'"`
	// SortWeight 分组排序
	SortWeight int `json:"sort_weight" gorm:"comment:'分组排序'" validate:"gte=0"`
}

// QueryContactWayGroupReq 查询渠道码分组列表请求参数
type QueryContactWayGroupReq struct {
	app.Pager
	app.Sorter
	// ID 渠道码分组ID
	ID string `form:"id" json:"id" gorm:"primaryKey;type:bigint;comment:'ID'" validate:"omitempty,int64" `
	// Name 渠道码分组名称
	Name string `form:"name" json:"name" gorm:"type:varchar(255);index;comment:'渠道码分组名称'"`
	// IsDefault 是否为默认分组
	IsDefault constants.Boolean `form:"is_default" json:"is_default" gorm:"default:2;comment:'是否为默认分组'"  validate:"omitempty,oneof=1 2" `
}

// DeleteContactWayGroupReq 删除渠道码分组请求参数
type DeleteContactWayGroupReq struct {
	// 渠道码分组ID
	IDs []string `json:"ids" validate:"gt=0,dive,int64"`
}
