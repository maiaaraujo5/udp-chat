package service

import (
	"context"
	"fmt"

	"github.com/maiaaraujo5/gostart/log/logger"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/repository"
)

type Saver interface {
	Save(ctx context.Context, message *model.Message) error
}

type saver struct {
	repository repository.Repository
	config     *Config
}

func NewSaver(repository repository.Repository, config *Config) Saver {
	return &saver{
		repository: repository,
		config:     config,
	}
}

func (r *saver) Save(ctx context.Context, message *model.Message) error {

	logger.Debug("recovering history of messages")
	messages, err := r.repository.FindAll(ctx)
	if err != nil {
		return err
	}

	logger.Debug(fmt.Sprintf("verifying if history is more than %d", r.config.MaxMessagesInHistory))
	if messages.Len() >= r.config.MaxMessagesInHistory {
		logger.Debug("Removing surplus history")
		messages.Remove(messages.Front())
	}

	logger.Debug(fmt.Sprintf("pushing new message from user %s in history", message.UserID))
	messages.PushBack(*message)

	logger.Debug("saving messages in repository")
	err = r.repository.SaveAll(ctx, messages)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("new message from user %s saved successfully", message.UserID))
	return nil
}
