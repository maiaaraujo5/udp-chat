package zerolog

import (
	"context"
	appconfig "github.com/maiaaraujo5/gostart/application/config"
	"github.com/maiaaraujo5/gostart/log"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"os"
	"strings"
)

type zeroLog struct {
	logger zerolog.Logger
}

func NewZeroLog() log.Log {
	config, err := newConfig()
	if err != nil {
		return nil
	}

	level, err := zerolog.ParseLevel(strings.ToLower(config.Level))
	if err != nil {
		return nil
	}
	return &zeroLog{
		logger: zerolog.New(os.Stdout).
			Level(level).
			With().
			Fields(appconfig.DefaultAppFields()).
			Timestamp().
			Logger(),
	}
}

func (z zeroLog) Fields(fields map[string]interface{}) log.Log {
	for key, value := range fields {
		z.logger = z.logger.With().Interface(key, value).Logger()
	}
	return z
}

func (z zeroLog) ToContext(ctx context.Context) context.Context {
	return z.logger.WithContext(ctx)
}

func (z zeroLog) FromContext(ctx context.Context) log.Log {
	z.logger = *zlog.Ctx(ctx)
	return z
}

func (z zeroLog) Panic(msg string) {
	z.logger.Panic().Msg(msg)
}

func (z zeroLog) Fatal(msg string) {
	z.logger.Fatal().Msg(msg)
}

func (z zeroLog) Error(msg string) {
	z.logger.Error().Msg(msg)
}

func (z zeroLog) Warn(msg string) {
	z.logger.Warn().Msg(msg)
}

func (z zeroLog) Info(msg string) {
	z.logger.Info().Msg(msg)
}

func (z zeroLog) Trace(msg string) {
	z.logger.Trace().Msg(msg)
}

func (z zeroLog) Debug(msg string) {
	z.logger.Debug().Msg(msg)
}
