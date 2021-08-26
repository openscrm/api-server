package requests

import (
	"openscrm/app/constants"
	"openscrm/common/app"
)

type SendMassMsgResp struct {
	MissionID     string                      `json:"mission_id"`
	MissionStatus constants.SendMassMsgStatus `json:"mission_status"`
	// 已发员工计数
	DeliveredNum int `json:"delivered_num"`
	// 已送达客户计数
	SuccessNum int `json:"success_num"`
	// 未发送员工计数
	UnDeliveredNum int `json:"undelivered_num"`
	// 未送达客户计数
	FailedNum int `json:"failed_num"`
}

type SendMassMsgReq struct {
	// 需要全部员工发送时，不带
	// 客户群群发时, ExtStaffIDs为群主IDs
	ExtStaffIDs constants.StringArrayField `json:"ext_staff_ids" form:"ext_staff_ids" validate:"omitempty,gt=0"`
	// 1-立即发送，2-定时发送
	SendType constants.SendMassMsgType `json:"send_type" validate:"required,oneof=1 2"`
	// 定时发送时间戳
	SendAt constants.DateTimeFiled `json:"send_at" validate:"omitempty,gt=0"`
	// 群发任务的类型，默认为single，表示发送给客户，group表示发送给客户群
	ChatType constants.ChatType `json:"chat_type" validate:"omitempty,oneof=single group"`
	// 需要发送消息的员工部门集合
	ExtDepartmentIDs constants.Int64ArrayField `json:"ext_department_ids" validate:"omitempty,gte=0"`
	// 是否有筛选条件
	ExtCustomerFilterEnable constants.Boolean           `json:"ext_customer_filter_enable" form:"ext_customer_filter_enable"`
	ExtCustomerFilter       constants.ExtCustomerFilter `json:"ext_customer_filter" validate:"required"`
	// 消息体
	Msg constants.AutoReplyField `json:"msg" validate:"omitempty,gte=0"`
}

type UpdateMassMsgReq struct {
	// 需要全部员工发送时，不带
	ExtStaffIDs constants.StringArrayField `json:"ext_staff_ids" form:"ext_staff_ids" validate:"omitempty,gt=0"`
	// 1-立即发送，2-定时发送
	SendType constants.SendMassMsgType `json:"send_type" validate:"required,oneof=1 2"`
	// 定时发送时间戳
	SendAt constants.DateTimeFiled `json:"send_at" validate:"omitempty,gt=0"`
	// 群发任务的类型，默认为single，表示发送给客户，group表示发送给客户群
	//ChatType constants.ChatType `json:"chat_type" validate:"omitempty,oneof=single group"`
	// 需要发送消息的员工部门集合
	ExtDepartmentIDs constants.Int64ArrayField `json:"ext_department_ids" validate:"omitempty,gte=0"`
	// 是否有筛选条件
	ExtCustomerFilterEnable constants.Boolean           `json:"ext_customer_filter_enable" form:"ext_customer_filter_enable"`
	ExtCustomerFilter       constants.ExtCustomerFilter `json:"ext_customer_filter" validate:"required"`
	// 消息体
	Msg constants.AutoReplyField `json:"msg" validate:"omitempty,gte=0"`
}

// QueryMassMsgReq 查询群发消息列表请求参数
type QueryMassMsgReq struct {
	app.Pager
	app.Sorter
}

type MassMsgNotifyReq struct {
	//  群发消息ids
	IDs constants.StringArrayField `json:"ids" form:"ids" validate:"gt=0"`
}

type CustomerNum struct {
	Total int64 `json:"total"`
}
type CountCustomer struct {
	// 筛选条件
	constants.ExtCustomerFilter
	// 是否有筛选条件
	ExtCustomerFilterEnable constants.Boolean `json:"ext_customer_filter_enable" form:"ext_customer_filter_enable"`
}
