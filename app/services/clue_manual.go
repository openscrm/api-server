package services

import (
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/id_generator"
	"openscrm/common/log"
)

type ClueManual struct {
	repo models.CustomerEvent
}

func NewClueManual() *ClueManual {
	return &ClueManual{repo: models.CustomerEvent{}}
}

// Create
// Description: 创建跟进事件
func (m ClueManual) Create(req requests.CreateClueManualReq, extCorpID string, extStaffID string) (models.CustomerEvent, error) {
	customerEvent := models.CustomerEvent{
		ExtCorpModel:  models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID, ExtCreatorID: extStaffID},
		Content:       req.Content,
		EventType:     constants.CustomerEventClueManualEvent,
		EventName:     constants.EventNameClueManualEvent,
		ExtCustomerID: req.ExtCustomerID,
		ExtStaffID:    req.ExtStaffID,
	}

	err := m.repo.Create(customerEvent)
	if err != nil {
		log.Sugar.Errorw("create clue_manual_event failed", "req", req)
		return customerEvent, err
	}
	return customerEvent, nil
}

func (m ClueManual) Delete(ids []string, extCorpID string) (int64, error) {
	return m.repo.Delete(ids, extCorpID)
}

// Update
// Description: 更新跟进事件
func (m ClueManual) Update(id string, req requests.UpdateClueManualReq, extCorpID string) (models.CustomerEvent, error) {
	customerEvent := models.CustomerEvent{
		ExtCorpModel: models.ExtCorpModel{ID: id, ExtCorpID: extCorpID},
		Content:      req.Content,
		EventType:    constants.CustomerEventClueManualEvent,
		EventName:    constants.EventNameClueManualEvent,
	}

	return m.repo.Update(customerEvent)
}
