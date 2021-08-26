package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
)

type CustomerGroupTag struct {
	Base
	srv *services.GroupChatTag
}

func NewGroupChatTag() *CustomerGroupTag {
	return &CustomerGroupTag{srv: services.NewCustomerGroupTag()}
}

// Create
// @tags 客户群标签
// @Summary 客户群标签新增
// @Produce  json
// @Accept json
// @Param params body requests.CreateGroupChatTagsReq true "客户群标签新增请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-group/tag [post]
func (o CustomerGroupTag) Create(c *gin.Context) {
	req := requests.CreateGroupChatTagsReq{}
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
// @tags 客户群标签
// @Summary 客户群标签更新
// @Produce  json
// @Accept json
// @Param params body requests.UpdateGroupChatTagReq true "客户群标签更新请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-group/tag [put]
func (o CustomerGroupTag) Update(c *gin.Context) {
	req := requests.UpdateGroupChatTagReq{}
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
// @tags 客户群标签
// @Summary 客户群标签删除
// @Produce  json
// @Accept json
// @Param params body requests.CommonDeleteReq true "客户群标签删除请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-group/tag/action/delete [post]
func (o CustomerGroupTag) Delete(c *gin.Context) {
	req := requests.CommonDeleteReq{}
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
