package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type CorpRiskMgr struct {
	Base
	srv *services.CorpRiskMgrService
}

// Update
// @tags 企业风控
// @Summary 设置员工删除客户的通知开关
// @Produce  json
// @Accept json
// @Param params body requests.UpdateStaffDeleteCustomerNotifierReq true "设置提醒开关请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff_admin/notify/delete-customer/status [put]
func (m CorpRiskMgr) Update(c *gin.Context) {
	req := requests.UpdateStaffDeleteCustomerNotifierReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := m.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	err = m.srv.Upsert(req, staffAdmin.ExtCorpID)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

// Query
// @tags 企业风控
// @Summary 员工删人记录
// @Produce  json
// @Accept json
// @Param params body requests.QueryStaffDeleteCustomerHistoryReq true "查询删人记录请求"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.StaffDeleteCustomer}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/action/notify/delete-customer [get]
func (m CorpRiskMgr) Query(c *gin.Context) {
	req := requests.QueryStaffDeleteCustomerHistoryReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := m.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	items, total, err := m.srv.Query(req, staffAdmin.ExtCorpID, &req.Sorter, &req.Pager)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(items, total)
}

// Get
// @tags 企业风控
// @Summary 获取员工删除客户的通知状态
// @Produce json
// @Accept json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff_admin/notify/delete_customer/status [get]
func (m CorpRiskMgr) Get(c *gin.Context) {
	handler := app.NewHandler(c)
	staffAdmin, err := m.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	items, err := m.srv.Get(staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "get notify status failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(items)
}

func (m CorpRiskMgr) QueryDeleteStaff(c *gin.Context) {

}

func NewCorpRiskMgr() *CorpRiskMgr {
	return &CorpRiskMgr{srv: services.NewCorpRiskMgrService()}
}
