package models

import (
	"openscrm/app/constants"
	"time"
)

// ChatSession 会话
// 开始对话时，对方若同意存档，则记录一条会话记录
// unique: ext_staff_id(wx员工id)-ext_id(wx员工、客户id)
type ChatSession struct {
	Id            int               `json:"id"`
	ExtId         string            `json:"ext_id"`
	Name          string            `json:"name"`
	AgreeMsgaudit constants.Boolean `json:"agree_msgaudit"`
	// 同意存档的时间
	AgreeTime time.Time `json:"agree_time"`
	// 客户/员工 头像
	Avatar string `json:"avatar"`
	// 聊天状态 idle
	ChatStatus string `json:"chat_status"`
	// 0-未定义 1-男 2-女 3-未知
	Gender    constants.UserGender `json:"gender"`
	GroupType string               `json:"group_type"`
	// 最新对话内容, 如 [文字消息] [语音消息]等
	LastMsgThumb string `json:"last_msg_thumb"`
	// 最新对话时间
	LastMsgTime time.Time `json:"last_msg_time"`
	// 群聊人员ID列表，extStaffID/extCustomerID
	RoomUserIDs string `json:"room_user_ids"`
	// external/inner/group
	SessionType string `json:"session_type"`
	ExtStaffId  string `json:"ext_staff_id"`
	// 员工的企业信息 saas 可用
	// todo
	//CorpFullName    string                 `json:"corp_full_name"`
	//CorpId          string                 `json:"corp_id"`
	//CorpName        string                 `json:"corp_name"`
	//ExternalProfile models.ExternalProfile `json:"external_profile"`
	//Position string `json:"position"`
	//Type    interface{} `json:"type"`
	//Unionid interface{} `json:"unionid"`
	//LastStartTime time.Time `json:"last_start_time"`
	//LastEndTime time.Time `json:"last_end_time"`
	Timestamp
}

// SessionMsgItem 会话中的消息列表条目
type SessionMsgItem struct {
	Id int `json:"id"`
	// 该消息的外部id
	MsgID string `json:"ext_id"`
	// 消息发送人的头像
	Avatar string `json:"avatar"`
	// 非文字类容信息
	Content string `json:"content"`
	// 文字信息
	TextContent string `json:"text_content"`
	// send/receive
	Direction string `json:"direction"`
	// 消息类型 image
	MsgType string `json:"msg_type"`
	// 发送消息的时间
	Msgtime time.Time `json:"msgtime"`
	FileUrl string    `json:"file_url"`
	// 群号
	RoomId string `json:"room_id"`
	// external/inner/group
	SessionType string `json:"session_type"`
	ExtStaffID  string `json:"ext_staff_id"`
	ToId        string `json:"to_id"`
	FromId      string `json:"from_id"`
	FromName    string `json:"from_name"`
	Name        string `json:"name"`

	// todo
	//SessionId string `json:"session_id"`
	//Action      interface{} `json:"action"`
	//CorpId      string      `json:"corp_id"`
	//UpdatedAt   int    `json:"updated_at"`
}
