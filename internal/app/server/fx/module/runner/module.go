package runner

import (
	"context"
	"net"

	"github.com/maiaaraujo5/gostart/config"
	logger "github.com/maiaaraujo5/gostart/log/instance"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/fx/module/server"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/handler"
	"go.uber.org/fx"
)

func Run() error {
	config.Load()
	logger.Load()
	return fx.New(
		server.Module(),
		fx.Provide(
			context.Background,
		),
		fx.Invoke(handle),
	).Start(context.Background())
}

func handle(lifecycle fx.Lifecycle, server *handler.Server, udpConn *net.UDPConn) {

	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				return server.Handle(ctx)
			},
			OnStop: func(ctx context.Context) error {
				return udpConn.Close()
			},
		},
	)
}
