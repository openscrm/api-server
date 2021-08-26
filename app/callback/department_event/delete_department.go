package department_event

import (
	"github.com/pkg/errors"
	"openscrm/app/models"
	gowx "openscrm/pkg/easywework"
)

// EventDeleteDepartment
// Description: 删除部门事件回调
// Detail: 删除DB, 更新tagGroup的可用部门
func EventDeleteDepartment(msg *gowx.RxMessage) error {
	if msg.MsgType != gowx.MessageTypeEvent ||
		msg.Event != gowx.EventTypeChangeContact ||
		msg.ChangeType != gowx.ChangeTypeDeleteParty {
		return errors.New("wrong handler for the callback event")
	}
	eventDeleteParty, ok := msg.EventDeleteParty()
	if !ok {
		return errors.New("msg.EventEditExternalContact failed")
	}
	err := models.DB.
		Where("ext_corp_id = ?", msg.ToUserID).
		Where("ext_id = ?", eventDeleteParty.GetID()).
		Delete(&models.Department{}).Error
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	//models.DB.Model(&models.TagGroup{}).Where("")
	return nil
}
