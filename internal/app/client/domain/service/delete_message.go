package service

import (
	"errors"
	"strings"

	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
)

type Deleter interface {
	Delete(messages []in.In, messageID, userID string) ([]in.In, error)
}

type deleter struct {
}

func NewDeleter() Deleter {
	return &deleter{}
}

func (d *deleter) Delete(messages []in.In, messageID, userID string) ([]in.In, error) {
	if !d.clientIsOwnerOfTheMessage(messages, messageID, userID) {
		return nil, errors.New("you can't delete a message that was not sent by you")
	}

	return d.removeMessageFromMessages(messages, messageID), nil
}

func (d *deleter) clientIsOwnerOfTheMessage(messages []in.In, messageID, userID string) bool {
	for _, message := range messages {
		if strings.EqualFold(message.ID, messageID) && strings.EqualFold(message.UserID, userID) {
			return true
		}
	}

	return false
}

func (d *deleter) removeMessageFromMessages(messages []in.In, messageID string) []in.In {
	var newMessages []in.In
	for _, msg := range messages {

		if !strings.EqualFold(msg.ID, messageID) {
			newMessages = append(newMessages, msg)
		}
	}

	return newMessages
}
