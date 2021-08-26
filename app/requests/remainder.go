package requests

import "openscrm/app/constants"

type CreateRemainderReq struct {
	// 发送提醒的时间
	SendAt constants.DateTimeFiled `json:"send_at" form:"send_at" validate:"required"` // todo later than now
	// 客户名
	CustomerName string `json:"customer_name" form:"customer_name"  validate:"required"`
	// 提醒内容
	Content string `json:"content" form:"content" validate:"required"`
	// 发送人
	ExtStaffID string `json:"ext_staff_id" form:"ext_staff_id" validate:"required"`
	// 客户外部id
	ExtCustomerID string `json:"ext_customer_id" form:"ext_customer_id" validate:"required"`
}

type UpdateRemainderReq struct {
	//SendAt  LocalTime `json:"send_at" validate:"required"` // todo later than now
	Content string `json:"content"  form:"content"  validate:"required"`
}
