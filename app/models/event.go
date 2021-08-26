package models

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
)

// EventNotify 删人提醒事件通知设置
type EventNotify struct {
	// ID
	ID string `json:"id" gorm:"primaryKey;type:bigint;comment:'ID'"`
	// ExtCorpID 外部企业ID
	ExtCorpID string `json:"ext_corp_id" gorm:"uniqueIndex;type:char(18);comment:外部企业ID"`
	// EventName 事件名称
	EventName constants.EventName `json:"event_name" gorm:"type:text;comment:时间名称"`
	// 是否开启通知, 员工删除客户时提醒管理员
	IsNotifyAdmins constants.EventNotifyStatus `gorm:"type:tinyint;comment:1-打开 2-关闭" json:"is_notify_admins"`
	// 是否通知员工，员工被客户删除时提醒
	IsNotifyStaff constants.EventNotifyStatus `gorm:"type:tinyint;comment:1-打开 2-关闭" json:"is_notify_staff"`
	// 发送通知的时间
	NotifyType constants.EventNotifyTime `gorm:"comment:通知类型 1-实时 2-定时" json:"notify_type"`
	// 接收通知的管理员
	ExtStaffIDs constants.StringArrayField `gorm:"type:json;comment:接收通知的管理员" json:"ext_staff_ids"`
	Timestamp
}

// Get
// Description: 获取删人提醒设置，没找到记录则返回默认关闭
// Param: 外部企业ID
func (e EventNotify) Get(extCorpID string) (notify EventNotify, err error) {
	err = DB.Model(&EventNotify{}).Where("ext_corp_id = ?", extCorpID).First(&notify).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return EventNotify{
				IsNotifyAdmins: constants.EventNotifyStatusOff,
				IsNotifyStaff:  constants.EventNotifyStatusOff,
			}, nil
		}
		return EventNotify{}, err
	}
	if notify.IsNotifyStaff != constants.EventNotifyStatusOn {
		notify.IsNotifyStaff = constants.EventNotifyStatusOff
	}
	if notify.IsNotifyAdmins != constants.EventNotifyStatusOn {
		notify.IsNotifyAdmins = constants.EventNotifyStatusOff
	}
	return notify, nil
}

// Upsert
// Description: 更新员工删人提醒事件通知设置
func (e EventNotify) Upsert(notify EventNotify) error {
	return DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ext_corp_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"is_notify_admins", "is_notify_staff", "notify_type", "ext_staff_ids"}),
	}).Create(&notify).Error
}

// UpdateDeleteEventNotify
// Description: 设置员工被客户删除时,是否删除员工
func (e EventNotify) UpdateDeleteEventNotify(notify EventNotify) error {
	return DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ext_corp_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"is_notify_staff"}),
	}).Create(&notify).Error
}
