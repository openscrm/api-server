package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type JsApiHandler struct {
	Base
	srv *services.JsApiService
}

// GetJsConfig
// @tags JS_SDK
// @Summary 获取JS_SDK企业级别config所需参数
// @Produce  json
// @Accept json
// @Param id path string true "组ID"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/action/get-js-config [any]
func (o JsApiHandler) GetJsConfig(c *gin.Context) {
	req := requests.GetJSConfigReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staff, err := o.GetStaffInfo(handler)
	if err != nil {
		log.TracedError("GetStaffInfo failed", err)
		return
	}

	resp, err := o.srv.GetJSConfig(req.URL, staff.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "GetJSConfig failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(resp)

}

// GetJsAgentConfig
// @tags JS_SDK
// @Summary 获取JS_SDK应用级别agentConfig所需参数
// @Produce  json
// @Accept json
// @Param id path string true "组ID"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/action/get-js-agent-config [any]
func (o JsApiHandler) GetJsAgentConfig(c *gin.Context) {
	req := requests.GetJSAgentConfigReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staff, err := o.GetStaffInfo(handler)
	if err != nil {
		log.TracedError("GetStaffInfo failed", err)
		return
	}

	resp, err := o.srv.GetJSAgentConfig(req.URL, staff.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "GetJSAgentConfig failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(resp)

}

func NewJsApiHandler() *JsApiHandler {
	return &JsApiHandler{srv: services.NewJsApiService()}
}
