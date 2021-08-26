package models

import (
	"encoding/json"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/app"
	"time"
)

type ChatMsgContent struct {
	ExtCorpModel
	ChatMsgID string `json:"chat_msg_id" gorm:"comment:内部消息ID"`
	// 消息类型
	ContentType string `json:"content_type" gorm:"comment:消息类型"`
	// 聊天的非文字内容
	Content string `json:"content" gorm:"comment: 非文字类型的消息内容"`
	// 文件地址
	FileURL string `json:"file_url" gorm:"comment: 文件下载地址"`
	// 文件名
	FileName string `json:"file_name" gorm:"comment: 文件名"`
	Timestamp
}

type ChatMsg struct {
	ExtCorpModel
	//消息id，消息的唯一标识，企业可以使用此字段进行消息去重。String类型
	MsgID string `gorm:"type:char(128);unique;comment:外部消息ID" json:"msgid"`
	//消息动作，目前有send(发送消息)/recall(撤回消息)/switch(切换企业日志)三种类型。String类型
	Action string `gorm:"type:char(8);comment:消息动作，目前有send(发送消息)/recall(撤回消息)/switch(切换企业日志)三种类型" json:"action"`
	//消息发送方id。同一企业内容为userid，非相同企业为external_userid。消息如果是机器人发出，也为external_userid。String类型
	From string `gorm:"type:char(32);comment:消息发送方id。同一企业内容为userid，非相同企业为external_userid。消息如果是机器人发出，也为external_userid" json:"from"`
	//消息接收方列表，可能是多个，同一个企业内容为userid，非相同企业为external_userid。数组，内容为string类型
	ToList constants.StringArrayField `gorm:"type:json;comment:消息接收方列表" json:"tolist"`
	//群聊消息的群id。如果是单聊则为空。String类型
	RoomID string `gorm:"type:char(128);comment:群聊消息的群id。如果是单聊则为空" json:"roomid"`
	//消息发送时间戳，utc时间，ms单位。
	MsgTime int64 `gorm:"type:bigint(64);comment:消息发送时间戳，utc时间，ms单位。" json:"msgtime"`
	//文本消息为：text。String类型
	MsgType string `gorm:"type:varchar(32);comment:文本消息为：text" json:"msgtype"`
	// 聊天的文本内容
	ContentText string `gorm:"class:FULLTEXT,option:WITH PARSER ngram;comment:聊天的文本内容" json:"content_text"`
	//消息的seq值，标识消息的序号。再次拉取需要带上上次回包中最大的seq。Uint64类型，范围0-pow(2,64)-1
	Seq uint64 `gorm:"type:bigint unsigned;comment:消息的seq值，标识消息的序号。再次拉取需要带上上次回包中最大的seq。Uint64类型，范围0-pow(2,64)-1" json:"seq"`
	// 发送-接收双方ID hash得到
	SessionID string `gorm:"type:char(128);comment:消息的会话ID,相同收发方的会话ID相同" json:"session_id"`
	// 会话类型
	SessionType    string         `gorm:"type:char(32);comment:会话类型" json:"session_type"`
	ChatMsgContent ChatMsgContent `gorm:"foreignKey:chat_msg_id;reference:ID" json:"chat_msg_content"`

	// 不用返回没有存在的类型
	Attachments interface{} `json:"attachments" gorm:"-"`
}

//type Text struct {
//	Content string `json:"content"`
//}

type Image struct {
	Md5sum    string `json:"md5sum"`
	Sdkfileid string `json:"sdkfileid"`
	Filesize  uint32 `json:"filesize"`
}

type File struct {
	Md5sum   string `json:"md5sum"`   //资源的md5值，供进行校验。String类型
	Filename string `json:"filename"` //文件名称。String类型
	Fileext  string `json:"fileext"`  //文件类型后缀。String类型
	Filesize uint32 `json:"filesize"` // 文件大小。Uint32类型
}
type Revoke struct {
	PreMsgid string `json:"pre_msgid"` // 标识撤回的原消息的msgid。String类型
}
type Agree struct {
	Userid    string `json:"userid"`     // 同意/不同意协议者的userid，外部企业默认为external_userid。String类型
	AgreeTime string `json:"agree_time"` // 同意/不同意协议的时间，utc时间，ms单位。
}
type Voice struct {
	Md5sum     string `json:"md5sum"`
	VoiceSize  uint32 `json:"voice_size"`
	PlayLength uint32 `json:"play_length"`
	Sdkfileid  string `json:"sdkfileid"`
}
type TODO struct {
}
type Redpacket struct {
	//红包消息类型。1 普通红包、2 拼手气群红包、3 激励群红包。Uint32类型
	Type uint32 `json:"type"`
	//wish	红包祝福语。String类型
	Wish string `json:"wish"`
	//totalcnt	红包总个数。Uint32类型
	Totalcnt uint32 `json:"totalcnt"`
	//totalamount	红包总金额。Uint32类型，单位为分。
	Totalamount uint32 `json:"redpacket"`
}

