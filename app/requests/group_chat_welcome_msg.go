package requests

import (
	"openscrm/app/constants"
	"openscrm/common/app"
)

type UpdateGroupChatWelcomeMsgReq struct {
	// 文字内容
	Content string `json:"content"`
	// 附件类型
	AttachmentType string `json:"attachment_type"`
	// 附件内容
	Attachment constants.GroupChatWelcomeMsgField `json:"attachment" validate:"omitempty"`
}

type CreateGroupChatWelcomeMsgReq struct {
	// 文字内容
	Content string `json:"content"`
	// 附件类型
	AttachmentType string `json:"attachment_type"`
	// 附件内容
	Attachment constants.GroupChatWelcomeMsgField `json:"attachment" validate:"omitempty"`
	// 开启后，新建该条欢迎语会通过「客户群」群发通知企业全部员工：“管理员创建了新的入群欢迎语”
	// 不可更改
	NotifyStaffsEnable constants.Boolean `json:"notify_staffs_enable"`
}

type QueryGroupChatWelcomeMsgReq struct {
	// 搜索内容
	Content string `json:"content" form:"content" validate:"required"`
	app.Pager
	app.Sorter
}
