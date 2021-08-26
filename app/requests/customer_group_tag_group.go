package requests

type CreateGroupChatTagGroupReq struct {
	Name string `json:"name" form:"name" validate:"required"`
	Tags []struct {
		Name string `json:"name"`
		//Id   string `json:"id,omitempty"`
	} `json:"tags"`
}

type UpdateGroupChatTagGroupReq struct {
	ID           string   `json:"id" validate:"required"`
	Name         string   `json:"name"`
	DeleteTagIDs []string `json:"delete_tag_ids"`
	Tags         []struct {
		Name string `json:"name"`
		Id   string `json:"id,omitempty"`
	} `json:"tags"`
}
