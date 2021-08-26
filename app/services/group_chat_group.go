package services

import (
	"github.com/pkg/errors"
	"openscrm/app/entities"
	"openscrm/app/models"
	"openscrm/common/app"
	"openscrm/common/id_generator"
)

type GroupChatGroupService struct {
	groupChatGroupModel models.GroupChatGroup
}

func NewGroupChatGroupService() *GroupChatGroupService {
	return &GroupChatGroupService{groupChatGroupModel: models.GroupChatGroup{}}
}

// Create  创建群聊分组
func (o GroupChatGroupService) Create(
	req entities.CreateGroupChatGroupReq, extCorpID string, creatorID string) (gcg models.GroupChatGroup, err error) {
	gcg = models.GroupChatGroup{
		ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID, ExtCreatorID: creatorID},
		Name:         req.Name,
	}
	err = o.groupChatGroupModel.Create(gcg)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

// Query 获取分组列表
func (o GroupChatGroupService) Query(
	req entities.QueryGroupChatGroupReq,
	extCorpID string, sorter *app.Sorter, pager *app.Pager) ([]models.GroupChatGroup, int64, error) {

	groupChatGroup := models.GroupChatGroup{
		ExtCorpModel: models.ExtCorpModel{ExtCorpID: extCorpID},
		Name:         req.Name,
	}
	return o.groupChatGroupModel.Query(groupChatGroup, extCorpID, sorter, pager)
}

// Update 更新群聊分组
func (o GroupChatGroupService) Update(id string, req entities.UpdateGroupChatGroupReq, extCorpID string) (gcg models.GroupChatGroup, err error) {
	gcg = models.GroupChatGroup{
		ExtCorpModel: models.ExtCorpModel{ID: id, ExtCorpID: extCorpID},
		Name:         req.Name,
	}
	err = o.groupChatGroupModel.Update(gcg)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

// Delete 删除群聊分组
func (o GroupChatGroupService) Delete(req entities.DeleteGroupChatGroupReq, extCorpID string) (int64, error) {
	return o.groupChatGroupModel.Delete(req.IDs, extCorpID)
}
