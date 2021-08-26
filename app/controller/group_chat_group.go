package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/entities"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type GroupChatGroup struct {
	Base
	srv *services.GroupChatGroupService
}

// Query
// @tags 自动拉群分组
// @Summary 自动拉群分组列表
// @Produce  json
// @Accept json
// @Param params body entities.QueryGroupChatGroupReq true "自动拉群分组列表请求"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.GroupChatGroup}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-group/auto-join/group [get]
func (o *GroupChatGroup) Query(c *gin.Context) {
	req := entities.QueryGroupChatGroupReq{}
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
// @tags 自动拉群分组
// @Summary 创建自动拉群分组
// @Produce  json
// @Accept json
// @Param params body entities.CreateGroupChatGroupReq true "创建自动拉群分组请求"
// @Success 200 {object} app.JSONResult{data=models.GroupChatGroup} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-group/auto-join/group [post]
func (o *GroupChatGroup) Create(c *gin.Context) {
	req := entities.CreateGroupChatGroupReq{}
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
// @tags 自动拉群分组
// @Summary 自动拉群分组
// @Produce  json
// @Accept json
// @Param id path string true "自动拉群分组ID"
// @Param params body entities.UpdateGroupChatGroupReq true "自动拉群分组更新请求"
// @Success 200 {object} app.JSONResult{data=models.GroupChatGroup} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-group/auto-join/group/{id} [put]
func (o *GroupChatGroup) Update(c *gin.Context) {
	req := entities.UpdateGroupChatGroupReq{}
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
// @tags 自动拉群分组
// @Summary 删除自动拉群分组
// @Produce  json
// @Accept json
// @Param params body entities.DeleteGroupChatGroupReq true "自动拉群分组请求"
// @Success 200 {object} app.JSONResult{data=bool} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-group/auto-join/action/delete [post]
func (o *GroupChatGroup) Delete(c *gin.Context) {
	req := entities.DeleteGroupChatGroupReq{}
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

	total, err := o.srv.Delete(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Delete failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(total)
}

func NewGroupChatGroup() *GroupChatGroup {
	return &GroupChatGroup{srv: services.NewGroupChatGroupService()}
}
