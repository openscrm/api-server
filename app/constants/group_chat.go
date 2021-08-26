package constants

type GroupChatStatus uint8

const (
	GroupChatStatusIsDismissed  = 1
	GroupChatStatusNotDismissed = 2
)

const (
	GroupChatChangeTypeAddMember    string = "add_member"    // 成员入群
	GroupChatChangeTypeDelMember    string = "del_member"    // 成员退群
	GroupChatChangeTypeChangeOwner  string = "change_owner"  //		change_owner : 群主变更
	GroupChatChangeTypeChangeName   string = "change_name"   //		change_name : 群名变更
	GroupChatChangeTypeChangeNotice string = "change_notice" //		change_notice : 群公告变更
)
