package tasks

import (
	"openscrm/app/models"
	"openscrm/common/log"
	"time"
)

type GroupChat struct {
	Base
}

// CleanGroupChatIncrement
// 每天删除客户群数量
func (o GroupChat) CleanGroupChatIncrement() {
	taskKey := "CleanCachedGroupChatNum"

	ok, err := o.Lock(taskKey, time.Minute)
	if err != nil {
		log.Sugar.Errorw("Lock failed", "err", err)
		return
	}
	if ok {
		defer o.Unlock(taskKey)
	}
	// 每天删除客户群数量
	err = (models.GroupChat{}).CleanGroupChatIncrement()
	if err != nil {
		log.Sugar.Errorw("UpdateMassMsgStatus failed", "err", err)
		return
	}
}
