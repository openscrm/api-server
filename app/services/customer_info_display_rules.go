package services

import (
	"openscrm/app/constants"
	"openscrm/app/entities"
	"openscrm/app/models"
	"openscrm/common/ecode"
)

type CustomerInfoDisplayRules struct {
	model models.CustomerInfoDisplayRule
}

func NewCustomerInfoDisplayRules() *CustomerInfoDisplayRules {
	return &CustomerInfoDisplayRules{model: models.CustomerInfoDisplayRule{}}
}

func (cs CustomerInfoDisplayRules) Get(extCorpID string) (*models.CustomerInfoDisplayRule, error) {
	rules := models.CustomerInfoDisplayRule{ExtCorpID: extCorpID}
	return cs.model.Get(&rules)
}

func (cs CustomerInfoDisplayRules) Update(extCorpID string, req *entities.UpdateCustomerInfoDisplayRulesReq) error {
	m := map[string]interface{}{}
	for _, filed := range req.DisplayFieldList {
		m[filed] = constants.True
	}
	for _, filed := range req.CancelDisplayFieldList {
		// 重复包含,参数校验
		if _, ok := m[filed]; ok {
			return ecode.InfoFieldDuplicateError
		}
		m[filed] = constants.False
	}
	return cs.model.Updates(extCorpID, m)
}
