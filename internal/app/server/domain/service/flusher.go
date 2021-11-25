package service

import (
	"context"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/repository"
	"log"
)

type Flusher interface {
	Execute(parentCtx context.Context) error
}

type FlusherImpl struct {
	repository repository.Repository
}

func NewFlusher(repository repository.Repository) Flusher {
	return &FlusherImpl{
		repository: repository,
	}
}

func (r *FlusherImpl) Execute(parentCtx context.Context) error {
	log.Println("Flushing database")
	return nil
}
