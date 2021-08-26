package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"net/http"
	"openscrm/app/constants"
	"openscrm/app/requests"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/storage"
	"openscrm/conf"
	"path"
	"strings"
)

type Util struct {
	Base
	srv *services.Util
}

func NewUtil() *Util {
	return &Util{srv: &services.Util{}}
}

// ParseLink
// @tags 话术库
// @Summary 解析链接
// @Produce  json
// @Accept json
// @Param params body requests.ParseLinkReq true "删除渠道码分组请求"
// @Success 200 {object} app.JSONResult{data=responses.ParseLinkResp} "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/common/action/parse-link [post]
func (o *Util) ParseLink(c *gin.Context) {
	req := requests.ParseLinkReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	item, err := o.srv.ParseLink(req)
	if err != nil {
		err = errors.Wrap(err, "parse-link failed")
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(item)
}

// GetUploadURL
// @tags 文件服务
// @Summary 调试接口-获取预签名URL
// @Produce  json
// @Accept json
// @Param params body requests.GetUploadURLReq true "获取预签名URL请求"
// @Success 200 {file} file "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/storage/action/get-signed-url [post]
func (o *Util) GetUploadURL(c *gin.Context) {
	req := requests.GetUploadURLReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	if conf.Settings.App.Env != constants.DEV && conf.Settings.App.Env != constants.TEST {
		handler.ResponseError(errors.WithStack(ecode.ForbiddenError))
		return
	}

	adminInfo, err := o.GetStaffAdminInfo(handler)
	if err != nil {
		handler.ResponseError(errors.WithStack(err))
		return
	}

	ext := req.Filename[strings.LastIndexByte(req.Filename, '.')+1:]
	fileExtends := []string{"jpg", "png", "bmp", "gif", "doc", "docx", "txt", "xls", "xlsx", "pdf", "ppt", "mp4", "rm", "rmvb", "mkv", "avi"}
	if !funk.Contains(fileExtends, ext) {
		err = ecode.UnsupportedFileTypeError
		handler.ResponseError(err)
		return
	}

	obj := path.Join(adminInfo.ExtCorpID, "/", "public", "/", req.Filename)
	uploadURL, err := storage.FileStorage.SignURL(obj, http.MethodPut, int64(3600))
	if err != nil {
		handler.ResponseError(err)
		return
	}

	downloadURL, err := storage.FileStorage.SignURL(obj, http.MethodGet, int64(10*356*24*3600))
	if err != nil {
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(requests.GetUploadURLResp{
		UploadURL:   uploadURL,
		DownloadURL: downloadURL,
	})
}

// UploadMedia
// @tags 文件服务
// @Summary 上传临时素材
// @Produce  json
// @Accept json
// @Param params body requests.UploadMediaReq true "上传临时素材请求"
// @Success 200 {file} file "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /api/v1/staff-admin/action/upload-media [post]
func (o *Util) UploadMedia(c *gin.Context) {
	req := requests.UploadMediaReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	item, err := o.srv.CachedUploadMedia(req, conf.Settings.WeWork.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "CachedUploadMedia failed")
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(item)
}
