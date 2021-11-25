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
}

func NewSaver(repository repository.Repository) Saver {
	return &SaverImpl{
		repository: repository,
	}
}

func (r *SaverImpl) Execute(ctx context.Context, message *model.Message) error {
	log.Println("saving new message - SERVICE")

	var messages []model.Message

	messages = append(messages, *message)

	err := r.repository.SaveAll(ctx, messages)
	if err != nil {
		return err
	}
	return nil
}
