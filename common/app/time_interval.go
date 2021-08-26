package app

type TimeInterval struct {
	// 开始时间
	StartTime int64 `json:"start_time" form:"start_time" validate:"omitempty,gt=0"`
	// 结束时间
	EndTime int64 `json:"end_time" form:"end_time" validate:"omitempty,gtefield=StartTime"`
}
