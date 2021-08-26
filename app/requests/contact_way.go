package requests

import (
	"openscrm/app/constants"
	"openscrm/common/app"
)

type ContactWayStaffParam struct {
	ID                  string `json:"id" gorm:"primaryKey;type:bigint;comment:'ID'" validate:"omitempty,int64"`
	DayAddCustomerLimit int    `json:"day_add_customer_limit" gorm:"comment:'员工每日添加客户上限'"`
	ExtStaffID          string `json:"ext_staff_id" gorm:"index;unique:contact_way_id_and_ext_staff_id;comment:'外部员工ID'" validate:"required,word"`
}

type ContactWayScheduleParam struct {
	ID                  string `json:"id" gorm:"primaryKey;type:bigint;comment:'ID'" validate:"omitempty,int64"`
	DayAddCustomerLimit int    `json:"day_add_customer_limit" gorm:"comment:'员工每日添加客户上限'"`
	// Weekdays 工作日
	Weekdays constants.StringArrayField `json:"weekdays" gorm:"json;comment:'工作日'" validate:"required,dive,weekday"`
	// StartTime 开始时间
	StartTime constants.TimeField `json:"start_time" gorm:"comment:开始时间" validate:"required"`
	// EndTime 结束时间
	EndTime constants.TimeField `json:"end_time" gorm:"comment:结束时间" validate:"required"`
	// Staffs 绑定员工参数
	Staffs []ContactWayStaffParam `json:"staffs" validate:"required"`
}

// CreateContactWayReq 创建渠道码
type CreateContactWayReq struct {
	Name string `gorm:"type:varchar(255);index;comment:'渠道码名称'" json:"name" validate:"required"`
	// GroupID 渠道码分组ID
	GroupID string `json:"group_id" gorm:"index;type:bigint;comment:活码分组ID" validate:"required,int64"`
	// AutoReplyType 欢迎语类型：1，渠道欢迎语；2, 渠道默认欢迎语；3，不送欢迎语；
	AutoReplyType constants.ContactWayAutoReplyType `json:"auto_reply_type" gorm:"index;default:1;comment:'欢迎语类型：1，渠道欢迎语；2, 渠道默认欢迎语；3，不送欢迎语；'"`
	// AutoReply 欢迎语策略
	AutoReply constants.AutoReplyField `json:"auto_reply" gorm:"comment:欢迎语策略" validate:"required_if=AutoReplyType 1,omitempty"`
	// CustomerDesc 客户描述
	CustomerDesc string `json:"customer_desc" gorm:"comment:客户描述" validate:"required_if=CustomerDescEnable 1"`
	// CustomerDescEnable 是否开启客户描述
	CustomerDescEnable constants.Boolean `json:"customer_desc_enable" gorm:"comment:是否开启客户描述" validate:"oneof=1 2"`
	// CustomerRemark 客户备注
	CustomerRemark string `json:"customer_remark" gorm:"comment:客户备注" validate:"required_if=CustomerRemarkEnable 1"`
	// CustomerRemarkEnable 是否开启客户备注
	CustomerRemarkEnable constants.Boolean `json:"customer_remark_enable" gorm:"comment:是否开启客户备注" validate:"oneof=1 2"`
	// DailyAddCustomerLimitEnable 是否开启员工每日添加上限
	DailyAddCustomerLimitEnable constants.Boolean `json:"daily_add_customer_limit_enable" gorm:"comment:是否开启员工每日添加上限" validate:"oneof=1 2"`
	// 员工每日添加上限
	DailyAddCustomerLimit int64 `json:"daily_add_customer_limit" gorm:"comment:员工每日添加上限" validate:"gte=0"`
	// ScheduleEnable 是否开启自动上下线
	ScheduleEnable constants.Boolean `json:"schedule_enable" gorm:"comment:是否开启自动上下线" validate:"oneof=1 2"`
	// Staffs 绑定员工参数
	Staffs []ContactWayStaffParam `json:"staffs" validate:"required_if=ScheduleEnable 2,dive"`
	// BackupStaffs 绑定备份员工参数
	BackupStaffs []ContactWayStaffParam `json:"backup_staffs" validate:"required"`
	// Schedules 员工调度设置
	Schedules []ContactWayScheduleParam `json:"schedules" validate:"required_if=ScheduleEnable 1,dive"`
	// AutoTagEnable 是否自动打标签
	AutoTagEnable constants.Boolean `json:"auto_tag_enable" gorm:"comment:'是否自动打标签'" validate:"oneof=1 2"`
	// CustomerTagExtIDs 自动打标签绑定的标签ExtID数组
	CustomerTagExtIDs constants.StringArrayField `json:"customer_tag_ext_ids" gorm:"type:json;comment:'自动打标签绑定的标签ExtID数组'" validate:"required_if=AutoTagEnable 1,dive,ext_id"`
	// SkipVerify 外部客户添加时是否无需验证，假布尔类型
	SkipVerify constants.Boolean `json:"skip_verify" gorm:"default:1;comment:外部客户添加时是否无需验证，假布尔类型" validate:"oneof=1 2"`
	// AutoSkipVerifyEnable 是否开启自动通过好友时段控制
	AutoSkipVerifyEnable constants.Boolean `json:"auto_skip_verify_enable" gorm:"default:1;comment:是否开启自动通过好友时段控制" validate:"required_if=SkipVerify 1,omitempty,oneof=1 2"`
	// SkipVerifyStartTime 自动通过好友开启时间
	SkipVerifyStartTime constants.TimeField `json:"skip_verify_start_time" gorm:"comment:自动通过好友开启时间" validate:"required_if=AutoSkipVerifyEnable 1,omitempty,time"`
	// SkipVerifyEndTime 自动通过好友结束时间
	SkipVerifyEndTime constants.TimeField `json:"skip_verify_end_time" gorm:"comment:自动通过好友结束时间" validate:"required_if=AutoSkipVerifyEnable 1,omitempty,time"`
	// Remark 渠道码的备注信息，用于助记
	Remark string `json:"remark" gorm:"comment:渠道码的备注信息"`
	// StaffControlEnable 员工自行上下线
	StaffControlEnable constants.Boolean `json:"staff_control_enable" gorm:"comment:员工自行上下线" validate:"oneof=1 2"`
	// NicknameBlockEnable 是否开启客户昵称屏蔽欢迎语
	NicknameBlockEnable constants.Boolean `json:"nickname_block_enable" gorm:"comment:是否开启客户昵称屏蔽欢迎语" validate:"oneof=1 2"`
	// NicknameBlockList 客户昵称屏蔽欢迎语列表
	NicknameBlockList constants.StringArrayField `json:"nickname_block_list" gorm:"type:json;comment:'客户昵称屏蔽欢迎语列表'" validate:"required_if=NicknameBlockEnable 1"`
}

