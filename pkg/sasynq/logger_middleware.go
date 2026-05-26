package sasynq

import (
	"context"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

var (
	// Print payload max length
	defaultMaxLength = 300
	// default zap log
	defaultLogger, _ = zap.NewProduction()
)

// LoggerOption set options.
type LoggerOption func(*loggerOptions)

type loggerOptions struct {
	logger    *zap.Logger
	zapSkip   int // default is 2
	maxLength int // default is 300
}

func (o *loggerOptions) apply(opts ...LoggerOption) { _ = "STUB: not implemented"; return }

func defaultLoggerOptions() *loggerOptions { _ = "STUB: not implemented"; return nil }

// WithLogger sets the logger to use for logging.
func WithLogger(l *zap.Logger) LoggerOption { _ = "STUB: not implemented"; return *new(LoggerOption) }

// WithMaxLength sets the maximum length of the payload to log.
func WithMaxLength(l int) LoggerOption { _ = "STUB: not implemented"; return *new(LoggerOption) }

// WithZapSkip sets the number of callers to skip when logging.
func WithZapSkip(s int) LoggerOption { _ = "STUB: not implemented"; return *new(LoggerOption) }

// LoggingMiddleware logs information about each processed task.
func LoggingMiddleware(opts ...LoggerOption) func(next asynq.Handler) asynq.Handler {
	_ = "STUB: not implemented"
	return nil
}

func getTaskID(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func getPayload(t *asynq.Task, maxLength int) zap.Field {
	_ = "STUB: not implemented"
	return *new(zap.Field)
}

// ------------------------------------------------------------------------------------------

type ZapLogger struct {
	zLog *zap.Logger
}

func NewZapLogger(l *zap.Logger, skip int) asynq.Logger {
	_ = "STUB: not implemented"
	return *new(asynq.Logger)
}

func (l *ZapLogger) Debug(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *ZapLogger) Info(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *ZapLogger) Warn(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *ZapLogger) Error(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *ZapLogger) Fatal(args ...interface{}) { _ = "STUB: not implemented"; return }
