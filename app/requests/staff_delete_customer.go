package requests

import (
	"openscrm/app/constants"
	"openscrm/common/app"
)

//QueryStaffDeleteCustomerHistoryReq 查找删除客户的员工记录
type QueryStaffDeleteCustomerHistoryReq struct {
	// 企微部门id
	ExtDepartmentID int64 `json:"department_id" form:"ext_department_id" validate:"omitempty,gte=0"`
	// 企微员工id
	ExtStaffIDs []string `json:"ext_staff_id" form:"ext_staff_id" validate:"omitempty"`
	// 添加好友时间起始
	ConnectionCreateStart constants.TimeField `json:"connection_create_start" form:"connection_create_start" validate:"omitempty"`
	ConnectionCreateEnd   constants.TimeField `json:"connection_create_end" form:"connection_create_end" validate:"omitempty"`
	// 删除员工时间起始
	DeleteCustomerStart constants.TimeField `json:"delete_customer_start" form:"delete_customer_start" validate:"omitempty"`
	DeleteCustomerEnd   constants.TimeField `json:"delete_customer_end" form:"delete_customer_end" validate:"omitempty"`
	app.Pager
	app.Sorter
}

// UpdateStaffDeleteCustomerNotifierReq 删人提醒设置开关
type UpdateStaffDeleteCustomerNotifierReq struct {
	// 是否开启通知, 员工删除客户时提醒管理员,2-关闭 1-打开
	IsNotifyStaff constants.EventNotifyStatus `json:"is_notify_staff" form:"is_notify_staff" validate:"oneof=1 2"`
	// 发送通知的时间,1-实时 2-每天早八点
	NotifyType constants.EventNotifyTime `json:"notify_type" form:"notify_type"  validate:"omitempty,required_if=IsNotifyAdmins 2,oneof=1 2"` //
	// 接收通知的管理员
	ExtStaffIDs constants.StringArrayField `json:"ext_staff_ids" form:"ext_staff_ids" validate:"omitempty,required_if=IsNotifyAdmins 2,unique,validAdmin,gt=0"` //
}

// UpdateCustomerDeleteStaffNotifierReq 流失提醒开关
type UpdateCustomerDeleteStaffNotifierReq struct {
	// 是否通知员工，员工被客户删除时提醒,1-打开 2-关闭
	IsNotifyStaff constants.EventNotifyStatus `json:"is_notify_staff" form:"is_notify_staff" validate:"oneof=1 2"`
}
