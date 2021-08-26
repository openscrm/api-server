package entities

import (
	"openscrm/app/constants"
	"openscrm/common/app"
)

// UpdateCorporationReq 更新企业请求参数
type UpdateCorporationReq struct {
	// ExtCorpID 外部企业ID
	ExtCorpID string `json:"ext_corp_id" gorm:"index;type:char(18);comment:外部企业ID" validate:"omitempty,required,corp_id" `
	// ContactSecret 通讯录secret
	ContactSecret string `json:"contact_secret" gorm:"comment:'通讯录secret'" validate:"omitempty,required"`
	// CustomerSecret 客户联系secret
	CustomerSecret string `json:"customer_secret" gorm:"comment:'客户联系secret'" validate:"omitempty,required"`
	// MainAgentID 企业主应用AgentID
	MainAgentID int64 `json:"main_agent_id" gorm:"comment:'企业主应用AgentID'" validate:"number,gte=0"`
	// MainAgentSecret 企业主应用secret
	MainAgentSecret string `json:"main_agent_secret" gorm:"comment:'企业主应用secret'"`
	// CallbackToken 企业微信事件回调Token
	CallbackToken string `json:"callback_token" gorm:"comment:'企业微信事件回调Token'" validate:"omitempty,required"`
	// CallbackAesKey 企业微信事件回调AesKey
	CallbackAesKey string `json:"callback_aes_key" gorm:"comment:'企业微信事件回调AesKey'" validate:"omitempty,required"`
}

// QueryCorporationReq 查询企业列表请求参数
type QueryCorporationReq struct {
	app.Pager
	app.Sorter
	// ID 企业ID
	ID string `form:"id" json:"id" gorm:"primaryKey;type:bigint;comment:'ID'" validate:"omitempty,int64" `
	// Name 企业名称
	Name string `form:"name" json:"name" gorm:"type:varchar(255);index;comment:'企业名称'"`
	// Enable 是否有效，1：是；2：否
	Enable constants.Boolean `json:"enable" gorm:"default:1;comment:'是否有效，1：是；2：否'" validate:"omitempty,oneof=1 2"`
	// ExtCorpID 外部企业ID
	ExtCorpID string `json:"ext_corp_id" gorm:"index;type:char(18);comment:外部企业ID" validate:"omitempty,required,corp_id"`
}
