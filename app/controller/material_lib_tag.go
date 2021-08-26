package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type MaterialLibTag struct {
	Base
	srv *services.MaterialLibTag
}

func NewMaterialLibTag() *MaterialLibTag {
	return &MaterialLibTag{srv: services.NewMaterialLibTag()}
}

// Create
// @tags 素材库
// @Summary 创建素材库素材标签
// @Produce  json
// @Accept json
// @Param params body requests.CreateMaterialLibTagReq true "创建素材库素材请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/material/lib/tag [post]
func (o MaterialLibTag) Create(c *gin.Context) {
	req := requests.CreateMaterialLibTagReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	item, err := o.srv.Create(req.Names, staffAdmin)
	if err != nil {
		err = errors.Wrap(err, "Create failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)

}

// Delete
// @tags 素材库
// @Summary 删除素材库标签
// @Produce  json
// @Accept json
// @Param params body requests.CommonDeleteReq true "删除素材库标签"
// @Success 200 {object} app.JSONResult{data=bool} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/material/lib/tag/action/delete [post]
func (o MaterialLibTag) Delete(c *gin.Context) {
	req := requests.CommonDeleteReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	total, err := o.srv.Delete(req.IDs)
	if err != nil {
		err = errors.Wrap(err, "Delete failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(total)
}

// Query
// @tags 素材库
// @Summary 查询素材库标签列表
// @Produce  json
// @Accept json
// @Param params body requests.QueryMaterialLibTagReq true "查询查询素材库标签列表请求"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.MaterialLibTag}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/material/lib/tags [get]
func (o MaterialLibTag) Query(c *gin.Context) {
	req := requests.QueryMaterialLibTagReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	items, total, err := o.srv.Query(req.Name, staffAdmin, &req.Pager, &req.Sorter)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(items, total)
}
