package entities

// CorpAdminLoginReq 企业系统管理员登录
type CorpAdminLoginReq struct {
	// Phone 手机号
	Phone string `json:"phone" gorm:"index;comment:'手机号'" validate:"required,phone"`
	// Password 密码
	Password string `json:"password" validate:"required,max=100"`
}

// SaasAdminLoginReq 企业超级管理员登录
type SaasAdminLoginReq struct {
	// Phone 手机号
	Phone string `json:"phone" gorm:"index;comment:'手机号'" validate:"required,phone"`
	// Password 密码
	Password string `json:"password" validate:"required,max=100"`
}

// StaffAdminLoginReq 企业员工后台扫码登录请求
type StaffAdminLoginReq struct {
	// ExtCorpID 外部企业ID
	ExtCorpID string `form:"ext_corp_id" json:"ext_corp_id" gorm:"index;comment:'外部企业ID'" validate:"omitempty,corp_id"`
	// SourceURL 登录来源页面URL，登录成功自动跳转过去
	SourceURL string `form:"source_url" json:"source_url" validate:"required"`
}

// StaffAdminLoginResp 企业员工后台扫码登录响应
type StaffAdminLoginResp struct {
	// AppID 外部企业ID
	AppID string `json:"app_id"`
	// AgentID 授权应用AgentID
	AgentID int64 `json:"agent_id"`
	// RedirectURI 重定向地址，需要进行UrlEncode，授权回调地址
	RedirectURI string `json:"redirect_uri"`
	// State
	State string `json:"state"`
	// LocationURL 完整的跳转URL
	LocationURL string `json:"location_url"`
}

// StaffAdminLoginCallbackReq 企业员工后台扫码登录回调请求
type StaffAdminLoginCallbackReq struct {
	// AppID 外部企业ID
	AppID string `form:"appid" json:"appid" validate:"required,corp_id"`
	// Code
	Code string `form:"code" json:"Code" validate:"required"`
	// State
	State string `form:"state" json:"state"`
}

// StaffAdminForceLoginReq 调试接口-指定普通管理员强制登录
type StaffAdminForceLoginReq struct {
	// ExtCorpID 外部企业ID
	ExtCorpID string `form:"ext_corp_id" json:"ext_corp_id" validate:"omitempty,corp_id"`
	// ExtStaffID 外部员工ID
	ExtStaffID string `form:"ext_staff_id" json:"ext_staff_id" validate:"required"`
}

// StaffForceLoginReq 调试接口-指定员工侧边栏强制登录
type StaffForceLoginReq struct {
	// ExtCorpID 外部企业ID
	ExtCorpID string `form:"ext_corp_id" json:"ext_corp_id" validate:"omitempty,corp_id"`
	// ExtStaffID 外部员工ID
	ExtStaffID string `form:"ext_staff_id" json:"ext_staff_id" validate:"required"`
}

// StaffLoginReq 企业员工H5登录请求
type StaffLoginReq struct {
	// ExtCorpID 外部企业ID
	ExtCorpID string `form:"ext_corp_id" json:"ext_corp_id" gorm:"index;comment:'外部企业ID'" validate:"omitempty,corp_id"`
	// SourceURL 登录来源页面URL，登录成功自动跳转过去
	SourceURL string `form:"source_url" json:"source_url" validate:"required"`
}

// StaffLoginResp 企业员工H5登录响应
type StaffLoginResp struct {
	// AppID 外部企业ID
	AppID string `json:"app_id"`
	// RedirectURI 重定向地址，需要进行UrlEncode，授权回调地址
	RedirectURI string `json:"redirect_uri"`
	// SourceURL 登录来源页面URL，登录成功自动跳转过去
	SourceURL string `form:"source_url" json:"source_url" validate:"required,url"`
	// State
	State string `json:"state"`
	// LocationURL 完整的跳转URL
	LocationURL string `json:"location_url"`
}

// StaffLoginCallbackReq 企业员工H5登录回调请求
type StaffLoginCallbackReq struct {
	// AppID 外部企业ID
	AppID string `form:"appid" json:"appid" validate:"required,corp_id"`
	// Code
	Code string `form:"code" json:"Code" validate:"required"`
	// SourceURL 登录来源页面URL，登录成功自动跳转过去
	SourceURL string `form:"source_url" json:"source_url" validate:"required,url"`
	// State
	State string `form:"state" json:"state"`
}

// CustomerLoginReq 客户H5登录请求
type CustomerLoginReq struct {
	// ExtCorpID 外部企业ID
	ExtCorpID string `form:"ext_corp_id" json:"ext_corp_id" gorm:"index;comment:'外部企业ID'" validate:"omitempty,corp_id"`
	// SourceURL 登录来源页面URL，登录成功自动跳转过去
	SourceURL string `form:"source_url" json:"source_url" validate:"required"`
}

// CustomerLoginResp 客户H5登录响应
type CustomerLoginResp struct {
	// AppID 外部企业ID
	AppID string `json:"app_id"`
	// RedirectURI 重定向地址，需要进行UrlEncode，授权回调地址
	RedirectURI string `json:"redirect_uri"`
	// SourceURL 登录来源页面URL，登录成功自动跳转过去
	SourceURL string `form:"source_url" json:"source_url" validate:"required,url"`
	// State
	State string `json:"state"`
	// LocationURL 完整的跳转URL
	LocationURL string `json:"location_url"`
}

// CustomerLoginCallbackReq 客户H5登录回调请求
type CustomerLoginCallbackReq struct {
	// Code
	Code string `form:"code" json:"Code" validate:"required"`
	// SourceURL 登录来源页面URL，登录成功自动跳转过去
	SourceURL string `form:"source_url" json:"source_url" validate:"required,url"`
	// State
	State string `form:"state" json:"state"`
}
