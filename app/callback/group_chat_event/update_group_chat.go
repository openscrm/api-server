package group_chat_event

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/we_work"
	"openscrm/conf"
	"openscrm/pkg/easywework"
	gowx "openscrm/pkg/easywework"
	"time"
)

// EventUpdateExternalChatHandler
// Description: 客户群更新事件
// Detail:
//	客户群被修改后（群名变更，群成员增加或移除，群主变更，群公告变更），回调该事件。收到该事件后，企业需要再调用获取客户群详情接口，以获取最新的群详情。
//	如果发生群信息变动，会立即收到此事件，但是部分信息是异步处理，可能需要等一段时间(例如2秒)调用获取客户群详情接口才能得到最新结果
// 	变更详情。目前有以下几种：
//		add_member : 成员入群
//		del_member : 成员退群
//		change_owner : 群主变更
//		change_name : 群名变更
//		change_notice : 群公告变更
func EventUpdateExternalChatHandler(msg *workwx.RxMessage) error {
	if msg.MsgType != gowx.MessageTypeEvent ||
		msg.Event != gowx.EventTypeChangeExternalChat ||
		msg.ChangeType != gowx.ChangeTypeUpdateChat {
		return errors.New("wrong handler for the callback event")
	}

	log.Sugar.Infow("group chat update callback", "msg", msg)
	eventCreateChat, ok := msg.EventChangeExternalChat()
	if !ok {
		return errors.New("msg.eventUpdateGroupChat failed")
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

	groupChat := models.GroupChat{
		ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: msg.ToUserID},
		ExtChatID:    eventCreateChat.GetChatID(),
		Name:         chat.Name,
		Owner:        chat.Owner,
		CreateTime:   time.Unix(int64(chat.CreateTime), 0),
		Notice:       chat.Notice,
		Total:        int64(len(chat.MemberList)),
	}

	repo := models.GroupChat{}
	err = repo.Update(groupChat)
	if err != nil {
		log.Sugar.Errorw("update group chat in callback failed", "err", err)
		return err
	}

	// 增删成员
	extCorpID := conf.Settings.WeWork.ExtCorpID
	extChatID := chat.ChatID
	updateDetail := eventCreateChat.GetUpdateDetail()
	switch updateDetail {
	case constants.GroupChatChangeTypeDelMember:
		// 成员减少, 删除不在ids中的成员
		memberIDs := make([]string, 0)
		for _, m := range chat.MemberList {
			memberIDs = append(memberIDs, m.Userid)
		}
		err = models.GroupChatMember{}.Delete(extCorpID, extChatID, memberIDs)
		if err != nil {
			return err
		}
	case constants.GroupChatChangeTypeAddMember:
		//	新增成员
		members := make([]models.GroupChatMember, 0)
		for _, m := range chat.MemberList {
			member := models.GroupChatMember{}
			err = copier.Copy(&member, m)
			if err != nil {
				return err
			}
			member.ExtChatID = extChatID
			member.Invitor = m.Invitor.Userid
			member.ID = id_generator.StringID()
			member.ExtCorpID = extCorpID
			members = append(members, member)
		}
		err = models.GroupChatMember{}.Upsert(members)
		if err != nil {
			return err
		}
	default:

	}

	// 更新群增量数据
	if eventCreateChat.GetUpdateDetail() != "add_member" || eventCreateChat.GetUpdateDetail() != "del_member" {
		err = repo.UpdateMemNum(chat.ChatID, eventCreateChat.GetUpdateDetail(), eventCreateChat.GetMemChangeCnt())
		if err != nil {
			return err
		}
	}

	return nil
}
