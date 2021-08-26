package requests

import "openscrm/common/app"

type CreateInternalTagReq struct {
	Names []string `json:"names" form:"names" validate:"required"`
}

type DeleteInternalTagReq struct {
	IDs []string `json:"ids" form:"ids" validate:"required,gt=0"`
}

type QueryInternalTagReq struct {
	ExtStaffID string `json:"ext_staff_id" form:"ext_staff_id" validate:"required"`
	app.Sorter
	app.Pager
}
