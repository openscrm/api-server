package tasks

import (
	"openscrm/app/models"
	"openscrm/common/log"
	"time"
)

type ContactWay struct {
	Base
}

// DailyClean 渠道码每日清理任务
func (o ContactWay) DailyClean() {
	taskKey := "ContactWayDailyClean"
	//获取分布式锁
	ok, err := o.Lock(taskKey, time.Minute)
	if err != nil {
		log.Sugar.Errorw("Lock failed", "err", err)
		return
	}
	if ok {
		defer o.Unlock(taskKey)
	}
	//每日清空渠道码关联员工添加人数统计
	err = (models.ContactWayStaff{}).DailyClean()
	if err != nil {
		log.Sugar.Errorw("DailyClean failed", "err", err)
		return
	}
}
