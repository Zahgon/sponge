// Package ws provides a websocket server implementation.
package ws

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// ServerOption is a functional option for the Server.
type ServerOption func(*serverOptions)

type serverOptions struct {
	responseHeader      http.Header
	upgrader            *websocket.Upgrader
	noClientPingTimeout time.Duration
	zapLogger           *zap.Logger
}

func defaultServerOptions() *serverOptions { _ = "STUB: not implemented"; return nil }

// default upgrader
// allow all origins

func (o *serverOptions) apply(opts ...ServerOption) { _ = "STUB: not implemented"; return }

// WithResponseHeader sets the response header for the WebSocket upgrade response.
func WithResponseHeader(header http.Header) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

// WithUpgrader sets the WebSocket upgrader for the server.
func WithUpgrader(upgrader *websocket.Upgrader) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

// WithMaxMessageWaitPeriod sets the maximum waiting period for a message before closing the connection.
// Deprecated: use WithNoClientPingTimeout instead.
func WithMaxMessageWaitPeriod(period time.Duration) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

// WithNoClientPingTimeout sets the timeout for the client to send a ping message, if timeout, the connection will be closed.
func WithNoClientPingTimeout(timeout time.Duration) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

// WithServerLogger sets the logger for the server.
func WithServerLogger(l *zap.Logger) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

// --------------------------------------------------------------------------------------

// Conn is a WebSocket connection.
type Conn = websocket.Conn

// LoopFn is the function that is called for each WebSocket connection.
type LoopFn func(ctx context.Context, conn *Conn)

// Server is a WebSocket server.
type Server struct {
	upgrader *websocket.Upgrader

	w              http.ResponseWriter
	r              *http.Request
	responseHeader http.Header

	// If it is greater than 0, it means that the message waiting timeout mechanism is enabled
	//and the connection will be closed after the timeout, if it is 0, it means that the message
	// waiting timeout mechanism is not enabled.
	noClientPingTimeout time.Duration

	loopFn LoopFn

	zapLogger *zap.Logger
}

// NewServer creates a new WebSocket server.
func NewServer(w http.ResponseWriter, r *http.Request, loopFn LoopFn, opts ...ServerOption) *Server {
	_ = "STUB: not implemented"
	return nil
}

// Run runs the WebSocket server.
func (s *Server) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

//nolint

// Set initial read deadline

// Set up Ping handling for the connection,
// when the client sends a ping message, the server side triggers this callback function

// IsClientClose returns true if the error is caused by client close.
func IsClientClose(err error) bool { _ = "STUB: not implemented"; return false }
