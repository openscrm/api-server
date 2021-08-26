package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type MsgArch struct {
	Base
	srv *services.MsgArch
}

func NewMsgArch() *MsgArch {
	return &MsgArch{srv: services.NewMsgArch()}
}

// QuerySessions
//@tags 会话存档
//@Summary 查询会话列表
//@Produce  json
//@Accept json
//@Param params body requests.QueryChatMsgReq true "查询会话列表请求"
//@Success 200 {object} app.JSONResult{} "成功"
//@Failure 400 {object} app.JSONResult{} "非法请求"
//@Failure 500 {object} app.JSONResult{} "内部错误"
//@Router /api/v1/staff-admin/customers/chat-sessions [get]
func (o MsgArch) QuerySessions(c *gin.Context) {
	req := requests.QuerySessionReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		handler.ResponseError(err)
		return
	}
	items, err := o.srv.QuerySessions(req, staffAdmin.ExtCorpID)
	if err != nil {
		log.TracedError("QuerySessions failed", err)
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(items)
}

// Sync
// @tags 会话存档
// @Summary 同步会话数据
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/chat-msg/sync [post]
func (o MsgArch) Sync(c *gin.Context) {
	handler := app.NewHandler(c)
	staffAdmin, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}
	err = o.srv.Sync(staffAdmin.ExtCorpID)
	if err != nil {
		log.TracedError("QuerySessions failed", err)
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(nil)
}

// QuerySessionMsgs
//@tags 会话存档
//@Summary 会话中的消息列表
//@Produce  json
//@Accept json
//@Param params body requests.QueryChatMsgReq true "查询会话中的消息列表请求"
//@Success 200 {object} app.JSONResult{} "成功"
//@Failure 400 {object} app.JSONResult{} "非法请求"
//@Failure 500 {object} app.JSONResult{} "内部错误"
//@Router /api/v1/staff-admin/customer/sessions-msgs [get]
func (o MsgArch) QuerySessionMsgs(c *gin.Context) {
	req := requests.QueryChatMsgReq{}
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
	res, err := o.srv.QueryMsgs(req, staffAdmin.ExtCorpID)
	if err != nil {
		log.TracedError("QuerySessions failed", err)
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(res)
}

// SearchMsgs
//@tags 会话存档
//@Summary 搜索消息
//@Produce  json
//@Accept json
//@Param params body requests.QueryChatMsgReq true "搜索消息请求"
//@Success 200 {object} app.JSONResult{} "成功"
//@Failure 400 {object} app.JSONResult{} "非法请求"
//@Failure 500 {object} app.JSONResult{} "内部错误"
//@Router /api/v1/staff-admin/customer/sessions-msg/action/search [post]
func (o MsgArch) SearchMsgs(c *gin.Context) {
	req := requests.SearchMsgReq{}
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
	res, err := o.srv.SearchMsgs(req, staffAdmin.ExtCorpID)
	if err != nil {
		log.TracedError("QuerySessions failed", err)
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(res)
}
