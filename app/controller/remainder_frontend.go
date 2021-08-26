package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type RemainderFrontend struct {
	Base
	srv *services.Remainder
}

func NewRemainderFrontend() *RemainderFrontend {
	return &RemainderFrontend{srv: services.NewRemainder()}
}

// Create
// @tags 客户画像
// @Summary H5创建提醒
// @Produce  json
// @Accept json
// @Param params body requests.CreateRemainderReq true "H5创建提醒请求"
// @Success 200 {object} app.JSONResult{data=models.Remainder} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-frontend/customer/remainder [post]
func (o RemainderFrontend) Create(c *gin.Context) {
	req := requests.CreateRemainderReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staff, err := o.GetStaffInfo(handler)
	if err != nil {
		log.TracedError("GetStaffInfo failed", err)
		return
	}

	item, err := o.srv.Create(req, staff.ExtCorpID, staff.ExtID)
	if err != nil {
		err = errors.Wrap(err, "Create failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// Delete
// @tags 客户画像
// @Summary H5删除提醒
// @Produce  json
// @Accept json
// @Param params body requests.CommonDeleteReq true "H5删除提醒请求"
// @Success 200 {object} app.JSONResult{} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-frontend/customer/remainder/action/delete/{id} [post]
func (o RemainderFrontend) Delete(c *gin.Context) {
	handler := app.NewHandler(c)
	req := requests.CommonDeleteReq{}
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}
	staff, err := o.GetStaffInfo(handler)
	if err != nil {
		log.TracedError("GetStaffInfo failed", err)
		return
	}

	res, err := o.srv.Delete(req.IDs[0], staff.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Delete failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(res)

}

// Update
// @tags 客户画像
// @Summary H5更新提醒
// @Produce  json
// @Accept json
// @Param id path string true "更新客户画像-提醒ID"
// @Param params body requests.UpdateRemainderReq true "更新客户画像-提醒请求"
// @Success 200 {object} app.JSONResult{data=models.Remainder} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-frontend/customer/remainder/{id} [put]
func (o RemainderFrontend) Update(c *gin.Context) {
	req := requests.UpdateRemainderReq{}
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

	staff, err := o.GetStaffInfo(handler)
	if err != nil {
		log.TracedError("-GetStaffInfo failed", err)
		return
	}

	item, err := o.srv.Update(id, req, staff.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Update failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}
