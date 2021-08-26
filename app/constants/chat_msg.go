package constants

type ChatSessionType string

const (
	ChatSessionTypeGroup    ChatSessionType = "group"
	ChatSessionTypeInternal ChatSessionType = "internal"
	ChatSessionTypeExternal ChatSessionType = "external"
)
