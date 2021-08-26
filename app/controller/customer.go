package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/entities"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
	"openscrm/common/we_work"
)

type CustomerHandler struct {
	Base
	srv *services.CustomerService
}

func NewCustomer() *CustomerHandler {
	return &CustomerHandler{srv: services.NewCustomer()}
}

func (o CustomerHandler) NotifyVerify(c *gin.Context) {
	we_work.Callback.EchoTestHandler(c.Writer, c.Request)
}

// Sync
// @tags 客户管理
// @Summary 同步企微客户数据
// @Produce json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/action/sync [post]
func (o CustomerHandler) Sync(c *gin.Context) {
	handler := app.NewHandler(c)
	staffAdmin, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		handler.ResponseError(err)
		return
	}

	err = o.srv.Sync(staffAdmin.ExtCorpID)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

// Get
// @tags 客户管理
// @Summary 客户详情
// @Param id path string true "客户id"
// @Produce json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/{extID} [get]
func (o CustomerHandler) Get(c *gin.Context) {
	handler := app.NewHandler(c)

	ExtID, err := handler.GetStringParam("ext_id")
	if err != nil {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	adminInfo, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}
	customer, errs := o.srv.Get(ExtID, adminInfo.ExtCorpID, true)
	if errs != nil {
		errs := errors.Wrap(errs, "get customer_event by id failed")
		handler.ResponseError(errs)
		return
	}

	handler.ResponseItem(customer)
}

// Query
// @tags 客户管理
// @Summary 客户列表
// @Produce  json
// @Param params body requests.QueryCustomerReq true "客户列表请求"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.Customer}} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customers [get]
func (o CustomerHandler) Query(c *gin.Context) {
	request := requests.QueryCustomerReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&request)
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

	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	customers, total, err := o.srv.Query(request, staffAdmin.ExtCorpID, &pager)
	if err != nil {
		err = errors.Wrap(err, "get customers failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(customers, total)
}

// ----------------------------------

// UpdateNotifyStaffRule
// @tags 客户管理
// @Summary 设置流失客户是否通知员工
// @Produce  json
// @Param params body entities.UpdateCustomerDeleteStaffNotifierReq true "设置流失客户是否通知员工请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/action/update-loss-notify-rule [put]
func (o CustomerHandler) UpdateNotifyStaffRule(c *gin.Context) {
	req := entities.UpdateCustomerDeleteStaffNotifierReq{}
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

	err = o.srv.UpdateDeleteEventNotify(req, staffAdmin.ExtCorpID)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

// GetNotifyStaffRule
// @tags 客户管理
// @Summary 获取流失客户是否通知员工选项
// @Produce  json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/delete-staff [get]
func (o CustomerHandler) GetNotifyStaffRule(c *gin.Context) {
	handler := app.NewHandler(c)
	staffAdmin, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	resp, err := o.srv.GetNotifyStaffRule(staffAdmin.ExtCorpID)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(resp)
}

// GetLossCustomers
// @tags 客户管理
// @Summary 获取流失客户列表
// @Produce  json
// @Param params body requests.QueryCustomerLossesReq true "查询客户流失请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/losses [get]
func (o CustomerHandler) GetLossCustomers(c *gin.Context) {
	handler := app.NewHandler(c)
	req := requests.QueryCustomerLossesReq{}
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

	items, total, err := o.srv.GetCustomerLosses(req, staffAdmin.ExtCorpID, &req.Pager, &req.Sorter)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(items, total)
}

// ExportCustomerLosses
// @tags 客户管理
// @Summary 流失提醒列表下载
// @Produce  json
// @Param params body requests.QueryCustomerLossesReq true "流失提醒列表下载请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/action/customers-losses-data-export [get]
func (o CustomerHandler) ExportCustomerLosses(c *gin.Context) {
	req := requests.QueryCustomerLossesReq{}
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
	buf, fileName, err := o.srv.ExportDeleteStaffWarningList(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseFile(buf, fileName)

	return
}

// Export
// @tags 客户管理
// @Summary 客户列表下载
// @Produce  json
// @Param params body requests.QueryCustomerReq true "客户列表下载请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customers/action/export [get]
func (o CustomerHandler) Export(c *gin.Context) {
	req := requests.QueryCustomerReq{}
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

	buf, filename, err := o.srv.Export(req, staffAdmin.ExtCorpID, &req.Pager)
	if err != nil {
		err = errors.Wrap(err, "get customers failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseFile(buf, filename)
}

// Statistic
// @tags 客户管理
// @Summary 客户统计
// @Produce  json
// @Param params body requests.QueryCustomerStatisticReq true "客户统计请求"
// @Success 200 {object} app.JSONResult{data=[]models.CustomerSummary} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customers/statistic [get]
func (o CustomerHandler) Statistic(c *gin.Context) {
	req := requests.QueryCustomerStatisticReq{}
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

	items, err := o.srv.CustomerStatistic(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "get customers failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(items)

}
