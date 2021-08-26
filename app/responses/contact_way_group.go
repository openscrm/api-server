package responses

import (
	"openscrm/app/models"
)

type ContactWayGroup struct {
	models.ContactWayGroup
	// 渠道码数量
	Count int64 `json:"count"`
}
