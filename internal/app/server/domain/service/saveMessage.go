package service

import (
	"context"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/repository"
	"log"
)

type Saver interface {
	Execute(ctx context.Context, message *model.Message) error
}

type SaverImpl struct {
	repository repository.Repository
	config     *config
}

func NewSaver(repository repository.Repository, config *config) Saver {
	return &SaverImpl{
		repository: repository,
		config:     config,
	}
}

func (r *SaverImpl) Execute(ctx context.Context, message *model.Message) error {
	log.Println("saving new message - SERVICE")

	messages, err := r.repository.List(ctx)
	if err != nil {
		return err
	}

	if messages.Len() >= r.config.maxMessagesInHistory {
		messages.Remove(messages.Front())
	}

	messages.PushBack(*message)

	err = r.repository.SaveAll(ctx, messages)
	if err != nil {
		return err
	}
	return nil
}
