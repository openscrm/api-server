package workwx

// rxMessageCommon 接收消息的公共部分
type rxMessageCommon struct {
	// ToUserName 企业微信CorpID
	ToUserName string `xml:"ToUserName"`
	// FromUserName 成员UserID
	FromUserName string `xml:"FromUserName"`
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// MsgType 消息类型
	MsgType MessageType `xml:"MsgType"`
	// MsgID 消息id，64位整型
	MsgID int64 `xml:"MsgId"`
	// AgentID 企业应用的id，整型。可在应用的设置页面查看
	AgentID int64 `xml:"AgentID"`
	// Event 事件类型 MsgType为event存在
	Event EventType `xml:"Event"`
	// ChangeType 变更类型 Event为change_external_contact存在
	ChangeType ChangeType `xml:"ChangeType"`
}

// MessageType 消息类型
type MessageType string

// MessageTypeText 文本消息
const MessageTypeText MessageType = "text"

// MessageTypeImage 图片消息
const MessageTypeImage MessageType = "image"

// MessageTypeVoice 语音消息
const MessageTypeVoice MessageType = "voice"

// MessageTypeVideo 视频消息
const MessageTypeVideo MessageType = "video"

// MessageTypeLocation 位置消息
const MessageTypeLocation MessageType = "location"

// MessageTypeLink 链接消息
const MessageTypeLink MessageType = "link"

// MessageTypeEvent 事件消息
const MessageTypeEvent MessageType = "event"

// EventType 事件类型
type EventType string

// EventTypeChangeExternalContact 企业客户事件
const EventTypeChangeExternalContact EventType = "change_external_contact"

// EventTypeChangeExternalTag 标签事件
const EventTypeChangeExternalTag EventType = "change_external_tag"

// EventTypeChangeContact 通讯论变更事件
const EventTypeChangeContact EventType = "change_contact"

// EventTypeChangeExternalChat 客户群变更事件
const EventTypeChangeExternalChat EventType = "change_external_chat"

// EventTypeSysApprovalChange 审批申请状态变化回调通知
const EventTypeSysApprovalChange EventType = "sys_approval_change"

// ChangeType 变更类型
type ChangeType string

// ChangeTypeAddExternalContact 添加企业客户事件
const ChangeTypeAddExternalContact ChangeType = "add_external_contact"

// ChangeTypeCreateUser 新增员工
const ChangeTypeCreateUser ChangeType = "create_user"

// ChangeTypeDelUser 删除员工事件
const ChangeTypeDelUser ChangeType = "delete_user"

// ChangeTypeUpdateUser 更新员工事件
const ChangeTypeUpdateUser ChangeType = "update_user"

// ChangeTypeEditExternalContact 编辑企业客户事件
const ChangeTypeEditExternalContact ChangeType = "edit_external_contact"

// ChangeTypeAddHalfExternalContact 外部联系人免验证添加成员事件
const ChangeTypeAddHalfExternalContact ChangeType = "add_half_external_contact"

// ChangeTypeDelExternalContact 删除企业客户事件
const ChangeTypeDelExternalContact ChangeType = "del_external_contact"

// ChangeTypeDelFollowUser 删除跟进成员事件
const ChangeTypeDelFollowUser ChangeType = "del_follow_user"

// ChangeTypeTransferFail 客户接替失败事件
const ChangeTypeTransferFail ChangeType = "transfer_fail"

// ChangeTypeCreateParty 添加部门事件
const ChangeTypeCreateParty ChangeType = "create_party"

// ChangeTypeUpdateParty 更新部门事件
const ChangeTypeUpdateParty ChangeType = "update_party"

// ChangeTypeDeleteParty 删除部门事件
const ChangeTypeDeleteParty ChangeType = "delete_party"

// ChangeTypeCreateTag 添加标签事件
const ChangeTypeCreateTag ChangeType = "create"

// ChangeTypeUpdateTag 更新标签事件
const ChangeTypeUpdateTag ChangeType = "update"

// ChangeTypeDeleteTag 删除标签事件
const ChangeTypeDeleteTag ChangeType = "delete"

// ChangeTypeCreateChat 创建群聊
const ChangeTypeCreateChat ChangeType = "create"

// ChangeTypeUpdateChat 群聊更新事件
const ChangeTypeUpdateChat ChangeType = "update"

// ChangeTypeDismissChat 客户群解散事件
const ChangeTypeDismissChat ChangeType = "dismiss"

//ChangeTypeMsgAuditApproved 添加外部联系人同意进行聊天内容存档时，回调该事件。
const ChangeTypeMsgAuditApproved = "msg_audit_approved"

// rxTextMessageSpecifics 接收的文本消息，特有字段
type rxTextMessageSpecifics struct {
	// Content 文本消息内容
	Content string `xml:"Content"`
}

