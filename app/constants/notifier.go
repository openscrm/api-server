package constants

type EventNotifyStatus int

const (
	EventNotifyStatusOn  = 1 // 打开通知
	EventNotifyStatusOff = 2 // 关闭通知
)

type EventNotifyTime int

const (
	EventNotifyTimeRealTime = 1
	EventNotifyTimeTimed    = 2 // 8:00 am
)

type EventNotifyType int

const (
	EventNotifyTypeStaffDeleteCustomer = 1 //员工删客户
	EventNotifyTypeCustomerDeleteStaff = 2 //客户删员工
)

type IsNotified int

const (
	NotifiedAdmin    = 1
	NotNotifiedAdmin = 2
)

type EventName string

const (
	EventNameStaffDeleteCustomer = "staff_delete_customer"
	EventNameCustomerDeleteStaff = "customer_delete_staff"
)

type TimedNotifyAdminMsg struct {
	ExtStaffID    string   `json:"ext_staff_id"`
	ExtCustomerID string   `json:"ext_customer_id"`
	Content       string   `json:"content"`
	AdminIDs      []string `json:"admin_ids"`
}
