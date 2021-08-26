package responses

import "openscrm/app/models"

type StaffMainInfo struct {
	models.StaffMainInfo
}
type StaffDetail struct {
	models.Staff
	Role models.Role `json:"role,omitempty"`
}
