package service

import (
	"errors"
	"github.com/maiaaraujo5/udp-chat/pkg/util"
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

	return util.RemoveMessageFromMessages(messages, messageID), nil
}

func (d *deleter) clientIsOwnerOfTheMessage(messages []in.In, messageID, userID string) bool {
	for _, message := range messages {
		if strings.EqualFold(message.ID, messageID) && strings.EqualFold(message.UserID, userID) {
			return true
		}
	}

	return false
}
