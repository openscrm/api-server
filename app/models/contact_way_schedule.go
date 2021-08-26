package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
)

// ContactWaySchedule 渠道码调度设置（根据时间自动上下线员工）
type ContactWaySchedule struct {
	ExtCorpModel
	// 渠道码ID
	ContactWayID          string `json:"contact_way_id" gorm:"index;type:bigint;comment:'渠道码ID'"`
	DailyAddCustomerLimit int    `json:"daily_add_customer_limit" gorm:"comment:'员工每日添加客户上限'"`
	// Weekdays 工作日
	Weekdays constants.StringArrayField `json:"weekdays" gorm:"json;comment:'工作日'" validate:"weekday"`
	// StartTime 开始时间
	StartTime constants.TimeField `json:"start_time" gorm:"comment:开始时间" validate:"time"`
	// EndTime 结束时间
	EndTime constants.TimeField `json:"end_time" gorm:"comment:结束时间" validate:"time"`
	// Staffs 绑定员工
	Staffs []ContactWayScheduleStaff `json:"staffs" gorm:"foreignKey:ContactWayScheduleID;references:ID"`
	Timestamp
}

// BeforeCreate 此表作为关联表时支持字段权限配置，没有更新权限的字段，不会被upsert更新
func (o ContactWaySchedule) BeforeCreate(tx *gorm.DB) (err error) {
	var colsNames []string
	// 设置冲突时需要更新的字段
	for _, field := range tx.Statement.Schema.Fields {
		if field.Updatable && field.Name != "Staffs" {
			colsNames = append(colsNames, field.DBName)
		}
	}
	tx.Statement.AddClause(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns(colsNames),
	})
	return nil
}
