package service

import (
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/service"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/fx/module/provider/redis"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		redis.Module(),
		fx.Provide(
			service.NewConfig,
			service.NewDelete,
			service.NewFlusher,
			service.NewRecover,
			service.NewSaver,
		),
	)
}
