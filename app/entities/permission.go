package entities

import (
	"openscrm/common/app"
)

// QueryPermissionReq 查询权限列表请求参数
type QueryPermissionReq struct {
	app.Pager
	app.Sorter
	// ID 权限ID
	ID string `form:"id" json:"id" gorm:"primaryKey;type:bigint;comment:'ID'" validate:"omitempty,int64" `
	// Name 权限名称
	Name string `json:"name" gorm:"index;comment:'权限名称'"`
	// BizIdentity 业务标识
	BizIdentity string `json:"biz_identity" gorm:"index;comment:'业务标识'" validate:"omitempty,alphanum"`
	// Operation 操作
	Operation string `json:"operation" gorm:"comment:'操作'" validate:"omitempty,oneof=read full"`
	// Identity 权限标识
	Identity string `json:"identity" gorm:"unique;size:50;comment:'权限标识'" validate:"omitempty,alphanum,max=50"`
}
