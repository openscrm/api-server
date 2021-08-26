package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
)

type CustomerGroupTagGroup struct {
	Base
	srv *services.GroupChatTagGroup
}

func NewGroupChatTagGroup() *CustomerGroupTagGroup {
	return &CustomerGroupTagGroup{srv: services.NewGroupChatTagGroup()}
}

// Create
// @tags 客户群标签组
// @Summary 客户群标签组新增
// @Produce  json
// @Accept json
// @Param params body requests.CreateGroupChatTagGroupReq true "客户群标签组新增请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-group/tag-group [post]
func (o CustomerGroupTagGroup) Create(c *gin.Context) {
	req := &requests.CreateGroupChatTagGroupReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}
	adminInfo, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	tags, err := o.srv.Create(req, adminInfo.ExtCorpID, adminInfo.ExtID)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(tags)
}

// Update
// @tags 客户群标签组
// @Summary 客户群标签组更新
// @Produce  json
// @Accept json
// @Param params body requests.UpdateGroupChatTagGroupReq true "客户群标签组更新请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-group/tag-group [put]
func (o CustomerGroupTagGroup) Update(c *gin.Context) {
	req := &requests.UpdateGroupChatTagGroupReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}
	adminInfo, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	tag, err := o.srv.Update(req, adminInfo.ExtCorpID)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(tag)
}

// Delete
// @tags 客户群标签组
// @Summary 客户群标签组删除
// @Produce  json
// @Accept json
// @Param params body requests.CommonDeleteReq true "客户群标签删除请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-group/tag-group/action/delete [post]
func (o CustomerGroupTagGroup) Delete(c *gin.Context) {
	req := &requests.CommonDeleteReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	tag, err := o.srv.Delete(req.IDs)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(tag)
}

// Query
// @tags 客户群标签组
// @Summary 客户群标签组查询
// @Produce  json
// @Accept json
// @Param params body requests.QueryCustomerGroupTagGroupReq true "客户群标签组查询请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-group/tag-group [get]
func (o CustomerGroupTagGroup) Query(c *gin.Context) {
	req := &requests.QueryCustomerGroupTagGroupReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}
	adminInfo, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	items, total, err := o.srv.Query(req.Name, adminInfo.ExtCorpID, &req.Pager, &req.Sorter)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	handler.ResponseItems(items, total)
}
