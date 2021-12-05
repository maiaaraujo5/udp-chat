package util

import (
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
	"strings"
)

func RemoveMessageFromMessages(messages []in.In, messageID string) []in.In {
	var newMessages []in.In
	for _, msg := range messages {

		if !strings.EqualFold(msg.ID, messageID) {
			newMessages = append(newMessages, msg)
		}
	}

	return newMessages
}
