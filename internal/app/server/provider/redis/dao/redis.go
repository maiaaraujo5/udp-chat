package dao

import (
	"container/list"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/repository"
	"strings"
)

type Redis struct {
	client *redis.Client
	config *Config
}

func NewRedis(client *redis.Client, config *Config) repository.Repository {
	return &Redis{
		client: client,
		config: config,
	}
}

func (r *Redis) SaveAll(parentCtx context.Context, messages *list.List) error {
	var values []string

	for element := messages.Front(); element != nil; element = element.Next() {
		message := element.Value.(model.Message)

		value := strings.Join([]string{message.ID, message.UserID, message.Message}, "-")
		values = append(values, value)
	}

	err := r.client.Del(parentCtx, r.config.Key).Err()
	if err != nil {
		return err
	}

	err = r.client.RPush(parentCtx, r.config.Key, values).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) List(parentCtx context.Context) (*list.List, error) {

	messages := list.New()

	values, err := r.client.LRange(parentCtx, r.config.Key, 0, -1).Result()
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

		messages.PushBack(message)
	}

	return messages, nil
}

func (r *Redis) Flush(parentCtx context.Context) error {
	err := r.client.FlushAllAsync(parentCtx).Err()
	if err != nil {
		return err
	}
	return nil
}
