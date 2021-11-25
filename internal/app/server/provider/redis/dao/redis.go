package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/repository"
	"time"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) repository.Repository {
	return &Redis{
		client: client,
	}
}

func (r *Redis) SaveAll(parentCtx context.Context, messages []model.Message) error {
	r.client.Set(parentCtx, "messages", messages, 0)
	return nil
}

func (r *Redis) List(parentCtx context.Context) ([]model.Message, error) {
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
