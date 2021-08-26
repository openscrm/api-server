package requests

import (
	"openscrm/app/constants"
	"openscrm/common/app"
)

// QueryGroupChatReq 查询群聊请求
type QueryGroupChatReq struct {
	// 群主列表
	Owners []string `json:"owners"  validate:"omitempty" form:"owners"`
	// 群名
	Name string `json:"name" validate:"omitempty" form:"name"`
	// 群状态 1-已解散 2-未解散
	Status constants.GroupChatStatus `json:"status" validate:"omitempty" form:"status"`
	// 创建群时间-开始
	CreateTimeStart constants.DateField `json:"create_time_start" validate:"omitempty,date" form:"create_time_start"`
	// 创建群时间-结束
	CreateTimeEnd constants.DateField `json:"create_time_end" validate:"omitempty,date" form:"create_time_end"`
	// 群标签ID列表
	GroupTagIDs constants.Int64ArrayField `json:"group_tag_ids" validate:"omitempty" form:"group_tag_ids"`
	// 群标签ID查询条件，and/or
	TagsUnionType string `json:"tags_union_type" validate:"omitempty" form:"tags_union_type"`
	app.Pager
	app.Sorter
}

type GetAllGroupChatReq struct {
	app.Pager
	app.Sorter
}
