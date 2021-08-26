package responses

import (
	"openscrm/app/constants"
	"openscrm/app/models"
	"time"
)

// ChatSessionItem 员工的会话列表条目
type ChatSessionItem struct {
	Id            int               `json:"id"`
	ExtId         string            `json:"ext_id"`
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
	//LastStartTime time.Time `json:"last_start_time"`
	//LastEndTime time.Time `json:"last_end_time"`
	// 最新对话内容
	LastMsgThumb string `json:"last_msg_thumb"`
	// 最新对话时间
	LastMsgTime time.Time `json:"last_msg_time"`
	Name        string    `json:"name"`
	// 群聊人员ID列表，extStaffID/extCustomerID
	RoomUserIDs string `json:"room_user_ids"`
	// external/inner/group
	SessionType string `json:"session_type"`
	// 内部会话时存在
	StaffExtId string `json:"staff_ext_id"`
	// 员工的企业信息 saas 可用
	//CorpFullName    string                 `json:"corp_full_name"`
	//CorpId          string                 `json:"corp_id"`
	//CorpName        string                 `json:"corp_name"`
	//ExternalProfile models.ExternalProfile `json:"external_profile"`
	// todo
	//Position string `json:"position"`
	//Type    interface{} `json:"type"`
	//Unionid interface{} `json:"unionid"`
}

// ChatMessage 会话中的消息列表条目
type ChatMessage struct {
	models.ChatMsg
	ExtStaffAvatar string `json:"ext_staff_avatar"`
	ToUserAvatar   string `json:"to_user_avatar"`
	ToUserName     string `json:"to_user_name"`
	// external/inner/group
	SessionType string `json:"session_type"`
}

type InnerMsgArchServSessionsResp struct {
	Items []ChatSessionItem
	Total int64 `json:"total"`
}

type InnerMsgArchSerMsgResp struct {
	Items []models.ChatMessage `json:"items"`
	Total int64                `json:"total"`
}
