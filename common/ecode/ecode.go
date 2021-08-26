package ecode

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"sync"
)

type Message struct {
	Msg    string
	Detail string
}

var (
	_codes                  = map[int]struct{}{}
	_mutex                  sync.Mutex
	_internalErrorCodeLimit = 5000
	Zh                      = "zh-CN"
	En                      = "en"
)

//GetMessages 获取所有错误码消息
func GetMessages() map[int]Message {
	return _messages
}

//RegisterMessages 注册错误码对应消息
func RegisterMessages(messages map[int]Message) {
	_mutex.Lock()
	defer _mutex.Unlock()
	for k, v := range messages {
		_messages[k] = v
	}
}

//New 创建一个错误
func New(e int) Code {
	if e <= _internalErrorCodeLimit {
		panic(fmt.Sprintf("business ecode must greater than %d", _internalErrorCodeLimit))
	}
	return add(e)
}

//add 添加一个内部错误
func add(e int) Code {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = struct{}{}
	return Int(e)
}

type Codes interface {
	Error() string
	Code() int
	Message() string
	Detail() string
	LocalizedMessage(lang string) string
	StatusCode() int
}

func (e Code) Detail() string {
	if m, ok := _messages[e.Code()]; ok {
		if m.Detail != "" {
			return m.Detail
		}
	}
	return strconv.FormatInt(int64(e.Code()), 10)
}

func (e Code) StatusCode() int {
	switch e.Code() {
	case OK.Code():
		return http.StatusOK
	case InternalError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case NoPermissionError.Code():
		fallthrough
	case TokDetailExpiredError.Code():
		fallthrough
	case InvalidTokDetailError.Code():
		fallthrough
	case TokDetailRequiredError.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	// 重试可解决的返回500
	// 需要更新请求再重试的200和具体业务error code
	return http.StatusOK
}

type Code int

//Error 返回错误信息
func (e Code) Error() string {
	return e.Message()
}

//Code 返回错误码
func (e Code) Code() int { return int(e) }

//LocalizedMessage 返回本地化的错误消息
func (e Code) LocalizedMessage(lang string) string {
	if m, ok := _messages[e.Code()]; ok {
		if lang == Zh && m.Msg != "" {
			return m.Msg
		}
		return m.Msg
	}
	return strconv.FormatInt(int64(e.Code()), 10)
}

//Message 返回英文错误消息，或错误码
func (e Code) Message() string {
	return e.LocalizedMessage(Zh)
}

//IsInternalError 检查是否是内部错误
func (e Code) IsInternalError() bool {
	if e.Code() <= _internalErrorCodeLimit && e.Code() > 0 {
		return true
	}
	return false
}

//Int 使用int类型创建错误码
func Int(i int) Code { return Code(i) }

//String 使用string类型创建错误码
func String(e string) Code {
	if e == "" {
		return OK
	}
	// try error string
	i, err := strconv.Atoi(e)
	if err != nil {
		return InternalError
	}
	return Code(i)
}

// Cause cause from error to ecode.
func Cause(e error) Codes {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(Codes)
	if ok {
		return ec
	}
	return UnknownError
}

// Equal equal a and b by code int.
func Equal(a, b Codes) bool {
	if a == nil {
		a = OK
	}
	if b == nil {
		b = OK
	}
	return a.Code() == b.Code()
}

// EqualError equal error
func EqualError(code Codes, err error) bool {
	return Cause(err).Code() == code.Code()
}
