package redis

import (
	"github.com/maiaaraujo5/gostart/cache/redis"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/provider/redis/dao"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			redis.Connect,
			dao.NewConfig,
			dao.NewRedis,
		),
	)
}
