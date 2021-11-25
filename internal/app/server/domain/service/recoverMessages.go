package service

import (
	"context"
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

	messages, err := r.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
