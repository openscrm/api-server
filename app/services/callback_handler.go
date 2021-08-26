package services

import "openscrm/pkg/easywework"

type Event struct {
	workwx.MessageType
	workwx.EventType
	workwx.ChangeType
}
type CallbackHandlerFunc func(message *workwx.RxMessage) error
