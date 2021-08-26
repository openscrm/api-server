package responses

import "openscrm/app/models"

type GroupChatMassMsgDetail struct {
	models.GroupChatMassMsg
	Creator models.StaffMainInfo `json:"creator"`
}
