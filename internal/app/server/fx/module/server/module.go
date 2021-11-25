package server

import (
	udp "github.com/maiaaraujo5/gostart/udp/server"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/fx/module/service"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/handler"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		service.Module(),
		fx.Provide(
			udp.Server,
			handler.NewServer,
		),
	)
}
