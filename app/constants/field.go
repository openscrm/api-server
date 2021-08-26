package constants

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/thoas/go-funk"
	gowx "openscrm/pkg/easywework"
	"strings"
	"time"
)

type JSONArrayField struct {
	V []string
}

func (f JSONArrayField) Value() (driver.Value, error) {
	b, err := json.Marshal(f.V)
	return string(b), err
}

func (f *JSONArrayField) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), &f.V)
}

// StringArrayField
// 使gorm支持[]string结构
type StringArrayField []string

func (o StringArrayField) Contains(item string) bool {
	return funk.ContainsString(o, item)
}

// Match 使用数组元素作为关键词去匹配传入的字符串
func (o StringArrayField) Match(paragraph string) bool {
	for _, keyword := range o {
		if strings.Contains(paragraph, keyword) {
			return true
		}
	}
	return false
}

func (o StringArrayField) Value() (driver.Value, error) {
	b, err := json.Marshal(o)
	return string(b), err
}

func (o *StringArrayField) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), o)
}

func (o StringArrayField) GormDataType() string {
	return "json"
}

func (o StringArrayField) ToStringArray() []string {
	r := make([]string, 0)
	for _, s := range o {
		r = append(r, s)
	}
	return r
}

// Int64ArrayField
// 使gorm支持[]int64结构
type Int64ArrayField []int64

func (o Int64ArrayField) ToInt64Array() []int64 {
	r := make([]int64, 0)
	for _, s := range o {
		r = append(r, s)
	}
	return r
}

func (o Int64ArrayField) Value() (driver.Value, error) {
	b, err := json.Marshal(o)
	return string(b), err
}

func (o *Int64ArrayField) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), o)
}

func (o Int64ArrayField) GormDataType() string {
	return "json"
}

// AttachmentField 欢迎语/群发消息的附件
type AttachmentField struct {
	// image、link、miniprogram或者video
	Msgtype AttachmentType `json:"msgtype" validate:"omitempty,oneof=image link miniprogram video"`
	// 图片
	Image Image `json:"image" validate:"omitempty,required_if=Msgtype image"`
	// 链接
	Link Link `json:"link" validate:"omitempty,required_if=Msgtype link"`
	// 视频
	Video Video `json:"video" validate:"omitempty,required_if=Msgtype video"`
	// 小程序
	Miniprogram MiniProgram `json:"miniprogram" validate:"omitempty,required_if=Msgtype miniprogram"`
}

func (o AttachmentField) Value() (driver.Value, error) {
	b, err := json.Marshal(o)
	return string(b), err
}

func (o *AttachmentField) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), o)
}

func (o AttachmentField) GormDataType() string {
	return "json"
}

// Image 图片
type Image struct {
	Title string `json:"title" validate:"omitempty"`
	// media_id 三天有效，暂时没有用
	MediaId string `json:"media_id" validate:"omitempty"`
	// 用获取到的signd URL
	PicUrl string `json:"pic_url" validate:"omitempty"`
}

// Video 视频
type Video struct {
	Title string `json:"title" validate:"omitempty"`
	//视频媒体文件id，可以通过素材管理接口获得
	MediaId string `json:"media_id" validate:"omitempty"`
}

// Link 图文消息
type Link struct {
	// 图文消息标题，最长为128字节
	Title string `json:"title" validate:"omitempty,lte=128"`
	// 图文消息的链接
	Url string `json:"url" validate:"omitempty,lte=2048"`
	// 图文消息封面的url
	Picurl string `json:"picurl" validate:"omitempty"`
	// 图文消息的描述，最长为512字节
	Desc string `json:"desc" validate:"omitempty"`
}

// MiniProgram	小程序
type MiniProgram struct {
	// 小程序消息标题，最长为64字节
	Title string `json:"title" validate:"lte=64"`
	// 小程序消息封面的mediaid，封面图建议尺寸为520*416
	PicMediaID string `json:"pic_media_id"`
	// 小程序appid，必须是关联到企业的小程序应用
	AppID string `json:"app_id"`
	// 小程序page路径
	Page string `json:"page"`
}

type AttachmentArrayField []gowx.Attachments

func (o AttachmentArrayField) Value() (driver.Value, error) {
	b, err := json.Marshal(o)
	return string(b), err
}

func (o *AttachmentArrayField) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), o)
}

func (o AttachmentArrayField) GormDataType() string {
	return "json"
}

type AutoReplyField struct {
	Text        string               `json:"text" validate:"omitempty"`
	Attachments AttachmentArrayField `json:"attachments" validate:"omitempty"`
}

func (o AutoReplyField) Value() (driver.Value, error) {
	b, err := json.Marshal(o)
	return string(b), err
}

func (o *AutoReplyField) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), o)
}

func (o AutoReplyField) GormDataType() string {
	return "json"
}

type TimeField string

