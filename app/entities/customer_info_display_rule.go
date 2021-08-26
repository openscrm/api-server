package entities

type GetCustomerInfoDisplayReq struct {
	CorpID string `json:"ext_corp_id" form:ext_corp_id" validate:"required"`
}
type UpdateCustomerInfoDisplayRulesReq struct {
	CorpID string `json:"ext_corp_id" form:ext_corp_id" validate:"required"`
	//todo validate the 2 slices
	DisplayFieldList       []string `json:"display_field_list" validate:"omitempty,gt=0"`
	CancelDisplayFieldList []string `json:"cancel_display_field_list" validate:"omitempty,gt=0"`
}
