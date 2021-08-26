package requests

import "openscrm/common/app"

// QueryWelcomeMsgReq 欢迎语查询请求
type QueryWelcomeMsgReq struct {
	// 欢迎语标题
	Name string `json:"name" form:"name" validate:"omitempty"`
	// 欢迎语可用员工
	ExtStaffIDs string `json:"ext_staff_ids" form:"ext_staff_ids" validate:"omitempty"`
	app.Pager
	app.Sorter
}
