package entities

type UpdateNotifyStaffReq struct {
	// 是否打开删人提醒 1-是 2-否
	NotifyFlag int `json:"notify_flag"  form:"notify_flag" validate:"oneof=1 2"`
}
