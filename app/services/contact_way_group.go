package services

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/app/responses"
	"openscrm/common/app"
	"openscrm/common/id_generator"
)

type ContactWayGroup struct {
	model models.ContactWayGroup
}

func NewContactWayGroup() *ContactWayGroup {
	return &ContactWayGroup{model: models.ContactWayGroup{}}
}

func (o *ContactWayGroup) Query(req requests.QueryContactWayGroupReq, extCorpID string, sorter *app.Sorter, pager *app.Pager) (items []responses.ContactWayGroup, total int64, err error) {
	param := models.ContactWayGroup{}
	err = copier.Copy(&param, req)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}

	var contactWayGroups []models.ContactWayGroup
	contactWayGroups, total, err = o.model.Query(param, extCorpID, sorter, pager)
	if err != nil {
		err = errors.Wrap(err, "o.model.Query failed")
		return
	}

	counts := make([]struct {
		GroupID string
		Total   int64
	}, 0)

	err = models.DB.Model(&models.ContactWay{}).Where("ext_corp_id = ?", extCorpID).Group("group_id").Select("count(*) as total,group_id").Scan(&counts).Error
	if err != nil {
		err = errors.Wrap(err, "count ContactWay failed")
		return
	}

	countMap := make(map[string]int64, 0)
	for _, count := range counts {
		countMap[count.GroupID] = count.Total
	}

	for _, contactWayGroup := range contactWayGroups {
		item := responses.ContactWayGroup{
			ContactWayGroup: contactWayGroup,
			Count:           countMap[contactWayGroup.ID],
		}
		items = append(items, item)
	}

	return
}

func (o *ContactWayGroup) Get(id string, extCorpID string) (item models.ContactWayGroup, err error) {
	return o.model.Get(id, extCorpID)
}

func (o *ContactWayGroup) Create(param requests.CreateContactWayGroupReq, extCorpID string, extCreatorID string) (item models.ContactWayGroup, err error) {
	err = copier.Copy(&item, param)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}
	item.ID = id_generator.StringID()
	item.ExtCorpID = extCorpID
	item.ExtCreatorID = extCreatorID
	return o.model.Create(item, extCorpID)
}

func (o *ContactWayGroup) Update(id string, param requests.UpdateContactWayGroupReq, extCorpID string) (item models.ContactWayGroup, err error) {
	err = copier.Copy(&item, param)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}
	item.ExtCorpID = ""
	return o.model.Update(id, item, extCorpID)
}

func (o *ContactWayGroup) Delete(ids []string, extCorpID string) (total int64, err error) {
	return o.model.Delete(ids, extCorpID)
}
