package tag_event

import (
	"github.com/pkg/errors"
	"openscrm/app/models"
	gowx "openscrm/pkg/easywework"
)

// EventDeleteExternalTagHandler
// Description: 删除客户标签/标签组组回调
// Detail: 同时删除已经对客户使用的标签
func EventDeleteExternalTagHandler(msg *gowx.RxMessage) error {
	if msg.MsgType != gowx.MessageTypeEvent ||
		msg.Event != gowx.EventTypeChangeExternalTag ||
		msg.ChangeType != gowx.ChangeTypeDeleteTag {
		return errors.New("wrong handler for the callback event")
	}

	eventEditCreateTag, ok := msg.EventDeleteTag()
	if !ok {
		return errors.New("msg.eventEditCreateTag failed")
	}

	ID := eventEditCreateTag.GetID()
	tagType := eventEditCreateTag.GetTagType()

	extTagIds := make([]string, 0)
	if tagType == "tag_group" {
		tags, err := (models.Tag{}).GetExistTag(ID, nil)
		if err != nil {
			return err
		}

		// 删除标签组
		_, err = (models.TagGroup{}).Delete([]string{ID})
		if err != nil {
			return err
		}

		for _, tag := range tags {
			extTagIds = append(extTagIds, tag.ExtID)
		}
	} else if tagType == "tag" {
		err := (models.Tag{}).Delete([]string{ID})
		if err != nil {
			return err
		}
		extTagIds = append(extTagIds, ID)
	}

	return models.CustomerStaffTag{}.Delete("", extTagIds, true)
}
