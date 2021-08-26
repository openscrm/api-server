package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/entities"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type Department struct {
	Base
	srv *services.Department
}

func NewDepartment() *Department {
	return &Department{srv: services.NewDepartment()}
}

// Query
// @tags 部门管理
// @Summary 查询部门信息
// @Produce json
// @Accept json
// @Param params body entities.GetSubDepartmentReq true "部门id列表"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.Department}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/departments [get]
func (d *Department) Query(c *gin.Context) {
	handler := app.NewHandler(c)
	req := entities.GetSubDepartmentReq{}
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := d.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}
	departments, total, err := d.srv.Query(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Query departments failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(departments, total)
}

// Sync
// @tags 部门管理
// @Summary 同步企微部门数据
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/department [post]
func (d *Department) Sync(c *gin.Context) {
	handler := app.NewHandler(c)
	staffAdminInfo, err := d.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}
	err = d.srv.Sync(staffAdminInfo.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "sync departments failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

// Get
// @tags 部门管理
// @Summary 获取根部门信息
// @Produce json
// @Accept json
// @Param params body requests.QueryDepartmentReq true "部门信息请求"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.Department}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/department [get]
func (d *Department) Get(c *gin.Context) {
	handler := app.NewHandler(c)
	req := requests.QueryDepartmentReq{}
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := d.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}
	departments, err := d.srv.Get(req.ExtID, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "get departments failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(departments)
}
