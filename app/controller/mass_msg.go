package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"openscrm/app/constants"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/log"
	"openscrm/common/storage"
	"openscrm/conf"
	"path"
)

type MassMsg struct {
	Base
	srv *services.MassMsgService
}

func NewDefaultMassMsg() *MassMsg {
	return &MassMsg{srv: services.NewDefaultMassMsgService()}
}

// Create
// @tags 客户群发
// @Summary 创建群发消息
// @Param pet body requests.SendMassMsgReq true "创建待发消息"
// @Produce json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/mass-msg [post]
func (ch MassMsg) Create(c *gin.Context) {
	req := requests.SendMassMsgReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		log.Sugar.Debugw("BadRequest", "err", err)
		handler.ResponseBadRequestError(err)
		return
	}

	staffAdminInfo, err := ch.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	msg, err := ch.srv.Create(req, staffAdminInfo.Name, staffAdminInfo.ExtCorpID)
	if err != nil {
		err := errors.Wrap(err, "send group msg failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(msg)
}

// Get
// @tags 客户群发
// @Summary 获取客户群发详情
// @Produce json
// @Success 200 {object} app.JSONResult{data=models.MassMsg} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/mass-msg/{id} [get]
func (ch MassMsg) Get(c *gin.Context) {
	handler := app.NewHandler(c)
	missionID, err := handler.GetIDParam()
	if err != nil {
		err = errors.Wrap(err, "handler.GetIDParam failed")
		handler.ResponseBadRequestError(err)
		return
	}

	info, err := ch.GetStaffAdminInfo(handler)
	if err != nil {
		return
	}

	result, err := ch.srv.Get(missionID, info.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Get failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(result)
}

// GetSendMassMsgResult
// @tags 客户群发
// @Summary 获取创建群发消息的结果
// @Produce json
// @Success 200 {object} app.JSONResult{data=requests.SendMassMsgResp} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/mass-msg/result/{id} [get]
func (ch MassMsg) GetSendMassMsgResult(c *gin.Context) {
	handler := app.NewHandler(c)
	missionID, err := handler.GetIDParam()
	if err != nil {
		err = errors.Wrap(err, "handler.GetIDParam failed")
		handler.ResponseBadRequestError(err)
		return
	}
	result, err := ch.srv.GetSendMassMsgResult(missionID)
	if err != nil {
		err = errors.Wrap(err, "get send group msg failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(result)
}

// Delete
// @tags 客户群发
// @Summary 删除定时群发的消息
// @Param id path int true "消息id"
// @Produce json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/mass-msg/action/delete/{id} [post]
func (ch MassMsg) Delete(c *gin.Context) {
	handler := app.NewHandler(c)

	req := requests.CommonDeleteReq{}
	_, err := handler.BindAndValidateReq(&req)
	if err != nil {
		err = errors.Wrap(err, "bind req failed")
		handler.ResponseBadRequestError(err)
		return
	}

	err = ch.srv.DeleteTimedMassMsg(req.IDs)
	if err != nil {
		err = errors.Wrap(err, "get send group msg failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

// Notify
// @tags 客户群发
// @Summary 提醒员工发送群发消息
// @Param id path int true "消息id"
// @Produce json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/mass-msg/notify [post]
func (ch MassMsg) Notify(c *gin.Context) {
	req := requests.MassMsgNotifyReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdminInfo, err := ch.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	err = ch.srv.Notify(req.IDs, staffAdminInfo.ExtCorpID)
	if err != nil {
		err := errors.Wrap(err, "send group msg failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

// Query
// @tags 客户管理
// @Summary 群发消息列表
// @Produce json
// @Param params body requests.QueryMassMsgReq true "群发消息列表请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/mass-msgs [get]
func (ch MassMsg) Query(c *gin.Context) {
	handler := app.NewHandler(c)
	req := requests.QueryMassMsgReq{}
	_, err := handler.BindAndValidateReq(&req)
	if err != nil {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	info, err := ch.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	items, total, err := ch.srv.Query(info.ExtCorpID, &req.Sorter, &req.Pager)
	if err != nil {
		err := errors.Wrap(err, "query group msg failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(items, total)
}

// CustomerFilter
// @tags 客户管理
// @Summary 群发消息-客户筛选
// @Produce json
// @Param params body  constants.ExtCustomerFilter true "群发消息-客户筛选请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/mass-msg/customer-filter [get]
func (ch MassMsg) CustomerFilter(c *gin.Context) {
	handler := app.NewHandler(c)
	req := requests.CountCustomer{}
	_, err := handler.BindAndValidateReq(&req)
	if err != nil {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	_, total, err := ch.srv.GetStaffsCustomers([]string{}, req.ExtCustomerFilterEnable, req.ExtCustomerFilter)
	if err != nil {
		err := errors.Wrap(err, "query group msg failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(total)
}

// GetUploadUrl
// @tags 客户管理
// @Summary 客户管理-消息群发-上传文件
// @Produce  json
// @Accept json
// @Param params body requests.GetUploadURLReq true "客户管理-消息群发-上传文件请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/mass-msg/action/get-upload-url [post]
func (ch MassMsg) GetUploadUrl(c *gin.Context) {
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

	info, err := ch.GetStaffAdminInfo(handler)
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

// Update
// @tags 客户管理
// @Summary 消息群发-修改
// @Produce json
// @Accept json
// @Param params body requests.UpdateMassMsgReq true "消息群发-修改请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/mass-msg/:id [get]
func (ch MassMsg) Update(c *gin.Context) {
	req := requests.UpdateMassMsgReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		log.Sugar.Debugw("BadRequest", "err", err)
		handler.ResponseBadRequestError(err)
		return
	}

	id, err := handler.GetIDParam()
	if err != nil {
		log.Sugar.Debugw("BadRequest", "err", err)
		handler.ResponseBadRequestError(err)
		return
	}

	staffAdminInfo, err := ch.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	msgID, err := ch.srv.UpdateMassMsg(req, id, staffAdminInfo.Name, staffAdminInfo.ExtCorpID)
	if err != nil {
		err := errors.Wrap(err, "send group msg failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(msgID)
}
