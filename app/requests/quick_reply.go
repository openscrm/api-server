package requests

import "openscrm/common/app"

// QueryQuickReplyReq 查询话术条目
type QueryQuickReplyReq struct {
	//  话术组id
	GroupID string `form:"group_id" json:"group_id"  validate:"omitempty,int64"`
	// 可用部门id
	DepartmentIDs []int64 `form:"department_ids" json:"department_ids"  validate:"omitempty,gt=0"`
	// 关键词
	Keyword string `form:"keyword" json:"keyword" validate:"omitempty"`
	app.Pager
	app.Sorter
}
