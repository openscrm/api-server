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

type QuickReplyGroupFrontend struct {
	Base
	srv                        *services.QuickReplyGroup
	QuickReplyGroupFrontendSrv *services.QuickReplyGroupFrontend
}

func NewQuickReplyGroupFrontend() *QuickReplyGroupFrontend {
	return &QuickReplyGroupFrontend{
		srv:                        services.NewQuickReplyGroup(),
		QuickReplyGroupFrontendSrv: services.NewQuickReplyGroupFrontend()}
}

// Query
// @tags  话术库组
// @Summary H5获取企业话术库组
// @Produce  json
// @Accept json
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.QuickReplyGroup}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-frontend/quick-reply-groups [get]
func (rg *QuickReplyGroupFrontend) Query(c *gin.Context) {
	req := requests.QueryQuickReplyGroupReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staff, err := rg.GetStaffInfo(handler)
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

// Search
// @tags  话术库组
// @Summary H5按关键字搜索企业话术库组
// @Produce  json
// @Accept json
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.QuickReplyGroup}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-frontend/quick-reply-group/action/search [post]
func (rg *QuickReplyGroupFrontend) Search(c *gin.Context) {
	req := requests.QueryQuickReplyReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		log.Sugar.Debugw("BadRequest", "err", err)
		handler.ResponseBadRequestError(err)
		return
	}

	staff, err := rg.GetStaffInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	item, total, err := rg.srv.QueryQuickReply(req, staff.ExtCorpID, &req.Pager)
	if err != nil {
		err = errors.Wrap(err, "QueryQuickReply failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(item, total)

}

// Create
// @tags 话术库组
// @Summary H5创建企业话术库组
// @Produce json
// @Accept json
// @Param params body entities.CreateQuickReplyGroupReq true "H5创建企业话术库组请求"
// @Success 200 {object} app.JSONResult{data=models.QuickReplyGroup} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-frontend/quick-reply-group [post]
func (rg *QuickReplyGroupFrontend) Create(c *gin.Context) {
	req := entities.CreateQuickReplyGroupReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staff, err := rg.GetStaffInfo(handler)
	if err != nil {
		log.TracedError("GetStaffInfo failed", err)
		return
	}

	item, err := rg.srv.Create(req, staff.ExtCorpID, staff.ExtID)
	if err != nil {
		err = errors.Wrap(err, "Create quick reply group failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// Update
// @tags 话术库组
// @Summary H5更新话术分组
// @Accept json
// @Param params body entities.CreateQuickReplyGroupReq true "H5更新话术分组请求"
// @Success 200 {object} app.JSONResult{data=models.QuickReplyGroup} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-frontend/quick-reply-group [put]
func (rg *QuickReplyGroupFrontend) Update(c *gin.Context) {
	req := entities.UpdateQuickReplyGroupReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staff, err := rg.GetStaffInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	group, err := rg.QuickReplyGroupFrontendSrv.Update(req, staff.ExtCorpID, staff.ExtID)
	if err != nil {
		err = errors.Wrap(err, "Create quick reply group failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(group)
}

// Delete
// @tags 话术库组
// @Summary H5删除话术库分组
// @Produce  json
// @Param params body entities.DeleteQuickReplyGroupReq true "H5删除话术库组请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-frontend/quick-reply-group/action/delete [post]
func (rg *QuickReplyGroupFrontend) Delete(c *gin.Context) {
	req := entities.DeleteQuickReplyGroupReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staff, err := rg.GetStaffInfo(handler)
	if err != nil {
		log.TracedError("GetStaffInfo failed", err)
		return
	}

	err = rg.QuickReplyGroupFrontendSrv.Delete(req.IDs, staff.ExtCorpID, staff.ExtID)
	if err != nil {
		err = errors.Wrap(err, "Delete failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}
