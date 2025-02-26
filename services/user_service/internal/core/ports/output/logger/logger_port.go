package port

import "context"

type Field struct {
	Key   string
	Value interface{}
}

type LoggerPort interface {
	Debug(ctx context.Context, msg string, fields ...Field)
	Info(ctx context.Context, msg string, fields ...Field)
	Warn(ctx context.Context, msg string, fields ...Field)
	Error(ctx context.Context, msg string, fields ...Field)
	Fatal(ctx context.Context, msg string, fields ...Field)

	WithContext(ctx context.Context) LoggerPort

	WithService(serviceName string) LoggerPort
	WithRequestID(requestID string) LoggerPort
	WithField(key string, value interface{}) LoggerPort

	Flush()
}
