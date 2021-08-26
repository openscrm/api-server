package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type MaterialLibTagFrontend struct {
	Base
	srv *services.MaterialLibTagFrontend
}

func NewMaterialLibTagFrontend() *MaterialLibTagFrontend {
	return &MaterialLibTagFrontend{srv: services.NewMaterialLibTagFrontend()}
}

// Query
// @tags 素材库
// @Summary H5查询素材库标签列表
// @Produce  json
// @Accept json
// @Param params body requests.QueryMaterialLibTagReq true "H5查询素材库标签列表请求"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.MaterialLibTag}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-frontend/material/lib/tags [get]
func (o MaterialLibTagFrontend) Query(c *gin.Context) {
	req := requests.QueryMaterialLibTagReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staff, err := o.GetStaffInfo(handler)
	if err != nil {
		log.TracedError("GetStaffInfo failed", err)
		return
	}

	items, total, err := o.srv.Query(req.Name, staff, &req.Pager, &req.Sorter)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(items, total)
}
