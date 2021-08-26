package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"openscrm/app/entities"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
)

type TempMaterial struct {
	Base
	srv *services.TempMaterial
}

func NewTempMaterial() *TempMaterial {
	return &TempMaterial{srv: services.NewTempMaterial()}
}

// Upload
// @tags 素材管理
// @Summary 上传临时素材
// @Produce  json
// @Accept json
// @Param params body entities.UploadMaterialReq true "上传临时素材请求"
// @Success 200 {object} app.JSONResult{data=workwx.MediaUploadResult} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/material/temp [post]
func (o TempMaterial) Upload(c *gin.Context) {
	req := entities.UploadMaterialReq{}
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
	resp, err := o.srv.UploadMedia(req, staffAdmin.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "upload failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(resp)
}
