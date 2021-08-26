package services

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/we_work"
	"openscrm/conf"
	"openscrm/pkg/easywework"
	"sort"
)

type TagService struct {
	TagGroupsRepo models.TagGroup
	TagsRepo      models.Tag
}

func NewTag() *TagService {
	return &TagService{TagsRepo: models.Tag{}, TagGroupsRepo: models.TagGroup{}}
}

// Sync
// Description: 同步标签/标签组
// Detail: 拉取所有标签组
func (t TagService) Sync(extCorpID string) error {
	client, err := GetCorpWxClient(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	// 若tag_id和group_id均为空，则返回所有标签。
	// 同时传递tag_id和group_id时，忽略tag_id，仅以group_id作为过滤条件。
	tagGroups, err := client.Customer.ListExternalContactCorpTags()
	if err != nil {
		log.Sugar.Error("err: ", err)
		return err
	}

	log.Sugar.Debugw("get all tagGroups", "tagGroups", tagGroups)

	// 拼接models.tagGroup,upsert
	err = t.UpsertTagGroups(tagGroups)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return nil
}

// UpsertTagGroups
// Description: 用微信的tag数据upsert DB中的标签
func (t TagService) UpsertTagGroups(tagGroups []workwx.ExternalContactCorpTagGroup) (err error) {

	extCorpID := conf.Settings.WeWork.ExtCorpID
	for _, group := range tagGroups {
		tagGroup := models.TagGroup{
			ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID},
			ExtID:        group.GroupID,
			Name:         group.GroupName,
			CreateTime:   group.CreateTime,
			Order:        group.Order,
		}

		tagGroup.Tags = make([]models.Tag, len(group.Tag))
		for i := 0; i < len(group.Tag); i++ {
			tagGroup.Tags[i] = models.Tag{
				ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID},
				ExtID:        group.Tag[i].ID,
				ExtGroupID:   group.GroupID, //外部组id
				Name:         group.Tag[i].Name,
				GroupName:    group.GroupName,
				CreateTime:   group.Tag[i].CreateTime,
				Order:        group.Tag[i].Order,
				Type:         1, // 1-企业设置, 2-用户自定义
			}
		}

		log.Logger.Debug("tagModel", zap.Reflect("group: ", tagGroup))

		err = t.TagGroupsRepo.Upsert(&tagGroup)
		if err != nil {
			log.Logger.Error("err", zap.Error(err))
			return err
		}
	}
	return nil
}

// Create tag 已有标签组中添加标签，需判重
func (t TagService) Create(req *requests.CreateTagReq, extCorpID string) (tagModels []models.Tag, err error) {
	if len(req.Names) > 1 {
		sort.SliceStable(req.Names, func(i, j int) bool {
			return true
		})
	}
	group, err := t.TagGroupsRepo.Get(req.ExtTagGroupId)
	if err != nil {
		err = errors.Wrap(err, "Get tag group failed")
		return
	}
	// 查找groupName下是否已有待建标签
	tags, err := t.TagsRepo.GetExistTag(req.ExtTagGroupId, req.Names)
	if err != nil || len(tags) > 0 {
		err = ecode.DuplicateTagError
		return
	}

	newTagReq := make([]workwx.ExternalContactCorpTag, 0)
	for _, tagName := range req.Names {
		newTag := workwx.ExternalContactCorpTag{Name: tagName}
		newTagReq = append(newTagReq, newTag)
	}

	tagGroup := workwx.ExternalContactCorpTagGroup{
		GroupID:   group.ExtID,
		GroupName: group.Name,
		Tag:       newTagReq,
	}
	client, err := GetCorpWxClient(extCorpID)
	if err != nil {
		return
	}
	createTagResp, err := client.Customer.AddExternalContactCorpTag(tagGroup)
	if err != nil {
		return
	}

	err = models.DB.Transaction(func(tx *gorm.DB) error {
		order, err := t.TagsRepo.GetCurMaxOrder(tx, req.ExtTagGroupId)
		if err != nil {
			return err
		}
		for _, tag := range createTagResp.Tag {
			order += 1
			newTagModel := models.Tag{
				ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: group.ExtCorpID},
				ExtGroupID:   group.ExtID,
				Name:         tag.Name,
				GroupName:    group.Name,
				Type:         1,
				ExtID:        tag.ID,
				CreateTime:   tag.CreateTime,
				Order:        uint32(order),
			}
			tagModels = append(tagModels, newTagModel)
		}
		return t.TagsRepo.CreateTags(tx, tagModels)
	})
	return
}

func (t TagService) DeleteTagGroups(req requests.DeleteTagGroupsReq) (int64, error) {
	err := we_work.App.DelExternalContactCorpTag(nil, req.ExtIDs)
	if err != nil {
		return 0, err
	}
	return t.TagGroupsRepo.Delete(req.ExtIDs)
}
