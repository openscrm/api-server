package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type GroupChatMassMsg struct {
	Base
	srv *services.GroupChatMassMsg
}

func NewDefaultGroupChatMassMsg() *GroupChatMassMsg {
	return &GroupChatMassMsg{srv: services.NewGroupChatMassMsg()}
}

// Create
// @tags 客户群群发
// @Summary 创建客户群群发消息
// @Param pet body requests.SendGroupChatMassMsgReq true "创建客户群群发消息请求"
// @Produce json
// @Success 200 {object} app.JSONResult{data=models.GroupChatMassMsg} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/group-chat/mass-msg [post]
func (ch GroupChatMassMsg) Create(c *gin.Context) {
	req := requests.SendGroupChatMassMsgReq{}
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
	msg, err := ch.srv.SendMassMsgAsync(req, staffAdminInfo.ExtID, staffAdminInfo.ExtCorpID)
	if err != nil {
		err := errors.Wrap(err, "send group msg failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(msg)
}

// Get
// @tags 客户群群发
// @Summary 获取客户群发详情
// @Produce json
// @Success 200 {object} app.JSONResult{data=models.MassMsg} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/mass-msg/{id} [get]
func (ch GroupChatMassMsg) Get(c *gin.Context) {
	handler := app.NewHandler(c)
	ID, err := handler.GetIDParam()
	if err != nil {
		err = errors.Wrap(err, "handler.GetIDParam failed")
		handler.ResponseBadRequestError(err)
		return
	}

	_, err = ch.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseBadRequestError(err)
		return
	}

	result, err := ch.srv.Get(ID)
	if err != nil {
		err = errors.Wrap(err, "Get failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(result)
}

//// GetSendMassMsgResult
//// @tags 客户管理-客户群群发
//// @Summary 获取创建群发消息的结果
//// @Produce json
//// @Success 200 {object} app.JSONResult{data=requests.SendMassMsgResp} "成功"
//// @Failure 400 {object} app.JSONResult{} "请求错误"
//// @Failure 500 {object} app.JSONResult{} "内部错误"
//// @Router /api/v1/staff-admin/customer/mass-msg/result/{id} [get]
//func (ch GroupChatMassMsg) GetSendMassMsgResult(c *gin.Context) {
//	handler := app.NewHandler(c)
//	missionID, err := handler.GetIDParam()
//	if err != nil {
//		err = errors.Wrap(err, "handler.GetIDParam failed")
//		handler.ResponseBadRequestError(err)
//		return
//	}
//	result, err := ch.srv.GetSendMassMsgResult(missionID)
//	if err != nil {
//		err = errors.Wrap(err, "get send group msg failed")
//		handler.ResponseError(err)
//		return
//	}
//	handler.ResponseItem(result)
//}

// Delete
// @tags 客户管理-客户群群发
// @Summary 删除定时群发的消息
// @Param id path int true "消息id"
// @Produce json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/mass-msg/action/delete [post]
func (ch GroupChatMassMsg) Delete(c *gin.Context) {
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

// Query
// @tags 客户群群发
// @Summary 群发消息列表
// @Produce json
// @Param params body requests.QueryMassMsgReq true "客户群群发消息列表请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/mass-msgs [get]
func (ch GroupChatMassMsg) Query(c *gin.Context) {
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
