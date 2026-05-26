// Package glog provides a gorm logger implementation based on zap.
package glog

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type gormLogger struct {
	gLog         *zap.Logger
	requestIDKey string
	logLevel     logger.LogLevel
}

// NewCustomGormLogger custom gorm logger
func NewCustomGormLogger(l *zap.Logger, requestIDKey string, logLevel logger.LogLevel) logger.Interface {
	_ = "STUB: not implemented"
	return *new(logger.Interface)
}

// LogMode log mode
func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	_ = "STUB: not implemented"
	return *new(logger.Interface)
}

// Info print info
func (l *gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Warn print warn messages
func (l *gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Error print error messages
func (l *gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Trace print sql message
func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	_ = "STUB: not implemented"
	return
}

func requestIDField(ctx context.Context, requestIDKey string) zap.Field {
	_ = "STUB: not implemented"
	return *new(zap.Field)
}
