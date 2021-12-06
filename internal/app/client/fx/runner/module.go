package runner

import (
	"context"
	"net"

	"github.com/maiaaraujo5/gostart/config"
	logger "github.com/maiaaraujo5/gostart/log/instance"
	udp "github.com/maiaaraujo5/gostart/udp/client"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/service"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/handler"
	"go.uber.org/fx"
)

func Run() error {
	config.Load()
	logger.Load()
	return fx.New(
		fx.Provide(
			udp.Connect,
			handler.NewClient,
			service.NewReceiver,
			service.NewDeleter,
			service.NewCreator,
		),
		fx.Invoke(handle),
	).Start(context.Background())
}

func handle(lifecycle fx.Lifecycle, client *handler.Client, udpConn *net.UDPConn) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				return client.Handle(ctx)
			},
			OnStop: func(ctx context.Context) error {
				return udpConn.Close()
			},
		},
	)
}
