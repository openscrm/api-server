package requests

import (
	"openscrm/app/constants"
)

type UpdateWelcomeMsgReq struct {
	// 标题
	Name string `json:"name" validate:"required"`
	// 主欢迎语内容
	WelcomeMsg constants.AutoReplyField `json:"welcome_msg" validate:"dive"`
	// 可用员工部门id
	ExtDepartmentIDs []int64 `json:"ext_department_ids" validate:"omitempty"`
	// 可用员工id列表
	ExtStaffIds []string `json:"ext_staff_ids" validate:"omitempty"`
	// 分时迎语内容
	TimePeriodMsg []TimePeriodMsg `json:"time_period_msg" validate:"dive"`
	// 启用分时欢迎语
	EnableTimePeriodMsg constants.Boolean `json:"enable_time_period_msg" validate:"oneof=1 2"`
}

type CreateWelcomeMsgReq struct {
	// 标题
	Name string `json:"name" validate:"required"`
	// 主欢迎语内容
	WelcomeMsg constants.AutoReplyField `json:"welcome_msg" validate:"required,dive"`
	// 可用员工组id
	ExtDepartmentIDs []int64 `json:"ext_department_ids" validate:"omitempty"`
	// 可用员工id列表,内部id,可用员工id列表和可用部门id列表均为空，则所有部门可用
	ExtStaffIds []string `json:"ext_staff_ids" validate:"omitempty"`
	// 启用分时欢迎语
	EnableTimePeriodMsg constants.Boolean `json:"enable_time_period_msg" validate:"oneof=1 2"`
	// 分时迎语内容
	TimePeriodMsg []TimePeriodMsg `json:"time_period_msg" validate:"dive"`
}

// TimePeriodMsg 分时欢迎语
type TimePeriodMsg struct {
	ID string `json:"id" form:"id" validate:"omitempty,int64"`
	// 分时欢迎语内容
	Attachments constants.AutoReplyField `json:"attachments"`
	// 生效时间,n-星期n
	EffectiveAt constants.Int64ArrayField `json:"effective_at"  validate:"gte=1,lte=7"`
	// 分时段欢迎语-开始时间
	StartTime constants.TimeField `json:"start_time"  validate:"omitempty,time"`
	// 分时段欢迎语-结束时间
	EndTime constants.TimeField `json:"end_time" validate:"omitempty,time"`
}

type DeleteReq struct {
	// 分组ID
	IDs []string `json:"ids" validate:"gt=0,dive,int64"`
}
