package entities

import "openscrm/common/app"

type CreateGroupChatGroupReq struct {
	Name string `json:"name" validate:"required"`
}

type UpdateGroupChatGroupReq struct {
	Name string `json:"name" validate:"required"`
}

type DeleteGroupChatGroupReq struct {
	IDs []string `json:"ids" validate:"gt=0,dive,int64"`
}

type QueryGroupChatGroupReq struct {
	Name string `json:"name" validate:"omitempty"`
	app.Sorter
	app.Pager
}
