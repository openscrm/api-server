package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type StaffPublic struct {
	Base
	staffService *services.StaffService
}

func NewStaffPublic() *Staff {
	return &Staff{staffService: services.NewStaffService()}
}

// GetStaff
// @tags 员工管理
// @Summary 员工详情
// @param ext_id path string true  "企业微信员工id"
// @Produce json
// @Success 200 {object} app.JSONResult{data=models.Staff} "成功"
// @Failure 400 {object} app.JSONResult{} "请求错误"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/staff/{ext_id} [get]
func (s Staff) GetStaff(c *gin.Context) {
	handler := app.NewHandler(c)
	extStaffID, err := handler.GetStringParam("ext_id")
	if err != nil {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	staffAdmin, err := s.GetStaffAdminInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}
	staff, err := s.staffService.Get(extStaffID, staffAdmin.ExtCorpID)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(staff)
}
