package goredis

import (
	"crypto/tls"
	"time"

	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Option set the redis options.
type Option func(*options)

type options struct {
	dialTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
	tlsConfig    *tls.Config

	// Note: this field is only used for Init and InitSingle, and the other parameters will be ignored.
	singleOptions *redis.Options

	// Note: this field is only used for InitSentinel, and the other parameters will be ignored.
	sentinelOptions *redis.FailoverOptions

	// Note: this field is only used for InitCluster, and the other parameters will be ignored.
	clusterOptions *redis.ClusterOptions

	// deprecated: use tp instead
	enableTrace    bool
	tracerProvider *trace.TracerProvider
}

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// default settings
func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

// whether to enable trace, default off

// WithEnableTrace use trace, redis v8
// Deprecated: use WithEnableTracer instead
func WithEnableTrace() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTracing set redis tracer provider, redis v9
func WithTracing(tp *trace.TracerProvider) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDialTimeout set dail timeout
func WithDialTimeout(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithReadTimeout set read timeout
func WithReadTimeout(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithWriteTimeout set write timeout
func WithWriteTimeout(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTLSConfig set TLS config
func WithTLSConfig(c *tls.Config) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSingleOptions set single redis options
func WithSingleOptions(opt *redis.Options) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSentinelOptions set redis sentinel options
func WithSentinelOptions(opt *redis.FailoverOptions) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithClusterOptions set redis cluster options
func WithClusterOptions(opt *redis.ClusterOptions) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
