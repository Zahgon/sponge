package etcdcli

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

// Option set the etcd client options.
type Option func(*options)

type options struct {
	dialTimeout time.Duration // connection timeout, unit(second)

	username string
	password string

	isSecure           bool
	serverNameOverride string // etcd domain
	certFile           string // path to certificate file

	autoSyncInterval time.Duration // automatic synchronization of member list intervals
	logger           *zap.Logger

	// if you set this parameter, all fields above are invalid
	config *clientv3.Config
}

func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// WithDialTimeout set dial timeout
func WithDialTimeout(duration time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAuth set authentication
func WithAuth(username string, password string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSecure set tls
func WithSecure(serverNameOverride string, certFile string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAutoSyncInterval set auto sync interval value
func WithAutoSyncInterval(duration time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLog set logger
func WithLog(l *zap.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithConfig set etcd client config
func WithConfig(c *clientv3.Config) Option { _ = "STUB: not implemented"; return *new(Option) }
