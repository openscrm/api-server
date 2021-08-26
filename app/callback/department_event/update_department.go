package department_event

import (
	"github.com/pkg/errors"
	"openscrm/app/models"
	"openscrm/common/id_generator"
	gowx "openscrm/pkg/easywework"
)

// EventUpdateDepartment
// Description: 更新部门事件回调
// Detail: 直接更新DB,不用再请求部门详情
func EventUpdateDepartment(msg *gowx.RxMessage) error {
	if msg.MsgType != gowx.MessageTypeEvent ||
		msg.Event != gowx.EventTypeChangeContact ||
		msg.ChangeType != gowx.ChangeTypeUpdateParty {
		return errors.New("wrong handler for the callback event")
	}
	eventUpdateParty, ok := msg.EventUpdateParty()
	if !ok {
		return errors.New("msg.EventUpdateParty failed")
	}
	department := models.Department{
		Model:       models.Model{ID: id_generator.StringID()},
		ExtCorpID:   msg.ToUserID,
		ExtID:       eventUpdateParty.GetID(),
		Name:        eventUpdateParty.GetName(),
		ExtParentID: eventUpdateParty.GetParentID(),
	}
	return models.DB.
		Where("ext_corp_id = ?", msg.ToUserID).
		Where("ext_id = ?", eventUpdateParty.GetID()).
		Updates(&department).Error
}
