package logger

import (
	"context"
	"github.com/maiaaraujo5/gostart/log"
	"github.com/maiaaraujo5/gostart/log/instance"
)

func Fields(fields map[string]interface{}) log.Log {
	i := instance.Load()
	return i.Fields(fields)
}

func FromContext(ctx context.Context) log.Log {
	i := instance.Load()
	return i.FromContext(ctx)
}

func ToContext(ctx context.Context) context.Context {
	i := instance.Load()
	return i.ToContext(ctx)
}

func Panic(msg string) {
	i := instance.Load()
	i.Panic(msg)
}
func Fatal(msg string) {
	i := instance.Load()
	i.Fatal(msg)
}
func Error(msg string) {
	i := instance.Load()
	i.Error(msg)
}
func Warn(msg string) {
	i := instance.Load()
	i.Warn(msg)
}
func Info(msg string) {
	i := instance.Load()
	i.Info(msg)
}
func Debug(msg string) {
	i := instance.Load()
	i.Debug(msg)
}

func Trace(msg string) {
	i := instance.Load()
	i.Trace(msg)
}
