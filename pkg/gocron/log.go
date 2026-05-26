package gocron

import (
	"go.uber.org/zap"
)

var (
	SecondType = 0
	MinuteType = 1
)

var defaultLog, _ = zap.NewProduction()

type options struct {
	zapLog           *zap.Logger
	isOnlyPrintError bool // default false

	granularity int // 0: second, 1: minute
}

func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// Option set the cron options.
type Option func(*options)

// WithGranularity set log
func WithGranularity(granularity int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLog set granularity
func WithLog(log *zap.Logger, isOnlyPrintError ...bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type zapLog struct {
	zapLog           *zap.Logger
	isOnlyPrintError bool
}

// Info print info
func (l *zapLog) Info(msg string, keysAndValues ...interface{}) { _ = "STUB: not implemented"; return }

// 忽略wake

// Error print error
func (l *zapLog) Error(err error, msg string, keysAndValues ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func parseKVs(kvs interface{}) []zap.Field { _ = "STUB: not implemented"; return nil }

//nolint

// replace id with task name
