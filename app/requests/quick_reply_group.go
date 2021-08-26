package requests

import "openscrm/common/app"

// QueryQuickReplyGroupReq 查询话术分组
// 企业话术无需查询条件
type QueryQuickReplyGroupReq struct {
	app.Pager
	app.Sorter
}
