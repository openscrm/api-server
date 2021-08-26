package tasks

import (
	"openscrm/app/services"
	"openscrm/common/log"
	"time"
)

type Staff struct {
	Base
}

// UpdateMsgArchStatus 每小时刷新员工开通会话存档的状态
func (o Staff) UpdateMsgArchStatus() {
	taskKey := "UpdateStaffMsgArchStatus"

	ok, err := o.Lock(taskKey, time.Minute)
	if err != nil {
		log.Sugar.Errorw("Lock failed", "err", err)
		return
	}
	if ok {
		defer o.Unlock(taskKey)
	}
	//每小时刷新员工开通会话存档的状态
	err = (services.NewStaffService()).UpdateStaffMsgArchStatus()
	if err != nil {
		log.Sugar.Errorw("UpdateStaffMsgArchStatus failed", "err", err)
		return
	}
}
