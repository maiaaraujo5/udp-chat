package service

import (
	"context"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"time"
)

type Recover interface {
	Execute(ctx context.Context) ([]model.Message, error)
}

type RecoverImpl struct {
}

func NewRecover() *RecoverImpl {
	return &RecoverImpl{}
}

func (r *RecoverImpl) Execute(ctx context.Context) ([]model.Message, error) {
	return []model.Message{
		{
			ID:      "1",
			UserID:  "123",
			Message: "Passei de ano!",
			Time:    time.Time{},
		},
		{
			ID:      "2",
			UserID:  "1234",
			Message: "Que legal!",
			Time:    time.Time{},
		},
		{
			ID:      "3",
			UserID:  "12345",
			Message: "Parabens",
			Time:    time.Time{},
		},
	}, nil
}
