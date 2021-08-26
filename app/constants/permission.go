package constants

type BizIdentity string

// 业务标识
const (

	//  企业员工后台业务权限标识

	BizQuickReply        BizIdentity = "BizQuickReply"
	BizQuickReplyGroup   BizIdentity = "BizQuickReplyGroup"
	BizDepartment        BizIdentity = "BizDepartment"
	BizMassMsg           BizIdentity = "BizMassMsg"
	BizCustomerRemark    BizIdentity = "BizCustomerRemark"
	BizCustomerTag       BizIdentity = "BizCustomerTag"
	BizCustomerInfo      BizIdentity = "BizCustomerInfo"
	BizStaffInfo         BizIdentity = "BizStaffInfo"
	BizContactWay        BizIdentity = "BizContactWay"
	BizDeleteCustomer    BizIdentity = "BizDeleteCustomer"
	BizCustomerLoss      BizIdentity = "BizCustomerLoss"
	BizWelcomeMsg        BizIdentity = "BizWelcomeMsg"
	BizCustomerGroupChat BizIdentity = "BizCustomerGroupChat"
	BizMediaMgr          BizIdentity = "BizMediaMgr"
	BizMsgArch           BizIdentity = "BizMsgArch"
	BizRole              BizIdentity = "BizRole"
)

type Operation string

// 操作标识
const (
	Read Operation = "Read"
	Full Operation = "Full"
)

type Permission struct {
	BizIdentity BizIdentity
	Operation   Operation
	Name        string
}

// StaffPermissions 普通员工权限
var StaffPermissions = []Permission{
	{
		BizIdentity: BizMediaMgr,
		Operation:   Full,
		Name:        "素材管理-完全",
	},
	{
		BizIdentity: BizMediaMgr,
		Operation:   Read,
		Name:        "素材管理-查看",
	},
	{
		BizIdentity: BizCustomerGroupChat,
		Operation:   Read,
		Name:        "群聊-查看",
	},
	{
		BizIdentity: BizWelcomeMsg,
		Operation:   Read,
		Name:        "欢迎语-查看",
	},
	{
		BizIdentity: BizQuickReply,
		Operation:   Read,
		Name:        "快捷回复-查看",
	},
	{
		BizIdentity: BizQuickReplyGroup,
		Operation:   Read,
		Name:        "话术库-查看",
	},
	{
		BizIdentity: BizMassMsg,
		Operation:   Read,
		Name:        "群发消息-查看",
	},
	{
		BizIdentity: BizCustomerRemark,
		Operation:   Read,
		Name:        "客户自定义信息-查看",
	},
	{
		BizIdentity: BizCustomerRemark,
		Operation:   Full,
		Name:        "客户自定义信息-完全",
	},
	{
		BizIdentity: BizCustomerTag,
		Operation:   Read,
		Name:        "客户标签-查看",
	},
	{
		BizIdentity: BizCustomerTag,
		Operation:   Full,
		Name:        "客户标签-完全",
	},
	{
		BizIdentity: BizCustomerInfo,
		Operation:   Read,
		Name:        "客户管理-查看",
	},
	{
		BizIdentity: BizCustomerInfo,
		Operation:   Full,
		Name:        "客户管理-完全",
	},
	{
		BizIdentity: BizStaffInfo,
		Operation:   Read,
		Name:        "员工管理-查看",
	},
	{
		BizIdentity: BizContactWay,
		Operation:   Read,
		Name:        "渠道码-查看",
	},
}

