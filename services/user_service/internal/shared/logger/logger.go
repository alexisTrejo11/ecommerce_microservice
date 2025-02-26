package logging

import (
	"context"
	"os"
	"sync"

	port "github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output/logger"
)

var (
	defaultLogger port.LoggerPort
	once          sync.Once
)

func InitLogger() port.LoggerPort {
	once.Do(func() {
		env := os.Getenv("ENVIRONMENT")
		isProduction := env == "production"

		defaultLogger = NewZapLogger(isProduction)
	})

	return defaultLogger
}

func GetLogger() port.LoggerPort {
	if defaultLogger == nil {
		return InitLogger()
	}
	return defaultLogger
}

func Debug(ctx context.Context, msg string, fields ...port.Field) {
	GetLogger().Debug(ctx, msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...port.Field) {
	GetLogger().Info(ctx, msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...port.Field) {
	GetLogger().Warn(ctx, msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...port.Field) {
	GetLogger().Error(ctx, msg, fields...)
}

func Fatal(ctx context.Context, msg string, fields ...port.Field) {
	GetLogger().Fatal(ctx, msg, fields...)
}

func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if requestID, ok := ctx.Value("requestID").(string); ok {
		return requestID
	}

	return ""
}
