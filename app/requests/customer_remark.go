package requests

type CustomerRemarkReq struct {
	// 微信客户ID
	ExtCustomerID string `json:"ext_customer_id" form:"ext_customer_id"`
	// 微信员工ID
	ExtStaffID string `form:"ext_staff_id" json:"ext_staff_id"`
}

// AddCustomerRemarkReq 添加自定义客户信息
type AddCustomerRemarkReq struct {
	// 自定义信息类型 option_text text timestamp
	FieldType string `json:"field_type" json:"field_type" validate:"required,oneof=option_text text timestamp"`
	FieldName string `json:"field_name" validate:"required"`
	// 多选类型的选项
	OptionNameList []string `json:"option_name_list" validate:"required_if=FieldType option_text"`
}
type UpdateRemarkReq struct {
	ID   string `json:"id" form:"id" validate:"required"`
	Name string `json:"name" form:"name" validate:"required"`
}
type DeleteCustomerRemarkReq struct {
	IDs []string `json:"ids" validate:"gt=0"`
}

// --------------------------------

type AddRemarkOptionReq struct {
	RemarkID string `json:"remark_id" form:"remark_id" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
}

type UpdateRemarkOptionReq struct {
	//// 自定义信息ID
	//RemarkID string `json:"remark_id" form:"remark_id" validate:"required"`
	// 选项类型的选项ID
	RemarkOptionID string `json:"remark_option_id" form:"remark_option_id" validate:"required"`
	// 信息名或选项名
	Name string `json:"name" form:"name" validate:"required"`
}

type DeleteRemarkOptionReq struct {
	IDs []string `json:"ids" form:"ids" validate:"required,gt=0"`
}
