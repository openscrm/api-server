package tasks

import (
	"github.com/gogf/gf/os/gcron"
	"openscrm/common/log"
)

//Start 用于运行定时任务
func Start() {
	var err error
	defer func() {
		if err := recover(); err != nil {
			log.Sugar.Errorw("task panic error", "err", err)
		}
	}()

	_, err = gcron.AddSingleton("@daily", (ContactWay{}).DailyClean, "ContactWayDailyClean")
	if err != nil {
		log.Sugar.Errorw("AddSingleton failed", "err", err)
	}

	_, err = gcron.AddSingleton("@hourly", (Staff{}).UpdateMsgArchStatus, "UpdateStaffMsgArchStatus")
	if err != nil {
		log.Sugar.Errorw("AddSingleton failed", "err", err)
	}

	_, err = gcron.AddSingleton("@hourly", (MassMsg{}).UpdateMassMsgStatus, "UpdateMassMsgStatus")
	if err != nil {
		log.Sugar.Errorw("AddSingleton failed", "err", err)
	}

	_, err = gcron.AddSingleton("@daily", (GroupChat{}).CleanGroupChatIncrement, "CleanGroupChatIncrement")
	if err != nil {
		log.Sugar.Errorw("AddSingleton failed", "err", err)
	}

	log.Sugar.Infow("Tasks Running")
}
