package services

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/id_generator"
)

type GroupChatWelcomeMsg struct {
	GroupChatWelcomeMsgRepo models.GroupChatWelcomeMsg
}

func (m GroupChatWelcomeMsg) Create(
	req requests.CreateGroupChatWelcomeMsgReq, staff models.Staff) (msg models.GroupChatWelcomeMsg, err error) {

	msg = models.GroupChatWelcomeMsg{
		ExtCorpModel: models.ExtCorpModel{
			ID:           id_generator.StringID(),
			ExtCorpID:    staff.ExtCorpID,
			ExtCreatorID: staff.ExtID,
		},
	}
	err = copier.CopyWithOption(&msg, req, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	err = m.GroupChatWelcomeMsgRepo.Create(msg)
	return
}

func (m GroupChatWelcomeMsg) Update(
	req requests.UpdateGroupChatWelcomeMsgReq, id string, staff models.Staff) (msg models.GroupChatWelcomeMsg, err error) {

	msg = models.GroupChatWelcomeMsg{
		ExtCorpModel: models.ExtCorpModel{ID: id},
		Content:      req.Content,
	}
	err = m.GroupChatWelcomeMsgRepo.Update(msg)
	return
}

func (m GroupChatWelcomeMsg) Delete(ids []string) (int64, error) {
	return m.GroupChatWelcomeMsgRepo.Delete(ids)
}

func (m GroupChatWelcomeMsg) Query(
	req requests.QueryGroupChatWelcomeMsgReq, staff models.Staff) (msgs []models.GroupChatWelcomeMsg, total int64, err error) {

	msg := models.GroupChatWelcomeMsg{
		ExtCorpModel: models.ExtCorpModel{ExtCorpID: staff.ExtCorpID},
		Content:      req.Content,
	}
	return m.GroupChatWelcomeMsgRepo.Query(msg, staff.ExtCorpID, &req.Pager, &req.Sorter)
}

func NewGroupChatWelcomeMsg() *GroupChatWelcomeMsg {
	return &GroupChatWelcomeMsg{GroupChatWelcomeMsgRepo: models.GroupChatWelcomeMsg{}}
}
