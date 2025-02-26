package logging

import (
	"context"
	"errors"
	"fmt"
	"time"

	port "github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormLoggerAdapter struct {
	logger        port.LoggerPort
	SlowThreshold time.Duration
	LogLevel      logger.LogLevel
}

func NewGormLogger(appLogger port.LoggerPort) *GormLoggerAdapter {
	return &GormLoggerAdapter{
		logger:        appLogger,
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Info,
	}
}

func (l *GormLoggerAdapter) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *GormLoggerAdapter) Info(ctx context.Context, msg string, args ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.logger.Info(ctx, fmt.Sprintf(msg, args...))
	}
}

func (l *GormLoggerAdapter) Warn(ctx context.Context, msg string, args ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.logger.Warn(ctx, fmt.Sprintf(msg, args...))
	}
}

func (l *GormLoggerAdapter) Error(ctx context.Context, msg string, args ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.logger.Error(ctx, fmt.Sprintf(msg, args...))
	}
}

func (l *GormLoggerAdapter) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error && !errors.Is(err, gorm.ErrRecordNotFound):
		sql, rows := fc()
		l.logger.Error(ctx, "SQL Error",
			port.Field{Key: "sql", Value: sql},
			port.Field{Key: "rows", Value: rows},
			port.Field{Key: "duration_ms", Value: elapsed.Milliseconds()},
			port.Field{Key: "error", Value: err.Error()},
		)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		l.logger.Warn(ctx, "Slow SQL Query",
			port.Field{Key: "sql", Value: sql},
			port.Field{Key: "rows", Value: rows},
			port.Field{Key: "duration_ms", Value: elapsed.Milliseconds()},
			port.Field{Key: "threshold_ms", Value: l.SlowThreshold.Milliseconds()},
		)
	case l.LogLevel >= logger.Info:
		sql, rows := fc()
		l.logger.Debug(ctx, "SQL Query",
			port.Field{Key: "sql", Value: sql},
			port.Field{Key: "rows", Value: rows},
			port.Field{Key: "duration_ms", Value: elapsed.Milliseconds()},
		)
	}
}
