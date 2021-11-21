package service

import (
	"context"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"log"
)

type Saver interface {
	Execute(ctx context.Context, message *model.Message) error
}

type SaverImpl struct {
}

func NewSaver() *SaverImpl {
	return &SaverImpl{}
}

func (r *SaverImpl) Execute(ctx context.Context, message *model.Message) error {
	log.Println("saving new message - SERVICE")
	return nil
}