// rxImageMessageSpecifics 接收的图片消息，特有字段
type rxImageMessageSpecifics struct {
	// PicURL 图片链接
	PicURL string `xml:"PicUrl"`
	// MediaID 图片媒体文件id，可以调用获取媒体文件接口拉取，仅三天内有效
	MediaID string `xml:"MediaId"`
}

// rxVoiceMessageSpecifics 接收的语音消息，特有字段
type rxVoiceMessageSpecifics struct {
	// MediaID 语音媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效
	MediaID string `xml:"MediaId"`
	// Format 语音格式，如amr，speex等
	Format string `xml:"Format"`
}

// rxVideoMessageSpecifics 接收的视频消息，特有字段
type rxVideoMessageSpecifics struct {
	// MediaID 视频媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效
	MediaID string `xml:"MediaId"`
	// ThumbMediaID 视频消息缩略图的媒体id，可以调用获取媒体文件接口拉取数据，仅三天内有效
	ThumbMediaID string `xml:"ThumbMediaId"`
}

// rxLocationMessageSpecifics 接收的位置消息，特有字段
type rxLocationMessageSpecifics struct {
	// Lat 地理位置纬度
	Lat float64 `xml:"Location_X"`
	// Lon 地理位置经度
	Lon float64 `xml:"Location_Y"`
	// Scale 地图缩放大小
	Scale int `xml:"Scale"`
	// Label 地理位置信息
	Label string `xml:"Label"`
	// AppType app类型，在企业微信固定返回wxwork，在微信不返回该字段
	AppType string `xml:"AppType"`
}

// rxLinkMessageSpecifics 接收的链接消息，特有字段
type rxLinkMessageSpecifics struct {
	// Title 标题
	Title string `xml:"Title"`
	// Description 描述
	Description string `xml:"Description"`
	// URL 链接跳转的url
	URL string `xml:"Url"`
	// PicURL 封面缩略图的url
	PicURL string `xml:"PicUrl"`
}

// rxEventAddExternalContact 接收的事件消息，添加企业客户事件
type rxEventAddExternalContact struct {
	// UserID 企业服务人员的UserID
	UserID string `xml:"UserID"`
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `xml:"ExternalUserID"`
	// State 添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道
	State string `xml:"State"`
	// WelcomeCode 欢迎语code，可用于发送欢迎语
	WelcomeCode string `xml:"WelcomeCode"`
}

// rxEventEditExternalContact 接收的事件消息，编辑企业客户事件
type rxEventEditExternalContact struct {
	// UserID 企业服务人员的UserID
	UserID string `xml:"UserID"`
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `xml:"ExternalUserID"`
	// State 添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道
	State string `xml:"State"`
}

// rxEventAddHalfExternalContact 接收的事件消息，外部联系人免验证添加成员事件
type rxEventAddHalfExternalContact struct {
	// UserID 企业服务人员的UserID
	UserID string `xml:"UserID"`
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `xml:"ExternalUserID"`
	// State 添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道
	State string `xml:"State"`
	// WelcomeCode 欢迎语code，可用于发送欢迎语
	WelcomeCode string `xml:"WelcomeCode"`
}

// rxEventDelExternalContact 接收的事件消息，删除企业客户事件
type rxEventDelExternalContact struct {
	// UserID 企业服务人员的UserID
	UserID string `xml:"UserID"`
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `xml:"ExternalUserID"`
}

// rxEventDelFollowUser 接收的事件消息，删除跟进成员事件
type rxEventDelFollowUser struct {
	// UserID 企业服务人员的UserID
	UserID string `xml:"UserID"`
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `xml:"ExternalUserID"`
}

// rxEventTransferFail 接收的事件消息，客户接替失败事件
type rxEventTransferFail struct {
	// FailReason 接替失败的原因, customer_refused-客户拒绝， customer_limit_exceed-接替成员的客户数达到上限
	FailReason string `xml:"FailReason"`
	// UserID 企业服务人员的UserID
	UserID string `xml:"UserID"`
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `xml:"ExternalUserID"`
}

// rxEventChangeExternalChat 接收的事件消息，客户群变更事件
type rxEventChangeExternalChat struct {
	// ToUserName 企业微信CorpID
	ToUserName string `xml:"ToUserName"`
	// FromUserName 此事件该值固定为sys，表示该消息由系统生成
	FromUserName string `xml:"FromUserName"`
	// FailReason 接替失败的原因, customer_refused-客户拒绝， customer_limit_exceed-接替成员的客户数达到上限
	FailReason string `xml:"FailReason"`
	// ChatID 群ID
	ChatID string `xml:"ChatId"`

	//UpdateDetail
	//变更详情。目前有以下几种： add_member : 成员入群 del_member : 成员退群 change_owner : 群主变更 change_name : 群名变更 change_notice : 群公告变更
	UpdateDetail string `xml:"UpdateDetail"`

	// JoinScene
	//当是成员入群时有值。表示成员的入群方式 0 - 由成员邀请入群（包括直接邀请入群和通过邀请链接入群） 3 - 通过扫描群二维码入群
	JoinScene int64 `xml:"JoinScene"`

	// QuitScene 当是成员退群时有值。表示成员的退群方式 0 - 自己退群 1 - 群主/群管理员移出
	QuitScene int64 `xml:"QuitScene"`

	//MemChangeCnt	当是成员入群或退群时有值。表示成员变更数量
	MemChangeCnt int64 `xml:MemChangeCnt"`
}

