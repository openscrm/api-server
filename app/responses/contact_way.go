package responses

import (
	"openscrm/app/models"
)

type ContactWay struct {
	models.ContactWay
	Group models.ContactWayGroup `json:"group"`
	// CustomerTags 自动打标签绑定的标签
	CustomerTags []models.Tag `json:"customer_tags" gorm:"type:json;comment:'自动打标签绑定的标签'"`
}
