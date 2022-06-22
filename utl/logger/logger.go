package logger

import (
	"context"
	"log"

	"go.uber.org/zap"
)

var Logger *ZapLogger

type ZapLogger struct {
	zapLogger *zap.Logger
}

// InitLogger initializes the logger globally
// this is called at the beginning of the program(service)
func InitLogger() {
	zapLogger, err := zap.NewProduction(
		zap.AddCallerSkip(1),
	)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	Logger = &ZapLogger{zapLogger: zapLogger}
}

// Info logs an info message
func (l *ZapLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	l.zapLogger.Info(msg, zap.Any("params", args))
}

// Debug logs a debug message
func (l *ZapLogger) Debug(ctx context.Context, msg string, args ...interface{}) {
	l.zapLogger.Debug(msg, zap.Any("params", args))
}

// Error logs an error message
func (l *ZapLogger) Error(ctx context.Context, msg string, err error, args ...interface{}) {
	l.zapLogger.Error(msg, zap.Error(err), zap.Any("params", args))
}