// DepartmentAdminPermissions 部门管理员权限
var DepartmentAdminPermissions = append(StaffPermissions, []Permission{
	{
		BizIdentity: BizRole,
		Operation:   Read,
		Name:        "角色权限-查看",
	},
	{
		BizIdentity: BizMsgArch,
		Operation:   Full,
		Name:        "会话存档-全部",
	},
	{
		BizIdentity: BizMsgArch,
		Operation:   Read,
		Name:        "会话存档-查看",
	},
	{
		BizIdentity: BizMediaMgr,
		Operation:   Full,
		Name:        "素材管理-完全",
	},
	{
		BizIdentity: BizMediaMgr,
		Operation:   Read,
		Name:        "素材管理-查看",
	},
	{
		BizIdentity: BizCustomerGroupChat,
		Operation:   Full,
		Name:        "群聊-完全",
	},
	{
		BizIdentity: BizCustomerGroupChat,
		Operation:   Read,
		Name:        "群聊-查看",
	},
	{
		BizIdentity: BizWelcomeMsg,
		Operation:   Read,
		Name:        "欢迎语-查看",
	},
	{
		BizIdentity: BizWelcomeMsg,
		Operation:   Full,
		Name:        "欢迎语-完全",
	},
	{
		BizIdentity: BizDeleteCustomer,
		Operation:   Read,
		Name:        "删人提醒-查看",
	},
	{
		BizIdentity: BizDeleteCustomer,
		Operation:   Full,
		Name:        "删人提醒-完全",
	},
	{
		BizIdentity: BizCustomerLoss,
		Operation:   Read,
		Name:        "流失提醒-查看",
	},
	{
		BizIdentity: BizCustomerLoss,
		Operation:   Full,
		Name:        "流失提醒-完全",
	},
	{
		BizIdentity: BizQuickReply,
		Operation:   Read,
		Name:        "快捷回复-查看",
	},
	{
		BizIdentity: BizQuickReply,
		Operation:   Full,
		Name:        "快捷回复-完全",
	},
	{
		BizIdentity: BizQuickReplyGroup,
		Operation:   Read,
		Name:        "话术库-查看",
	},
	{
		BizIdentity: BizQuickReplyGroup,
		Operation:   Full,
		Name:        "话术库-完全",
	},
	{
		BizIdentity: BizDepartment,
		Operation:   Read,
		Name:        "部门管理-查看",
	},
	{
		BizIdentity: BizDepartment,
		Operation:   Full,
		Name:        "部门管理-完全",
	},
	{
		BizIdentity: BizMassMsg,
		Operation:   Read,
		Name:        "群发消息-查看",
	},
	{
		BizIdentity: BizMassMsg,
		Operation:   Full,
		Name:        "群发消息-完全",
	},
	{
		BizIdentity: BizCustomerRemark,
		Operation:   Read,
		Name:        "客户自定义信息-查看",
	},
	{
		BizIdentity: BizCustomerRemark,
		Operation:   Full,
		Name:        "客户自定义信息-完全",
	},
	{
		BizIdentity: BizCustomerTag,
		Operation:   Read,
		Name:        "客户标签-查看",
	},
	{
		BizIdentity: BizCustomerTag,
		Operation:   Full,
		Name:        "客户标签-完全",
	},
	{
		BizIdentity: BizCustomerInfo,
		Operation:   Read,
		Name:        "客户管理-查看",
	},
	{
		BizIdentity: BizCustomerInfo,
		Operation:   Full,
		Name:        "客户管理-完全",
	},
	{
		BizIdentity: BizStaffInfo,
		Operation:   Read,
		Name:        "员工管理-查看",
	},
	{
		BizIdentity: BizContactWay,
		Operation:   Read,
		Name:        "渠道码-查看",
	},
	{
		BizIdentity: BizContactWay,
		Operation:   Full,
		Name:        "渠道码-完全",
	},
}...)

