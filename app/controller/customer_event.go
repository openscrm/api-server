package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type CustomerEvent struct {
	Base
	srv *services.CustomerEvent
}

func NewCustomerEvent() *CustomerEvent {
	return &CustomerEvent{srv: services.NewCustomerEvent()}
}

// Query
// @tags 客户管理
// @Summary 客户动态列表
// @Produce  json
// @Accept json
// @Param params body requests.QueryEventListReq true "查询客户动态列表请求"
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.CustomerEvent}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/customer/events [get]
func (o CustomerEvent) Query(c *gin.Context) {
	req := requests.QueryEventListReq{}
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

	items, total, err := o.srv.Query(req, staffAdmin.ExtCorpID, &req.Sorter, &req.Pager)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(items, total)
}
