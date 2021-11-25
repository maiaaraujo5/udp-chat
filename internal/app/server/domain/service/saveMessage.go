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

	messages, err := r.repository.List(ctx)
	if err != nil {
		return err
	}

	if len(messages) >= 3 {
		for i := 0; i <= len(messages); i++ {
			messages = append(messages[:i], messages[i+1:]...)
		}
	}

	messages = append(messages, *message)

	err = r.repository.SaveAll(ctx, messages)
	if err != nil {
		return err
	}
	return nil
}
