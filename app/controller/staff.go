package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

// Staff staff controller
type Staff struct {
	Base
	staffService *services.StaffService
}

func NewStaff() *Staff {
	return &Staff{staffService: services.NewStaffService()}
}

// Sync 同步企微员工数据
// @tags 员工管理
// @Summary 同步企微员工数据
// @Produce json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/staff [post]
func (s Staff) Sync(c *gin.Context) {
	handler := app.NewHandler(c)
	adminInfo, err := s.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseBadRequestError(err)
		return
	}
	err = s.staffService.Sync(adminInfo.ExtCorpID)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

// Get
// @tags 员工管理
// @Summary 员工详情
// @param ext_id path string true  "企业微信员工id"
// @Produce json
// @Success 200 {object} app.JSONResult{data=models.Staff} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/staff/{ext_id} [get]
func (s Staff) Get(c *gin.Context) {
	handler := app.NewHandler(c)
	extStaffID, err := handler.GetStringParam("ext_id")
	if err != nil {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := s.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	staff, err := s.staffService.Get(extStaffID, staffAdmin.ExtCorpID)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(staff)
}

// Update
// @tags 员工管理
// @Summary 批量启用/禁用员工
// @Produce json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/staff [put]
func (s Staff) Update(c *gin.Context) {
	req := requests.EnableStaffs{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := s.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	err = s.staffService.Enable(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

// GetCurrent
// @tags 员工管理
// @Summary 获取当前登录员工详情
// @Produce json
// @Success 200 {object} app.JSONResult{data=responses.StaffDetail} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/action/get-current-staff [get]
func (s Staff) GetCurrent(c *gin.Context) {
	handler := app.NewHandler(c)
	staffAdmin, err := s.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}
	item, err := s.staffService.GetDetail(staffAdmin.ExtID, staffAdmin.ExtCorpID)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(item)
}

// Query
// @tags 员工管理
// @Summary  员工列表
// @Param params body requests.QueryStaffReq true "查询员工列表请求"
// @Produce json
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.Staff}} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/staffs [get]
func (s Staff) Query(c *gin.Context) {
	req := requests.QueryStaffReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := s.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	items, total, err := s.staffService.Query(
		req, staffAdmin.ExtCorpID, &req.Sorter, &req.Pager)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(items, total)
}

// GetStaffStaticsInfo
// @tags 员工管理
// @Summary 统计员工客户数量和变化情况
// @param ext_id path string true  "企业微信员工id"
// @Produce json
// @Success 200 {object} app.JSONResult{data=models.StaffCustomerCount} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/staff-event/statistics/{ext_id} [get]
func (s Staff) GetStaffStaticsInfo(c *gin.Context) {
	handler := app.NewHandler(c)
	extStaffID, err := handler.GetStringParam("ext_id")
	if err != nil {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := s.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}
	staff, err := s.staffService.StaffInfoStatistics(extStaffID, staffAdmin.ExtCorpID)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(staff)
}

// UpdateCustomerTag
// @tags 客户管理
// @Summary 更新客户标签
// @Param params body requests.UpdateCustomerTagsReq true "更新客户标签请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/action/update-tags [post]
func (s Staff) UpdateCustomerTag(c *gin.Context) {
	req := requests.UpdateCustomerTagsReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := s.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	var extStaffID = staffAdmin.ExtID
	//if req.ExtStaffID != "" {
	//	extStaffID = req.ExtStaffID
	//}
	err = s.staffService.UpdateCustomerTag(req, extStaffID, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

// UpdateCustomerInternalTag
// @tags 客户管理
// @Summary 更新对客户的个人标签
// @Param params body requests.UpdateCustomerInternalTagsReq true "更新对客户的个人标签请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/action/update-internal-tags [post]
func (s Staff) UpdateCustomerInternalTag(c *gin.Context) {
	req := requests.UpdateCustomerInternalTagsReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := s.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	err = s.staffService.UpdateCustomerInternalTag(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

// QueryMainInfo
// @tags 员工管理
// @Summary 员工主要信息列表
// @Produce json
// @Success 200 {object} app.JSONResult{data=models.StaffMainInfo} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff_admin/staff/action/get-all [get]
func (s Staff) QueryMainInfo(c *gin.Context) {
	req := requests.QueryMainStaffInfoReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := s.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	res, err := s.staffService.QueryMainInfo(req, staffAdmin.ExtCorpID, &req.Pager)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(res.Staffs, res.Total)
}

// ExportDeleteCustomers
// @tags 员工管理
// @Summary 员工删除客户记录导出
// @Produce json
// @Success 200 {object} app.JSONResult{} "员工删除客户记录导出"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff_admin/staff/action/delete-customers-data-export [get]
func (s Staff) ExportDeleteCustomers(c *gin.Context) {
	req := requests.QueryStaffDeleteCustomerHistoryReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := s.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}
	file, fileName, err := s.staffService.ExportDeleteCustomerList(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseFile(file, fileName)

	return
}

// GetCurrentFrontendStaff
// @tags 侧边栏-员工会话
// @Summary 获取当前登录员工详情
// @Produce json
// @Success 200 {object} app.JSONResult{data=responses.StaffDetail} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-frontend/action/get-current-staff [get]
func (s Staff) GetCurrentFrontendStaff(c *gin.Context) {
	handler := app.NewHandler(c)
	staff, err := s.GetStaffInfo(handler)
	if err != nil {
		log.TracedError("GetStaffInfo failed", err)
		return
	}
	item, err := s.staffService.GetDetail(staff.ExtID, staff.ExtCorpID)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(item)
}
