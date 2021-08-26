package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type QuickReplyFrontend struct {
	Base
	srv *services.QuickReplyFrontend
}

func NewQuickReplyFrontend() *QuickReplyFrontend {
	return &QuickReplyFrontend{srv: services.NewQuickReplyFrontend()}
}

// Query
// @tags  话术库
// @Summary  H5查询企业话术库
// @Produce  json
// @Accept json
// @Success 200 {object} app.JSONResult{data=app.ItemsData{items=[]models.QuickReply}} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-frontend/quick-replies [get]
func (r *QuickReplyFrontend) Query(c *gin.Context) {
	req := requests.QueryQuickReplyReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		log.Sugar.Debugw("BadRequest", "err", err)
		handler.ResponseBadRequestError(err)
		return
	}

	staff, err := r.GetStaffInfo(handler)
	if err != nil {
		log.TracedError("GetStaffAdminInfo failed", err)
		return
	}

	item, total, err := r.srv.QueryQuickReply(req, staff.ExtCorpID, &req.Pager)
	if err != nil {
		err = errors.Wrap(err, "QueryQuickReply failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItems(item, total)
}
