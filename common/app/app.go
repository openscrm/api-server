package app

import (
	"bytes"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"net/http"
	"openscrm/app/constants"
	"openscrm/common/ecode"
	"openscrm/common/log"
	"openscrm/common/util"
	"openscrm/conf"
	"openscrm/pkg/easywework"
)

type Handler struct {
	Ctx *gin.Context
	// CustomerSession 前台客户会话，主要用于企业微信内h5
	CustomerSession sessions.Session
	// StaffSession 前台员工会话，主要用于企业微信侧边栏
	StaffSession sessions.Session
	// StaffAdminSession 企业员工后台会话，主要用于企业微信扫码登录后台
	StaffAdminSession sessions.Session
	// CorpAdminSession 企业超级管理员会话，主要用于企业超级管理员修改企业信息
	CorpAdminSession sessions.Session
	// SaasAdminSession Saas管理员会话，主要用于管理租户
	SaasAdminSession sessions.Session
}

func NewDummyHandler(c *gin.Context) *Handler {
	return &Handler{Ctx: c}
}

func NewHandler(ctx *gin.Context) *Handler {
	handler := &Handler{
		Ctx:               ctx,
		CustomerSession:   sessions.DefaultMany(ctx, string(constants.CustomerSessionName)),
		StaffSession:      sessions.DefaultMany(ctx, string(constants.StaffSessionName)),
		StaffAdminSession: sessions.DefaultMany(ctx, string(constants.StaffAdminSessionName)),
		CorpAdminSession:  sessions.DefaultMany(ctx, string(constants.CorpAdminSessionName)),
		SaasAdminSession:  sessions.DefaultMany(ctx, string(constants.SaasAdminSessionName)),
	}

	handler.CustomerSession.Options(sessions.Options{
		MaxAge: 86400 * 7,
		Path:   "/",
	})
	handler.StaffSession.Options(sessions.Options{
		MaxAge: 86400 * 30, // 侧边栏会话时效长一些
		Path:   "/",
	})
	handler.StaffAdminSession.Options(sessions.Options{
		MaxAge: 86400 * 7,
		Path:   "/",
	})
	handler.CorpAdminSession.Options(sessions.Options{
		MaxAge: 86400 * 7,
		Path:   "/",
	})
	handler.SaasAdminSession.Options(sessions.Options{
		MaxAge: 86400 * 7,
		Path:   "/",
	})

	return handler
}

// GetSessionVal 从会话中获取指定字段数据
// result 接收结果
func (r *Handler) GetSessionVal(sessionName constants.SessionName, sessionField constants.SessionField, result interface{}) (err error) {
	sess := sessions.DefaultMany(r.Ctx, string(sessionName))
	jsonData, ok := (sess.Get(string(sessionField))).(string)
	if !ok {
		err = errors.WithStack(ecode.InvalidSessionError)
		return
	}

	if jsonData == "" {
		err = errors.WithStack(ecode.InvalidSessionError)
		return
	}

	err = jsoniter.UnmarshalFromString(jsonData, &result)
	if err != nil {
		err = errors.Wrap(err, "UnmarshalFromString failed")
		return
	}

	return
}

func (r *Handler) GetExtDeptIDInt64() (id int64, err error) {
	idStr := r.Ctx.Param("ext_dept_id")
	if idStr == "" {
		err = errors.New("企业微信部门id为空")
		return
	}

	id, err = util.ShouldInt64ID(idStr)
	if err != nil {
		err = errors.Wrap(err, "util.ShouldInt64ID failed")
		return
	}

	return
}

// GetIDParam 从路径中获取ID，并校验是否为int64
func (r *Handler) GetIDParam() (id string, err error) {
	id, err = r.GetStringParam("id")
	if err != nil {
		return
	}

	int64ID, err := util.ShouldInt64ID(id)
	if err != nil {
		err = errors.Wrap(err, "util.ShouldInt64ID failed")
		return
	}

	id = fmt.Sprintf("%d", int64ID)

	return
}

