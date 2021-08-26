package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type CustomerRemarkHandler struct {
	Base
	srv *services.CustomerRemarkService
}

// Create
// @tags 客户管理
// @Summary 添加自定义客户信息
// @Produce  json
// @Param params body requests.AddCustomerRemarkReq true "添加自定义客户信息请求"
// @Success 200 {object} app.JSONResult{data=models.CustomerRemark} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/remark  [post]
func (ch CustomerRemarkHandler) Create(c *gin.Context) {
	req := requests.AddCustomerRemarkReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		log.Sugar.Debugw("BadRequest", "err", err)
		handler.ResponseBadRequestError(err)
		return
	}

	info, err := ch.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	item, err := ch.srv.Create(&req, info.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Upsert failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// Get
// @tags 客户管理
// @Summary 客户信息设置中remark和info 所有信息
// @Produce  json
// @Success 200 {object} app.JSONResult{data=services.InfoRemark} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/remark [get]
func (ch CustomerRemarkHandler) Get(c *gin.Context) {
	handler := app.NewHandler(c)
	staff, err := ch.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	remark, err := ch.srv.Get(staff.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "get info & remark failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(remark)
}

// Delete
// @tags 客户管理
// @Summary 删除自定义客户信息
// @Produce  json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/remark/action/delete [post]
func (ch CustomerRemarkHandler) Delete(c *gin.Context) {
	handler := app.NewHandler(c)
	req := requests.DeleteCustomerRemarkReq{}
	_, err := handler.BindAndValidateReq(&req)
	if err != nil {
		handler.ResponseBadRequestError(err)
		return
	}

	info, err := ch.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	err = ch.srv.Delete(req.IDs, info.ExtCorpID)
	if err != nil {
		log.Sugar.Errorf("delete remark errs: %v", err)
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

// Update
// @tags 客户管理
// @Summary 修改自定义客户信息
// @Param params body  requests.UpdateRemarkReq true "修改自定义客户信息"
// @Produce json
// @Success 200 {object} app.JSONResult{data=models.CustomerRemark} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/remark [put]
func (ch CustomerRemarkHandler) Update(c *gin.Context) {
	req := requests.UpdateRemarkReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		log.Sugar.Debugw("BadRequest", "err", err)
		handler.ResponseBadRequestError(err)
		return
	}

	item, err := ch.srv.Update(&req)
	if err != nil {
		err = errors.Wrap(err, "Delete failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// ExchangeOrder
// @tags 客户管理
// @Summary 自定义客户信息-调整排序
// @Param params body  requests.ExchangeOrderReq true "调整排序请求"
// @Produce json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/remark/action/exchange-order [put]
func (ch CustomerRemarkHandler) ExchangeOrder(c *gin.Context) {
	req := requests.ExchangeOrderReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		log.Sugar.Debugw("BadRequest", "err", err)
		handler.ResponseBadRequestError(err)
		return
	}

	err = ch.srv.ExchangeOrder(&req)
	if err != nil {
		err = errors.Wrap(err, "Delete failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

// AddRemarkOption
// @tags 客户管理
// @Summary 增加自定义客户类型-多选类型的选项
// @Param params body requests.AddRemarkOptionReq true "新增选项请求"
// @Produce json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/remark/option [post]
func (ch CustomerRemarkHandler) AddRemarkOption(c *gin.Context) {
	req := requests.AddRemarkOptionReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		log.Sugar.Debugw("BadRequest", "err", err)
		handler.ResponseBadRequestError(err)
		return
	}
	err = ch.srv.AddRemarkOption(&req)
	if err != nil {
		log.Sugar.Errorf("app.AddRemarkOption errs: %v", err)
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
	return
}

// UpdateRemarkOption
// @tags 客户管理
// @Summary 更新自定义客户类型-多选类型的选项
// @Param params body  requests.UpdateRemarkOptionReq true "选项名称"
// @Produce json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/remark/option [put]
func (ch CustomerRemarkHandler) UpdateRemarkOption(c *gin.Context) {
	req := requests.UpdateRemarkOptionReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		log.Sugar.Debugw("BadRequest", "err", err)
		handler.ResponseBadRequestError(err)
		return
	}

	err = ch.srv.UpdateRemarkOption(&req)
	if err != nil {
		log.Sugar.Errorf("app.UpdateRemarkOption errs: %v", err)
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
	return
}

// DeleteRemarkOption
// @tags 客户管理
// @Summary 删除自定义客户类型-多选类型的选项
// @Param field_id body int true "remark中的FieldID"
// @Param id body int true "option中的ID"
// @Produce json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/remark/option/action/delete [post]
func (ch CustomerRemarkHandler) DeleteRemarkOption(c *gin.Context) {
	req := requests.DeleteRemarkOptionReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		log.Sugar.Debugw("BadRequest", "err", err)
		handler.ResponseBadRequestError(err)
		return
	}

	err = ch.srv.DeleteRemarkOption(req.IDs)
	if err != nil {
		log.Sugar.Errorf("app.DeleteRemarkOption errs: %v", err)
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
	return
}

func NewCustomerRemarkHandler() *CustomerRemarkHandler {
	return &CustomerRemarkHandler{srv: services.NewCustomerRemarkService()}
}
