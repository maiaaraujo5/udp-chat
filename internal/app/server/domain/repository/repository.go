package repository

import (
	"context"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
)

type Repository interface {
	SaveAll(parentCtx context.Context, messages []model.Message) error
	List(parentCtx context.Context) ([]model.Message, error)
}
