package service

import (
	"context"
	"github.com/maiaaraujo5/gostart/log/logger"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/repository"
)

type Recover interface {
	Execute(ctx context.Context) ([]model.Message, error)
}

type RecoverImpl struct {
	repository repository.Repository
}

func NewRecover(repository repository.Repository) Recover {
	return &RecoverImpl{
		repository: repository,
	}
}

func (r *RecoverImpl) Execute(ctx context.Context) ([]model.Message, error) {

	var messages []model.Message

	logger.Debug("recovering messages from repository")
	list, err := r.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	logger.Debug("converting messages to array of domain model")
	for element := list.Front(); element != nil; element = element.Next() {

		message := element.Value.(model.Message)

		messages = append(messages, message)
	}

	logger.Info("messages successfully recovered")
	return messages, nil
}
