package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/services"
	"openscrm/common/app"
)

type CustomerFrontend struct {
	Base
	srv *services.CustomerService
}

// Get
// @tags 员工前台
// @Summary 客户画像
// @Param id path string true "客户外部ID"
// @Produce json
// @Success 200 {object} app.JSONResult{data=responses.FullCustomerInfo} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-frontend/customer/{ext_id} [get]
func (o *CustomerFrontend) Get(c *gin.Context) {
	handler := app.NewHandler(c)
	extCustomerID, err := handler.GetStringParam("ext_id")
	if err != nil {
		handler.ResponseBadRequestError(err)
		return
	}

	staff, err := o.GetStaffInfo(handler)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	item, err := o.srv.GetFullCustomerInfo(extCustomerID, staff.ExtID, staff.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Query failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

func NewCustomerFrontend() *CustomerFrontend {
	return &CustomerFrontend{srv: services.NewCustomer()}
}