type Emotion struct {
	Type      int    `json:"type"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Imagesize int    `json:"imagesize"`
	Md5Sum    string `json:"md5sum"`
	Sdkfileid string `json:"sdkfileid"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Address   string  `json:"address"`
	Title     string  `json:"title"`
	Zoom      int     `json:"zoom"`
}

//名片
type Card struct {
	CorpName string `json:"corpname"` //名片所有者所在的公司名称
	UserID   string `json:"userid"`   //名片所有者的id，同一公司是userid，不同公司是external_userid。String类型
}
type DocMsg struct {
	Msgtype    string `json:"msgtype"`     //类型, 标识在线文档消息类型
	Title      string `json:"title"`       //       在线文档名称
	LinkURL    string `json:"link_url"`    //    在线文档链接
	DocCreator string `json:"doc_creator"` // 在线文档创建者。本企业成员创建为userid；外部企业成员创建为external_userid
}
type Meeting struct {
	MsgID   string `json:"msgid"`   //	msgid 消息id，消息的唯一标识，企业可以使用此字段进行消息去重。String类型
	Action  string `json:"action"`  //	action 消息动作，目前有send(发送消息)/recall(撤回消息)/switch (切换企业日志)三种类型。String类型
	From    string `json:"from"`    //	from   消息发送方id。同一企业内容为userid，非相同企业为external_userid。消息如果是机器人发出，也为external_userid。String类型
	Tolist  string `json:"tolist"`  //  消息接收方列表，可能是多个，同一个企业内容为userid，非相同企业为external_userid。数组，内容为string类型
	Msgtime string `json:"msgtime"` //  消息发送时间戳，utc时间, 单位毫秒。
	Msgtype string `json:"msgtype"` //  meeting_voice_call。String类型, 标识音频存档消息类型
	Voiceid string `json:"voiceid"` // String类型, 音频id
	//Meeting_voice_call
	// 音频消息内容。包括结束时间、fileid，可能包括多个demofiledata、sharescreendata消息，demofiledata表示文档共享信息，sharescreendata表示屏幕共享信息。Object类型
	//Endtime    音频结束时间。uint32类型
	//Sdkfileid    sdkfileid。音频媒体下载的id。String类型
	//Demofiledata    文档分享对象，Object类型
	//Filename    文档共享名称。String类型
	//Demooperator    文档共享操作用户的id。String类型
	//Starttime    文档共享开始时间。Uint32类型
	//Endtime    文档共享结束时间。Uint32类型
	//Sharescreendata    屏幕共享对象，Object类型
	//Share    屏幕共享用户的id。String类型
	//Starttime    屏幕共享开始时间。Uint32类型
	//Endtime    屏幕共享结束时间。Uint32类型
}
type Link struct {
	Title       string `json:"title"`       // 消息标题。String类型
	Description string `json:"description"` // 消息描述。String类型
	LinkUrl     string `json:"link_url"`    // 链接url地址。String类型
	ImageUrl    string `json:"image_url"`   //  链接图片url。String类型
}

// ChatMessage 会话中的消息列表条目
type ChatMessage struct {
	ChatMsg
	// 发送方名字
	SenderName string `json:"sender_name"`
	// 发送方头像
	SenderAvatar string `json:"sender_avatar"`
}

// ChatSessions
// 会话列表
type ChatSessions struct {
	ExtCorpModel
	//消息id，消息的唯一标识，企业可以使用此字段进行消息去重。String类型
	MsgID string `gorm:"type:char(128);unique" json:"msgid"`
	//消息动作，目前有send(发送消息)/recall(撤回消息)/switch(切换企业日志)三种类型。String类型
	Action string `gorm:"type:char(8)" json:"action"`
	//消息发送方id。同一企业内容为userid，非相同企业为external_userid。消息如果是机器人发出，也为external_userid。String类型
	From string `gorm:"type:char(32)" json:"from"`
	//消息接收方列表，可能是多个，同一个企业内容为userid，非相同企业为external_userid。数组，内容为string类型
	ToList constants.StringArrayField `gorm:"type:json" json:"tolist"`
	//群聊消息的群id。如果是单聊则为空。String类型
	RoomID string `gorm:"type:char(128)" json:"roomid"`
	//消息发送时间戳，utc时间，ms单位。
	MsgTime int64 `gorm:"type:bigint(64)" json:"msgtime"`
	//文本消息为：text。String类型
	MsgType string `gorm:"type:varchar(32)" json:"msgtype"`
	// 聊天的文本内容
	ContentText string `gorm:"class:FULLTEXT,option:WITH PARSER ngram" json:"content_text"`
	//消息的seq值，标识消息的序号。再次拉取需要带上上次回包中最大的seq。Uint64类型，范围0-pow(2,64)-1
	Seq uint64 `gorm:"type:bigint unsigned" json:"seq"`
	// 发送-接收双方ID hash得到
	SessionID string `json:"session_id"`
	// 会话类型
	SessionType string `json:"session_type"`
	// 会话对方头像
	PeerAvatar string `json:"peer_avatar"`
	// 会话对方头像
	PeerName string `json:"peer_name"`
	// 会话对方外部ID
	PeerExtID string `json:"peer_ext_id"`
	// 群名
	GroupChatName string `json:"group_chat_name"`
}

