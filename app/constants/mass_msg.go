package constants

const NotifyStaffSendMassMsg = `【管理员】提醒你发送群发任务
任务创建于%s，将群发给%s等%d个客户，可前往【客户联系】中确认发送`

// ChatType 创建企业群发消息的发送类型
type ChatType string

const (
	// Single 表示发送给客户
	Single ChatType = "single"
	// Group 表示发送给客户群
	Group ChatType = "group"
)

// SendMassMsgType 创建企业群发消息的发送时间类型
type SendMassMsgType uint8

const (
	// Instant 立即发送
	Instant SendMassMsgType = 1
	// Timed 定时发送
	Timed SendMassMsgType = 2
)

// 客户信息info字段
const (
	PhoneNumber = "phone_number"
	Age         = "age"
	Email       = "email"
	Birthday    = "birthday"
	Weibo       = "weibo"
	Address     = "address"
	Description = "description"
	QQ          = "qq"
)

// QuickReplyContentType 话术库/快速恢复的类型
type QuickReplyContentType uint8

const (
//QuickReplyContentTypeText  = 1
//QuickReplyContentTypePic   = 2
//QuickReplyContentTypeNews  = 3 // website,we_work:news
//QuickReplyContentTypePDF   = 4
//QuickReplyContentTypeVideo = 5
)

type SendMassMsgStatus uint8

const (
	NotActive SendMassMsgStatus = 1 // 定时发送，尚未到指定时间
	Sending   SendMassMsgStatus = 2 // 发送中, 已经提交给微信，部分员工已发送，部分员工还没发送
	Sent      SendMassMsgStatus = 3 // 发送成功，所有员工均已发送
	Deleted   SendMassMsgStatus = 4 // 任务已取消
	Failed    SendMassMsgStatus = 5 // 提交给微信时失败
)
