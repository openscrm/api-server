package constants

import (
	"database/sql/driver"
	"encoding/json"
)

type ExtCustomerFilter struct {
	// 客户性别 1-男 2-女 3-未知
	Gender UserGender `json:"gender" form:"gender" validate:"omitempty,oneof=0 1 2 3"`
	// 群聊外部ID
	ExtGroupChatIDs StringArrayField `json:"ext_group_chat_ids" form:"ext_group_chat_ids" validate:"omitempty"`
	// 客户标签
	ExtTagIDs StringArrayField `json:"ext_tag_ids" form:"ext_tag_ids" validate:"omitempty"`
	// 使用ExtTagIDs的逻辑条件 and-且 or-或 none-无标签客户
	TagLogicalCondition string `json:"tag_logical_condition" form:"tag_logical_condition" validate:"omitempty,oneof=and or none"`
	// 排除客户标签
	ExcludeExtTagIDs StringArrayField `json:"exclude_ext_tag_ids" form:"exclude_ext_tag_ids" validate:"omitempty"`
	// 添加好友,开始时间
	StartTime DateField `json:"start_time" form:"start_time"`
	// 添加好友,结束时间
	EndTime DateField `json:"end_time" form:"end_time"`
}

func (o ExtCustomerFilter) Value() (driver.Value, error) {
	b, err := json.Marshal(o)
	return string(b), err
}

func (o *ExtCustomerFilter) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), o)
}

func (o ExtCustomerFilter) GormDataType() string {
	return "json"
}
