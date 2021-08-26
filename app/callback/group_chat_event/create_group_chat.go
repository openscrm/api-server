package group_chat_event

import (
	"github.com/pkg/errors"
	"openscrm/app/models"
	"openscrm/app/services"
	"openscrm/common/we_work"
	"openscrm/conf"
	"openscrm/pkg/easywework"
	gowx "openscrm/pkg/easywework"
)

// EventCreateExternalChatHandler
// Description: 创建含有外部联系人的群聊事件回调
// Detail:
//	群中至少有一个外部联系人
//	更新今日入群数
func EventCreateExternalChatHandler(msg *workwx.RxMessage) error {
	if msg.MsgType != gowx.MessageTypeEvent ||
		msg.Event != gowx.EventTypeChangeExternalChat ||
		msg.ChangeType != gowx.ChangeTypeCreateChat {
		return errors.New("wrong handler for the callback event")
	}

	eventCreateChat, ok := msg.EventChangeExternalChat()
	if !ok {
		return errors.New("msg.eventCreateChat failed")
	}
	client, err := we_work.Clients.Get(msg.ToUserID)
	if err != nil {
		return err
	}

	groupChatInfo, err := client.Customer.GetGroupChat(gowx.GetGroupChatReq{ChatId: eventCreateChat.GetChatID()})
	if err != nil {
		return err
	}

	chat := groupChatInfo.GroupChat

	groupChatService := services.NewGroupChatService()
	extCorpID := conf.Settings.WeWork.ExtCorpID
	err = groupChatService.Syncs(chat.ChatID, extCorpID)
	if err != nil {
		return err
	}

	// 更新今日入群人数
	// 更新群增量数据
	repo := models.GroupChat{}
	return repo.UpdateMemNum(chat.ChatID, "add_member", int64(len(chat.MemberList)))
}
