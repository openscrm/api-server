package constants

type SessionName string

const (
	// CustomerSessionName 前台用户Session
	CustomerSessionName SessionName = "OpenSCRMCustomerSession"
	// StaffSessionName 前台员工Session
	StaffSessionName SessionName = "OpenSCRMStaffSession"
	// StaffAdminSessionName 企业员工后台session
	StaffAdminSessionName SessionName = "OpenSCRMStaffAdminSession"
	// CorpAdminSessionName 企业超级管理员session
	CorpAdminSessionName SessionName = "OpenSCRMCorpAdminSession"
	// SaasAdminSessionName Saas管理员Session
	SaasAdminSessionName SessionName = "OpenSCRMSaasAdminSession"
)

type SessionField string

const (
	// StaffID 员工ID
	StaffID SessionField = "ExtStaffID"
	// CustomerID 客户ID
	CustomerID SessionField = "CustomerID"
	// ExtCorpID 外部企业ID
	ExtCorpID SessionField = "ExtCorpID"
	// ExtStaffID 外部员工ID
	ExtStaffID SessionField = "ExtStaffID"
	// ExtCustomerID 外部客户ID
	ExtCustomerID SessionField = "ExtCustomerID"
	// StaffInfo 会话中的员工信息
	StaffInfo SessionField = "StaffInfo"
	// CustomerInfo 会话中的客户信息
	CustomerInfo SessionField = "CustomerInfo"
	// QrcodeAuthState 扫码登录state
	QrcodeAuthState SessionField = "QrcodeAuthState"
)
