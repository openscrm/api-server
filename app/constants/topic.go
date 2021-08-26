package constants

type Topic string

func (o Topic) String() string {
	return string(o)
}

const (
	TagAdminTopic          Topic = "topic:TagAdminTopic"
	DataExportTopic        Topic = "topic:DataExportTopic"
	RemainderTopic         Topic = "topic:RemainderTopic"
	MassMsgTopic           Topic = "topic:MassMsgTopic"
	GroupChatMassMsgTopic  Topic = "topic:GroupChatMassMsgTopic"
	SyncCustomerDataTopic  Topic = "topic:SyncCustomerDataTopic"
	RefreshContactWayTopic Topic = "topic:RefreshContactWayTopic"
)

type JobPrefix string

func (o JobPrefix) String() string {
	return string(o)
}

const (
	ContactWayJobPrefix JobPrefix = "job:contactWay:"
)
