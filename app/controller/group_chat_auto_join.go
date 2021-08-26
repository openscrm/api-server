package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type GroupChatAutoJoin struct {
	Base
	srv *services.GroupChatAutoJoin
}

func NewGroupChatAutoJoin() *GroupChatAutoJoin {
	return &GroupChatAutoJoin{srv: services.NewGroupChatAutoJoin()}
}

// Query
// @tags 自动拉群码
// @Summary 查询自动拉群码列表
// @Produce  json
// @Accept json
// @Param params body requests.QueryGroupChatAutoJoinReq true "查询自动拉群码列表请求"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.GroupChatAutoJoinCode}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/auto-join-code [get]
func (o *GroupChatAutoJoin) Query(c *gin.Context) {
	req := requests.QueryGroupChatAutoJoinReq{}
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

// Create
// @tags 自动拉群码
// @summary 创建自动拉群码
// @produce  json
// @accept json
// @param params body requests.CreateGroupChatAutoJoinCodeReq true "创建自动拉群码请求"
// @success 200 {object} app.JSONResult{data=models.GroupChatAutoJoinCode} "成功"
// @failure 400 {object} app.JSONResult{} "非法请求"
// @failure 500 {object} app.JSONResult{} "内部错误"
// @router /api/v1/staff-admin/customer-group/auto-join-code [post]
func (o *GroupChatAutoJoin) Create(c *gin.Context) {
	req := requests.CreateGroupChatAutoJoinCodeReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("get staff admin info failed", err)
		return
	}

	item, err := o.srv.Create(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "create failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// Update
// @tags 自动拉群码
// @summary 更新自动拉群码（必须全量更新）
// @produce  json
// @accept json
// @param id path string true "自动拉群码id"
// @param params body requests.UpdateGroupChatAutoJoinQrCodeReq true "更新自动拉群码请求"
// @success 200 {object} app.JSONResult{data=models.GroupChatAutoJoinCode} "成功"
// @failure 400 {object} app.JSONResult{} "非法请求"
// @failure 500 {object} app.JSONResult{} "内部错误"
// @router /api/v1/staff-admin/customer-group/auto-join-code/{id} [put]
func (o *GroupChatAutoJoin) Update(c *gin.Context) {
	req := requests.UpdateGroupChatAutoJoinQrCodeReq{}
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
		err = errors.Wrap(err, "update failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// Delete
// @tags 自动拉群码
// @summary 删除自动拉群码
// @produce  json
// @accept json
// @param params body requests.DeleteGroupChatAutoJoinReq true "删除自动拉群码请求"
// @success 200 {object} app.JSONResult{} "成功"
// @failure 400 {object} app.JSONResult{} "非法请求"
// @failure 500 {object} app.JSONResult{} "内部错误"
// @router /api/v1/staff-admin/customer-group/auto-join-code/action/delete [post]
func (o *GroupChatAutoJoin) Delete(c *gin.Context) {
	req := requests.DeleteGroupChatAutoJoinReq{}
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
		err = errors.Wrap(err, "delete failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(total)
}

// BatchRegroup
// @tags 自动拉群码
// @summary 批量分组
// @produce  json
// @accept json
// @param params body requests.BatchRegroupReq true "批量分组"
// @success 200 {object} app.JSONResult{} "成功"
// @failure 400 {object} app.JSONResult{} "非法请求"
// @failure 500 {object} app.JSONResult{} "内部错误"
// @router /api/v1/staff-admin/customer-group/auto-join-code/action/batch-regroup [post]
func (o *GroupChatAutoJoin) BatchRegroup(c *gin.Context) {
	req := requests.BatchRegroupReq{}
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

	err = o.srv.BatchRegroup(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "delete failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}
