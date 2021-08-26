package constants

import (
	"database/sql/driver"
	"encoding/json"
)

type QuickReplyType int64

const (
	QuickReplyTypeCollection QuickReplyType = 1
	QuickReplyTypeText       QuickReplyType = 2
	QuickReplyTypePic        QuickReplyType = 3
	QuickReplyTypeNews       QuickReplyType = 4 // website
	QuickReplyTypePDF        QuickReplyType = 5
	QuickReplyTypeVideo      QuickReplyType = 6
)

const (
	QuickReplyModuleName = "quick-reply"
	MassMsgModuleName    = "mass-msg"
	WelcomeMsgModuleName = "welcome-msg"
)

func (o QuickReplyField) Value() (driver.Value, error) {
	b, err := json.Marshal(o)
	return string(b), err
}

func (o *QuickReplyField) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), o)
}

func (o QuickReplyField) GormDataType() string {
	return "json"
}

// QuickReplyField 欢迎语/群发消息的附件
type QuickReplyField struct {
	MsgType string `json:"msg_type"`
	// 图片
	Image Img `json:"image" validate:"omitempty,required_if=FieldType image"`
	// 链接
	Link Link `json:"link" validate:"omitempty,required_if=FieldType link"`
	// 视频
	Video Vid `json:"video" validate:"omitempty,required_if=FieldType video"`
	// PDF
	Pdf PDF `json:"pdf"  validate:"omitempty,required_if=FieldType video"`
	// text
	Text Text `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

// Img 图片
type Img struct {
	Title string `json:"title" validate:"omitempty"`
	Size  string `json:"size"`
	// 用获取到的signd URL
	PicUrl  string `json:"picurl" validate:"omitempty"`
	MediaID string `json:"media_id"`
}

// Vid  视频
type Vid struct {
	Title string `json:"title"`
	Size  string `json:"size"`
	// 用获取到的signd URL
	PicUrl  string `json:"picurl"`
	MediaID string `json:"media_id"`
}

type PDF struct {
	Title   string `json:"title"`
	Size    string `json:"size"`
	FileURL string `json:"fileurl"`
	MediaID string `json:"media_id"`
}
