package msg_arch_event

import (
	"github.com/pkg/errors"
	"openscrm/pkg/easywework"
)

func EventCustomerAgreeMsgArchHandler(msg *workwx.RxMessage) error {
	if msg.MsgType != workwx.MessageTypeEvent ||
		msg.Event != workwx.EventTypeChangeExternalContact ||
		msg.ChangeType != workwx.ChangeTypeMsgAuditApproved {
		return errors.New("wrong handler for the callback event")
	}
	//todo 是添加时的回调还是同意存档时的回调
	//eventCreateChat, ok := msg.EventChangeExternalChat()
	//if !ok {
	//	return errors.New("msg.eventCreateChat failed")
	//}
	//client, err := we_work.Clients.Get(msg.ReceiverID)
	//if err != nil {
	//	return err
	//}
	//
	//groupChatInfo, err := client.Customer.GetGroupChat(workwx.ReqGetGroupChat{ExtChatID: eventCreateChat.GetChatID()})
	//if err != nil {
	//	return err
	//}
	//
	//chat := groupChatInfo.GroupChat
	//
	//groupChat := models.GroupChat{
	//	ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: msg.ReceiverID},
	//	ExtChatID:       eventCreateChat.GetChatID(),
	//	Name:         chat.Name,
	//	Owner:        chat.Owner,
	//	CreateTime:   time.Unix(int64(chat.CreateTime), 0),
	//	Notice:       chat.Notice,
	//	//MemberList:   chat.MemberList,
	//	//AdminList:    nil,
	//}
	//repo := models.GroupChat{}
	//return repo.Upsert(groupChat)
	return nil
}
