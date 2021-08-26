package requests

import "openscrm/common/app"

type QueryEventListReq struct {
	// 客户动态列表分类
	// customer_action 客户动态
	// integral_record 积分记录
	// manual_event 跟进记录
	// moment_interaction 朋友圈互动
	// reminder_event 提醒事件
	// template_event 模板事件
	// update_remark 修改信息
	EventType     string `json:"event_type" validate:"omitempty,oneof=customer_action integral_record manual_event moment_interaction reminder_event template_event update_remark" form:"event_type"`
	ExtStaffID    string `json:"ext_staff_id" validate:"required" form:"ext_staff_id"`
	ExtCustomerID string `json:"ext_customer_id" validate:"required" form:"ext_customer_id"`
	app.Pager     `form:"app_pager"`
	app.Sorter    `form:"app_sorter"`
}
