package workwx

// ChatInfo 群聊信息
type ChatInfo struct {
	// ChatID 群聊唯一标志
	ChatID string `json:"chatid"`
	// Name 群聊名
	Name string `json:"name"`
	// OwnerUserID 群主id
	OwnerUserID string `json:"owner"`
	// MemberUserIDs 群成员id列表
	MemberUserIDs []string `json:"userlist"`
}
