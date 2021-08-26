package constants

// Boolean 为避免系统中出现零值，这里定义了一个自定义Boolean类型
// 1, true
// 2, false
type Boolean int

const (
	Enable  Boolean = 1
	Disable Boolean = 2
	True    Boolean = 1
	False   Boolean = 2
)

func (o Boolean) Bool() bool {
	return o == True
}

type AsyncTaskStatus string

const (
	AsyncTaskStatusCreating AsyncTaskStatus = "creating"
	AsyncTaskStatusSuccess  AsyncTaskStatus = "success"
	AsyncTaskStatusFailed   AsyncTaskStatus = "failed"
)

const (
	WeWorkPicHost = "wework.qpic.cn"
	WxPicHost     = "wx.qlogo.cn"
)

const (
	// MsgArchSrvPathSync todo
	MsgArchSrvPathSync     = "/chat-msg/sync"
	MsgArchSrvPathSessions = "/chat-msg/sessions"
	MsgArchSrvPathMsgs     = "/chat-msg/session-msgs"
	MsgArchSrvSearchMsgs   = "/chat-msg/search"
)

type LogicalCondition string

const (
	LogicalConditionAND  string = "and"
	LogicalConditionOR   string = "or"
	LogicalConditionNone string = "none"
)

type JsonResult struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
