package services

import (
	"github.com/pkg/errors"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/pkg/easywework"
)

type TagGroupService struct {
	groupRepo models.TagGroup
	tagRepo   models.Tag
}

func NewTagGroupService() *TagGroupService {
	return &TagGroupService{groupRepo: models.TagGroup{}, tagRepo: models.Tag{}}
}

func (t TagGroupService) Query(req requests.TagListReq, extCorpID string) ([]*models.TagGroup, int64, error) {
	return t.groupRepo.Query(req, extCorpID)
}

func (t TagGroupService) Create(req requests.CreateTagGroupReq, extCorpID string) (tagGroup *models.TagGroup, err error) {
	total := int64(0)
	err = models.DB.Model(&models.TagGroup{}).Where(&models.TagGroup{Name: req.Name}).Count(&total).Error
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if total > 0 {
		err = ecode.DuplicateTagGroupError
		return
	}

	tags := make([]workwx.ExternalContactCorpTag, 0)
	// wx post 参数
	for i, tag := range req.Tags {
		tag := workwx.ExternalContactCorpTag{Name: tag.Name, Order: tag.Order}
		// 标签默认按数组里的排序生成order值
		if tag.Order == 0 {
			tag.Order = uint32(10000 - i)
		}
		tags = append(tags, tag)
	}

	//http request
	client, err := GetCorpWxClient(extCorpID)
	if err != nil {
		return
	}
	group := workwx.ExternalContactCorpTagGroup{GroupName: req.Name, Tag: tags, Order: req.Order}
	res, err := client.Customer.AddExternalContactCorpTag(group)
	if err != nil {
		return
	}

	// 更新ext_group_id
	tagGroup = &models.TagGroup{
		ExtCorpModel:   models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID},
		ExtID:          res.GroupID,
		CreateTime:     res.CreateTime,
		Order:          res.Order,
		Name:           req.Name,
		DepartmentList: req.DepartmentList,
		Tags:           nil,
	}

	tagsModel := make([]models.Tag, 0)
	for _, tag := range res.Tag {
		tagModel := models.Tag{
			ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID},
			ExtGroupID:   res.GroupID,
			//GroupID:      tagGroup.ID,
			Name:       tag.Name,
			GroupName:  res.GroupName,
			ExtID:      tag.ID,
			CreateTime: tag.CreateTime,
			Order:      tag.Order,
			Type:       1,
		}
		tagsModel = append(tagsModel, tagModel)
	}
	tagGroup.Tags = tagsModel

	return t.groupRepo.Create(tagGroup)
}

func (t TagGroupService) Delete(req requests.DeleteTagGroupsReq, extCorpID string) (int64, error) {
	client, err := GetCorpWxClient(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return 0, err
	}
	err = client.Customer.DelExternalContactCorpTag([]string{}, req.ExtIDs)
	if err != nil {
		err = errors.WithStack(err)
		return 0, err
	}
	return t.groupRepo.Delete(req.ExtIDs)
}

// Update
// Description: 更新标签组
// Detail: 同时更新标签组和标签
// Param:
//	标签组-ext_id 标签组名-Name 排序权重-Order 标签可用部门列表, 缺省所有部门可用-DepartmentList
//	删除的标签id-RemoveExtTagIDs
// 	标签列表-Tags 用于新增或更新标签
func (t TagGroupService) Update(req requests.UpdateTagGroupReq, extCorpID string) (tg models.TagGroup, err error) {
	client, err := GetCorpWxClient(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if len(req.DepartmentList) == 0 {
		req.DepartmentList = []int64{0}
	}

	newTags := make([]workwx.ExternalContactCorpTag, 0)
	for i, tag := range req.Tags {
		tag.Order = uint32(10000 - i)
		// new tag
		if tag.ExtId == "" && tag.Name != "" {
			corpTag := workwx.ExternalContactCorpTag{Name: tag.Name, Order: tag.Order}
			newTags = append(newTags, corpTag)
		}
		// update tag
		if tag.ExtId != "" && tag.Name != "" {
			err = client.Customer.EditExternalContactCorpTag(tag.ExtId, tag.Name, tag.Order)
			if err != nil {
				log.Sugar.Error(err)
				// err = errors.WithStack(err)
				return
			}
			err = t.tagRepo.Update(models.Tag{ExtID: tag.ExtId, ExtGroupID: req.ExtID, Name: tag.Name, Order: tag.Order})
			if err != nil {
				err = errors.WithStack(err)
				return
			}
		}
	}

	// 在给定组下批量添加
	if len(newTags) > 0 {
		corpTag := workwx.ExternalContactCorpTagGroup{
			GroupID: req.ExtID,
			Tag:     newTags,
		}
		tags, err := client.Customer.AddExternalContactCorpTag(corpTag)
		if err != nil {
			return tg, errors.WithStack(err)
		}

		for _, tag := range tags.Tag {
			newTag := models.Tag{
				ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID},
				ExtID:        tag.ID,
				ExtGroupID:   tags.GroupID,
				Name:         tag.Name,
				GroupName:    tags.GroupName,
				CreateTime:   tag.CreateTime,
				Order:        tag.Order,
				Type:         1,
			}
			err = t.tagRepo.Upsert(newTag)
			if err != nil {
				return tg, errors.WithStack(err)
			}
			log.Sugar.Info(tag)
		}
	}

	// rm tag
	if len(req.RemoveExtTagIDs) > 0 {
		err = client.Customer.DelExternalContactCorpTag(req.RemoveExtTagIDs, nil)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = t.tagRepo.Delete(req.RemoveExtTagIDs)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
	}

	// update tagGroupName
	err = client.Customer.EditExternalContactCorpTag(req.ExtID, req.Name, req.Order)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	err = t.groupRepo.Update(
		&models.TagGroup{
			ExtCorpModel:   models.ExtCorpModel{ID: id_generator.StringID()},
			ExtID:          req.ExtID,
			Name:           req.Name,
			Order:          req.Order,
			DepartmentList: req.DepartmentList,
		},
	)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	tg, err = t.groupRepo.Get(req.ExtID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

func (t TagGroupService) ExchangeOrder(e *requests.ExchangeOrderReq) error {
	return t.groupRepo.ExchangeOrder(e.ExchangeOrderID, e.ID)
}
