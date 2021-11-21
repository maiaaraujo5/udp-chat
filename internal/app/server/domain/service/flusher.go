package service

import (
	"context"
	"log"
)

type Flusher interface {
	Execute(parentCtx context.Context) error
}

type FlusherImpl struct {
}

func NewFlusher() *FlusherImpl {
	return &FlusherImpl{}
}

func (r *FlusherImpl) Execute(parentCtx context.Context) error {
	log.Println("Flushing database")
	return nil
}
