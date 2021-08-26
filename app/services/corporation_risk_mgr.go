package services

import (
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/app"
	"openscrm/common/id_generator"
	"time"
)

type CorpRiskMgrService struct {
	customerStaffRepo models.CustomerStaff
	eventNotifyRepo   models.EventNotify
	relationHistory   models.CustomerStaffRelationHistory
}

func NewCorpRiskMgrService() *CorpRiskMgrService {
	return &CorpRiskMgrService{
		customerStaffRepo: models.CustomerStaff{},
		eventNotifyRepo:   models.EventNotify{},
		relationHistory:   models.CustomerStaffRelationHistory{},
	}
}

func (s CorpRiskMgrService) Query(
	req requests.QueryStaffDeleteCustomerHistoryReq, extCorpID string, sorter *app.Sorter, pager *app.Pager) ([]models.StaffDeleteCustomer, int64, error) {

	return s.relationHistory.QueryStaffDeleteCustomer(req, extCorpID, pager, sorter)
}

func (s CorpRiskMgrService) Upsert(req requests.UpdateStaffDeleteCustomerNotifierReq, extCorpID string) error {

	notify := models.EventNotify{
		ID:             id_generator.StringID(),
		IsNotifyAdmins: req.IsNotifyStaff,
		ExtStaffIDs:    req.ExtStaffIDs,
		NotifyType:     req.NotifyType,
		IsNotifyStaff:  req.IsNotifyStaff,
	}

	notify.ExtCorpID = extCorpID
	notify.ID = id_generator.StringID()
	notify.UpdatedAt = time.Now()
	return s.eventNotifyRepo.Upsert(notify)
}

func (s CorpRiskMgrService) Get(extCorpID string) (models.EventNotify, error) {
	return s.eventNotifyRepo.Get(extCorpID)
}
