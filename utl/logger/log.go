package logger

import "context"

type Log interface {
	Info(ctx context.Context, msg string, args ...interface{})
	Debug(ctx context.Context, msg string, args ...interface{})
	Error(ctx context.Context, msg string, err error, args ...interface{})
}
