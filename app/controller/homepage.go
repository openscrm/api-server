package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type HomePageHandler struct {
	Base
	customerSrv *services.CustomerService
	staffSrv    *services.StaffService
}

func NewHomePageHandler() *HomePageHandler {
	return &HomePageHandler{
		customerSrv: services.NewCustomer(),
		staffSrv:    services.NewStaffService(),
	}
}

// GetCustomerSummary
// @tags 首页
// @Summary 客户统计
// @Produce json
// @Accept json
// @Param params body requests.CreateClueManualReq true "客户统计请求"
// @Success 200 {object} app.JSONResult{data=models.CustomerSummary} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/action/get-summary [get]
func (o HomePageHandler) GetCustomerSummary(c *gin.Context) {
	handler := app.NewHandler(c)
	admin, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}
	customer, errs := o.staffSrv.GetCustomerSummary(admin.ExtID, admin.ExtCorpID)
	if errs != nil {
		errs := errors.Wrap(errs, "get customer_event by id failed")
		handler.ResponseError(errs)
		return
	}
	handler.ResponseItem(customer)
}

// GetCustomersTrend
// @tags 首页
// @Summary 客户数曲线图统计
// @Produce  json
// @Param params body requests.QueryCustomerStatisticReq true "客户数曲线图统计请求"
// @Success 200 {object} app.JSONResult{data=[]models.CustomerTrend} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/action/get-trend [get]
func (o HomePageHandler) GetCustomersTrend(c *gin.Context) {
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

	items, err := o.customerSrv.CustomerStatistic(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "get customers failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(items)

}
