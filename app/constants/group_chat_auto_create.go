package constants

type GroupChatAutoCreateType uint8

const (
	GroupChatAutoCreateTypeGroupQRCode = 1
	GroupChatAutoCreateTypeLiveCode    = 2
)

// GroupChatAutoCreateCodeScene 自动拉群码场景
// 1-在小程序中联系
// 2-通过二维码联系
type GroupChatAutoCreateCodeScene int

const (
	// GroupChatAutoCreateCodeSceneMicroApp 在小程序中联系
	GroupChatAutoCreateCodeSceneMicroApp GroupChatAutoCreateCodeScene = 1
	// GroupChatAutoCreateCodeSceneQrcode 通过二维码联系
	GroupChatAutoCreateCodeSceneQrcode GroupChatAutoCreateCodeScene = 2
)

// GroupChatAutoCreateCodeType 自动拉群码联系方式
// 1-单人
// 2-多人
type GroupChatAutoCreateCodeType int

const (
	// GroupChatAutoCreateCodeTypeSingle 单人
	GroupChatAutoCreateCodeTypeSingle GroupChatAutoCreateCodeType = 1
	// GroupChatAutoCreateCodeTypeMultiple 多人
	GroupChatAutoCreateCodeTypeMultiple GroupChatAutoCreateCodeType = 2
)

// GroupChatAutoCreateCodeAutoReplyType 自动拉群码欢迎语类型
// 1, 自动拉群欢迎语
// 2, 自动拉群默认欢迎语
// 3, 不发送欢迎语
type GroupChatAutoCreateCodeAutoReplyType int

const (
	// GroupChatAutoCreateCodeAutoReplyTypeCustom 自动拉群欢迎语
	GroupChatAutoCreateCodeAutoReplyTypeCustom GroupChatAutoCreateCodeAutoReplyType = 1
	// GroupChatAutoCreateCodeAutoReplyTypeDefault 自动拉群默认欢迎语
	GroupChatAutoCreateCodeAutoReplyTypeDefault GroupChatAutoCreateCodeAutoReplyType = 2
	// GroupChatAutoCreateCodeAutoReplyTypeDisable 不发送欢迎语
	GroupChatAutoCreateCodeAutoReplyTypeDisable GroupChatAutoCreateCodeAutoReplyType = 3
)

const GroupChatAutoCreateCodeStatePrefix = "join_group:"

type GroupChatAutoJoinQRCodeStatus int

const (
	GroupChatAutoJoinQRCodeStatusInEffect   = 1
	GroupChatAutoJoinQRCodeStatusTerminated = 2
)