// AdminPermissions 普通管理员权限
var AdminPermissions = append(DepartmentAdminPermissions, []Permission{
	{
		BizIdentity: BizRole,
		Operation:   Read,
		Name:        "角色权限-查看",
	},
	{
		BizIdentity: BizMsgArch,
		Operation:   Full,
		Name:        "会话存档-全部",
	},
	{
		BizIdentity: BizMsgArch,
		Operation:   Read,
		Name:        "会话存档-查看",
	},
	{
		BizIdentity: BizMediaMgr,
		Operation:   Full,
		Name:        "素材管理-完全",
	},
	{
		BizIdentity: BizMediaMgr,
		Operation:   Read,
		Name:        "素材管理-查看",
	},
	{
		BizIdentity: BizCustomerGroupChat,
		Operation:   Full,
		Name:        "群聊-完全",
	},
	{
		BizIdentity: BizCustomerGroupChat,
		Operation:   Read,
		Name:        "群聊-查看",
	},
	{
		BizIdentity: BizWelcomeMsg,
		Operation:   Read,
		Name:        "欢迎语-查看",
	},
	{
		BizIdentity: BizWelcomeMsg,
		Operation:   Full,
		Name:        "欢迎语-完全",
	},
	{
		BizIdentity: BizDeleteCustomer,
		Operation:   Read,
		Name:        "删人提醒-查看",
	},
	{
		BizIdentity: BizDeleteCustomer,
		Operation:   Full,
		Name:        "删人提醒-完全",
	},
	{
		BizIdentity: BizCustomerLoss,
		Operation:   Read,
		Name:        "流失提醒-查看",
	},
	{
		BizIdentity: BizCustomerLoss,
		Operation:   Full,
		Name:        "流失提醒-完全",
	},
	{
		BizIdentity: BizQuickReply,
		Operation:   Read,
		Name:        "快捷回复-查看",
	},
	{
		BizIdentity: BizQuickReply,
		Operation:   Full,
		Name:        "快捷回复-完全",
	},
	{
		BizIdentity: BizQuickReplyGroup,
		Operation:   Read,
		Name:        "话术库-查看",
	},
	{
		BizIdentity: BizQuickReplyGroup,
		Operation:   Full,
		Name:        "话术库-完全",
	},
	{
		BizIdentity: BizDepartment,
		Operation:   Read,
		Name:        "部门管理-查看",
	},
	{
		BizIdentity: BizDepartment,
		Operation:   Full,
		Name:        "部门管理-完全",
	},
	{
		BizIdentity: BizMassMsg,
		Operation:   Read,
		Name:        "群发消息-查看",
	},
	{
		BizIdentity: BizMassMsg,
		Operation:   Full,
		Name:        "群发消息-完全",
	},
	{
		BizIdentity: BizCustomerRemark,
		Operation:   Read,
		Name:        "客户自定义信息-查看",
	},
	{
		BizIdentity: BizCustomerRemark,
		Operation:   Full,
		Name:        "客户自定义信息-完全",
	},
	{
		BizIdentity: BizCustomerTag,
		Operation:   Read,
		Name:        "客户标签-查看",
	},
	{
		BizIdentity: BizCustomerTag,
		Operation:   Full,
		Name:        "客户标签-完全",
	},
	{
		BizIdentity: BizCustomerInfo,
		Operation:   Read,
		Name:        "客户管理-查看",
	},
	{
		BizIdentity: BizCustomerInfo,
		Operation:   Full,
		Name:        "客户管理-完全",
	},
	{
		BizIdentity: BizStaffInfo,
		Operation:   Read,
		Name:        "员工管理-查看",
	},
	{
		BizIdentity: BizStaffInfo,
		Operation:   Full,
		Name:        "员工管理-完全",
	},
	{
		BizIdentity: BizContactWay,
		Operation:   Read,
		Name:        "渠道码-查看",
	},
	{
		BizIdentity: BizContactWay,
		Operation:   Full,
		Name:        "渠道码-完全",
	},
}...)

// SuperAdminPermissions 超级管理员权限
var SuperAdminPermissions = append(AdminPermissions, []Permission{
	{
		BizIdentity: BizRole,
		Operation:   Full,
		Name:        "角色权限-全部",
	},
	{
		BizIdentity: BizRole,
		Operation:   Read,
		Name:        "角色权限-查看",
	},
}...)
