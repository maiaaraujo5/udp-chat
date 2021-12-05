package repository

import (
	"container/list"
	"context"
)

type Repository interface {
	SaveAll(parentCtx context.Context, messages *list.List) error
	FindAll(parentCtx context.Context) (*list.List, error)
	Flush(parentCtx context.Context) error
}
