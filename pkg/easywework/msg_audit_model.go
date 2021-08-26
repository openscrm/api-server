package workwx

import (
	"time"
)

// CheckMsgAuditSingleAgreeUserInfo 获取会话同意情况（单聊）内外成员
type CheckMsgAuditSingleAgreeUserInfo struct {
	// UserID 内部成员的userid
	UserID string `json:"userid"`
	// ExternalOpenID 外部成员的externalopenid
	ExternalOpenID string `json:"exteranalopenid"`
}

// CheckMsgAuditSingleAgreeInfo 获取会话同意情况（单聊）同意信息
type CheckMsgAuditSingleAgreeInfo struct {
	CheckMsgAuditSingleAgreeUserInfo
	// AgreeStatus 同意:”Agree”，不同意:”Disagree”，默认同意:”Default_Agree”
	AgreeStatus MsgAuditAgreeStatus
	// StatusChangeTime 同意状态改变的具体时间
	StatusChangeTime time.Time
}

// CheckMsgAuditRoomAgreeInfo 获取会话同意情况（群聊）同意信息
type CheckMsgAuditRoomAgreeInfo struct {
	// StatusChangeTime 同意状态改变的具体时间
	StatusChangeTime time.Time
	// AgreeStatus 同意:”Agree”，不同意:”Disagree”，默认同意:”Default_Agree”
	AgreeStatus MsgAuditAgreeStatus
	// ExternalOpenID 群内外部联系人的externalopenid
	ExternalOpenID string
}

// MsgAuditGroupChatMember 获取会话内容存档内部群成员
type MsgAuditGroupChatMember struct {
	// MemberID roomid群成员的id，userid
	MemberID int
	// JoinTime roomid群成员的入群时间
	JoinTime time.Time
}

// MsgAuditGroupChat 获取会话内容存档内部群信息
type MsgAuditGroupChat struct {
	// Members roomid对应的群成员列表
	Members []MsgAuditGroupChatMember
	// RoomName roomid对应的群名称
	RoomName string
	// Creator roomid对应的群创建者，userid
	Creator string
	// RoomCreateTime roomid对应的群创建时间
	RoomCreateTime time.Time
	// Notice roomid对应的群公告
	Notice string
}
