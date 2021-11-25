package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/repository"
	"strings"
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
	var values []string

	for _, message := range messages {
		value := strings.Join([]string{message.ID, message.UserID, message.Message}, "-")
		values = append(values, value)
	}

	err := r.client.Del(parentCtx, "messages").Err()
	if err != nil {
		return err
	}

	err = r.client.RPush(parentCtx, "messages", values).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) List(parentCtx context.Context) ([]model.Message, error) {

	var messages []model.Message

	values, err := r.client.LRange(parentCtx, "messages", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	for _, value := range values {
		v := strings.Split(value, "-")
		message := model.Message{
			ID:      v[0],
			UserID:  v[1],
			Message: v[2],
		}

		messages = append(messages, message)
	}

	return messages, nil
}
