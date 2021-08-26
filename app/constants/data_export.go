package constants

type DataExportType string

const (
	DataExportTypeGroupChat             DataExportType = "group_chat_list"
	DataExportTypeCustomer              DataExportType = "customer_list"
	DataExportTypeDeleteStaffWarning    DataExportType = "delete_staff_warning"
	DataExportTypeDeleteCustomerWarning DataExportType = "delete_customer_warning"
)

const (
	DataExportGroupCustomerListPrefix      = "xjyk-CustomerList"
	DataExportGroupChatListPrefix          = "xjyk-GroupChatList"      //"小橘有客-群聊列表"
	DataExportDeleteCustomerFilenamePrefix = "xjyk-DeleteCustomerList" //"小橘有客-删人提醒"
	DataExportDeleteStaffFilenamePrefix    = "xjyk-DeleteStaffList"    //"小橘有客-客户流失提醒提醒"
)

const (
	DataExportCustomerListSheetName       = "客户列表"   //"小橘有客-客户列表"
	DataExportGroupChatListSheetName      = "客户群列表"  //"小橘有客-群聊列表"
	DataExportDeleteCustomerListSheetName = "删人提醒列表" //"小橘有客-删人提醒"
	DataExportDeleteStaffListSheetName    = "流失提醒列表" //"小橘有客-客户流失提醒提醒"
)
