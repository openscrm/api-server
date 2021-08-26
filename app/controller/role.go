package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/conf"
)

type Role struct {
	Base
	srv *services.Role
}

func NewRole() *Role {
	return &Role{srv: services.NewRole()}
}

// Query
// @tags 角色
// @Summary 查询角色列表
// @Produce  json
// @Accept json
// @Param params body requests.QueryRoleReq true "查询角色列表请求"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]responses.Role}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/corp_admin/roles [get]
func (o *Role) Query(c *gin.Context) {
	req := requests.QueryRoleReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	items, total, err := o.srv.Query(req, conf.Settings.WeWork.ExtCorpID, &req.Sorter, &req.Pager)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}

	handler.ResponseItems(items, total)
}

// Get
// @tags 角色
// @Summary 获取角色详情
// @Produce  json
// @Accept json
// @Param id path string true "角色ID"
// @Success 200 {object} app.JSONResult{data=responses.Role} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/corp_admin/role/{id} [get]
func (o *Role) Get(c *gin.Context) {
	handler := app.NewHandler(c)
	id, err := handler.GetIDParam()
	if err != nil {
		err = errors.Wrap(err, "handler.GetIDParam failed")
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	item, err := o.srv.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Get failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// Create
// @tags 角色
// @Summary 创建角色
// @Produce  json
// @Accept json
// @Param params body requests.CreateRoleReq true "创建角色请求"
// @Success 200 {object} app.JSONResult{data=models.Role} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/corp_admin/role [post]
func (o *Role) Create(c *gin.Context) {
	req := requests.CreateRoleReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	item, err := o.srv.Create(req, conf.Settings.WeWork.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Create failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// Update
// @tags 角色
// @Summary 更新角色
// @Produce  json
// @Accept json
// @Param id path string true "角色ID"
// @Param params body requests.UpdateRoleReq true "更新角色请求"
// @Success 200 {object} app.JSONResult{data=models.Role} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/corp_admin/role/{id} [put]
func (o *Role) Update(c *gin.Context) {
	req := requests.UpdateRoleReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	id, err := handler.GetIDParam()
	if err != nil {
		err = errors.Wrap(err, "handler.GetIDParam failed")
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	item, err := o.srv.Update(id, req)
	if err != nil {
		err = errors.Wrap(err, "Update failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// AssignToStaffs
// @tags 角色
// @Summary 授权角色给员工
// @Produce  json
// @Accept json
// @Param params body requests.AssignToStaffsReq true "授权角色给员工请求"
// @Success 200 {object} app.JSONResult{data=int} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/corp_admin/role/action/assign-to-staffs [post]
func (o *Role) AssignToStaffs(c *gin.Context) {
	req := requests.AssignToStaffsReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	total, err := o.srv.AssignToStaffs(req.ExtStaffIDs, req.RoleID)
	if err != nil {
		err = errors.Wrap(err, "AssignToStaffs failed")
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(total)
}

// QueryStaffs
// @tags 角色
// @Summary 查询授权员工列表
// @Produce  json
// @Accept json
// @Param params body requests.QueryRoleStaffsReq true "查询授权员工列表请求参数"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.Staff}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/corp-admin/role/action/query-staffs [get]
func (o *Role) QueryStaffs(c *gin.Context) {
	req := requests.QueryRoleStaffsReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	items, total, err := o.srv.QueryStaffs(req, conf.Settings.WeWork.ExtCorpID, &req.Sorter, &req.Pager)
	if err != nil {
		err = errors.Wrap(err, "QueryStaffs failed")
		handler.ResponseError(err)
		return
	}

	handler.ResponseItems(items, total)
}
