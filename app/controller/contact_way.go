package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type ContactWay struct {
	Base
	srv *services.ContactWay
}

func NewContactWay() *ContactWay {
	return &ContactWay{srv: services.NewContactWay()}
}

// Query
// @tags 渠道码
// @Summary 查询渠道码列表
// @Produce  json
// @Accept json
// @Param params body requests.QueryContactWayReq true "查询渠道码列表请求"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]responses.ContactWay}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff_admin/contact_ways [get]
func (o *ContactWay) Query(c *gin.Context) {
	req := requests.QueryContactWayReq{}
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

	items, total, err := o.srv.Query(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(items, total)
}

// Get
// @tags 渠道码
// @Summary 获取渠道码详情
// @Produce  json
// @Accept json
// @Param id path string true "渠道码ID"
// @Success 200 {object} app.JSONResult{data=models.ContactWay} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff_admin/contact_way/{id} [get]
func (o *ContactWay) Get(c *gin.Context) {
	handler := app.NewHandler(c)
	id, err := handler.GetIDParam()
	if err != nil {
		err = errors.Wrap(err, "handler.GetIDParam failed")
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	item, err := o.srv.Get(id, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Get failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// Create
// @tags 渠道码
// @Summary 创建渠道码
// @Produce  json
// @Accept json
// @Param params body requests.CreateContactWayReq true "创建渠道码请求"
// @Success 200 {object} app.JSONResult{data=models.ContactWay} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff_admin/contact_way [post]
func (o *ContactWay) Create(c *gin.Context) {
	req := requests.CreateContactWayReq{}
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

	item, err := o.srv.Create(req, staffAdmin.ExtCorpID, staffAdmin.ExtID)
	if err != nil {
		err = errors.Wrap(err, "Create failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// Update
// @tags 渠道码
// @Summary 更新渠道码（必须全量更新）
// @Produce  json
// @Accept json
// @Param id path string true "渠道码ID"
// @Param params body requests.UpdateContactWayReq true "更新渠道码请求"
// @Success 200 {object} app.JSONResult{data=models.ContactWay} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff_admin/contact_way/{id} [put]
func (o *ContactWay) Update(c *gin.Context) {
	req := requests.UpdateContactWayReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	id, err := handler.GetIDParam()
	if err != nil {
		err = errors.Wrap(err, "handler.GetIDParam failed")
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	item, err := o.srv.Update(id, req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Update failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// Delete
// @tags 渠道码
// @Summary 删除渠道码
// @Produce  json
// @Accept json
// @Param params body requests.DeleteContactWayReq true "删除渠道码请求"
// @Success 200 {object} app.JSONResult{data=bool} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff_admin/contact_way/action/delete [post]
func (o *ContactWay) Delete(c *gin.Context) {
	req := requests.DeleteContactWayReq{}
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

// BatchUpdate
// @tags 渠道码
// @Summary 批量更新渠道码
// @Produce  json
// @Accept json
// @Param params body requests.BatchUpdateContactWayReq true "批量更新渠道码请求"
// @Success 200 {object} app.JSONResult{data=integer} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff_admin/contact_way/action/batch-update [post]
func (o *ContactWay) BatchUpdate(c *gin.Context) {
	req := requests.BatchUpdateContactWayReq{}
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

	total, err := o.srv.BatchUpdate(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "BatchUpdate failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(total)
}
