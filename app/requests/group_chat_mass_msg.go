package requests

import "openscrm/app/constants"

type SendGroupChatMassMsgReq struct {
	// ExtStaffIDs为群主IDs
	ExtStaffIDs constants.StringArrayField `json:"ext_staff_ids" form:"ext_staff_ids" validate:"omitempty"`
	// 1-立即发送，2-定时发送
	SendType constants.SendMassMsgType `json:"send_type" validate:"required,oneof=1 2"`
	// 定时发送时间戳
	SendAt constants.DateTimeFiled `json:"send_at" validate:"omitempty,gt=0"`
	// 群发任务的类型，默认为single，表示发送给客户，group表示发送给客户群
	//ChatType constants.ChatType `json:"chat_type" validate:"omitempty,oneof=single group"`
	// 消息体
	Msg constants.AutoReplyField `json:"msg" validate:"omitempty,gte=0"`
}
