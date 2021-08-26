package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"net/http"
	"openscrm/common/ecode"
	"openscrm/common/log"
	"openscrm/pkg/easywework"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

var vt *ut.Translator

func NewBindingValidator(trans *ut.Translator) {
	vt = trans
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v interface{}) ValidErrors {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			errs = append(errs, &ValidError{
				Key:     "",
				Message: err.Error(),
			})
			return errs
		}

		trans, _ := (*vt).(ut.Translator)
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		//return false, errs
		return errs
	}

	return nil
}

func ResponseErr(c *gin.Context, err error) {
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
			Response(c, http.StatusInternalServerError, e.Code(), nil, e.Message())
			log.TracedError("InternalServerError", err)
		} else {
			//自定义非内部错误，响应http 200
			Response(c, http.StatusOK, e.Code(), nil, e.Message())
			log.TracedError("BizError", err)
		}
		return
	}

	//如果根错误是系统错误（不可控错误）
	if _, ok := rootErr.(error); ok {
		Response(c, http.StatusInternalServerError, 500, nil, err.Error())
		log.TracedError("InternalServerError", err)
		return
	}

	//没有wrap过的系统错误
	c.Error(err)
	Response(c, http.StatusInternalServerError, ecode.InternalError.Code(), nil, err.Error())
	log.TracedError("InternalServerError", err)

}
func Response(c *gin.Context, httpCode int, errorCode int, data interface{}, msg string) {
	c.Header("Content-Type", "application/json")
	c.JSON(httpCode, JSONResult{
		Code:    errorCode,
		Message: msg,
		Data:    data,
	})
	return
}
func ResponseItems(c *gin.Context, items interface{}, totalRows int64) {
	c.JSON(http.StatusOK, JSONResult{
		Code:    0,
		Message: "",
		Data: ItemsData{
			Items: items,
			Pager: Pager{Page: GetPage(c), PageSize: GetPageSize(c), TotalRows: totalRows},
		}},
	)
}
func ResponseItem(c *gin.Context, item interface{}) {
	c.JSON(http.StatusOK, JSONResult{
		Code:    0,
		Message: "ok",
		Data:    item,
	},
	)
}
