package entities

import (
	"openscrm/app/constants"
	"openscrm/common/app"
)

//QueryStaffDeleteCustomerHistoryReq 查找删除客户的员工记录
type QueryStaffDeleteCustomerHistoryReq struct {
	// 企微部门id
	ExtDepartmentID int64 `json:"department_id" form:"ext_department_id" validate:"omitempty,gte=0"`
	// 企微员工id
	ExtStaffID string `json:"ext_staff_id" form:"ext_staff_id" validate:"omitempty"`
	app.Pager
	app.Sorter
	app.TimeInterval
}
type EventList struct {
	CorpId             string `json:"ext_corp_id"`
	EventType          string `json:"event_type"`
	ExternalAvatar     string `json:"external_avatar"`
	ExternalUserExtId  string `json:"external_user_ext_id"`
	ExternalUserName   string `json:"external_user_name"`
	RelationCreatetime int    `json:"relation_createtime"`
	StaffAvatar        string `json:"staff_avatar"`
	StaffExtId         string `json:"staff_ext_id"`
	StaffId            int    `json:"staff_id"`
	StaffName          string `json:"staff_name"`
	//HasChat            bool   `json:"has_chat"`
	//Read               bool   `json:"read"`
	//StaffUnionid       interface{} `json:"staff_unionid"`
	//EuUnionid          interface{} `json:"eu_unionid"`
}

// QueryStaffDeleteCustomerHistoryResp 员工删除客户记录响应
type QueryStaffDeleteCustomerHistoryResp struct {
	// 客户id
	ExtCustomerID string `json:"ext_customer_id"`
	// 客户头像 url
	ExtCustomerAvatar string `json:"ext_customer_avatar"`
	// 客户名
	ExtCustomerName string `json:"ext_customer_name"`
	// 添加好友时间
	RelationCreateAt int `json:"relation_create_at"`
	// 删除好友时间
	RelationDeleteAt int `json:"relation_delete_at"`
	// 员工头像url
	ExtStaffAvatar string `json:"ext_staff_avatar"`
	// 企微员工id
	ExtStaffId string `json:"ext_staff_id"`
	// 员工id
	StaffId int `json:"staff_id"`
	// 员工名字
	StaffName string `json:"staff_name"`
}

// UpdateStaffDeleteCustomerNotifierReq 删人提醒设置开关
type UpdateStaffDeleteCustomerNotifierReq struct {
	// 是否开启通知, 员工删除客户时提醒管理员,2-关闭 1-打开
	IsNotifyAdmins constants.EventNotifyStatus `json:"is_notify_admins" form:"is_notify_admins" validate:"oneof=1 2"`
	// 发送通知的时间,1-实时 2-每天早八点
	NotifyType constants.EventNotifyTime `json:"notify_type" form:"notify_type"  validate:"omitempty,required_if=IsNotifyAdmins 2,oneof=1 2"` //
	// 接收通知的管理员
	AdminIDs constants.StringArrayField `json:"admin_ids" form:"admin_ids" validate:"omitempty,required_if=IsNotifyAdmins 2,unique,validAdmin,gt=0"` //
}

// UpdateCustomerDeleteStaffNotifierReq 流失提醒开关
type UpdateCustomerDeleteStaffNotifierReq struct {
	// 是否通知员工，员工被客户删除时提醒,1-打开 2-关闭
	IsNotifyStaff constants.EventNotifyStatus `json:"is_notify_staff" form:"is_notify_staff" validate:"oneof=1 2"`
}

// UpdateCustomerDeleteStaffNotifierResp 流失提醒开关
type UpdateCustomerDeleteStaffNotifierResp struct {
	// 是否通知员工，员工被客户删除时提醒,1-打开 2-关闭
	IsNotifyStaff constants.EventNotifyStatus `json:"is_notify_staff" form:"is_notify_staff"`
}
