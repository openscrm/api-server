package constants

import (
	"database/sql/driver"
	"encoding/json"
)

type CustomerRemarkField []CustomerRemarkContent

// CustomerRemarkContent CustomerRemark的字段内容，json格式，放在CustomerInfo中
type CustomerRemarkContent struct {
	RemarkID string `json:"remark_id"`
	//RemarkOptionID string `json:"remark_option_id"`
	RemarkType string `json:"remark_type"`
	// 多选-optionID 时间-字符串 text-字符串
	RemarkValue string `json:"remark_value"`
}

func (o CustomerRemarkField) Value() (driver.Value, error) {
	b, err := json.Marshal(o)
	return string(b), err
}

func (o *CustomerRemarkField) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), o)
}

func (o CustomerRemarkField) GormDataType() string {
	return "json"
}
