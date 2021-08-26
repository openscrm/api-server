package services

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/app"
)

type CustomerEvent struct {
	customerEventRepo models.CustomerEvent
}

func NewCustomerEvent() *CustomerEvent {
	return &CustomerEvent{customerEventRepo: models.CustomerEvent{}}
}

// Query
// Description: 查客户动态列表
func (e CustomerEvent) Query(req requests.QueryEventListReq, extCorpID string, sorter *app.Sorter, pager *app.Pager) (res []models.CustomerEvent, total int64, err error) {

	res = make([]models.CustomerEvent, 0)
	customerEvent := models.CustomerEvent{ExtCorpModel: models.ExtCorpModel{ExtCorpID: extCorpID}}
	err = copier.Copy(&customerEvent, req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	res, total, err = e.customerEventRepo.Query(customerEvent, pager, sorter)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
