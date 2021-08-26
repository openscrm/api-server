package requests

import (
	"openscrm/app/constants"
	"openscrm/common/app"
)

// QueryCustomerLossesReq  查询客户流失记录
type QueryCustomerLossesReq struct {
	// 企微员工id
	ExtStaffIDs []string `json:"ext_staff_ids" form:"ext_staff_ids" validate:"omitempty,dive,word"`
	// 流失起止时间
	LossStart constants.DateField `json:"loss_start" form:"loss_start" validate:"omitempty"`
	LossEnd   constants.DateField `json:"loss_end" form:"loss_end" validate:"omitempty"`
	// 添加好友起止时间
	ConnectionCreateStart constants.DateField `json:"connection_create_start" form:"connection_create_start" validate:"omitempty"`
	ConnectionCreateEnd   constants.DateField `json:"connection_create_end" form:"connection_create_end" validate:"omitempty"`
	// 好友关系时长, 单位-天
	TimeSpanLowerLimit int64 `json:"time_span_lower_limit" form:"time_span_lower_limit"  validate:"omitempty,gt=0,numeric"`
	TimeSpanUpperLimit int64 `json:"time_span_upper_limit" form:"time_span_upper_limit" validate:"omitempty,gtefield=TimeSpanLowerLimit,numeric"`
	app.Pager
	app.Sorter
}
