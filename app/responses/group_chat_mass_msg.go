package responses

import "openscrm/app/models"

type MassMsgDetail struct {
	models.MassMsg
	Creator models.StaffMainInfo `json:"creator"`
}
