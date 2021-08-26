package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/entities"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type QuickReplyGroup struct {
	Base
	srv *services.QuickReplyGroup
}

func NewQuickReplyGroup() *QuickReplyGroup {
	return &QuickReplyGroup{srv: services.NewQuickReplyGroup()}
}

// Create
// @tags  话术库组
// @Summary 创建企业话术库组
// @Produce  json
// @Accept json
// @Param params body entities.CreateQuickReplyGroupReq true "创建企业话术库组"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/quick-reply-group [post]
func (rg *QuickReplyGroup) Create(c *gin.Context) {
	req := entities.CreateQuickReplyGroupReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := rg.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	item, err := rg.srv.Create(req, staffAdmin.ExtCorpID, staffAdmin.ExtID)
	if err != nil {
		err = errors.Wrap(err, "Create quick reply group failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// Get
// @tags  话术库组
// @Summary 获取企业话术库组
// @Produce  json
// @Accept json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/quick-reply-groups [get]
func (rg *QuickReplyGroup) Get(c *gin.Context) {
	req := requests.QueryQuickReplyGroupReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staff, err := rg.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	items, total, err := rg.srv.Query(staff.ExtCorpID, &req.Sorter, &req.Pager)
	if err != nil {
		err = errors.Wrap(err, "get quick-reply group failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(items, total)
}

// Delete
// @tags 话术库组
// @Summary 删除话术库分组
// @Produce  json
// @Param params body entities.DeleteQuickReplyGroupReq true "删除话术库组请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/quick-reply-group/action/delete [post]
func (rg *QuickReplyGroup) Delete(c *gin.Context) {
	req := entities.DeleteQuickReplyGroupReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := rg.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	err = rg.srv.Delete(req.IDs, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Delete failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

// Update
// @tags 话术库组
// @Summary 更新话术分组
// @Accept json
// @Param params body entities.CreateQuickReplyGroupReq true "删除话术库组请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/quick-reply-group/{id} [put]
func (rg *QuickReplyGroup) Update(c *gin.Context) {
	req := entities.UpdateQuickReplyGroupReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := rg.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	group, err := rg.srv.Update(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Create quick reply group failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(group)
}
