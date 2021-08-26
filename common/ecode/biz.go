package ecode

// 定义业务错误 20000000 - 30000000
var (
	DuplicatedPhoneError        = add(20000001) //重复手机号
	InvalidLoginError           = add(20000002) //账号或密码错误
	DisabledUserError           = add(20000003) //用户被禁用
	ForbiddenError              = add(20000004) //无权访问
	DuplicatedCorpIDError       = add(20000005) //重复CorpID
	DoNotDeleteYourSelfError    = add(20000006) //不要删除自己
	InvalidSignError            = add(20000007) //非法签名
	ExpiredSignError            = add(20000008) //签名已过期
	InvalidPathError            = add(20000009) //非法路径
	DoNotUpdateDefaultRoleError = add(20000010) //禁止修改默认角色
	InvalidCorpConfError        = add(20000011) //不正确的企业配置信息

	DuplicateQuickReplyGroupNameError = add(20000100) // 话术库组名重复,话术库业务错误 20000100 - 20000199
	NotifyTypeError                   = add(20000200) // 删人通知提醒,通知提醒错误
	DuplicateTagError                 = add(20000300) // 重复标签, <标签>错误 20000300 - 20000399
	DuplicateTagGroupError            = add(20000301) // 重复标签组
	UnsupportedMsgError               = add(20000400) // 不支持的消息类型, <群消息>错误 20000400 - 20000499
	EarlierThanNowError               = add(20000401) // 定时发送消息不能比当前时间早
	TimedMsgUnchangeableError         = add(20000402)
	NoMassMsgReceiversErr             = add(20000403) // 群发消息未找到有效接收人
	UnsupportedFileTypeError          = add(20000404) // 不支持的上传文件类型
	InfoFieldDuplicateError           = add(20000500) // 客户信息字段重复, 客户信息错误 20000500 - 20000599
	DuplicateRemarkNameError          = add(20000600) // 自定义客户信息字段名重复, 客户自定义信息错误 20000600 - 20000699
	GroupChatNotExistsError           = add(20000700) // 自动拉群 20000700
	CheckSignFailed                   = add(20000800) // 内部服务调用错误码
	NoStaffError                      = add(20000900) // 内部服务调用错误码
	TimedMsgEarlierThanNowErr         = add(20000901) // 定时发送消息不能比当前时间早
	IllegalURL                        = add(20001000)
	ParseFileUrlErr                   = add(20001001)
	FileNotExistsErr                  = add(20001002)
	NotImageFile                      = add(20002000)
	CustomerNumErr                    = add(20003001) // 客户统计，数量错误
	QuickReplyGroupNotFoundErr        = add(20004001)
	UpdateOtherRecordNotAllowedErr    = add(20004002)
	DeleteOtherRecordNotAllowedErr    = add(20004003)
	EmptyExternalContactInfoErr       = add(20005001) // 同步员工数据为空
	UnknownEventTypeErr               = add(20006001)
)

func init() {
	_commonMessage := map[int]Message{
		DuplicatedPhoneError.Code(): {
			Msg: "重复手机号",
		},
		InvalidLoginError.Code(): {
			Msg: "账号或密码错误",
		},
		DisabledUserError.Code(): {
			Msg: "用户被禁用",
		},
		ForbiddenError.Code(): {
			Msg: "无权访问",
		},
		DuplicatedCorpIDError.Code(): {
			Msg: "系统已存在此CorpID",
		},
		DoNotDeleteYourSelfError.Code(): {
			Msg: "不要删除自己",
		},
		InvalidSignError.Code(): {
			Msg: "非法签名",
		},
		ExpiredSignError.Code(): {
			Msg: "签名已过期",
		},
		InvalidPathError.Code(): {
			Msg: "非法路径",
		},
		DoNotUpdateDefaultRoleError.Code(): {
			Msg: "禁止修改默认角色",
		},
		InvalidCorpConfError.Code(): {
			Msg: "不正确的企业配置信息",
		},
		NotifyTypeError.Code(): {
			Msg: "通知时间类型错误",
		},
		DuplicateTagGroupError.Code(): {
			Msg: "标签组重复",
		},
		DuplicateTagError.Code(): {
			Msg: "标签重复",
		},
		UnsupportedMsgError.Code(): {
			Msg: "不支持的消息类型",
		},
		InfoFieldDuplicateError.Code(): {
			Msg: "取消展示/确认展示 包含重复字段",
		},
		DuplicateRemarkNameError.Code(): {
			Msg: "自定义字段名重复",
		},
		GroupChatNotExistsError.Code(): {
			Msg: "自动拉群分组不存在",
		},
		CheckSignFailed.Code(): {
			Msg: "验签失败",
		},
		NoStaffError.Code(): {
			Msg: "员工列表或者员工分组列表至少需要一个不为空",
		},
		IllegalURL.Code(): {
			Msg: "URL 不正确",
		},
		NotImageFile.Code(): {
			Msg: "文件不是图片格式",
		},
		EarlierThanNowError.Code(): {
			Msg: "延迟发送时间不能比当前时间早",
		},
		ParseFileUrlErr.Code(): {
			Msg: "上传url解析错误",
		},
		FileNotExistsErr.Code(): {
			Msg: "文件不存在",
		},
		TimedMsgUnchangeableError.Code(): {
			Msg: "立即发送的消息不支持修改",
		},
		UnsupportedFileTypeError.Code(): {
			Msg: "不支持的上传文件类型",
		},
		CustomerNumErr.Code(): {
			Msg: "客户数量错误",
		},
		QuickReplyGroupNotFoundErr.Code(): {
			Msg: "未找到话术分组",
		},
		UpdateOtherRecordNotAllowedErr.Code(): {
			Msg: "不能更新别人的话术分组",
		},
		DeleteOtherRecordNotAllowedErr.Code(): {
			Msg: "不能删除别人的话术分组",
		},
		NoMassMsgReceiversErr.Code(): {
			Msg: "群发消息未找到有效接收人",
		},
		EmptyExternalContactInfoErr.Code(): {
			Msg: "空员工数据",
		},
		UnknownEventTypeErr.Code(): {
			Msg: "未知事件类型错误",
		},
	}

	for code, message := range _commonMessage {
		_messages[code] = message
	}
}