// GetStringParam 从路径中获取字符串参数
func (r *Handler) GetStringParam(name string) (id string, err error) {
	id = r.Ctx.Param(name)
	if id == "" {
		err = errors.New(name + " param is required")
		return
	}
	return
}

// BindAndValidateReq 绑定并校验请求参数
// req 必须是指针
func (r *Handler) BindAndValidateReq(req interface{}) (bool, error) {
	err := BindAndValid(r.Ctx, req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Handler) Response(httpCode int, errorCode int, data interface{}, msg string) {
	r.Ctx.Header("Content-Type", "application/json")
	r.Ctx.JSON(httpCode, JSONResult{
		Code:    errorCode,
		Message: msg,
		Data:    data,
	})
	return
}

func (r *Handler) ResponseRawData(data interface{}) {
	if data == nil {
		data = JSONResult{Code: 0, Message: "ok"}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Handler) ResponseItem(item interface{}) {
	r.Ctx.JSON(http.StatusOK, JSONResult{
		Code:    0,
		Message: "ok",
		Data:    item,
	},
	)
}

func (r *Handler) ResponseItems(items interface{}, totalRows int64) {
	r.Ctx.JSON(http.StatusOK, JSONResult{
		Code:    0,
		Message: "",
		Data: ItemsData{
			Items: items,
			Pager: Pager{Page: GetPage(r.Ctx), PageSize: GetPageSize(r.Ctx), TotalRows: totalRows},
		}},
	)
}

// ResponseBadRequestError 响应400非法请求错误
func (r *Handler) ResponseBadRequestError(err error) {
	if conf.Settings.App.Env == constants.DEV || conf.Settings.App.Env == constants.TEST {
		log.TracedError("BadRequestError", err)
	}

	//如果根错误是自定义错误码（可控错误）
	rootErr := errors.Cause(err) //获取根错误
	if e, ok := rootErr.(ecode.Code); ok {
		r.Response(http.StatusOK, e.Code(), nil, err.Error())
		return
	}

	r.Response(http.StatusOK, ecode.BadRequest.Code(), nil, err.Error())
}

func (r *Handler) ResponseError(err error) {

	//检查wrap过的错误
	rootErr := errors.Cause(err) //获取根错误

	// 微信返回的错误转换为ecode错误码
	if clientError, ok := rootErr.(*workwx.ClientError); ok {
		rootErr = ecode.Int(int(clientError.Code))
	}

	//如果根错误是自定义错误码（可控错误）
	if e, ok := rootErr.(ecode.Code); ok {
		//当根错误是自定义错误时
		//自定义内部错误，响应http 500
		if e.IsInternalError() {
			r.Response(http.StatusInternalServerError, e.Code(), nil, e.Message())
			log.TracedError("InternalServerError", err)
		} else {
			//自定义非内部错误，响应http 200
			r.Response(http.StatusOK, e.Code(), nil, e.Message())
			log.TracedError("BizError", err)
		}
		return
	}

	//如果根错误是系统错误（不可控错误）
	if _, ok := rootErr.(error); ok {
		r.Response(http.StatusInternalServerError, 500, nil, err.Error())
		log.TracedError("InternalServerError", err)
		return
	}

	//没有wrap过的系统错误
	r.Ctx.Error(err)
	r.Response(http.StatusInternalServerError, ecode.InternalError.Code(), nil, err.Error())
	log.TracedError("InternalServerError", err)
}

func (r *Handler) ResponseFile(buf *bytes.Buffer, fileName string) {
	//
	//r.Ctx.Header("Content-Type", "text/csv")
	//r.Ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	//r.Ctx.Header("Content-Transfer-Encoding", "binary")
	//r.Ctx.Header("Cache-Control", "no-cache")
	contentLength := buf.Len()
	contentType := "text/csv"
	extraHeaders := map[string]string{"Content-Disposition": `attachment; filename=` + fileName}

	r.Ctx.DataFromReader(http.StatusOK, int64(contentLength), contentType, buf, extraHeaders)
}
