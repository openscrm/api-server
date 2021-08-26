package services

import (
	"github.com/pkg/errors"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/app"
)

type QuickReplyFrontend struct {
	quickReplyRepo      models.QuickReply
	quickReplyGroupRepo models.QuickReplyGroup
}

func NewQuickReplyFrontend() *QuickReplyFrontend {
	return &QuickReplyFrontend{quickReplyRepo: models.QuickReply{}}
}

func (o QuickReplyFrontend) QueryQuickReply(
	req requests.QueryQuickReplyReq, extCorpID string, pager *app.Pager) (groups []models.QuickReplyGroup, total int64, err error) {

	groups = make([]models.QuickReplyGroup, 0)

	replies, err := o.quickReplyRepo.QueryByKeyword(req.Keyword, extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	groupIDs := make([]string, 0)
	repliesIDs := make([]string, 0)
	for _, reply := range replies {
		groupIDs = append(groupIDs, reply.GroupID)
		repliesIDs = append(repliesIDs, reply.ID)
	}

	groups, err = o.quickReplyGroupRepo.QueryByID(extCorpID, groupIDs, repliesIDs)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	groupsByQuickReplyGroup, err := o.quickReplyGroupRepo.QueryByKeyword(req.Keyword, extCorpID, pager)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	groups = append(groups, groupsByQuickReplyGroup...)
	total = int64(len(groups))
	return
}
