package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"openscrm/app/callback"
	c "openscrm/app/constants"
	"openscrm/app/controller"
	m "openscrm/app/middleware"
	"openscrm/app/services"
	"openscrm/common/session"
	"openscrm/conf"
	_ "openscrm/docs"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if conf.Settings.App.Env == c.DEV {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(m.AccessLog())
		r.Use(m.Recovery())
	}

	r.Use(m.Tracing())
	//r.Use(m.Translations())
	r.Use(sessions.SessionsMany([]string{
		string(c.CustomerSessionName),
		string(c.StaffSessionName),
		string(c.StaffAdminSessionName),
		string(c.CorpAdminSessionName),
		string(c.SaasAdminSessionName),
	}, session.Store))

	// 基础服务API
	{
		// 仅开发和测试环境可用的API
		if conf.Settings.App.Env == c.DEV || conf.Settings.App.Env == c.TEST {
			r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		// 本地存储服务路由
		if conf.Settings.Storage.Type == string(c.LocalStorage) {
			storage := controller.NewStorage()
			localStorageServerPath := conf.Settings.Storage.ServerRootPath + "/*path"
			// 获取本地存储文件
			r.GET(localStorageServerPath, storage.GetLocalFile)
			// 上传本地存储文件
			r.PUT(localStorageServerPath, storage.PutLocalFile)
			// 调试接口-获取预签名URL
			r.POST("/storage/action/get-signed-url", storage.GetSignedURL)
		}

	}

	apiV1 := r.Group("/api/v1")

	tagSvc := services.NewTag()
	tag := controller.NewTag(tagSvc)

	staff := controller.NewStaff()
	customer := controller.NewCustomer()
	department := controller.NewDepartment()
	loginHandler := controller.NewLogin()
	callbackHandler := callback.NewHandler()
	util := controller.NewUtil()

	// 公开可访问的API
	{
		// 回调事件处理
		apiV1.Any("/callback", callbackHandler.HandleCallback)
	}

	//客户前台（H5）
	customerPublicApiV1 := apiV1.Group("/customer-frontend")
	{
		//公开可访问的Api
		//customerPublicApiV1.Any("/action/login", loginHandler.CustomerLogin)
		customerPublicApiV1.GET("/action/login-callback", loginHandler.CustomerLoginCallback)

		//登录后才可访问的Api
		//staffApiV1 := staffPublicApiV1.Use(m.RequireStaffLogin())

	}

	//企业员工前台（侧边栏,H5）
	staffPublicApiV1 := apiV1.Group("/staff-frontend")
	{
		//公开可访问的Api
		staffPublicApiV1.Any("/action/login", loginHandler.StaffLogin)
		staffPublicApiV1.GET("/action/login-callback", loginHandler.StaffLoginCallback)
		staffPublicApiV1.Any("/action/force-login", loginHandler.StaffForceLogin)

		//登录后才可访问的Api
		staffApiV1 := staffPublicApiV1.Use(m.RequireStaffLogin())

		// 上传临时素材
		staffApiV1.POST("/action/upload-media", util.UploadMedia)

		// 获取当前登录员工
		staffApiV1.GET("/action/get-current-staff", staff.GetCurrentFrontendStaff)

		// 获取ticket
		jsApiHandler := controller.NewJsApiHandler()
		staffApiV1.Any("/action/get-js-config", jsApiHandler.GetJsConfig)
		staffApiV1.Any("/action/get-js-agent-config", jsApiHandler.GetJsAgentConfig)

		// 客户画像
		customerFrontendHandler := controller.NewCustomerFrontend()
		staffApiV1.GET("/customer/:ext_id", customerFrontendHandler.Get)

		// 侧边栏-提醒
		remainderHandler := controller.NewRemainderFrontend()
		staffApiV1.POST("/customer/remainder", remainderHandler.Create)
		staffApiV1.POST("/customer/remainder/action/delete", remainderHandler.Delete)
		staffApiV1.PUT("/customer/remainder/:id", remainderHandler.Update)

		// 侧边栏-跟进
		clueManualHandler := controller.NewClueManualFrontend()
		staffApiV1.POST("/customer/clue-manual", clueManualHandler.Create)
		staffApiV1.POST("/customer/clue-manual/action/delete", clueManualHandler.Delete)
		staffApiV1.PUT("/customer/clue-manual/:id", clueManualHandler.Update)

		// 话术库
		quickReplyFrontend := controller.NewQuickReplyFrontend()
		staffApiV1.GET("/quick-replies", quickReplyFrontend.Query)

		// 话术库分组
		quickReplyGroupFrontend := controller.NewQuickReplyGroupFrontend()
		staffApiV1.GET("/quick-reply-groups", quickReplyGroupFrontend.Query)
		staffApiV1.POST("/quick-reply-group/action/search", quickReplyGroupFrontend.Search)
		staffApiV1.POST("/quick-reply-group/action/delete", quickReplyGroupFrontend.Delete)
		staffApiV1.PUT("/quick-reply-group", quickReplyGroupFrontend.Update)

		// 素材
		materialLibFrontend := controller.NewMaterialLibFrontend()
		staffApiV1.GET("/material/lib", materialLibFrontend.Query)

		// 素材标签
		materialLibTagsFrontend := controller.NewMaterialLibTagFrontend()
		staffApiV1.GET("/material/lib/tags", materialLibTagsFrontend.Query)
	}

	//企业普通管理员后台
	staffAdminPublicApiV1 := apiV1.Group("/staff-admin")
	{
		// 公开可访问的Api
		staffAdminPublicApiV1.Any("/action/login", loginHandler.StaffAdminLogin)
		staffAdminPublicApiV1.GET("/action/login-callback", loginHandler.StaffAdminLoginCallback)
		staffAdminPublicApiV1.POST("/action/force-login", loginHandler.StaffAdminForceLogin)

		//登录后才可访问的Api
		staffAdminApiV1 := staffAdminPublicApiV1.Use(m.RequireStaffAdminLogin())

		utilHandler := controller.NewUtil()
		staffAdminApiV1.POST("/common/action/parse-link", utilHandler.ParseLink)
		staffAdminApiV1.POST("/common/action/get-signed-url", utilHandler.GetUploadURL)

		contactWayGroupHandler := controller.NewContactWayGroup()
		staffAdminApiV1.GET("/contact-way-groups", m.Guard(c.BizContactWay, c.Read), contactWayGroupHandler.Query)
		staffAdminApiV1.GET("/contact-way-group/:id", m.Guard(c.BizContactWay, c.Read), contactWayGroupHandler.Get)
		staffAdminApiV1.POST("/contact-way-group", m.Guard(c.BizContactWay, c.Full), contactWayGroupHandler.Create)
		staffAdminApiV1.PUT("/contact-way-group/:id", m.Guard(c.BizContactWay, c.Full), contactWayGroupHandler.Update)
		staffAdminApiV1.POST("/contact-way-group/action/delete", m.Guard(c.BizContactWay, c.Full), contactWayGroupHandler.Delete)

		contactWayHandler := controller.NewContactWay()
		staffAdminApiV1.GET("/contact-ways", m.Guard(c.BizContactWay, c.Read), contactWayHandler.Query)
		staffAdminApiV1.GET("/contact-way/:id", m.Guard(c.BizContactWay, c.Read), contactWayHandler.Get)
		staffAdminApiV1.POST("/contact-way", m.Guard(c.BizContactWay, c.Full), contactWayHandler.Create)
		staffAdminApiV1.PUT("/contact-way/:id", m.Guard(c.BizContactWay, c.Full), contactWayHandler.Update)
		staffAdminApiV1.POST("/contact-way/action/delete", m.Guard(c.BizContactWay, c.Full), contactWayHandler.Delete)
		staffAdminApiV1.POST("/contact-way/action/batch-update", m.Guard(c.BizContactWay, c.Full), contactWayHandler.BatchUpdate)

		// 企业管理-部门
		staffAdminApiV1.POST("/department", m.Guard(c.BizDepartment, c.Full), department.Sync)
		staffAdminApiV1.GET("/department", m.Guard(c.BizDepartment, c.Read), department.Get)
		staffAdminApiV1.GET("/departments", m.Guard(c.BizDepartment, c.Read), department.Query)

		// 企业管理-员工
		staffAdminApiV1.POST("/staff", m.Guard(c.BizStaffInfo, c.Full), staff.Sync)
		staffAdminApiV1.GET("/staffs", m.Guard(c.BizStaffInfo, c.Read), staff.Query)
		staffAdminApiV1.GET("/staff/action/get-all", m.Guard(c.BizStaffInfo, c.Read), staff.QueryMainInfo)
		staffAdminApiV1.PUT("/staff", m.Guard(c.BizStaffInfo, c.Full), staff.Update)
		staffAdminApiV1.GET("/staff/:ext-staff-id", m.Guard(c.BizStaffInfo, c.Read), staff.Get)
		staffAdminApiV1.GET("/statistics/:ext-staff-id", m.Guard(c.BizStaffInfo, c.Read), staff.GetStaffStaticsInfo)
		staffAdminApiV1.POST("/customer/action/update-tags", m.Guard(c.BizCustomerTag, c.Full), staff.UpdateCustomerTag)
		staffAdminApiV1.POST("/customer/action/update-internal-tags", m.Guard(c.BizCustomerTag, c.Full), staff.UpdateCustomerInternalTag)
		staffAdminApiV1.GET("/staff/action/delete-customers-data-export", staff.ExportDeleteCustomers)

		// 客户管理
		staffAdminApiV1.POST("/customer/action/sync", m.Guard(c.BizCustomerInfo, c.Full), customer.Sync)
		//staffAdminApiV1.POST("/customer/async", m.Guard(c.BizCustomerInfo, c.Full), customer.Syncs)
		staffAdminApiV1.GET("/customer/:ext_id", m.Guard(c.BizCustomerInfo, c.Read), customer.Get)
		staffAdminApiV1.GET("/customers", m.Guard(c.BizCustomerInfo, c.Read), customer.Query)
		staffAdminApiV1.GET("/customers/action/export", m.Guard(c.BizCustomerInfo, c.Read), customer.Export)
		staffAdminApiV1.GET("/customers/statistic", m.Guard(c.BizCustomerInfo, c.Read), customer.Statistic)

		homePageHandler := controller.NewHomePageHandler()
		staffAdminApiV1.GET("/action/get-summary", m.Guard(c.BizCustomerInfo, c.Full), homePageHandler.GetCustomerSummary)
		staffAdminApiV1.GET("/action/get-trend", m.Guard(c.BizCustomerInfo, c.Full), homePageHandler.GetCustomersTrend)

		// 侧边栏-提醒
		remainder := controller.NewRemainder()
		staffAdminApiV1.POST("/customer/remainder", m.Guard(c.BizCustomerInfo, c.Full), remainder.Create)
		staffAdminApiV1.POST("/customer/remainder/action/delete", m.Guard(c.BizCustomerInfo, c.Full), remainder.Delete)
		staffAdminApiV1.PUT("/customer/remainder/:id", m.Guard(c.BizCustomerInfo, c.Full), remainder.Update)

		// 侧边栏-跟进
		cm := controller.NewClueManual()
		staffAdminApiV1.POST("/customer/clue-manual", m.Guard(c.BizCustomerInfo, c.Full), cm.Create)
		staffAdminApiV1.POST("/customer/clue-manual/action/delete", m.Guard(c.BizCustomerInfo, c.Full), cm.Delete)
		staffAdminApiV1.PUT("/customer/clue-manual/:id", m.Guard(c.BizCustomerInfo, c.Full), cm.Update)

		// 客户管理-事件列表
		customerEventHandler := controller.NewCustomerEvent()
		staffAdminApiV1.GET("/customer/events", m.Guard(c.BizCustomerInfo, c.Read), customerEventHandler.Query)

		customerInfoHandler := controller.NewCustomerInfoHandler()
		staffAdminApiV1.GET("/customer/info", m.Guard(c.BizCustomerInfo, c.Read), customerInfoHandler.Get)
		staffAdminApiV1.PUT("/customer/info", m.Guard(c.BizCustomerInfo, c.Full), customerInfoHandler.Update)

		rulesHandler := controller.NewCustomerInfoDisplayRules()
		staffAdminApiV1.GET("/customer/info/displays", m.Guard(c.BizCustomerInfo, c.Read), rulesHandler.Get)
		staffAdminApiV1.PUT("/customer/info/displays", m.Guard(c.BizCustomerInfo, c.Full), rulesHandler.Update)

		// 客户管理-被删提醒
		// 获取被删提醒规则
		staffAdminApiV1.GET("/customer/action/get-loss-notify-rule", m.Guard(c.BizCustomerLoss, c.Read), customer.GetNotifyStaffRule)
		// 设置被删提醒规则
		staffAdminApiV1.POST("/customer/action/update-loss-notify-rule", m.Guard(c.BizCustomerLoss, c.Full), customer.UpdateNotifyStaffRule)
		// 获取企业流失客户记录
		staffAdminApiV1.GET("/customer/losses", m.Guard(c.BizCustomerLoss, c.Read), customer.GetLossCustomers)
		// 流失提醒记录下载
		staffAdminApiV1.GET("/customer/action/customers-losses-data-export", m.Guard(c.BizCustomerLoss, c.Read), customer.ExportCustomerLosses)

		// 企微客户标签
		staffAdminApiV1.POST("/customer/tag/action/sync", m.Guard(c.BizCustomerTag, c.Full), tag.Sync)
		staffAdminApiV1.POST("/customer/tag", m.Guard(c.BizCustomerTag, c.Full), tag.Create)

		// 内部客户标签
		internalTag := controller.NewInternalTag()
		staffAdminApiV1.POST("/customer/internal-tag", m.Guard(c.BizCustomerTag, c.Full), internalTag.Create)
		staffAdminApiV1.GET("/customer/internal-tags", m.Guard(c.BizCustomerTag, c.Read), internalTag.Query)
		staffAdminApiV1.POST("/customer/internal-tag/action/delete", m.Guard(c.BizCustomerTag, c.Full), internalTag.Delete)

		tagGroupHandler := controller.NewTagGroup()
		staffAdminApiV1.POST("/customer/tag-group", m.Guard(c.BizCustomerTag, c.Full), tagGroupHandler.Create)
		staffAdminApiV1.GET("/customer/tag-groups", m.Guard(c.BizCustomerTag, c.Read), tagGroupHandler.Query)
		staffAdminApiV1.PUT("/customer/tag-group/:ext_id", m.Guard(c.BizCustomerTag, c.Read), tagGroupHandler.Update)
		staffAdminApiV1.POST("/customer/tag-group/action/exchange-order", m.Guard(c.BizCustomerTag, c.Full), tagGroupHandler.ExchangeOrder)
		staffAdminApiV1.POST("/customer/tag-group/action/delete", m.Guard(c.BizCustomerTag, c.Full), tagGroupHandler.Delete)

		// 客户自定义信息
		remarkHandler := controller.NewCustomerRemarkHandler()
		staffAdminApiV1.POST("/customer/remark", m.Guard(c.BizCustomerRemark, c.Full), remarkHandler.Create)
		staffAdminApiV1.GET("/customer/remark", m.Guard(c.BizCustomerRemark, c.Read), remarkHandler.Get)
		staffAdminApiV1.PUT("/customer/remark", m.Guard(c.BizCustomerRemark, c.Full), remarkHandler.Update)
		staffAdminApiV1.PUT("/customer/remark/action/exchange-order", m.Guard(c.BizCustomerRemark, c.Full), remarkHandler.ExchangeOrder)
		staffAdminApiV1.POST("/customer/remark/action/delete", m.Guard(c.BizCustomerRemark, c.Full), remarkHandler.Delete)
		staffAdminApiV1.POST("/customer/remark/option", m.Guard(c.BizCustomerRemark, c.Full), remarkHandler.AddRemarkOption)
		staffAdminApiV1.PUT("/customer/remark/option", m.Guard(c.BizCustomerRemark, c.Full), remarkHandler.UpdateRemarkOption)
		staffAdminApiV1.POST("/customer/remark/option/action/delete", m.Guard(c.BizCustomerRemark, c.Full), remarkHandler.DeleteRemarkOption)

		// 企业素材库
		materialLib := controller.NewMaterialLib()
		staffAdminApiV1.POST("/material/lib", m.Guard(c.BizCustomerInfo, c.Full), materialLib.Create)
		staffAdminApiV1.POST("/material/lib/action/delete", m.Guard(c.BizCustomerInfo, c.Full), materialLib.Delete)
		staffAdminApiV1.PUT("/material/lib/:id", m.Guard(c.BizCustomerInfo, c.Full), materialLib.Update)
		staffAdminApiV1.GET("/material/libs", m.Guard(c.BizCustomerInfo, c.Read), materialLib.Query)
		staffAdminApiV1.GET("/material/lib/sidebar-status", m.Guard(c.BizCustomerInfo, c.Read), materialLib.GetSidebarStatus)
		staffAdminApiV1.PUT("/material/lib/sidebar-status", m.Guard(c.BizCustomerInfo, c.Full), materialLib.UpdateSidebarStatus)

		// 素材库标签管理
		materialLibTag := controller.NewMaterialLibTag()
		staffAdminApiV1.POST("/material/lib/tag", m.Guard(c.BizCustomerInfo, c.Full), materialLibTag.Create)
		staffAdminApiV1.GET("/material/lib/tags", m.Guard(c.BizCustomerInfo, c.Read), materialLibTag.Query)
		staffAdminApiV1.POST("/material/lib/tag/action/delete", m.Guard(c.BizCustomerInfo, c.Full), materialLibTag.Delete)

		// 客户转化-群发消息
		massMsgHandler := controller.NewDefaultMassMsg()
		staffAdminApiV1.POST("/customer/mass-msg", m.Guard(c.BizMassMsg, c.Full), massMsgHandler.Create)
		staffAdminApiV1.POST("/customer/mass-msg/action/delete", m.Guard(c.BizMassMsg, c.Full), massMsgHandler.Delete)
		staffAdminApiV1.GET("/customer/mass-msg/:id", m.Guard(c.BizMassMsg, c.Read), massMsgHandler.Get)
		staffAdminApiV1.PUT("/customer/mass-msg/:id", m.Guard(c.BizQuickReply, c.Full), massMsgHandler.Update)
		staffAdminApiV1.GET("/customer/mass-msgs", m.Guard(c.BizMassMsg, c.Read), massMsgHandler.Query)
		staffAdminApiV1.POST("/customer/mass-msg/action/notify", m.Guard(c.BizMassMsg, c.Full), massMsgHandler.Notify)
		staffAdminApiV1.GET("/customer/mass-msg/result/:id", m.Guard(c.BizMassMsg, c.Read), massMsgHandler.GetSendMassMsgResult)
		staffAdminApiV1.GET("/customer/mass-msg/customer-filter", m.Guard(c.BizMassMsg, c.Read), massMsgHandler.CustomerFilter)
		staffAdminApiV1.POST("/customer/mass-msg/action/get-upload-url", m.Guard(c.BizQuickReply, c.Full), massMsgHandler.GetUploadUrl)

		//// 客户群-群发
		groupChatMassMsgHandler := controller.NewDefaultGroupChatMassMsg()
		staffAdminApiV1.POST("/group-chat/mass-msg", m.Guard(c.BizMassMsg, c.Full), groupChatMassMsgHandler.Create)
		staffAdminApiV1.POST("/group-chat/mass-msg/action/delete", m.Guard(c.BizMassMsg, c.Full), groupChatMassMsgHandler.Delete)
		staffAdminApiV1.GET("/group-chat/mass-msg/:id", m.Guard(c.BizMassMsg, c.Read), groupChatMassMsgHandler.Get)
		//staffAdminApiV1.PUT("/group-chat/mass-msg/:id", m.Guard(c.BizQuickReply, c.Full), groupChatMassMsgHandler.Update)
		staffAdminApiV1.GET("/group-chat/mass-msgs", m.Guard(c.BizMassMsg, c.Read), groupChatMassMsgHandler.Query)
		//staffAdminApiV1.POST("/group-chat/mass-msg/action/notify", m.Guard(c.BizMassMsg, c.Full), groupChatMassMsgHandler.Notify)
		//staffAdminApiV1.GET("/group-chat/mass-msg/result/:id", m.Guard(c.BizMassMsg, c.Read), groupChatMassMsgHandler.GetSendMassMsgResult)
		//staffAdminApiV1.GET("/group-chat/mass-msg/customer-filter", m.Guard(c.BizMassMsg, c.Read), groupChatMassMsgHandler.CustomerFilter)

		// 客户转化-欢迎语
		welcomeMsgHandler := controller.NewWelComeMsg()
		staffAdminApiV1.POST("/customer/welcome-msg", m.Guard(c.BizWelcomeMsg, c.Full), welcomeMsgHandler.Create)
		staffAdminApiV1.GET("/customer/welcome-msgs", m.Guard(c.BizWelcomeMsg, c.Read), welcomeMsgHandler.Query)
		staffAdminApiV1.GET("/customer/welcome-msg/:id", m.Guard(c.BizWelcomeMsg, c.Read), welcomeMsgHandler.Get)
		staffAdminApiV1.PUT("/customer/welcome-msg/:id", m.Guard(c.BizWelcomeMsg, c.Full), welcomeMsgHandler.Update)
		staffAdminApiV1.POST("/customer/welcome-msg/action/delete", m.Guard(c.BizWelcomeMsg, c.Full), welcomeMsgHandler.Delete)
		staffAdminApiV1.POST("/customer/welcome-msg/action/upload-image", m.Guard(c.BizWelcomeMsg, c.Full), welcomeMsgHandler.UploadFileUrl)

		// 话术库分组
		quickReplyGroupHandler := controller.NewQuickReplyGroup()
		staffAdminApiV1.POST("/quick-reply-group", m.Guard(c.BizQuickReplyGroup, c.Full), quickReplyGroupHandler.Create)
		staffAdminApiV1.PUT("/quick-reply-group/:id", m.Guard(c.BizQuickReplyGroup, c.Full), quickReplyGroupHandler.Update)
		staffAdminApiV1.GET("/quick-reply-groups", m.Guard(c.BizQuickReplyGroup, c.Read), quickReplyGroupHandler.Get) // 所有组的层级关系
		staffAdminApiV1.POST("/quick-reply-group/action/delete", m.Guard(c.BizQuickReplyGroup, c.Full), quickReplyGroupHandler.Delete)

		// 话术库
		quickReply := controller.NewQuickReply()
		staffAdminApiV1.POST("/quick-reply", m.Guard(c.BizQuickReply, c.Full), quickReply.Create)
		staffAdminApiV1.PUT("/quick-reply", m.Guard(c.BizQuickReply, c.Read), quickReply.Update)
		staffAdminApiV1.GET("/quick-replies", m.Guard(c.BizQuickReply, c.Read), quickReply.Query)
		staffAdminApiV1.POST("/quick-reply/action/delete", m.Guard(c.BizQuickReply, c.Read), quickReply.Delete)
		staffAdminApiV1.POST("/quick-reply/action/get-upload-url", m.Guard(c.BizQuickReply, c.Full), quickReply.GetUploadURL)

		// 企业风控
		corpRiskMgr := controller.NewCorpRiskMgr()
		// 设置企业删人提醒的通知规则
		staffAdminApiV1.PUT("/notify/delete-customer/status", m.Guard(c.BizDeleteCustomer, c.Full), corpRiskMgr.Update)
		// 获取企业删人提醒的通知规则
		staffAdminApiV1.GET("/notify/delete-customer/status", m.Guard(c.BizDeleteCustomer, c.Read), corpRiskMgr.Get)
		// 获取删客户的列表
		staffAdminApiV1.GET("/notify/delete-customers", m.Guard(c.BizDeleteCustomer, c.Read), corpRiskMgr.Query)

		// 客户群管理
		CustomerGroupChatHandler := controller.NewGroupChat()
		staffAdminApiV1.GET("/group-chats", m.Guard(c.BizCustomerGroupChat, c.Read), CustomerGroupChatHandler.Query)
		staffAdminApiV1.GET("/group-chat/action/export", m.Guard(c.BizCustomerGroupChat, c.Read), CustomerGroupChatHandler.Export)
		staffAdminApiV1.GET("/group-chat/owners", m.Guard(c.BizCustomerGroupChat, c.Read), CustomerGroupChatHandler.GetAllOwners)
		staffAdminApiV1.POST("/group-chat/action/get-all", m.Guard(c.BizCustomerGroupChat, c.Read), CustomerGroupChatHandler.GetAll)
		staffAdminApiV1.POST("/group-chat/action/update-tags", m.Guard(c.BizCustomerGroupChat, c.Read), CustomerGroupChatHandler.UpdateTags)

		// 客户群标签
		CustomerGroupTagHandler := controller.NewGroupChatTag()
		staffAdminApiV1.POST("/group-chat/tag", m.Guard(c.BizCustomerGroupChat, c.Full), CustomerGroupTagHandler.Create)
		staffAdminApiV1.PUT("/group-chat/tag", m.Guard(c.BizCustomerGroupChat, c.Full), CustomerGroupTagHandler.Update)

		// 客户群标签组
		CustomerGroupTagGroupHandler := controller.NewGroupChatTagGroup()
		staffAdminApiV1.POST("/group-chat/tag-group", m.Guard(c.BizCustomerGroupChat, c.Full), CustomerGroupTagGroupHandler.Create)
		staffAdminApiV1.PUT("/group-chat/tag-group", m.Guard(c.BizCustomerGroupChat, c.Full), CustomerGroupTagGroupHandler.Update)
		staffAdminApiV1.GET("/group-chat/tag-groups", m.Guard(c.BizCustomerGroupChat, c.Read), CustomerGroupTagGroupHandler.Query)
		staffAdminApiV1.POST("/group-chat/tag-group/action/delete", m.Guard(c.BizCustomerGroupChat, c.Full), CustomerGroupTagGroupHandler.Delete)

		// 客户群管理-自动拉群-分组
		groupChatGroupHandler := controller.NewGroupChatGroup()
		staffAdminApiV1.POST("/group-chat/auto-join/group", m.Guard(c.BizCustomerGroupChat, c.Full), groupChatGroupHandler.Create)
		staffAdminApiV1.GET("/group-chat/auto-join/groups", m.Guard(c.BizCustomerGroupChat, c.Read), groupChatGroupHandler.Query)
		staffAdminApiV1.PUT("/group-chat/auto-join/group/:id", m.Guard(c.BizCustomerGroupChat, c.Full), groupChatGroupHandler.Update)
		staffAdminApiV1.POST("/group-chat/auto-join/action/delete", m.Guard(c.BizCustomerGroupChat, c.Full), groupChatGroupHandler.Delete)

		// 客户群管理-自动拉群
		groupChatAutoJoinHandler := controller.NewGroupChatAutoJoin()
		staffAdminApiV1.POST("/group-chat/auto-join-code", m.Guard(c.BizCustomerGroupChat, c.Full), groupChatAutoJoinHandler.Create)
		staffAdminApiV1.PUT("/group-chat/auto-join-code/:id", m.Guard(c.BizCustomerGroupChat, c.Full), groupChatAutoJoinHandler.Update)
		staffAdminApiV1.GET("/group-chat/auto-join-codes", m.Guard(c.BizCustomerGroupChat, c.Read), groupChatAutoJoinHandler.Query)
		staffAdminApiV1.POST("/group-chat/auto-join-code/action/delete", m.Guard(c.BizCustomerGroupChat, c.Full), groupChatAutoJoinHandler.Delete)
		staffAdminApiV1.POST("/group-chat/auto-join-code/action/batch-regroup", m.Guard(c.BizCustomerGroupChat, c.Full), groupChatAutoJoinHandler.BatchRegroup)

		// 客户群管理-入群欢迎语
		groupChatWelcomeMsgHandler := controller.NewGroupChatWelcomeMsg()
		staffAdminApiV1.POST("/group-chat/welcome-msg", m.Guard(c.BizCustomerGroupChat, c.Full), groupChatWelcomeMsgHandler.Create)
		staffAdminApiV1.PUT("/group-chat/welcome-msg/:id", m.Guard(c.BizCustomerGroupChat, c.Full), groupChatWelcomeMsgHandler.Update)
		staffAdminApiV1.GET("/group-chat/welcome-msgs", m.Guard(c.BizCustomerGroupChat, c.Read), groupChatWelcomeMsgHandler.Query)
		staffAdminApiV1.POST("/group-chat/welcome-msg/action/delete", m.Guard(c.BizCustomerGroupChat, c.Full), groupChatWelcomeMsgHandler.Delete)

		// 聊天记录
		msgArchHandler := controller.NewMsgArch()
		staffAdminApiV1.POST("/chat-msg/sync", m.Guard(c.BizMsgArch, c.Full), msgArchHandler.Sync)
		staffAdminApiV1.GET("/customer/chat-sessions", m.Guard(c.BizMsgArch, c.Read), msgArchHandler.QuerySessions)
		staffAdminApiV1.GET("/customer/session-msgs", m.Guard(c.BizMsgArch, c.Read), msgArchHandler.QuerySessionMsgs)
		staffAdminApiV1.POST("/customer/session-msg/action/search", m.Guard(c.BizMsgArch, c.Read), msgArchHandler.SearchMsgs)

		//权限信息
		permissionHandler := controller.NewPermission()
		staffAdminApiV1.GET("/permissions", m.Guard(c.BizRole, c.Read), permissionHandler.Query)
		staffAdminApiV1.GET("/permission/:id", m.Guard(c.BizRole, c.Read), permissionHandler.Get)

		//角色管理
		roleHandler := controller.NewRole()
		staffAdminApiV1.GET("/roles", m.Guard(c.BizRole, c.Read), roleHandler.Query)
		staffAdminApiV1.GET("/role/:id", m.Guard(c.BizRole, c.Read), roleHandler.Get)
		staffAdminApiV1.POST("/role", m.Guard(c.BizRole, c.Full), roleHandler.Create)
		staffAdminApiV1.PUT("/role/:id", m.Guard(c.BizRole, c.Full), roleHandler.Update)
		staffAdminApiV1.POST("/role/action/assign-to-staffs", m.Guard(c.BizRole, c.Full), roleHandler.AssignToStaffs)
		staffAdminApiV1.GET("/role/action/query-staffs", m.Guard(c.BizRole, c.Read), roleHandler.QueryStaffs)

		// 获取当前登录员工
		staffAdminApiV1.GET("/action/get-current-staff", staff.GetCurrent)

	}

	return r
}