// UpdateContactWayReq 更新渠道码（必须全量更新）
type UpdateContactWayReq struct {
	CreateContactWayReq
}

// QueryContactWayReq 查询渠道码列表请求参数
type QueryContactWayReq struct {
	app.Pager
	app.Sorter
	// ID 渠道码ID
	ID string `form:"id" json:"id" gorm:"primaryKey;type:bigint;comment:'ID'" validate:"omitempty,int64" `
	// 企微员工id
	ExtStaffIDs []string `json:"ext_staff_ids" form:"ext_staff_ids" validate:"omitempty,dive,word"`
	// Name 渠道码名称
	Name string `form:"name" json:"name" gorm:"type:varchar(255);index;comment:'渠道码名称'"`
	// ConfigID 渠道码配置ID
	ConfigID string `json:"config_id" form:"config_id" gorm:"comment:渠道码配置ID" validate:"omitempty,word"`
	// GroupID 渠道码分组ID
	GroupID string `json:"group_id" form:"group_id" gorm:"type:bigint;comment:活码分组ID" validate:"omitempty,int64"`
	// CreatedAtStart 创建时间范围开始
	CreatedAtStart constants.DateField `json:"created_at_start" form:"created_at_start" gorm:"-" validate:"omitempty,date"`
	// CreatedAtEnd 创建时间范围结束
	CreatedAtEnd constants.DateField `json:"created_at_end" form:"created_at_end" gorm:"-" validate:"omitempty,date"`
}

// DeleteContactWayReq 删除渠道码请求参数
type DeleteContactWayReq struct {
	// 渠道码ID
	IDs []string `json:"ids" validate:"gt=0,dive,int64"`
}

// BatchUpdateContactWayReq 批量修改渠道码请求参数
type BatchUpdateContactWayReq struct {
	// GroupID 渠道码分组ID
	GroupID string `json:"group_id" gorm:"index;type:bigint;comment:活码分组ID" validate:"omitempty,int64"`
	// 渠道码ID
	IDs []string `json:"ids" validate:"gt=0,dive,int64"`
}
