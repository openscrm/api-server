package tag_event

import (
	"github.com/pkg/errors"
	"openscrm/app/models"
	"openscrm/app/services"
	"openscrm/common/id_generator"
	"openscrm/common/we_work"
	gowx "openscrm/pkg/easywework"
)

// EventUpdateExternalTagHandler
// Description: 当企业客户标签/标签组被修改时，回调此事件。收到该事件后，企业需要调用获取企业标签库来获取标签/标签组的详细信息。
// Detail: 同步和创建tag的回调字段相同
func EventUpdateExternalTagHandler(msg *gowx.RxMessage) error {
	if msg.MsgType != gowx.MessageTypeEvent ||
		msg.Event != gowx.EventTypeChangeExternalTag ||
		msg.ChangeType != gowx.ChangeTypeUpdateTag {
		return errors.New("wrong handler for the callback event")
	}

	var ID, tagType string
	eventEditUpdateTag, ok := msg.EventUpdateTag()
	if !ok {
		return errors.New("msg.eventEditUpdateTag failed")
	}

	ID = eventEditUpdateTag.GetID()
	tagType = eventEditUpdateTag.GetTagType()
	extCorpID := msg.ToUserID

	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	tagGroups, err := client.Customer.ListExternalContactCorpTags(ID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	if tagType == "tag" {
		tag := models.Tag{
			ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID},
			ExtID:        tagGroups[0].Tag[0].ID,
			ExtGroupID:   tagGroups[0].GroupID,
			Name:         tagGroups[0].Tag[0].Name,
			GroupName:    tagGroups[0].GroupName,
			CreateTime:   tagGroups[0].Tag[0].CreateTime,
			Order:        tagGroups[0].Tag[0].Order,
			Type:         1,
		}
		err = models.Tag{}.Upsert(tag)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
	} else if tagType == "tag_group" {
		// 暂时没有获取tagGroups 的方法,这里用同步数据的方式
		tagService := services.NewTag()
		err = tagService.Sync(extCorpID)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
	}
	return nil
}
