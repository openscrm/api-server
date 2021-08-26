package requests

type QueryDepartmentReq struct {
	ExtID int64 `json:"ext_id" form:"ext_id" validate:"omitempty"`
}
