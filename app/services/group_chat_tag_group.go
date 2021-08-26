package services

import (
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
)

type GroupChatTagGroup struct {
	model   models.GroupChatTagGroup
	tagRepo models.GroupChatTag
}

func NewGroupChatTagGroup() *GroupChatTagGroup {
	return &GroupChatTagGroup{
		model:   models.GroupChatTagGroup{},
		tagRepo: models.GroupChatTag{},
	}
}

func (o GroupChatTagGroup) Create(req *requests.CreateGroupChatTagGroupReq, extCorpID string, extStaffID string) (models.GroupChatTagGroup, error) {
	tg := models.GroupChatTagGroup{
		ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID(), ExtCreatorID: extStaffID, ExtCorpID: extCorpID},
		Name:         req.Name,
	}
	tags := make([]models.GroupChatTag, 0)
	for _, tag := range req.Tags {
		tag := models.GroupChatTag{
			ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID(), ExtCreatorID: extStaffID, ExtCorpID: extCorpID},
			Name:         tag.Name,
		}
		tags = append(tags, tag)
	}
	tg.Tags = tags
	err := o.model.Create(tg)

	mysqlErr := &mysql.MySQLError{}
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return tg, ecode.DuplicateTagGroupError
	}
	return tg, err
}

func (o GroupChatTagGroup) Delete(ids []string) (int64, error) {
	return o.model.Delete(ids)
}

func (o GroupChatTagGroup) Update(req *requests.UpdateGroupChatTagGroupReq, extCorpID string) (models.GroupChatTagGroup, error) {
	group := models.GroupChatTagGroup{
		ExtCorpModel: models.ExtCorpModel{ID: req.ID, ExtCorpID: extCorpID},
		Name:         req.Name,
	}

	newTags := make([]models.GroupChatTag, 0)
	updateTags := make([]models.GroupChatTag, 0)
	for _, tag := range req.Tags {
		// new tag
		if tag.Id == "" && tag.Name != "" {
			newTag := models.GroupChatTag{
				ExtCorpModel:        models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID},
				Name:                tag.Name,
				GroupChatTagGroupID: req.ID,
			}
			newTags = append(newTags, newTag)
		}
		// update tag
		if tag.Id != "" && tag.Name != "" {
			updateTag := models.GroupChatTag{
				ExtCorpModel:        models.ExtCorpModel{ID: tag.Id, ExtCorpID: extCorpID},
				Name:                tag.Name,
				GroupChatTagGroupID: req.ID}
			updateTags = append(updateTags, updateTag)
		}
	}

	// 在给定组下批量添加
	if len(newTags) > 0 {
		err := o.tagRepo.Create(newTags)
		if err != nil {
			return group, errors.WithStack(err)
		}
	}

	if len(updateTags) > 0 {
		err := o.tagRepo.Upsert(updateTags)
		if err != nil {
			return group, errors.WithStack(err)
		}
	}
	// rm tag
	if len(req.DeleteTagIDs) > 0 {
		_, err := o.tagRepo.Delete(req.DeleteTagIDs)
		if err != nil {
			err = errors.WithStack(err)
			return group, err
		}
	}

	err := o.model.Update(group)
	if err != nil {
		err = errors.WithStack(err)
		return group, err
	}

	group, err = o.model.Get(req.ID)
	if err != nil {
		err = errors.WithStack(err)
		return group, err
	}

	return group, nil
}

func (o GroupChatTagGroup) Query(name string, extCorpID string, pager *app.Pager, sorter *app.Sorter) ([]models.GroupChatTagGroup, int64, error) {
	return o.model.Query(
		models.GroupChatTagGroup{
			ExtCorpModel: models.ExtCorpModel{ExtCorpID: extCorpID},
			Name:         name,
		},
		pager, sorter)
}
