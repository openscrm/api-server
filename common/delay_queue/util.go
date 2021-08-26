package delay_queue

import "openscrm/app/constants"

func topicsToStrings(topics []constants.Topic) []string {
	strings := make([]string, 0)
	for _, topic := range topics {
		strings = append(strings, string(topic))
	}
	return strings
}
