package requests

import "openscrm/app/constants"

type UpdateCustomerGroupReq struct {
	// 群聊ID
	GroupChatIDs []string `json:"group_chat_ids" validate:"required,gt=0"`
	// 新增标签ID列表
	AddTagIDs constants.StringArrayField `json:"add_tag_ids" validate:"omitempty,gte=0"`
	// 删除标签的ID列表
	RemoveTagIDs constants.StringArrayField `json:"remove_tag_ids" validate:"omitempty,gte=0"`
}
