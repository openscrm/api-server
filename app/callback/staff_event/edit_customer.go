package staff_event

import (
	"github.com/pkg/errors"
	"openscrm/common/log"
	gowx "openscrm/pkg/easywework"
)

// EventEditExternalContactHandler
// Description: 更新客户数据回调事件
func EventEditExternalContactHandler(msg *gowx.RxMessage) error {
	if msg.MsgType != gowx.MessageTypeEvent ||
		msg.Event != gowx.EventTypeChangeExternalContact ||
		msg.ChangeType != gowx.ChangeTypeEditExternalContact {
		return errors.New("wrong handler for the callback event")
	}

	eventEditExternalContact, ok := msg.EventEditExternalContact()
	if !ok {
		return errors.New("msg.EventEditExternalContact failed")
	}
	extStaffID := eventEditExternalContact.GetUserID()
	extCustomerID := eventEditExternalContact.GetExternalUserID()

	log.Sugar.Debugw("sync customer data")

	return SyncExternalContact(extStaffID, extCustomerID)
}
