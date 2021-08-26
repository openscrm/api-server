package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/http"
	"openscrm/app/constants"
	"openscrm/app/entities"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/log"
	"openscrm/common/storage"
	"openscrm/conf"
	"path"
)

type QuickReply struct {
	Base
	srv  *services.QuickReply
	uuid uuid.UUID
}

func NewQuickReply() *QuickReply {
	return &QuickReply{srv: services.NewQuickReply(), uuid: uuid.New()}
}

// Create
// @tags 话术库
// @Summary 创建企业话术
// @Produce  json
// @Accept json
// @Param params body entities.CreateQuickReplyReq true "创建企业话术"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/quick-reply [post]
func (r *QuickReply) Create(c *gin.Context) {
	req := entities.CreateQuickReplyReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		log.Sugar.Debugw("BadRequest", "err", err)
		handler.ResponseBadRequestError(err)
		return
	}

	staffAdmin, err := r.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	item, err := r.srv.Create(req, staffAdmin)
	if err != nil {
		err = errors.Wrap(err, "Create quick-reply failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// Query
// @tags  话术库
// @Summary  查询企业话术库
// @Produce  json
// @Accept json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-frontend/quick-replies [get]
func (r *QuickReply) Query(c *gin.Context) {
	req := requests.QueryQuickReplyReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		log.Sugar.Debugw("BadRequest", "err", err)
		handler.ResponseBadRequestError(err)
		return
	}

	staffAdmin, err := r.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	item, total, err := r.srv.QueryQuickReply(req, staffAdmin.ExtCorpID, &req.Sorter, &req.Pager)
	if err != nil {
		err = errors.Wrap(err, "QueryQuickReply failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(item, total)
}

// Delete
// @tags 话术库
// @Summary 删除企业话术库
// @Produce  json
// @Accept json
// @Param params body entities.DeleteQuickReplyReq true "删除话术库请求"
// @Success 200 {object} app.JSONResult{data=bool} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/quick-reply/action/delete [post]
func (r *QuickReply) Delete(c *gin.Context) {
	req := entities.DeleteQuickReplyReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := r.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	total, err := r.srv.Delete(req.IDs, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Delete failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(total)

}

// Update
// @tags 话术库
// @Summary 更新企业话术库
// @Produce  json
// @Accept json
// @Param params body entities.UpdateQuickReplyReq true "更新企业话术库"
// @Success 200 {object} app.JSONResult{data=bool} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/quick-reply [put]
func (r *QuickReply) Update(c *gin.Context) {
	req := entities.UpdateQuickReplyReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := r.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	item, err := r.srv.Update(req, staffAdmin)
	if err != nil {
		err = errors.Wrap(err, "update quick-reply failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// GetUploadURL
// @tags 话术库
// @Summary 获取上传文件地址
// @Produce  json
// @Accept json
// @Param params body requests.GetUploadURLReq true "获取上传文件地址请求"
// @Success 200 {object} app.JSONResult{data=requests.GetUploadURLResp} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/quick-reply/action/get-upload-url [get]
//上传图片得到图片URL，该URL永久有效
//返回的图片URL，仅能用于图文消息正文中的图片展示，或者给客户发送欢迎语等
func (r *QuickReply) GetUploadURL(c *gin.Context) {
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

	expiredInSec := int64(3600)
	fileName := path.Join(
		info.ExtCorpID, "/",
		constants.QuickReplyModuleName, "/",
		info.Name, "/",
		req.Filename)

	signedURl, err := storage.FileStorage.SignURL(fileName, http.MethodPut, expiredInSec)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	downloadURl, err := storage.FileStorage.SignURL(fileName, http.MethodGet, expiredInSec)
	if err != nil {
		err = errors.Wrap(err, "Bucket.SignURL failed")
		return
	}

	handler.ResponseItem(requests.GetUploadURLResp{UploadURL: signedURl, DownloadURL: downloadURl})
}
