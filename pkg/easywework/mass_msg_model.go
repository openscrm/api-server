package workwx

// AddMsgTemplateReq 创建企业群发
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/92135#创建企业群发
type AddMsgTemplateReq struct {
	ChatType       string        `json:"chat_type"`
	ExternalUserid []string      `json:"external_userid"`
	Sender         string        `json:"sender"`
	Text           Text          `json:"text"`
	Attachments    []Attachments `json:"attachments,omitempty"`
}

// addMsgTemplateResp 创建企业群发
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/92135#创建企业群发
type addMsgTemplateResp struct {
	CommonResp
	// FailList 无效或无法发送的external_userid列表
	FailList []string `json:"fail_list"`
	// MsgID 企业群发消息的id，可用于获取企业群发成员执行结果
	MsgID string `json:"msgid"`
}

// GetGroupMsgSendResultExternalContactReq 获取企业群发成员执行结果请求
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/93338#获取企业群发成员执行结果
type GetGroupMsgSendResultExternalContactReq struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 返回的最大记录数，整型，最大值1000，默认值500，超过最大值时取默认值
	Limit int `json:"limit,omitempty"`
	// Msgid 群发消息的id，通过<a href="#获取群发记录列表">获取群发记录列表</a>接口返回，必填
	Msgid  string `json:"msgid"`
	Userid string `json:"userid"`
}

// GetGroupMsgSendResultExternalContactResp 获取企业群发成员执行结果响应
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/93338#获取企业群发成员执行结果
type GetGroupMsgSendResultExternalContactResp struct {
	CommonResp
	NextCursor string `json:"next_cursor"`
	SendList   []struct {
		ChatID         string `json:"chat_id"`
		ExternalUserid string `json:"external_userid"`
		SendTime       int    `json:"send_time"`
		Status         int    `json:"status"`
		Userid         string `json:"userid"`
	} `json:"send_list"`
}

// reqGetGroupmsgTaskExternalcontact 获取群发成员发送任务列表请求
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/93338#获取群发成员发送任务列表
type reqGetGroupmsgTaskExternalcontact struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 返回的最大记录数，整型，最大值1000，默认值500，超过最大值时取默认值
	Limit int `json:"limit,omitempty"`
	// Msgid 群发消息的id，通过<a href="#获取群发记录列表">获取群发记录列表</a>接口返回，必填
	Msgid string `json:"msgid"`
}

// getGroupMsgTaskExternalContactResp 获取群发成员发送任务列表响应
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/93338#获取群发成员发送任务列表
type getGroupMsgTaskExternalContactResp struct {
	CommonResp
	NextCursor string `json:"next_cursor"`
	TaskList   []struct {
		SendTime int    `json:"send_time"`
		Status   int    `json:"status"`
		Userid   string `json:"userid"`
	} `json:"task_list"`
}

// getGroupMsgListV2ExternalContactReq 获取群发记录列表请求
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/93338#获取群发记录列表
type getGroupMsgListV2ExternalContactReq struct {
	// ChatType 群发任务的类型，默认为single，表示发送给客户，group表示发送给客户群，必填
	ChatType string `json:"chat_type"`
	// Creator 群发任务创建人企业账号id
	Creator string `json:"creator,omitempty"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// EndTime 群发任务记录结束时间，必填
	EndTime int `json:"end_time"`
	// FilterType 创建人类型。0:企业发表 1:个人发表 2:所有，包括个人创建以及企业创建，默认情况下为所有类型
	FilterType int `json:"filter_type,omitempty"`
	// Limit 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取默认值
	Limit int `json:"limit,omitempty"`
	// StartTime 群发任务记录开始时间，必填
	StartTime int `json:"start_time"`
}

// getGroupMsgListV2ExternalContactResp 获取群发记录列表响应
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/93338#获取群发记录列表
type getGroupMsgListV2ExternalContactResp struct {
	CommonResp
	GroupMsgList []struct {
		Attachments []struct {
			Image struct {
				MediaID string `json:"media_id"`
				PicURL  string `json:"pic_url"`
			} `json:"image"`
			Link struct {
				Desc   string `json:"desc"`
				Picurl string `json:"picurl"`
				Title  string `json:"title"`
				URL    string `json:"url"`
			} `json:"link"`
			Miniprogram struct {
				Appid      string `json:"appid"`
				Page       string `json:"page"`
				PicMediaID string `json:"pic_media_id"`
				Title      string `json:"title"`
			} `json:"miniprogram"`
			Msgtype string `json:"msgtype"`
			Video   struct {
				MediaID string `json:"media_id"`
			} `json:"video"`
		} `json:"attachments"`
		CreateTime string `json:"create_time"`
		CreateType int    `json:"create_type"`
		Creator    string `json:"creator"`
		Msgid      string `json:"msgid"`
		Text       struct {
			Content string `json:"content"`
		} `json:"text"`
	} `json:"group_msg_list"`
	NextCursor string `json:"next_cursor"`
}

// GetGroupMsgTaskExternalContactResp 获取群发成员发送任务列表响应
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/93338#获取群发成员发送任务列表
type GetGroupMsgTaskExternalContactResp struct {
	CommonResp
	NextCursor string `json:"next_cursor"`
	TaskList   []struct {
		SendTime int    `json:"send_time"`
		Status   int    `json:"status"`
		Userid   string `json:"userid"`
	} `json:"task_list"`
}
