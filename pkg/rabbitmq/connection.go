// Package rabbitmq is a go wrapper for github.com/rabbitmq/amqp091-go
//
// producer and consumer using the five types direct, topic, fanout, headers, x-delayed-message.
// publisher and subscriber using the fanout message type.
package rabbitmq

import (
	"crypto/tls"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

// DefaultURL default rabbitmq url
const DefaultURL = "amqp://guest:guest@localhost:5672/"

var defaultLogger, _ = zap.NewProduction()

// ConnectionOption connection option.
type ConnectionOption func(*connectionOptions)

type connectionOptions struct {
	tlsConfig     *tls.Config   // tls config, if the url is amqps this field must be set
	reconnectTime time.Duration // reconnect time interval, default is 3s

	zapLog *zap.Logger
}

func (o *connectionOptions) apply(opts ...ConnectionOption) { _ = "STUB: not implemented"; return }

// default connection settings
func defaultConnectionOptions() *connectionOptions { _ = "STUB: not implemented"; return nil }

// WithTLSConfig set tls config option.
func WithTLSConfig(tlsConfig *tls.Config) ConnectionOption {
	_ = "STUB: not implemented"
	return *new(ConnectionOption)
}

// WithReconnectTime set reconnect time interval option.
func WithReconnectTime(d time.Duration) ConnectionOption {
	_ = "STUB: not implemented"
	return *new(ConnectionOption)
}

// WithLogger set logger option.
func WithLogger(zapLog *zap.Logger) ConnectionOption {
	_ = "STUB: not implemented"
	return *new(ConnectionOption)
}

// -------------------------------------------------------------------------------------------

// Connection rabbitmq connection
type Connection struct {
	mutex sync.Mutex

	url           string
	tlsConfig     *tls.Config
	reconnectTime time.Duration
	exit          chan struct{}
	zapLog        *zap.Logger

	conn        *amqp.Connection
	blockChan   chan amqp.Blocking
	closeChan   chan *amqp.Error
	isConnected bool
}

// NewConnection rabbitmq connection
func NewConnection(url string, opts ...ConnectionOption) (*Connection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func connect(url string, tlsConfig *tls.Config) (*amqp.Connection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CheckConnected rabbitmq connection
func (c *Connection) CheckConnected() bool { _ = "STUB: not implemented"; return false }

func (c *Connection) monitor() { _ = "STUB: not implemented"; return }

// wait for reconnect

// set new connection

// Close rabbitmq connection
func (c *Connection) Close() { _ = "STUB: not implemented"; return }

func (c *Connection) closeConn() error { _ = "STUB: not implemented"; return nil }
