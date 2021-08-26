package requests

import (
	"openscrm/app/constants"
	"openscrm/common/app"
	"time"
)

type CommonDeleteReq struct {
	IDs []string `json:"ids" form:"ids" validate:"required,gt=0"`
}

type LocalTime time.Time

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}

	loc, _ := time.LoadLocation("Asia/Chongqing")
	now, err := time.ParseInLocation(`"`+constants.DateTimeLayout+`"`, string(data), loc)
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(constants.DateTimeLayout)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, constants.DateTimeLayout)
	b = append(b, '"')
	return b, nil
}

type QueryCustomerGroupTagGroupReq struct {
	Name string `json:"name" form:"name" validate:"omitempty"`
	app.Sorter
	app.Pager
}
