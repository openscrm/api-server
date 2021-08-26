package requests

import "openscrm/common/app"

type CreateMaterialLibTagReq struct {
	Names []string `json:"names" form:"names" validate:"required"`
}

type QueryMaterialLibTagReq struct {
	Name       string `json:"name" form:"name"`
	app.Sorter `form:"app_sorter"`
	app.Pager  `form:"app_pager"`
}
