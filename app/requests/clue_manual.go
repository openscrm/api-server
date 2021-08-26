package requests

// CreateClueManualReq 创建跟进请求
type CreateClueManualReq struct {
	// 员工外部ID
	ExtStaffID string `json:"ext_staff_id" form:"ext_staff_id" validate:"required"`
	// 客户外部ID
	ExtCustomerID string `json:"ext_customer_id"   form:"ext_customer_id"  validate:"required"`
	// 跟进事件的内容
	Content string `json:"content"  form:"content"  validate:"required"`
}

// UpdateClueManualReq 更新跟进请求
type UpdateClueManualReq struct {
	// 跟进事件的内容
	Content string `json:"content"  form:"content" validate:"required"`
}
