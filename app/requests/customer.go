package requests

import (
	"openscrm/app/constants"
	"openscrm/common/app"
)

// QueryCustomerReq 查询客户条件
type QueryCustomerReq struct {
	Name          string              `form:"name" json:"name"`                   // 客户名字，支持模糊查询
	ExtStaffIDs   []string            `json:"ext_staff_ids" form:"ext_staff_ids"` // 所属客服
	ExtTagIDs     []string            `form:"ext_tag_ids" json:"ext_tag_ids"`     // 企业标签
	TagUnionType  string              `form:"tag_union_type" json:"tag_union_type" `
	ChannelType   int                 `form:"channel_type" json:"channel_type"`                      // 添加渠道
	Gender        int                 `form:"gender" json:"gender" validate:"omitempty,oneof=0 1 2"` // 性别
	OutFlowStatus int                 `form:"out_flow_status" json:"out_flow_status "`               // 流失状态 1-已经流失 2-未流失
	Type          int                 `form:"type" json:"type" validate:"omitempty,oneof=0 1 2"`     // 客户类型 1-微信用户, 2-企业微信用户
	StartTime     constants.DateField `form:"start_time" json:"start_time"`                          // 添加客户的时间
	EndTime       constants.DateField `form:"end_time" json:"end_time"`                              // 添加客户的时间
	app.Pager
	app.Sorter
}
