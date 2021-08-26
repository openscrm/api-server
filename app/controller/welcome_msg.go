package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/constants"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/log"
	"openscrm/conf"
)

type WelComeMsg struct {
	Base
	srv *services.WelcomeMsgService
}

func NewWelComeMsg() *WelComeMsg {
	return &WelComeMsg{srv: services.NewWelcomeMsgService()}
}

// Create 创建欢迎语
// @tags 欢迎语
// @Summary 创建欢迎语
// @Produce  json
// @Accept json
// @Param params body requests.CreateWelcomeMsgReq true "创建欢迎语请求"
// @Success 200 {object} app.JSONResult{data=models.WelcomeMsg} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/customer/welcome-msg [post]
func (o WelComeMsg) Create(c *gin.Context) {
	req := requests.CreateWelcomeMsgReq{}
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
		err = errors.Wrap(err, "Create welcome msg failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// Query 查询欢迎语
// @tags  欢迎语
// @Summary 查询欢迎语
// @Produce  json
// @Accept json
// @Param params body requests.QueryWelcomeMsgReq true "查询欢迎语列表请求"
// @Success 200 {object} app.JSONResult{data=models.WelcomeMsg} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/welcome-msg [get]
func (o WelComeMsg) Query(c *gin.Context) {
	handler := app.NewHandler(c)
	req := requests.QueryWelcomeMsgReq{}
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

// Update 更新欢迎语
// @tags 欢迎语
// @Summary 更新欢迎语
// @Produce  json
// @Accept json
// @Param params body requests.UpdateWelcomeMsgReq true "更新欢迎语请求"
// @Success 200 {object} app.JSONResult{data=models.WelcomeMsg} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/welcome-msg/{id} [put]
func (o WelComeMsg) Update(c *gin.Context) {
	req := requests.UpdateWelcomeMsgReq{}
	handler := app.NewHandler(c)
	ID, err := handler.GetIDParam()
	if err != nil {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}
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

	item, err := o.srv.Update(ID, req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Create welcome msg failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// Delete 删除欢迎语
// @tags 欢迎语
// @Summary 删除欢迎语
// @Produce  json
// @Accept json
// @Param params body requests.CommonDeleteReq true "删除欢迎语请求"
// @Success 200 {object} app.JSONResult{data=bool} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/welcome-msg/action/delete [post]
func (o WelComeMsg) Delete(c *gin.Context) {
	req := requests.CommonDeleteReq{}
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

	err = o.srv.Delete(req.IDs, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Delete failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

// UploadFileUrl
// @tags 欢迎语
// @Summary 上传文件地址
// @Produce  json
// @Accept json
// @Param params body requests.GetUploadURLReq true "获取上传文件地址请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/welcome-msg/action/get-upload-url [get]
func (r *WelComeMsg) UploadFileUrl(c *gin.Context) {
	req := requests.GetUploadURLReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	if conf.Settings.App.Env != constants.DEV && conf.Settings.App.Env != constants.TEST {
		handler.ResponseError(errors.WithStack(ecode.ForbiddenError))
		return
	}

	info, err := r.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	url, err := r.srv.UploadImg(handler.Ctx.Request.Body, req.Filename, info.ExtCorpID)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(url)
}

// Get
// @tags 欢迎语
// @Summary 欢迎语详情
// @Produce  json
// @Accept json
// @Param id path string true "欢迎语ID"
// @Success 200 {object} app.JSONResult{data=models.WelcomeMsgWithDeptAndStaff} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/welcome-msg/{id} [get]
func (o *WelComeMsg) Get(c *gin.Context) {
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
