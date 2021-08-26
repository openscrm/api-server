package requests

import (
	"openscrm/app/constants"
)

type DataExportReq struct {
	// 导出格式
	Format string `json:"format" form:"format" validate:"oneof=excel"`
	// 数据类型 group_chat_list-群聊天列表  delete_customer_warning-删人提醒 delete_staff_warning-流失提醒
	Type string `json:"type" form:"type" validate:"oneof=group_chat_list delete_customer_warning delete_staff_warning"`
	// 筛选条件
	DataFilter struct {
		QueryGroupChatListReq
		QueryStaffDeleteCustomerHistoryReq
		QueryCustomerLossesReq
	} `json:"data_filter"`
	TaskId string `json:"task_id"`
}

type DataExportJob struct {
	DataExportReq
	ExtCorpID string `json:"ext_corp_id"`
}

type QueryDataExportResultReq struct {
}

type DataExportTaskCreateResult struct {
	// taskID
	ID string `json:"id"`
}

type DataExportTaskExeResult struct {
	Status     string `json:"status"`
	ExportTime string `json:"export_time"`
	URL        string `json:"url"`
}

type QueryGroupChatListReq struct {
	// 群名
	Name string `json:"name"`
	// 群标签
	TagIdList []int `json:"tag_id_list"`
	// 使用标签方式 AND/OR
	TagsUnionType string `json:"tags_union_type"`
	// 群主
	Owner string `json:"owner"`
	// 群状态
	IsDismissed string `json:"is_dismissed"`
	// 排序方法
	SortType string `json:"sort_type"`
	// 时间 todo
	Date string `json:"date"`
	// 群状态
	Status constants.GroupChatStatus `json:"status"`
}