func (o ChatMsg) QueryMsg(
	msg ChatMsg, sendAtStart, sendAtEnd *time.Time, sorter *app.Sorter, pager *app.Pager) ([]ChatMessage, int64, error) {

	db := DB.Table("chat_msg").
		Joins("left join customer c on c.ext_id = chat_msg.ext_creator_id").
		Joins("left join staff s on s.ext_id = chat_msg.ext_creator_id")

	if len(msg.ToList) > 0 && msg.From != "" {
		db = db.Where("json_contains(to_list, json_array(?)) and `from` = ? ", msg.ToList[0], msg.From).
			Or("json_contains(to_list, json_array(?)) and `from` = ? ", msg.From, msg.ToList[0])
	}
	if msg.MsgType != "" {
		db = db.Where("msg_type = ?", msg.MsgType)
	}
	if !sendAtStart.IsZero() {
		db = db.Where("msg_time between ? and ?", sendAtStart.Unix()*1000, sendAtEnd.Unix()*1000)
	}
	if msg.ContentText != "" {
		db = db.Where("match (content_text) against (?)", msg.ContentText)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count chat msg failed")
		return nil, 0, err
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	var msgs []ChatMessage
	// ChatMessage 会话中的消息列表条目
	err = db.Preload("ChatMsgContent").Select("chat_msg.*, " +
		" IF(c.avatar is not null, c.avatar, s.avatar_url) as sender_avatar," +
		" IF(c.name is not null, c.name, s.name)           as sender_name").
		Find(&msgs).Error
	if err != nil {
		err = errors.Wrap(err, "Find chat msg failed")
		return nil, 0, err
	}

	for i, message := range msgs {
		content, err := message.unmarshalContent()
		if err != nil {
			return nil, 0, err
		}
		msgs[i].Attachments = content
	}
	return msgs, total, nil
}

func (o ChatMessage) unmarshalContent() (res interface{}, err error) {
	switch o.MsgType {
	case constants.MsgTypeTextMsg:
		return nil, nil
	case constants.MsgTypeRevoke:
		var revoke Revoke
		err = json.Unmarshal([]byte(o.ChatMsgContent.Content), &revoke)
		if err != nil {
			return
		}
		return revoke, nil
	case constants.MsgTypeVoice:
		var voice Voice
		err = json.Unmarshal([]byte(o.ChatMsgContent.Content), &voice)
		if err != nil {
			return
		}
		return voice, nil
	case constants.MsgTypeCard:
		var card Card
		err = json.Unmarshal([]byte(o.ChatMsgContent.Content), &card)
		if err != nil {
			return
		}
		return card, nil
	case constants.MsgTypeLocation:
		var location Location
		err = json.Unmarshal([]byte(o.ChatMsgContent.Content), &location)
		if err != nil {
			return
		}
		return location, nil
	case constants.MsgTypeEmotion:
		var emotion Emotion
		err = json.Unmarshal([]byte(o.ChatMsgContent.Content), &emotion)
		if err != nil {
			return
		}
		return emotion, nil
	case constants.MsgTypeFile:
		var file File
		err = json.Unmarshal([]byte(o.ChatMsgContent.Content), &file)
		if err != nil {
			return
		}
		return file, nil

	case constants.MsgTypeLink:
	case constants.MsgTypeChatRecord:
	case constants.MsgTypeWeApp:
	case constants.MsgTypeDocMsg:
	case constants.MsgTypeTodo:
	case constants.MsgTypeMarkdown:
	case constants.MsgTypeMeeting:
	case constants.MsgTypeCollect:
	case constants.MsgTypeVote:
	case constants.MsgTypeNews:
	case constants.MsgTypeCalendar:
	case constants.MsgTypeImage:
		var image Image
		err = json.Unmarshal([]byte(o.ChatMsgContent.Content), &image)
		if err != nil {
			return
		}
		return image, nil

	default:
		err = errors.New("unknown type")
	}
	return
}

func (o ChatMsg) Creat(msgs []ChatMsg) error {
	return DB.Model(&ChatMsg{}).Preload("ChatMsgContent").CreateInBatches(&msgs, 10).Error
}

func (o ChatMsg) GetLatestSeq(extCorpID string) (maxSeq int64, err error) {
	var chatMsg ChatMsg
	err = DB.Model(&ChatMsg{}).Where("ext_corp_id = ?", extCorpID).
		Order("seq desc").Limit(1).First(&chatMsg).Error
	if err == gorm.ErrRecordNotFound {
		maxSeq = 0
		return maxSeq, nil
	}
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return int64(chatMsg.Seq), nil
}

func (o ChatMsg) QuerySessions(
	extStaffID string, sessionType string, extCorpID string, sorter *app.Sorter, pager *app.Pager) (chatSessions []ChatSessions, total int64, err error) {

	countSessionsSQL := "select count(1) as total from (" +
		" select *, rank() over (partition by session_id order by id desc) as r" +
		" from chat_msg where ext_corp_id = ? and (`from` = ? or json_contains(to_list, json_array(?)))" +
		" and session_type = ?  " +
		" ) t where t.r = 1"

	err = DB.Raw(countSessionsSQL, extCorpID, extStaffID, extStaffID, sessionType).Take(&total).Error
	if err != nil {
		err = errors.Wrap(err, "Count Sessions failed")
		return
	}

	querySessionsSQL := "select * from (" +
		" select *, rank() over (partition by session_id order by id desc) as r" +
		" from chat_msg where ext_corp_id = ? and (`from` = ? or json_contains(to_list, json_array(?)))" +
		" and session_type = ?" +
		" ) t where t.r = 1 order by id desc limit ?, ?"

	pager.SetDefault()
	err = DB.Raw(querySessionsSQL, extCorpID, extStaffID, extStaffID, sessionType, pager.GetOffset(), pager.GetLimit()).Find(&chatSessions).Error
	if err != nil {
		err = errors.Wrap(err, "query Sessions failed")
		return
	}

	// 除了自己外的ExtIDs
	extIDs := make([]string, 0)
	for _, item := range chatSessions {
		if item.From != extStaffID {
			extIDs = append(extIDs, item.From)
		}
		if item.ToList[0] != extStaffID {
			extIDs = append(extIDs, item.ToList[0])
		}
	}
	// 取这些IDs的基本信息
	type Avatars struct {
		ExtID  string `json:"ext_id"`
		Avatar string `json:"avatar"`
		Name   string `json:"name"`
	}
	var avatars []Avatars

	switch constants.ChatSessionType(sessionType) {
	case constants.ChatSessionTypeExternal:
		err = DB.Model(&Customer{}).Where("ext_id in (?)", extIDs).Find(&avatars).Error
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		avatarMap := make(map[string]Avatars, 0)
		for _, avatar := range avatars {
			avatarMap[avatar.ExtID] = avatar
		}

		for i, item := range chatSessions {
			// 发送方为员工，则取接收方的头像，反之则取发送方的头像
			if item.From == extStaffID {
				chatSessions[i].PeerName = avatarMap[item.ToList[0]].Name
				chatSessions[i].PeerAvatar = avatarMap[item.ToList[0]].Avatar
				chatSessions[i].PeerExtID = avatarMap[item.ToList[0]].ExtID
			} else if item.ToList[0] == extStaffID {
				chatSessions[i].PeerName = avatarMap[item.From].Name
				chatSessions[i].PeerAvatar = avatarMap[item.From].Avatar
				chatSessions[i].PeerExtID = avatarMap[item.From].ExtID
			}
		}
	case constants.ChatSessionTypeInternal:
		err = DB.Model(&Staff{}).Where("ext_id in (?)", extIDs).Find(&avatars).Error
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		avatarMap := make(map[string]Avatars, 0)
		for _, avatar := range avatars {
			avatarMap[avatar.ExtID] = avatar
		}

		for i, item := range chatSessions {
			// 发送方为员工，则取接收方的头像，反之则取发送方的头像
			if item.From == extStaffID {
				chatSessions[i].PeerName = avatarMap[item.ToList[0]].Name
				chatSessions[i].PeerAvatar = avatarMap[item.ToList[0]].Avatar
				chatSessions[i].PeerExtID = avatarMap[item.ToList[0]].ExtID
			} else if item.ToList[0] == extStaffID {
				chatSessions[i].PeerName = avatarMap[item.From].Name
				chatSessions[i].PeerAvatar = avatarMap[item.From].Avatar
				chatSessions[i].PeerExtID = avatarMap[item.From].ExtID
			}
		}
	case constants.ChatSessionTypeGroup:
		// todo 先来条数据测测
	default:
	}
	return
}
