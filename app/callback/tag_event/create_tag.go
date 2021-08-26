package tag_event

import (
	"github.com/pkg/errors"
	"openscrm/app/services"
	"openscrm/common/log"
	gowx "openscrm/pkg/easywework"
)

// EventCreateExternalTagHandler
// Description: 企业/管理员创建客户标签/标签组时，回调此事件。收到该事件后，企业需要调用获取企业标签库来获取标签/标签组的详细信息。
// Detail:
//	Id	标签或标签组的ID
//	TagType	变更标签时，此项为tag，变更标签组时，此项为tag_group
func EventCreateExternalTagHandler(msg *gowx.RxMessage) error {
	if msg.MsgType != gowx.MessageTypeEvent ||
		msg.Event != gowx.EventTypeChangeExternalTag ||
		(msg.ChangeType != gowx.ChangeTypeCreateTag && msg.ChangeType != gowx.ChangeTypeUpdateTag) {
		return errors.New("wrong handler for the callback event")
	}

	var ID, tagType string
	if msg.ChangeType == gowx.ChangeTypeCreateTag {
		eventCreateTag, ok := msg.EventCreateTag()
		if !ok {
			return errors.New("msg.eventCreateTag failed")
		}

		ID = eventCreateTag.GetID()
		tagType = eventCreateTag.GetTagType()
		log.Sugar.Debugw("callback tag msg", "id", ID, "tagType", tagType)
	}

	tagService := services.NewTag()
	err := tagService.Sync(msg.ToUserID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return nil
}
