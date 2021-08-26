package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type InternalTag struct {
	Base
	srv *services.InternalTag
}

// Create
// @tags 内部标签
// @Summary 创建内部标签
// @Produce  json
// @Accept json
// @Param params body requests.CreateInternalTagReq true "创建创建内部标签请求"
// @Success 200 {object} app.JSONResult{data=models.InternalTag} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/internal-tag [post]
func (o InternalTag) Create(c *gin.Context) {
	req := requests.CreateInternalTagReq{}
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

	tag, err := o.srv.Create(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Create failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(tag)
}

func NewInternalTag() *InternalTag {
	return &InternalTag{srv: services.NewInternalTag()}
}

// Delete
// @tags 内部标签
// @Summary 删除内部标签
// @Produce  json
// @Accept json
// @Param params body requests.DeleteInternalTagReq true "删除内部标签请求"
// @Success 200 {object} app.JSONResult{data=bool} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/internal-tag/action/delete [post]
func (o *InternalTag) Delete(c *gin.Context) {
	req := requests.DeleteInternalTagReq{}
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

	total, err := o.srv.Delete(req.IDs, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Delete failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(total)
}

// Query
// @tags 内部标签
// @Summary 内部标签列表
// @Produce json
// @Accept json
// @Param params body requests.QueryInternalTagReq true "内部标签列表"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.InternalTag}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/internal-tags [get]
func (o *InternalTag) Query(c *gin.Context) {
	req := requests.QueryInternalTagReq{}
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

	iterms, total, err := o.srv.Query(req, staffAdmin.ExtCorpID, &req.Sorter, &req.Pager)
	if err != nil {
		err = errors.Wrap(err, "Delete failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(iterms, total)
}
