package ecode

var (
	// 定义系统错误 0 - 5000

	OK            = add(0)   // 正确
	InternalError = add(500) //内部错误
	UnknownError  = add(404) //未知错误

	// 定义通用错误 10000000 - 20000000
	InternalServiceInvalidSignError = add(10000001) //内部服务验签失败
	MissingUserIDError              = add(10000002) //缺失用户ID
	UserBusyError                   = add(10000003) //此用户正忙
	AccountNotFound                 = add(10000004) //账户未找到
	InvalidParams                   = add(10000005) //非法参数
	BadRequest                      = add(10000400) //非法请求
	NoPermissionError               = add(10000401) //无权访问
	TokDetailExpiredError           = add(10000402) //tokDetail过期
	InvalidTokDetailError           = add(10000403) //非法tokDetail
	TooManyRequests                 = add(10000429) //请求过多
	SessExpiredError                = add(10000501) //会话过期
	InvalidCipherError              = add(10000502) //无效密文
	NoFollowError                   = add(10000503) //未关注公众号
	TooManyRequestsError            = add(10000504) //请求过于频繁
	TokDetailRequiredError          = add(10000404) //无效token
	ItemNotFoundError               = add(10000886) //条目未找到
	InvalidSessionError             = add(10000887) //无效会话
)

func init() {
	_commonMessage := map[int]Message{
		OK.Code(): {
			Msg:    "成功",
			Detail: "success",
		},
		InternalError.Code(): {
			Msg:    "内部错误",
			Detail: "Internal Error",
		},
		UnknownError.Code(): {
			Msg:    "未知错误",
			Detail: "Unknown Error",
		},
		InternalServiceInvalidSignError.Code(): {
			Msg:    "内部服务验签失败",
			Detail: "Internal Service Invalid Signature Error",
		},
		MissingUserIDError.Code(): {
			Msg:    "缺失用户ID",
			Detail: "Missing UserID Error",
		},
		UserBusyError.Code(): {
			Msg:    "此用户正忙",
			Detail: "User Busy Error",
		},
		UserBusyError.Code(): {
			Msg:    "账户未找到",
			Detail: "Account Not Found Error",
		},
		BadRequest.Code(): {
			Msg:    "非法请求",
			Detail: "Bad Request",
		},
		NoPermissionError.Code(): {
			Msg:    "无权访问",
			Detail: "ForbiddDetail",
		},
		InvalidParams.Code(): {
			Msg:    "非法参数",
			Detail: "Invalid Params",
		},
		TokDetailExpiredError.Code(): {
			Msg:    "TokDetail过期",
			Detail: "TokDetail Expired",
		},
		InvalidTokDetailError.Code(): {
			Msg:    "无效TokDetail",
			Detail: "Invalid TokDetail",
		},
		SessExpiredError.Code(): {
			Msg:    "会话已过期",
			Detail: "Session Expired",
		},
		InvalidCipherError.Code(): {
			Msg:    "无效密文",
			Detail: "Invalid Cipher",
		},
		NoFollowError.Code(): {
			Msg:    "请先关注公众号",
			Detail: "No Follow Error",
		},
		TooManyRequestsError.Code(): {
			Msg:    "您的请求过于频繁，请休息一会儿",
			Detail: "Too Many Requests Error",
		},
		ItemNotFoundError.Code(): {
			Msg: "未找到指定条目",
		},
		TokDetailRequiredError.Code(): {
			Msg: "无效token",
		},
		InvalidSessionError.Code(): {
			Msg: "无效会话",
		},
	}

	for code, message := range _commonMessage {
		_messages[code] = message
	}
}
