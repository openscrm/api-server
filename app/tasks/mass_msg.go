package tasks

import (
	"openscrm/app/services"
	"openscrm/common/log"
	"time"
)

type MassMsg struct {
	Base
}

// UpdateMassMsgStatus 更新群发消息的状态
func (o MassMsg) UpdateMassMsgStatus() {
	taskKey := "UpdateMassMsgStatus"

	ok, err := o.Lock(taskKey, time.Minute)
	if err != nil {
		log.Sugar.Errorw("Lock failed", "err", err)
		return
	}
	if ok {
		defer o.Unlock(taskKey)
	}

	// 每小时更新群发状态
	err = (services.GroupChatMassMsg{}).UpdateGroupMsgSentStatus()
	if err != nil {
		log.Sugar.Errorw("UpdateMassMsgStatus failed", "err", err)
		return
	}
}