func (o *TimeField) Scan(value interface{}) (err error) {
	switch v := value.(type) {
	case []byte:
		return o.UnmarshalJSON(v)
	case string:
		return o.UnmarshalJSON([]byte(v))
	case time.Time:
		*o = TimeField(o.NewTimeFiled(v))
	case nil:
		*o = ""
	default:
		return fmt.Errorf("cannot sql.Scan() TimeField from: %#v", v)
	}
	return nil
}

func (o TimeField) Value() (driver.Value, error) {
	return string(o), nil
}

func (o TimeField) Time() (time.Time, error) {
	return time.Parse(TimeLayout, string(o))
}

func (o TimeField) MustTime() time.Time {
	t, err := time.Parse(TimeLayout, string(o))
	if err != nil {
		panic(err)
	}
	return t
}

// Seconds 获取从今天凌晨至今的秒数
func (o TimeField) Seconds() int64 {
	t, err := o.Time()
	if err != nil {
		return 0
	}
	return int64(3600*t.Hour() + 60*t.Minute() + t.Second())
}

func (o TimeField) Duration() time.Duration {
	return time.Duration(o.Seconds()) * time.Second
}

// GormDataType gorm common data type
func (o TimeField) GormDataType() string {
	return "time(3)"
}

func (o TimeField) MarshalJSON() ([]byte, error) {
	return []byte(`"` + o + `"`), nil
}

func (o *TimeField) UnmarshalJSON(b []byte) error {
	if b == nil || len(b) == 0 || string(b) == `""` {
		return nil
	}
	if string(b) == "null" {
		return nil
	}
	var err error
	var t time.Time
	t, err = time.Parse(TimeLayout, strings.Trim(string(b), `"`))
	if err != nil {
		return err
	}

	*o = TimeField(t.Format(TimeLayout))
	return nil
}

func (o *TimeField) NewTimeFiled(t time.Time) string {
	return time.Date(0, time.January, 1, t.Hour(), t.Minute(), t.Second(), 0, time.UTC).Format(TimeLayout)
}

type DateField string

func (o *DateField) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	if !nullTime.Valid {
		return nil
	}
	err = nullTime.Scan(value)
	*o = DateField(nullTime.Time.Format(DateLayout))
	return
}

func (o DateField) Value() (driver.Value, error) {
	return string(o), nil
}

func (o DateField) Time() (time.Time, error) {
	return time.Parse(DateLayout, string(o))
}

func (o DateField) MustTime() time.Time {
	t, err := time.Parse(DateLayout, string(o))
	if err != nil {
		return time.Time{}
	}
	return t
}

// GormDataType gorm common data type
func (o DateField) GormDataType() string {
	return "date"
}

func (o DateField) MarshalJSON() ([]byte, error) {
	return []byte(`"` + o + `"`), nil
}

func (o *DateField) UnmarshalJSON(b []byte) error {
	if b == nil || len(b) == 0 || string(b) == `""` {
		return nil
	}
	if string(b) == "null" {
		return nil
	}
	var err error
	var t time.Time
	t, err = time.Parse(DateLayout, strings.Trim(string(b), `"`))
	if err != nil {
		return err
	}

	*o = DateField(t.Format(DateLayout))
	return nil
}

type DateTimeFiled string

func (o DateTimeFiled) MarshalJSON() ([]byte, error) {
	return []byte(`"` + o + `"`), nil
}

func (o *DateTimeFiled) UnmarshalJSON(b []byte) error {
	if b == nil || len(b) == 0 || string(b) == `""` {
		return nil
	}
	if string(b) == "null" {
		return nil
	}
	var err error
	var t time.Time
	t, err = time.Parse(DateTimeLayout, strings.Trim(string(b), `"`))
	if err != nil {
		return err
	}

	*o = DateTimeFiled(t.Format(DateTimeLayout))
	return nil
}

func (o DateTimeFiled) ToInt64() int64 {
	location, err := time.LoadLocation("Local")
	if err != nil {
		return 0
	}
	t, err := time.ParseInLocation(DateTimeLayout, string(o), location)
	if err != nil {
		return 0
	}
	return t.Unix()
	//return t.Unix()
}

// GroupChatWelcomeMsgField 入群欢迎语附件
type GroupChatWelcomeMsgField struct {
	// 图片
	Image Image `json:"image" validate:"omitempty"`
	// 链接
	Link Link `json:"link" validate:"omitempty"`
	// 小程序
	Miniprogram MiniProgram `json:"miniprogram" validate:"omitempty"`
}

func (o GroupChatWelcomeMsgField) Value() (driver.Value, error) {
	b, err := json.Marshal(o)
	return string(b), err
}

func (o *GroupChatWelcomeMsgField) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), o)
}

func (o GroupChatWelcomeMsgField) GormDataType() string {
	return "json"
}

type Time sql.NullTime

// Scan implements the Scanner interface.
func (n *Time) Scan(value interface{}) error {
	return (*sql.NullTime)(n).Scan(value)
}

// Value implements the driver Valuer interface.
func (n Time) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}

func (n Time) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Time)
	}
	return json.Marshal(nil)
}

func (n *Time) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &n.Time)
	if err == nil {
		n.Valid = true
	}
	return err
}
