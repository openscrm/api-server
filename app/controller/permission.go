package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/entities"
	"openscrm/app/services"
	"openscrm/common/app"
)

type Permission struct {
	Base
	srv *services.Permission
}

func NewPermission() *Permission {
	return &Permission{srv: services.NewDefaultPermission()}
}

// Query
// @tags 企业系统管理-权限
// @Summary 查询权限列表
// @Produce  json
// @Accept json
// @Param params body entities.QueryPermissionReq true "查询权限列表请求"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.Permission}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/corp-admin/permissions [get]
func (o *Permission) Query(c *gin.Context) {
	req := entities.QueryPermissionReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	items, total, err := o.srv.Query(req, &req.Sorter, &req.Pager)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(items, total)
}

// Get
// @tags 企业系统管理-权限
// @Summary 获取权限详情
// @Produce  json
// @Accept json
// @Param id path string true "权限ID"
// @Success 200 {object} app.JSONResult{data=models.Permission} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/corp-admin/permission/{id} [get]
func (o *Permission) Get(c *gin.Context) {
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
