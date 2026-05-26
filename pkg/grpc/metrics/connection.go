package metrics

import (
	"net"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

// ConnectionOption set connection option
type ConnectionOption func(*connectionOptions)

type connectionOptions struct {
	zapLogger       *zap.Logger
	connectionGauge prometheus.Gauge
}

func defaultConnectionOptions() *connectionOptions { _ = "STUB: not implemented"; return nil }

func (o *connectionOptions) apply(opts ...ConnectionOption) { _ = "STUB: not implemented"; return }

// WithConnectionsLogger set logger for connection
func WithConnectionsLogger(l *zap.Logger) ConnectionOption {
	_ = "STUB: not implemented"
	return *new(ConnectionOption)
}

// WithConnectionsGauge set prometheus gauge for connections
func WithConnectionsGauge() ConnectionOption {
	_ = "STUB: not implemented"
	return *new(ConnectionOption)
}

// ------------------------------------------------------------------------------------------

// CustomConn custom connections, intercept disconnected behavior
type CustomConn struct {
	net.Conn
	listener *CustomListener
}

// CustomListener custom listener for counting connections
type CustomListener struct {
	net.Listener
	activeConnections int
	mu                sync.Mutex
	zapLogger         *zap.Logger
	connectionGauge   prometheus.Gauge
}

// Accept waits for and returns the next connection to the listener.
func (l *CustomListener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// GetActiveConnections returns the number of active connections.
func (l *CustomListener) GetActiveConnections() int { _ = "STUB: not implemented"; return 0 }

// closes the connection and decrements the active connections count.
func (l *CustomListener) closeConnection(clientAddr string) { _ = "STUB: not implemented"; return }

// Close closes the listener, any blocked except operations will be unblocked and return errors.
func (c *CustomConn) Close() error { _ = "STUB: not implemented"; return nil }

// NewCustomListener creates a new custom listener.
func NewCustomListener(listener net.Listener, opts ...ConnectionOption) *CustomListener {
	_ = "STUB: not implemented"
	return nil
}
