package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type GroupChat struct {
	Base
	srv *services.GroupChatService
}

func NewGroupChat() *GroupChat {
	return &GroupChat{srv: services.NewGroupChatService()}
}

// Query
// @tags 客户群
// @Summary 客户群列表
// @Produce  json
// @Accept json
// @Param params body requests.QueryGroupChatReq true "查询客户群列表请求"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.GroupChat}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-group [get]
func (o GroupChat) Query(c *gin.Context) {
	req := requests.QueryGroupChatReq{}
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

	items, total, err := o.srv.Query(req, staffAdmin.ExtCorpID, &req.Pager, &req.Sorter)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}

	handler.ResponseItems(items, total)
}

// UpdateTags
// @tags 客户群
// @Summary 客户群打标签
// @Produce  json
// @Accept json
// @Param params body requests.UpdateCustomerGroupReq true "客户群打标签请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-group/action/update-tags [post]
func (o GroupChat) UpdateTags(c *gin.Context) {
	req := requests.UpdateCustomerGroupReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	err = o.srv.UpdateTags(req)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

// Export
// @tags 客户群
// @Summary 查询客户群列表导出
// @Produce  json
// @Accept json
// @Param params body requests.QueryGroupChatReq true "查询客户群列表请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-groups/action/export [get]
func (o GroupChat) Export(c *gin.Context) {
	req := requests.QueryGroupChatReq{}
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

	buf, filename, err := o.srv.Export(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseFile(buf, filename)
}

// GetAllOwners
// @tags 客户群
// @Summary 群主列表
// @Produce json
// @Accept json
// @Success 200 {object} app.JSONResult{data=[]models.StaffMainInfo} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-group/tags [put]
func (o GroupChat) GetAllOwners(c *gin.Context) {
	handler := app.NewHandler(c)

	staffAdmin, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	items, err := o.srv.GetAllOwners(staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(items)
}

// GetAll
// @tags 客户群
// @Summary 全部列表
// @Produce json
// @Accept json
// @Param params body requests.GetAllGroupChatReq true "查询客户群列表请求"
// @Success 200 {object} app.JSONResult{data=[]models.StaffMainInfo} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer-group/action/get-all [post]
func (o GroupChat) GetAll(c *gin.Context) {
	req := requests.QueryGroupChatReq{}
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

	items, total, err := o.srv.GetAll(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(items, total)
}
