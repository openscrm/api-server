package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type ContactWayGroup struct {
	Base
	srv *services.ContactWayGroup
}

func NewContactWayGroup() *ContactWayGroup {
	return &ContactWayGroup{srv: services.NewContactWayGroup()}
}

// Query
// @tags 渠道码分组
// @Summary 查询渠道码分组列表
// @Produce  json
// @Accept json
// @Param params body requests.QueryContactWayGroupReq true "查询渠道码分组列表请求"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.ContactWayGroup}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff_admin/contact_way_groups [get]
func (o *ContactWayGroup) Query(c *gin.Context) {
	req := requests.QueryContactWayGroupReq{}
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

	items, total, err := o.srv.Query(req, staffAdmin.ExtCorpID, &req.Sorter, &req.Pager)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(items, total)
}

// Get
// @tags 渠道码分组
// @Summary 获取渠道码分组详情
// @Produce  json
// @Accept json
// @Param id path string true "渠道码分组ID"
// @Success 200 {object} app.JSONResult{data=models.ContactWayGroup} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff_admin/contact_way_group/{id} [get]
func (o *ContactWayGroup) Get(c *gin.Context) {
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
// @tags 渠道码分组
// @Summary 创建渠道码分组
// @Produce  json
// @Accept json
// @Param params body requests.CreateContactWayGroupReq true "创建渠道码分组请求"
// @Success 200 {object} app.JSONResult{data=models.ContactWayGroup} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff_admin/contact_way_group [post]
func (o *ContactWayGroup) Create(c *gin.Context) {
	req := requests.CreateContactWayGroupReq{}
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
// @tags 渠道码分组
// @Summary 更新渠道码分组
// @Produce  json
// @Accept json
// @Param id path string true "渠道码分组ID"
// @Param params body requests.UpdateContactWayGroupReq true "更新渠道码分组请求"
// @Success 200 {object} app.JSONResult{data=models.ContactWayGroup} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff_admin/contact_way_group/{id} [put]
func (o *ContactWayGroup) Update(c *gin.Context) {
	req := requests.UpdateContactWayGroupReq{}
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
// @tags 渠道码分组
// @Summary 删除渠道码分组
// @Produce  json
// @Accept json
// @Param params body requests.DeleteContactWayGroupReq true "删除渠道码分组请求"
// @Success 200 {object} app.JSONResult{data=bool} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff_admin/contact_way_group/action/delete [post]
func (o *ContactWayGroup) Delete(c *gin.Context) {
	req := requests.DeleteContactWayGroupReq{}
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
