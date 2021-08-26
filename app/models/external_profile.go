package models

import (
	"database/sql/driver"
	"encoding/json"
)

func (o ExternalProfile) Value() (driver.Value, error) {
	b, err := json.Marshal(o)
	return string(b), err
}

func (o *ExternalProfile) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), o)
}

func (o ExternalProfile) GormDataType() string {
	return "json"
}

// ExternalProfile 使gorm支持 ExternalProfile 转为json读取
type ExternalProfile struct {
	ExternalCorpName string         `json:"external_corp_name"`
	ExternalAttr     []ExternalAttr `json:"external_attr"`
}

type Text struct {
	Value string `json:"value"`
}
type Web struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}
type Miniprogram struct {
	Appid    string `json:"appid"`
	Pagepath string `json:"pagepath"`
	Title    string `json:"title"`
}
type ExternalAttr struct {
	Type        int         `json:"type"`
	Name        string      `json:"name"`
	Text        Text        `json:"text,omitempty"`
	Web         Web         `json:"web,omitempty"`
	Miniprogram Miniprogram `json:"miniprogram,omitempty"`
}
