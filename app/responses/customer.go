package responses

import "openscrm/app/models"

type FullCustomerInfo struct {
	models.Customer
	models.CustomerInfo
	CustomerEvents
}
type CustomerEvents struct {
	Events []models.CustomerEvent `json:"events"`
	Total  int64                  `json:"total"`
}
