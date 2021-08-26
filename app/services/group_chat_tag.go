package services

import (
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
)

type GroupChatTag struct {
	model models.GroupChatTag
}

func NewCustomerGroupTag() *GroupChatTag {
	return &GroupChatTag{model: models.GroupChatTag{}}
}

func (o GroupChatTag) Create(req requests.CreateGroupChatTagsReq, extCorpID string, extStaffID string) ([]models.GroupChatTag, error) {
	tags := make([]models.GroupChatTag, 0)
	for _, name := range req.Names {
		tag := models.GroupChatTag{
			ExtCorpModel:        models.ExtCorpModel{ID: id_generator.StringID(), ExtCreatorID: extStaffID, ExtCorpID: extCorpID},
			GroupChatTagGroupID: req.GroupID,
			Name:                name,
		}
		tags = append(tags, tag)
	}
	err := o.model.Create(tags)

	mysqlErr := &mysql.MySQLError{}
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return tags, ecode.DuplicateTagError
	}
	return tags, err
}

func (o GroupChatTag) Delete(ids []string) (int64, error) {
	return o.model.Delete(ids)
}

func (o GroupChatTag) Update(req requests.UpdateGroupChatTagReq, extCorpID string) (models.GroupChatTag, error) {
	tag := models.GroupChatTag{ExtCorpModel: models.ExtCorpModel{ID: req.ID, ExtCorpID: extCorpID}, Name: req.Name}
	err := o.model.Updates(tag)
	return tag, err
}
