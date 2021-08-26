package constants

// ContactWayScene 渠道码场景
// 1-在小程序中联系
// 2-通过二维码联系
type ContactWayScene int

const (
	// ContactWaySceneMicroApp 在小程序中联系
	ContactWaySceneMicroApp ContactWayScene = 1
	// ContactWaySceneQrcode 通过二维码联系
	ContactWaySceneQrcode ContactWayScene = 2
)

// ContactWayType 渠道码联系方式
// 1-单人
// 2-多人
type ContactWayType int

const (
	// ContactWayTypeSingle 单人
	ContactWayTypeSingle ContactWayType = 1
	// ContactWayTypeMultiple 多人
	ContactWayTypeMultiple ContactWayType = 2
)

// ContactWayAutoReplyType 渠道码欢迎语类型
// 1, 渠道欢迎语
// 2, 渠道默认欢迎语
// 3, 不发送欢迎语
type ContactWayAutoReplyType int

const (
	// ContactWayAutoReplyTypeCustom 渠道欢迎语
	ContactWayAutoReplyTypeCustom ContactWayAutoReplyType = 1
	// ContactWayAutoReplyTypeDefault 渠道默认欢迎语
	ContactWayAutoReplyTypeDefault ContactWayAutoReplyType = 2
	// ContactWayAutoReplyTypeDisable 不发送欢迎语
	ContactWayAutoReplyTypeDisable ContactWayAutoReplyType = 3
)

const ContactWayStatePrefix = "ixj:"
