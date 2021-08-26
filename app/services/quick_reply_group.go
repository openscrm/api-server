package services

import (
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"openscrm/app/constants"
	"openscrm/app/entities"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
)

type QuickReplyGroup struct {
	QuickReplyGroupRepo models.QuickReplyGroup
	QuickReplyRepo      models.QuickReply
}

func NewQuickReplyGroup() *QuickReplyGroup {
	return &QuickReplyGroup{QuickReplyGroupRepo: models.QuickReplyGroup{}, QuickReplyRepo: models.QuickReply{}}
}

func (o QuickReplyGroup) Create(req entities.CreateQuickReplyGroupReq, extCorpID, extCreatorID string) (models.QuickReplyGroup, error) {
	topGroupID := id_generator.StringID()
	groups := make([]models.QuickReplyGroup, 0)
	for i, subGroup := range req.SubGroups {
		group := models.QuickReplyGroup{
			ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID, ExtCreatorID: extCreatorID},
			Name:         subGroup.Name,
			ParentID:     &topGroupID,
			IsTopGroup:   constants.False,
			Departments:  req.Departments,
			Order:        int64(10000 - i),
		}
		groups = append(groups, group)
	}
	topGroup := models.QuickReplyGroup{
		ExtCorpModel: models.ExtCorpModel{ID: topGroupID, ExtCorpID: extCorpID, ExtCreatorID: extCreatorID},
		Name:         req.Name,
		Departments:  req.Departments,
		IsTopGroup:   constants.True,
		SubGroups:    groups,
	}

	err := o.QuickReplyGroupRepo.Create(topGroup)

	mysqlErr := &mysql.MySQLError{}
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return models.QuickReplyGroup{}, ecode.DuplicateQuickReplyGroupNameError
	}
	return topGroup, err
}

func (o QuickReplyGroup) Query(
	extCorpID string, sorter *app.Sorter, pager *app.Pager) (qr []models.QuickReplyGroup, total int64, err error) {

	return o.QuickReplyGroupRepo.Query(extCorpID, sorter, pager)
}

func (o QuickReplyGroup) Delete(ids []string, extCorpID string) error {
	return models.DB.Transaction(func(tx *gorm.DB) error {
		err := o.QuickReplyRepo.DeleteByGroupIDs(tx, ids)
		if err != nil {
			return err
		}

		err = o.QuickReplyGroupRepo.Delete(tx, ids, extCorpID)
		if err != nil {
			return err
		}
		return nil
	})
}

func (o QuickReplyGroup) Update(req entities.UpdateQuickReplyGroupReq, extCorpID string) (models.QuickReplyGroup, error) {
	group := models.QuickReplyGroup{
		ExtCorpModel: models.ExtCorpModel{ID: req.ID, ExtCorpID: extCorpID},
		Name:         req.Name,
		Departments:  req.Departments,
	}

	newGroups := make([]models.QuickReplyGroup, 0)
	updateGroups := make([]models.QuickReplyGroup, 0)
	for i, subGroup := range req.SubGroups {
		// new tag
		if subGroup.ID == "" && subGroup.Name != "" {
			newGroups = append(newGroups, models.QuickReplyGroup{
				ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID},
				Name:         subGroup.Name,
				ParentID:     &req.ID,
				Order:        int64(10000 - i),
			})
		}
		// update tag
		if subGroup.ID != "" && subGroup.Name != "" {
			updateGroups = append(updateGroups, models.QuickReplyGroup{
				ExtCorpModel: models.ExtCorpModel{ID: subGroup.ID, ExtCorpID: extCorpID},
				Name:         subGroup.Name,
				ParentID:     &req.ID,
				Order:        int64(10000 - i),
			})
		}
	}

	// 在给定组下批量添加
	if len(newGroups) > 0 {
		err := o.QuickReplyGroupRepo.CreateInBatches(newGroups)
		if err != nil {
			return group, err
		}
	}

	if len(updateGroups) > 0 {
		err := o.QuickReplyGroupRepo.Upsert(updateGroups)
		if err != nil {
			return group, err
		}
	}
	// rm tag
	if len(req.DeleteGroupIDs) > 0 {
		err := o.QuickReplyGroupRepo.Delete(models.DB, req.DeleteGroupIDs, extCorpID)
		if err != nil {
			return group, err
		}
	}

	err := o.QuickReplyGroupRepo.Update([]models.QuickReplyGroup{group})
	if err != nil {
		return group, err
	}

	g, err := o.QuickReplyGroupRepo.Get(req.ID, extCorpID)
	if err != nil {
		return g[0], err
	}

	return g[0], nil
}

func (o QuickReplyGroup) QueryQuickReply(
	req requests.QueryQuickReplyReq, extCorpID string, pager *app.Pager) (groups []models.QuickReplyGroup, total int64, err error) {

	replies, err := o.QuickReplyRepo.QueryByKeyword(req.Keyword, extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	groupIDs := make([]string, 0)
	replyIDs := make([]string, 0)
	for _, reply := range replies {
		groupIDs = append(groupIDs, reply.GroupID)
		replyIDs = append(replyIDs, reply.ID)
	}

	groups, err = o.QuickReplyGroupRepo.QueryByID(extCorpID, groupIDs, replyIDs)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	groupsByQuickReplyGroup, err := o.QuickReplyGroupRepo.QueryByKeyword(req.Keyword, extCorpID, pager)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	groups = append(groups, groupsByQuickReplyGroup...)
	total = int64(len(groups))
	return
}
