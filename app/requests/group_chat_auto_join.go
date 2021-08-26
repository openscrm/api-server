package requests

import (
	"openscrm/app/constants"
	"openscrm/common/app"
)

type GroupChatAutoJoinCodeStaffParam struct {
	// 员工外部ID
	ExtStaffID string `json:"staff_id"  validate:"required"`
	// 每日添加客户数上线
	DayAddCustomerLimit int `json:"day_add_customer_limit" validate:"omitempty"`
}

type CreateGroupChatAutoJoinCodeReq struct {
	// 拉群方式,1-群二维码 2-企微活码
	CreateType constants.GroupChatAutoCreateType `json:"create_type" validate:"required,oneof=1 2"`
	// 二维码名称
	Remark string `json:"remark" validate:"required"`
	// 群ID
	GroupID string `json:"group_id" validate:"required,int64"`
	// 关联的外部员工ID列表
	ExtStaffIDs constants.StringArrayField `json:"ext_staff_ids" validate:"required,gt=0"`
	// 备份员工ID列表
	BackupExtStaffIDs constants.StringArrayField `json:"backup_ext_staff_ids" validate:"omitempty,gte=0"`
	// 员工参数
	Staffs []GroupChatAutoJoinCodeStaffParam `json:"staffs" validate:"required,gt=0"`
	// 备份员工
	BackupStaffs []GroupChatAutoJoinCodeStaffParam `json:"backup_staffs" validate:"omitempty"`
	// 客户添加员工是否需要员工确认,默认-否
	SkipVerify constants.Boolean `json:"skip_verify" validate:"omitempty,oneof=1 2"`
	// 是否自动打标签
	AutoTagEnable constants.Boolean `json:"auto_tag_enable"  validate:"oneof=1 2"`
	// 自动打标签绑定的标签ID数组
	ExtTagIDs constants.StringArrayField `json:"ext_tag_ids"  validate:"omitempty"`
	// 是否开启员工每日添加上限
	DailyAddCustomerLimitEnable constants.Boolean `json:"daily_add_customer_limit_enable"  validate:"oneof=1 2"`
	// 员工被客户扫码添加的自动回复
	AutoReply string `json:"auto_reply" validate:"required"`
	// 自动拉群码绑定的群二维码
	GroupChatQRCode []GroupChatQRCode `json:"group_chat_qr_code" validate:"required,dive"`
}

type GroupChatQRCode struct {
	// 排序
	Order int `json:"order"`
	// 二维码mediaID
	QrMediaId string `json:"qr_media_id" validate:"required"`
	// 二维码url
	QrUrl string `json:"qr_url" validate:"required"`
	// 群二维码加好友人数上限
	UserLimit int `json:"user_limit" validate:"required"`
	// 二维码状态 1-启用中 2-已停用
	Status constants.GroupChatAutoJoinQRCodeStatus `json:"status" validate:"omitempty"`
}

type DeleteGroupChatAutoJoinReq struct {
	// 自动拉群码ID
	IDs []string `json:"ids" validate:"gt=0,dive,int64"`
}

type QueryGroupChatAutoJoinReq struct {
	// 群二维码分组ID
	GroupID string `json:"group_id"   validate:"omitempty" form:"group_id"`
	// 备注
	Remark string `json:"remark" validate:"omitempty" form:"remark"`
	app.Pager
	app.Sorter
}

type UpdateGroupChatAutoJoinQrCodeReq struct {
	CreateGroupChatAutoJoinCodeReq
}

// BatchRegroupReq 批量分组
type BatchRegroupReq struct {
	// 自动拉群码ID
	IDs []string `json:"ids" form:"ids" validate:"gt=0,dive,int64"`
	// 新组ID
	NewGroupID string `json:"new_group_id" form:"new_group_id" validate:"required"`
}
