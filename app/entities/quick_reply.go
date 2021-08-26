package entities

import (
	"openscrm/app/constants"
)

// CreateQuickReplyReq 新建话术请求
type CreateQuickReplyReq struct {
	// 话术标题
	Name string `json:"name" form:"name" validate:"required"`
	// 所属分组
	GroupID string `json:"group_id" form:"group_id" validate:"required"`
	// 话术内容
	QuickReplyDetails []QuickReplyDetail `json:"reply_details" form:"reply_details" validate:"required,dive"`
}

// UpdateQuickReplyReq 更新话术请求
type UpdateQuickReplyReq struct {
	// ID
	ID string `json:"id" form:"id"`
	// 话术标题
	Name string `json:"name" form:"name" validate:"required"`
	// 所属分组,如果和原分组不同，则跟新为当前值
	GroupID string `json:"group_id" form:"group_id" validate:"required"`
	// 具体内容
	QuickReplyDetails []QuickReplyDetail `json:"reply_details" form:"reply_details" validate:"required,dive"`
	// 删除的条目
	DeletedIDs []string `json:"deleted_ids" form:"deleted_ids"`
}

// QuickReplyDetail 话术条目详情
type QuickReplyDetail struct {
	// 更新时需要,没有id表示新建，有id为更新
	ID string `json:"id" form:"id"`
	//  2-文字 3-图片 4-网页 5-pdf 6-视频
	ContentType       constants.QuickReplyType  `json:"content_type" form:"content_type" validate:"required,oneof=2 3 4 5 6"`
	QuickReplyContent constants.QuickReplyField `json:"quick_reply_content"  form:"quick_reply_content"`
}

// DeleteQuickReplyReq 删除话术
type DeleteQuickReplyReq struct {
	// 话术id
	IDs []string `json:"ids" validate:"gt=0,dive,int64"`
}
