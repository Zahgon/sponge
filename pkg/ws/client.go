package ws

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var (
	pingData = []byte("ping")
)

// ClientOption is a functional option for the client.
type ClientOption func(*clientOptions)

type clientOptions struct {
	dialer           *websocket.Dialer
	requestHeader    http.Header
	pingDialInterval time.Duration

	zapLogger *zap.Logger
}

func defaultClientOptions() *clientOptions { _ = "STUB: not implemented"; return nil }

func (o *clientOptions) apply(opts ...ClientOption) { _ = "STUB: not implemented"; return }

// WithDialer sets the dialer for the client.
func WithDialer(dialer *websocket.Dialer) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// WithRequestHeader sets the request header for the client.
func WithRequestHeader(header http.Header) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// WithPing sets the interval for sending ping message to the server.
func WithPing(interval time.Duration) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// WithClientLogger sets the logger for the client.
func WithClientLogger(l *zap.Logger) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// ----------------------------------------------------------------------------------

// Client is a wrapper of gorilla/websocket.
type Client struct {
	dialer        *websocket.Dialer
	requestHeader http.Header
	url           string
	conn          *websocket.Conn

	pingInterval time.Duration
	ctx          context.Context
	cancel       context.CancelFunc
	zapLogger    *zap.Logger
}

// NewClient creates a new client.
func NewClient(url string, opts ...ClientOption) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetConn returns the connection of the client.
func (c *Client) GetConn() *websocket.Conn { _ = "STUB: not implemented"; return nil }

// connect the websocket server.
func (c *Client) connect() error { _ = "STUB: not implemented"; return nil }

// TryReconnect tries to reconnect the websocket server.
func (c *Client) TryReconnect() error { _ = "STUB: not implemented"; return nil }

// ping websocket server, try to reconnect if connection failed.
func (c *Client) ping() { _ = "STUB: not implemented"; return }

// exit

// GetCtx returns the context of the client.
func (c *Client) GetCtx() context.Context {
	_ = "STUB: not implemented"

	// Close closes the connection.
	// Note: if set pingDialInterval, the Close method must be called, otherwise it will cause the goroutine to leak
	return *new(context.Context)
}

func (c *Client) Close() error { _ = "STUB: not implemented"; return nil }
