package responses

import "openscrm/app/models"

type Role struct {
	models.Role
	// Count 成员数量
	Count       int64               `json:"count" gorm:"default:0;comment:'成员数量'" validate:"gte=0"`
	Permissions []models.Permission `json:"permissions"`
}
