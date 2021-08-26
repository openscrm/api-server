package entities

import (
	"openscrm/app/constants"
)

type UpdateCustomerInfoReq struct {
	ExtStaffID    string                        `json:"ext_staff_id" validate:"required"`
	ExtCustomerID string                        `json:"ext_customer_id" validate:"required"`
	Age           int                           `form:"age" json:"age" validate:"omitempty,gt=0,lt=120"`
	Description   string                        `form:"description" json:"description" validate:"omitempty,gt=0"`
	Email         string                        `form:"email" json:"email" validate:"omitempty,email"`
	PhoneNumber   string                        `form:"phone_number" json:"phone_number" validate:"omitempty"`
	RemarkField   constants.CustomerRemarkField `json:"remark_field" form:"remark_field" validate:"omitempty,dive"`
}

type GetCustomerInfoReq struct {
	// 微信客户ID
	ExtCustomerID string `json:"ext_customer_id" form:"ext_customer_id" validate:"required"`
	// 微信员工ID
	ExtStaffID string `form:"ext_staff_id" json:"ext_staff_id" validate:"required"`
}
