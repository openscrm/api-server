package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"openscrm/app/constants"
	"openscrm/app/entities"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/storage"
	"openscrm/conf"
)

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

// GetLocalFile
// @tags 文件服务
// @Summary 获取本地存储文件
// @Produce  json
// @Accept json
// @Param params body entities.LocalStorageFileReq true "获取本地存储文件请求"
// @Success 200 {file} file "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /storage/public [get]
func (o *Storage) GetLocalFile(c *gin.Context) {
	req := entities.LocalStorageFileReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	localStorage, err := storage.NewLocalStorage(conf.Settings.Storage)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	err = localStorage.CheckSignedURL(handler.Ctx.Request.URL.String(), handler.Ctx.Request.Method, req.ExpireAt, req.Signature)
	if err != nil {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	objectKey := handler.Ctx.Param("path")
	if objectKey == "" {
		err = ecode.BadRequest
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	r, err := localStorage.Get(objectKey)
	if err != nil {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	contentType, err := storage.GetContentType(objectKey)
	if err != nil {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	defer r.Close()

	handler.Ctx.Data(http.StatusOK, contentType, data)
}

// PutLocalFile
// @tags 文件服务
// @Summary 上传本地存储文件
// @Produce  json
// @Accept json
// @Param params body entities.LocalStorageFileReq true "上传本地存储文件请求"
// @Success 200 {file} file "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /storage/public [put]
func (o *Storage) PutLocalFile(c *gin.Context) {
	req := entities.LocalStorageFileReq{}
	handler := app.NewHandler(c)
	ok, err := handler.BindAndValidateReq(&req)
	if !ok {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	localStorage, err := storage.NewLocalStorage(conf.Settings.Storage)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	err = localStorage.CheckSignedURL(handler.Ctx.Request.URL.String(), handler.Ctx.Request.Method, req.ExpireAt, req.Signature)
	if err != nil {
		handler.ResponseBadRequestError(errors.WithStack(err))
		return
	}

	objectKey := handler.Ctx.Param("path")
	if objectKey == "" {
		handler.ResponseBadRequestError(errors.WithStack(ecode.BadRequest))
		return
	}

	err = localStorage.Put(objectKey, handler.Ctx.Request.Body)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(nil)
}

// GetSignedURL
// @tags 文件服务
// @Summary 调试接口-获取预签名URL
// @Produce  json
// @Accept json
// @Param params body entities.GetSignedURLReq true "获取预签名URL请求"
// @Success 200 {file} file "成功"
// @Failure 400 {object} app.JSONResult{} "非法请求"
// @Failure 500 {object} app.JSONResult{} "内部错误"
// @Router /storage/action/get_signed_url [post]
func (o *Storage) GetSignedURL(c *gin.Context) {
	req := entities.GetSignedURLReq{}
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

	signedURl, err := storage.FileStorage.SignURL(req.ObjectKey, constants.HTTPMethod(req.Method), req.ExpiredInSec)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	handler.ResponseItem(signedURl)
}
