package workwx

import (
	"fmt"
	"openscrm/pkg/easywework/errcodes"
)

// ClientError  企业微信客户端 SDK 的响应错误
type ClientError struct {
	// Code 错误码，0表示成功，非0表示调用失败。
	// 开发者需根据errcode是否为0判断是否调用成功(errcode意义请见全局错误码)。
	Code errcodes.ErrCode
	// Msg 错误信息，调用失败会有相关的错误信息返回。
	// 仅作参考，后续可能会有变动，因此不可作为是否调用成功的判据。
	Msg string
}

var _ error = (*ClientError)(nil)

func (e *ClientError) Error() string {
	return fmt.Sprintf(
		"ClientError { Code: %d, Msg: %#v }",
		e.Code,
		e.Msg,
	)
}
