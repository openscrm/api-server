package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type TagGroup struct {
	Base
	srv *services.TagGroupService
}

func NewTagGroup() *TagGroup {
	return &TagGroup{srv: services.NewTagGroupService()}
}

// Query
// @tags 客户标签
// @Summary 获取多个标签组及其所含标签
// @Produce json
// @Param params body requests.TagListReq true "添加标签组"
// @Success 200 {object} app.JSONResult{data=models.TagGroupSwagger} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/tag-group [get]
func (t TagGroup) Query(c *gin.Context) {
	request := requests.TagListReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&request)
	if !ok {
		handler.ResponseBadRequestError(err)
		return
	}

	adminInfo, err := t.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	tags, total, err := t.srv.Query(request, adminInfo.ExtCorpID)
	if err != nil {
		log.Sugar.Errorf("Query errs: %v", err)
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(tags, total)
}

// Create
// @tags 客户标签
// @Param params body requests.CreateTagGroupReq true "添加标签组"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/tag-group [post]
func (t TagGroup) Create(c *gin.Context) {
	req := requests.CreateTagGroupReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	adminInfo, err := t.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	item, err := t.srv.Create(req, adminInfo.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "add group failed")
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(item)
}

// Delete
// @tags 客户标签
// @Summary 删除某个标签组下的标签
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/tag-group/action/delete [post]
func (t TagGroup) Delete(c *gin.Context) {
	req := requests.DeleteTagGroupsReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}
	adminInfo, err := t.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(errors.WithStack(err))
		return
	}
	n, err := t.srv.Delete(req, adminInfo.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "delete group failed")
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(n)
}

// Update
// @tags 客户标签
// @Summary 更新标签组，组内新增标签，组内删除标签，更改某个标签名字
// @Param params body requests.UpdateTagGroupReq true "更新标签组"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/tag-group/{ext_id} [post]
func (t TagGroup) Update(c *gin.Context) {
	var req = requests.UpdateTagGroupReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	ExtID, err := handler.GetStringParam("ext_id")
	if err != nil {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}
	adminInfo, err := t.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(errors.WithStack(err))
		return
	}

	req.ExtID = ExtID
	tagGroup, err := t.srv.Update(req, adminInfo.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "update group failed")
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(tagGroup)
}

// ExchangeOrder
// @tags 客户标签
// @Summary 客户标签-调整排序
// @Param params body  requests.ExchangeOrderReq true "调整排序请求"
// @Produce json
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/tag-group/action/exchange-order [post]
func (t TagGroup) ExchangeOrder(c *gin.Context) {
	req := requests.ExchangeOrderReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		log.Sugar.Debugw("BadRequest", "err", err)
		handler.ResponseBadRequestError(err)
		return
	}

	err = t.srv.ExchangeOrder(&req)
	if err != nil {
		err = errors.Wrap(err, "exchanger order failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}
