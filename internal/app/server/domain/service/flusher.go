package service

import (
	"context"
	"github.com/maiaaraujo5/gostart/log/logger"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/repository"
)

type Flusher interface {
	Flush(parentCtx context.Context) error
}

type flusher struct {
	repository repository.Repository
}

func NewFlusher(repository repository.Repository) Flusher {
	return &flusher{
		repository: repository,
	}
}

func (r *flusher) Flush(parentCtx context.Context) error {
	logger.Debug("flushing database")
	err := r.repository.Flush(parentCtx)
	if err != nil {
		return err
	}

	logger.Info("repository flushed")
	return nil
}
