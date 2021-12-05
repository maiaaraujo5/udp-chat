package service

import (
	"strings"

	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
)

type Receiver interface {
	Receive(messages []in.In, message *in.In) []in.In
}
type receiver struct {
}

func NewReceiver() Receiver {
	return &receiver{}
}

func (r *receiver) Receive(messages []in.In, message *in.In) []in.In {
	if r.isDeletedMessage(message) {
		return r.removeMessageFromMessages(messages, message.ID)
	}

	messages = append(messages, *message)
	return messages
}

func (r *receiver) isDeletedMessage(msg *in.In) bool {
	return msg.ID != "" && msg.Message == ""
}

func (r *receiver) removeMessageFromMessages(messages []in.In, messageID string) []in.In {
	var newMessages []in.In

	for _, msg := range messages {

		if !strings.EqualFold(msg.ID, messageID) {
			newMessages = append(newMessages, msg)
		}
	}

	return newMessages
}
