package requests

type UpdateGroupChatTagReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateGroupChatTagsReq struct {
	GroupID string   `json:"group_id" validate:"required"`
	Names   []string `json:"names" validate:"required,gt=0"`
}
