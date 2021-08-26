package entities

import "openscrm/pkg/easywework"

// DeleteCustomerNotifyText 定时推送
type DeleteCustomerNotifyText struct {
	Recipient workwx.Recipient `json:"recipient"`
	Content   string           `json:"content"`
}
