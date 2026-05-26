package mysql

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Option set the mysql options.
type Option func(*options)

type options struct {
	isLog         bool
	slowThreshold time.Duration

	maxIdleConns    int
	maxOpenConns    int
	connMaxLifetime time.Duration

	disableForeignKey bool
	enableTrace       bool

	requestIDKey string
	gLog         *zap.Logger
	logLevel     logger.LogLevel

	slavesDsn  []string
	mastersDsn []string

	plugins []gorm.Plugin
}

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// default settings
func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

// whether to output logs, default off
// if greater than 0, only print logs that are longer than the threshold, higher priority than isLog

// set the maximum number of connections in the idle connection pool
// set the maximum number of open database connections
// sets the maximum amount of time a connection can be reused

// disables the use of foreign keys, true is recommended for production environments, enabled by default
// whether to enable link tracing, default is off

// request id key
// custom logger
// default logLevel

// WithLogging set log sql, If l=nil, the gorm log library will be used
func WithLogging(l *zap.Logger, level ...logger.LogLevel) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSlowThreshold Set sql values greater than the threshold
func WithSlowThreshold(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxIdleConns set max idle conns
func WithMaxIdleConns(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxOpenConns set max open conns
func WithMaxOpenConns(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithConnMaxLifetime set conn max lifetime
func WithConnMaxLifetime(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableForeignKey use foreign keys
func WithEnableForeignKey() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableTrace use trace
func WithEnableTrace() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLogRequestIDKey log request id
func WithLogRequestIDKey(key string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRWSeparation setting read-write separation
func WithRWSeparation(slavesDsn []string, mastersDsn ...string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGormPlugin setting gorm plugin
func WithGormPlugin(plugins ...gorm.Plugin) Option { _ = "STUB: not implemented"; return *new(Option) }
