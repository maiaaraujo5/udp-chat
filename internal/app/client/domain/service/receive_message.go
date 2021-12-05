package service

import (
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
	"github.com/maiaaraujo5/udp-chat/pkg/util"
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
		return util.RemoveMessageFromMessages(messages, message.ID)
	}

	messages = append(messages, *message)
	return messages
}

func (r *receiver) isDeletedMessage(msg *in.In) bool {
	return msg.ID != "" && msg.Message == ""
}