func (r *rxEventChangeExternalChat) GetJoinScene() int64 {
	return r.JoinScene
}

func (r *rxEventChangeExternalChat) GetQuitScene() int64 {
	return r.QuitScene
}

func (r *rxEventChangeExternalChat) GetMemChangeCnt() int64 {
	return r.MemChangeCnt
}

func (r *rxEventChangeExternalChat) GetUpdateDetail() string {
	return r.UpdateDetail
}

// rxEventSysApprovalChange 接收的事件消息，审批申请状态变化回调通知
type rxEventSysApprovalChange struct {
	// ApprovalInfo 审批信息、
	ApprovalInfo OAApprovalInfo `xml:"ApprovalInfo"`
}

// rxEventCreateParty 接收的事件消息，添加部门
type rxEventCreateParty struct {
	ID       int64  `xml:"Id"`       //	部门Id
	Name     string `xml:"Name"`     // 部门名称
	ParentId int64  `xml:"ParentId"` // 父部门id
	Order    int64  `xml:"Order"`    //    部门排序
}

// rxEventUpdateParty 接收的事件消息，更新部门
type rxEventUpdateParty struct {
	ID       int64  `xml:"Id"`       //	部门Id
	Name     string `xml:"Name"`     // 部门名称
	ParentId int64  `xml:"ParentId"` // 父部门id
}

// rxEventDeleteParty 接收的事件消息，删除部门
type rxEventDeleteParty struct {
	ID int64 `xml:"Id"` //	部门Id
}

// rxEventCreateTag 接收的事件消息，新建标签
type rxEventCreateTag struct {
	ID      string `xml:"Id"`      // 标签或标签组的ID
	TagType string `xml:"TagType"` // 删除标签时，此项为tag，删除标签组时，此项为tag_group
}

// rxEventUpdateTag 接收的事件消息，更新标签
type rxEventUpdateTag struct {
	ID      string `xml:"Id"`      // 标签或标签组的ID
	TagType string `xml:"TagType"` // 删除标签时，此项为tag，删除标签组时，此项为tag_group
}

// rxEventDeleteTag 接收的事件消息，删除标签
type rxEventDeleteTag struct {
	ID      string `xml:"Id"`      // 标签或标签组的ID
	TagType string `xml:"TagType"` // 删除标签时，此项为tag，删除标签组时，此项为tag_group
}
type rxEventUpdateUser struct {
	UserID string `xml:"UserID"`
}

type rxEventDeleteUser struct {
	UserID string `xml:"UserID"`
}

type rxEventCreateUser struct {
	//XMLName        xml.Name `xml:"xml"`
	//Text           string   `xml:",chardata"`
	//ToUserName     string   `xml:"ToUserName"`
	//FromUserName   string   `xml:"FromUserName"`
	//CreateTime     string   `xml:"CreateTime"`
	//MsgType        string   `xml:"MsgType"`
	//Event          string   `xml:"Event"`
	//ChangeType     string   `xml:"ChangeType"`
	UserID string `xml:"UserID"`
	//Name           string   `xml:"Name"`
	//Department     string   `xml:"Department"`
	//MainDepartment string   `xml:"MainDepartment"`
	//IsLeaderInDept string   `xml:"IsLeaderInDept"`
	//Position       string   `xml:"Position"`
	//Mobile         string   `xml:"Mobile"`
	//Gender         string   `xml:"Gender"`
	//Email          string   `xml:"Email"`
	//Status         string   `xml:"Status"`
	//Avatar         string   `xml:"Avatar"`
	//Alias          string   `xml:"Alias"`
	//Telephone      string   `xml:"Telephone"`
	//Address        string   `xml:"Address"`
	//ExtAttr        struct {
	//	Text string `xml:",chardata"`
	//	Item []struct {
	//		Chardata string `xml:",chardata"`
	//		Name     string `xml:"Name"`
	//		Type     string `xml:"Type"`
	//		Text     struct {
	//			Text  string `xml:",chardata"`
	//			Value string `xml:"Value"`
	//		} `xml:"Text"`
	//		Web struct {
	//			Text  string `xml:",chardata"`
	//			Title string `xml:"Title"`
	//			URL   string `xml:"Url"`
	//		} `xml:"Web"`
	//	} `xml:"Item"`
	//} `xml:"ExtAttr"`
}
