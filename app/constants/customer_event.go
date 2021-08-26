package constants

type CustomerEvent string

type EventType string

// EventType
// 事件类型，用于筛选 event_type, 有以下几个大类
const (
	CustomerEventRemainder         EventType = "reminder_event"     // 提醒事件
	CustomerEventCustomerAction              = "customer_action"    // 客户事件
	CustomerEventManualEvent                 = "manual_event"       //
	CustomerEventMomentInteraction           = "moment_interaction" // 朋友圈
	CustomerEventTemplateEvent               = "template_event"     // 模板事件
	CustomerEventUpdateRemark                = "update_remark"      // 更新备注
	CustomerEventClueManualEvent             = "clue_manual_event"  // 跟进事件
)

// EventName
// EventType 中 constomer_action 中 细分类别
const (
	EventNameAutoMerge             EventName = "AutoMerge"
	EventNameAbandon                         = "abandon"
	EventNameAddBlacklistWxUser              = "add_blacklist_wx_user"
	EventNameAddContact                      = "add_contact"
	EventNameAddCustomOrder                  = "add_custom_order"
	EventNameAddDistribute                   = "add_distribute"
	EventNameAddIntegral                     = "add_integral"
	EventNameAddManual                       = "add_manual"
	EventNameAddOcsContact                   = "add_ocs_contact"
	EventNameAddScore                        = "add_score"
	EventNameAddShopOrder                    = "add_shop_order"
	EventNameAddTags                         = "add_tags"
	EventNameAddToOpenSea                    = "add_to_open_sea"
	EventNameAddWeidianOrder                 = "add_weidian_order"
	EventNameAddXiaoeTechOrder               = "add_xiaoe_tech_order"
	EventNameAddYouzanOrder                  = "add_youzan_order"
	EventNameAfAdd                           = "af_add"
	EventNameAfEnd                           = "af_end"
	EventNameAutoMarkTagByBacNew             = "auto_mark_tag_by_bac_new"
	EventNameAutoMarkTagByBacOld             = "auto_mark_tag_by_bac_old"
	EventNameAutoMarkTagByGroup              = "auto_mark_tag_by_group"
	EventNameAutoMarkTagByKeyword            = "auto_mark_tag_by_keyword"
	EventNameAutoMarkTagByMpAfOld            = "auto_mark_tag_by_mp_af_old"
	EventNameBatchAddScore                   = "batch_add_score"
	EventNameBatchAddTags                    = "batch_add_tags"
	EventNameBatchAddTagsByGroup             = "batch_add_tags_by_group"
	EventNameBatchReduceScore                = "batch_reduce_score"
	EventNameBatchRemoveTags                 = "batch_remove_tags"
	EventNameCall                            = "call"
	EventNameCheckIn                         = "check_in"
	EventNameClueManualEvent                 = "clue_manual_event"
	EventNameDelete                          = "delete"
	EventNameDeleteExternalUser              = "delete_external_user"
	EventNameAddExternalUser                 = "add_external_user"
	EventNameDistribute                      = "distribute"
	EventNameEndChat                         = "end_chat"
	EventNameEuTransfer                      = "eu_transfer"
	EventNameExternalUserOutflow             = "external_user_outflow"
	EventNameFileOpen                        = "file_open"
	EventNameGafEnd                          = "gaf_end"
	EventNameImportDistribute                = "import_distribute"
	EventNameImportManual                    = "import_manual"
	EventNameImportToOpenSea                 = "import_to_open_sea"
	EventNameInviteUser                      = "invite_user"
	EventNameJoinGroup                       = "join_group"
	EventNameManualEvent                     = "manual_event"
	EventNameMarkTagByUnionid                = "mark_tag_by_unionid"
	EventNameMerge                           = "merge"
	EventNameModifyTags                      = "modify_tags"
	EventNameMomentOpen                      = "moment_open"
	EventNameMove                            = "move"
	EventNameOpenRadar                       = "open_radar"
	EventNameOpenSmartForm                   = "open_smart_form"
	EventNameOutboundCall                    = "outbound_call"
	EventNamePartRaffleAct                   = "part_raffle_act"
	EventNamePreTagByMobile                  = "pre_tag_by_mobile"
	EventNamePreTagByOpenapi                 = "pre_tag_by_openapi"
	EventNamePreTagByUnionid                 = "pre_tag_by_unionid"
	EventNameQuitGroup                       = "quit_group"
	EventNameRadarTag                        = "radar_tag"
	EventNameReadRadar                       = "read_radar"
	EventNameReceive                         = "receive"
	EventNameReceiveRedpacket                = "receive_redpacket"
	EventNameReceiveYouzanCoupon             = "receive_youzan_coupon"
	EventNameReduceIntegral                  = "reduce_integral"
	EventNameReduceScore                     = "reduce_score"
	EventNameReminderEvent                   = "reminder_event"
	EventNameRemoveBlacklistWxUser           = "remove_blacklist_wx_user"
	EventNameRemoveTags                      = "remove_tags"
	EventNameSendSms                         = "send_sms"
	EventNameSendYouzanCoupon                = "send_youzan_coupon"
	EventNameSmartFromMarkTag                = "smart_from_mark_tag"
	EventNameSpanAutoMarkTag                 = "span_auto_mark_tag"
	EventNameStartChat                       = "start_chat"
	EventNameSubmitSmartForm                 = "submit_smart_form"
	EventNameSyncYouzanUserTag               = "sync_youzan_user_tag"
	EventNameUpdateRemark                    = "update_remark"
	EventNameViewRaffleAct                   = "view_raffle_act"
	EventNameWonRaffleAct                    = "won_raffle_act"
	EventNameWriteSmartForm                  = "write_smart_form"
	EventNameWxMomentComment                 = "wx_moment_comment"
	EventNameWxMomentLike                    = "wx_moment_like"
)

// RemainderContent 提醒类型事件的内容
const (
	RemainderContent string = `[ %s ], 您收到一条关于客户[ %s ]的提醒事件。 提醒内容: [ %s ] `
)

const (
	RemoveTagEvent string = "%s 取消了 %s 的标签 [%s]"
	AddTagEvent           = "%s 给 %s 添加了标签 [%s]"
)
