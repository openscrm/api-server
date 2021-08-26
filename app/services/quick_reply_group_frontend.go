package services

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"openscrm/app/entities"
	"openscrm/app/models"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
)

type QuickReplyGroupFrontend struct {
	QuickReplyGroupRepo models.QuickReplyGroup
	QuickReplyRepo      models.QuickReply
}

func NewQuickReplyGroupFrontend() *QuickReplyGroupFrontend {
	return &QuickReplyGroupFrontend{QuickReplyGroupRepo: models.QuickReplyGroup{}, QuickReplyRepo: models.QuickReply{}}
}

func (r QuickReplyGroupFrontend) Update(
	req entities.UpdateQuickReplyGroupReq, extCorpID, extStaffID string) (g models.QuickReplyGroup, err error) {
	groups, err := r.QuickReplyGroupRepo.Get(req.ID, extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if len(groups) == 0 {
		err = ecode.QuickReplyGroupNotFoundErr
		err = errors.WithStack(err)
		return
	}
	if groups[0].ExtCreatorID != extStaffID {
		err = ecode.UpdateOtherRecordNotAllowedErr
		err = errors.WithStack(err)
		return
	}
	group := models.QuickReplyGroup{
		ExtCorpModel: models.ExtCorpModel{ID: req.ID, ExtCorpID: extCorpID},
		Name:         req.Name,
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
		err := r.QuickReplyGroupRepo.CreateInBatches(newGroups)
		if err != nil {
			err = errors.WithStack(err)
			return group, err
		}
	}

	if len(updateGroups) > 0 {
		err := r.QuickReplyGroupRepo.Upsert(updateGroups)
		if err != nil {
			err = errors.WithStack(err)
			return group, err
		}
	}
	// rm tag
	if len(req.DeleteGroupIDs) > 0 {
		err := r.QuickReplyGroupRepo.Delete(models.DB, req.DeleteGroupIDs, extCorpID)
		if err != nil {
			err = errors.WithStack(err)
			return group, err
		}
	}

	err = r.QuickReplyGroupRepo.Update([]models.QuickReplyGroup{group})
	if err != nil {
		err = errors.WithStack(err)
		return group, err
	}

	groups, err = r.QuickReplyGroupRepo.Get(req.ID, extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return groups[0], err
	}

	return groups[0], nil
}

func (r QuickReplyGroupFrontend) Delete(ids []string, extCorpID, extStaffID string) error {
	groups, err := r.QuickReplyGroupRepo.GetByIDs(ids)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	for _, g := range groups {
		if g.ExtCreatorID != extStaffID {
			err = ecode.DeleteOtherRecordNotAllowedErr
			return err
		}
	}
	return models.DB.Transaction(func(tx *gorm.DB) error {

		err := r.QuickReplyRepo.DeleteByGroupIDs(tx, ids)
		if err != nil {
			return err
		}

		err = r.QuickReplyGroupRepo.Delete(tx, ids, extCorpID)
		if err != nil {
			return err
		}
		return nil
	})
}
