package log

import "context"

type Log interface {
	Fields(fields map[string]interface{}) Log
	ToContext(ctx context.Context) context.Context
	FromContext(ctx context.Context) Log
	Panic(msg string)
	Fatal(msg string)
	Error(msg string)
	Warn(msg string)
	Info(msg string)
	Trace(msg string)
	Debug(msg string)
}
