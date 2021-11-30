package service

import (
	"container/list"
	"context"
	"github.com/maiaaraujo5/gostart/log/logger"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/repository"
	"strings"
)

type DeleteMessage interface {
	Execute(parentCtx context.Context, message *model.Message) error
}

type DeleteImpl struct {
	repository repository.Repository
}

func NewDelete(repository repository.Repository) DeleteMessage {
	return &DeleteImpl{
		repository: repository,
	}
}

func (d *DeleteImpl) Execute(parentCtx context.Context, message *model.Message) error {
	logger.Debug("recovering messages from repository")
	messages, err := d.repository.List(parentCtx)
	if err != nil {
		return err
	}

	logger.Debug("finding message to delete in list")
	element := d.findMessageInList(messages, message)
	if element == nil {
		return nil
	}

	logger.Debug("removing message from list of messages")
	messages.Remove(element)

	logger.Debug("saving messages in repository")
	err = d.repository.SaveAll(parentCtx, messages)
	if err != nil {
		return err
	}

	logger.Info("message successfully deleted")
	return nil
}

func (d *DeleteImpl) findMessageInList(list *list.List, message *model.Message) *list.Element {
	for element := list.Front(); element != nil; element = element.Next() {
		m := element.Value.(model.Message)

		if strings.EqualFold(m.ID, message.ID) && strings.EqualFold(m.UserID, message.UserID) {
			return element
		}
	}
	return nil
}
