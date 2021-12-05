package dao

import (
	"container/list"
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/maiaaraujo5/gostart/log/logger"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/repository"
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

	logger.Trace("converting messages to array of strings")
	for element := messages.Front(); element != nil; element = element.Next() {
		message := element.Value.(model.Message)

		value := strings.Join([]string{message.ID, message.UserID, message.Message}, "-")
		values = append(values, value)
	}

	logger.Trace("deleting old messages from redis")
	err := r.client.Del(parentCtx, r.config.Key).Err()
	if err != nil {
		logger.Warn(fmt.Sprintf("error to delete old messages from redis: %s", err))
		return nil
	}

	logger.Trace("saving new messages in redis")
	err = r.client.RPush(parentCtx, r.config.Key, values).Err()
	if err != nil {
		logger.Warn(fmt.Sprintf("error to save new messages in redis %s", err))
		return nil
	}

	logger.Trace("messages saved successfully")
	return nil
}

func (r *Redis) List(parentCtx context.Context) (*list.List, error) {

	messages := list.New()
	logger.Trace("recovering messages from redis")
	values, err := r.client.LRange(parentCtx, r.config.Key, 0, -1).Result()
	if err != nil {
		logger.Warn(fmt.Sprintf("error to recover old messages from redis: %s", err))
		return messages, nil
	}

	logger.Trace("converting messages to domain model")
	for _, value := range values {
		v := strings.Split(value, "-")
		message := model.Message{
			ID:      v[0],
			UserID:  v[1],
			Message: v[2],
		}

		messages.PushBack(message)
	}

	logger.Trace("messages found successfully")
	return messages, nil
}

func (r *Redis) Flush(parentCtx context.Context) error {
	err := r.client.FlushAllAsync(parentCtx).Err()
	if err != nil {
		logger.Warn(fmt.Sprintf("error to flush db in redis: %s", err))
	}
	return nil
}
