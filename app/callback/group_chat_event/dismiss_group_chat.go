package group_chat_event

import (
	"github.com/pkg/errors"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/common/log"
	"openscrm/pkg/easywework"
	gowx "openscrm/pkg/easywework"
)

// EventDismissExternalChatHandler
// Description: 解散客户群聊事件
// Detail: 更新群聊状态为解散. 删除首页缓存数据
func EventDismissExternalChatHandler(msg *workwx.RxMessage) error {
	if msg.MsgType != gowx.MessageTypeEvent ||
		msg.Event != gowx.EventTypeChangeExternalChat ||
		msg.ChangeType != gowx.ChangeTypeDismissChat {
		return errors.New("wrong handler for the callback event")
	}

	log.Sugar.Infow("group chat dismiss callback", "msg", msg)

	eventCreateChat, ok := msg.EventChangeExternalChat()
	if !ok {
		return errors.New("msg.eventUpdateGroupChat failed")
	}

	err := models.GroupChat{}.Update(models.GroupChat{
		ExtCorpModel: models.ExtCorpModel{ExtCorpID: msg.ToUserID},
		ExtChatID:    eventCreateChat.GetChatID(),
		Status:       constants.GroupChatStatusIsDismissed,
	})
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	err = models.Staff{}.CleanStaffSummaryCache("", msg.ToUserID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return nil
}
