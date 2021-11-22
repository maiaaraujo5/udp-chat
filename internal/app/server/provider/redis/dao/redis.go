package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) *Redis {
	return &Redis{
		client: client,
	}
}

func (r *Redis) SaveAll(parentCtx context.Context, messages []model.Message) error {
	r.client.Set(parentCtx, "messages", messages, 0)
	return nil
}

func (r *Redis) List(parentCtx context.Context) ([]model.Message, error) {
	return nil, nil
}
