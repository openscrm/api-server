package requests

import (
	"openscrm/app/constants"
)

// QueryCustomerStatisticReq  客户统计
type QueryCustomerStatisticReq struct {
	// 数据类型 total increase decrease net_increase
	StatisticType string `json:"statistic_type" form:"statistic_type" validate:"oneof=total increase decrease net_increase"`
	// 员工外部ID
	ExtStaffIDs []string `json:"ext_staff_ids" form:"ext_staff_ids" validate:"omitempty"`
	// 开始时间
	StartTime constants.DateField `form:"start_time" json:"start_time" validate:"required"`
	// 结束时间
	EndTime constants.DateField `form:"end_time" json:"end_time"  validate:"required"`
}
