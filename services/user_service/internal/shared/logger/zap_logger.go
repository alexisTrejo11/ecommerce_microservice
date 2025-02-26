package logging

import (
	"context"
	"os"

	port "github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
	fields []port.Field
}

func NewZapLogger(isProduction bool) *ZapLogger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Determinar nivel de log
	logLevel := os.Getenv("LOG_LEVEL")
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		if isProduction {
			level = zapcore.InfoLevel
		} else {
			level = zapcore.DebugLevel
		}
	}

	var config zap.Config
	if isProduction {
		config = zap.Config{
			Level:             zap.NewAtomicLevelAt(level),
			Development:       false,
			DisableCaller:     false,
			DisableStacktrace: false,
			Sampling:          nil,
			Encoding:          "json",
			EncoderConfig:     encoderConfig,
			//portPaths:         []string{"stdout"},
			//ErrorportPaths:    []string{"stderr"},
		}
	} else {
		config = zap.Config{
			Level:             zap.NewAtomicLevelAt(level),
			Development:       true,
			DisableCaller:     false,
			DisableStacktrace: false,
			Sampling:          nil,
			Encoding:          "console",
			EncoderConfig:     encoderConfig,
			//portPaths:         []string{"stdout"},
			//ErrorportPaths:    []string{"stderr"},
		}
	}

	// Crear logger
	logger, err := config.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	// Agregar metadatos comunes a todos los logs
	hostname, _ := os.Hostname()
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	logger = logger.With(
		zap.String("service", "go-api"),
		zap.String("type", "go-app"),
		zap.String("host", hostname),
		zap.String("environment", env),
	)

	return &ZapLogger{
		logger: logger,
		fields: []port.Field{},
	}
}

// convertFields convierte port.Field a zap.Field
func (l *ZapLogger) convertFields(fields []port.Field) []zap.Field {
	// Combinar campos predefinidos con campos adicionales
	allFields := append(l.fields, fields...)
	zapFields := make([]zap.Field, 0, len(allFields))

	for _, field := range allFields {
		zapFields = append(zapFields, zap.Any(field.Key, field.Value))
	}

	return zapFields
}

// addRequestIDFromContext agrega el ID de solicitud desde el contexto si está disponible
func (l *ZapLogger) addRequestIDFromContext(ctx context.Context, fields []port.Field) []port.Field {
	if ctx == nil {
		return fields
	}

	requestID := GetRequestID(ctx)
	if requestID != "" {
		return append(fields, port.Field{Key: "request_id", Value: requestID})
	}

	return fields
}

// Implementación de LoggerPort

func (l *ZapLogger) Debug(ctx context.Context, msg string, fields ...port.Field) {
	fields = l.addRequestIDFromContext(ctx, fields)
	l.logger.Debug(msg, l.convertFields(fields)...)
}

func (l *ZapLogger) Info(ctx context.Context, msg string, fields ...port.Field) {
	fields = l.addRequestIDFromContext(ctx, fields)
	l.logger.Info(msg, l.convertFields(fields)...)
}

func (l *ZapLogger) Warn(ctx context.Context, msg string, fields ...port.Field) {
	fields = l.addRequestIDFromContext(ctx, fields)
	l.logger.Warn(msg, l.convertFields(fields)...)
}

func (l *ZapLogger) Error(ctx context.Context, msg string, fields ...port.Field) {
	fields = l.addRequestIDFromContext(ctx, fields)
	l.logger.Error(msg, l.convertFields(fields)...)
}

func (l *ZapLogger) Fatal(ctx context.Context, msg string, fields ...port.Field) {
	fields = l.addRequestIDFromContext(ctx, fields)
	l.logger.Fatal(msg, l.convertFields(fields)...)
}

func (l *ZapLogger) WithContext(ctx context.Context) port.LoggerPort {
	newLogger := &ZapLogger{
		logger: l.logger,
		fields: make([]port.Field, len(l.fields)),
	}
	copy(newLogger.fields, l.fields)

	requestID := GetRequestID(ctx)
	if requestID != "" {
		newLogger.fields = append(newLogger.fields, port.Field{Key: "request_id", Value: requestID})
	}

	return newLogger
}

func (l *ZapLogger) WithService(serviceName string) port.LoggerPort {
	newLogger := &ZapLogger{
		logger: l.logger,
		fields: make([]port.Field, len(l.fields)),
	}
	copy(newLogger.fields, l.fields)
	newLogger.fields = append(newLogger.fields, port.Field{Key: "service", Value: serviceName})

	return newLogger
}

func (l *ZapLogger) WithRequestID(requestID string) port.LoggerPort {
	newLogger := &ZapLogger{
		logger: l.logger,
		fields: make([]port.Field, len(l.fields)),
	}
	copy(newLogger.fields, l.fields)
	newLogger.fields = append(newLogger.fields, port.Field{Key: "request_id", Value: requestID})

	return newLogger
}

func (l *ZapLogger) WithField(key string, value interface{}) port.LoggerPort {
	newLogger := &ZapLogger{
		logger: l.logger,
		fields: make([]port.Field, len(l.fields)),
	}
	copy(newLogger.fields, l.fields)
	newLogger.fields = append(newLogger.fields, port.Field{Key: key, Value: value})

	return newLogger
}

func (l *ZapLogger) Flush() {
	_ = l.logger.Sync()
}
