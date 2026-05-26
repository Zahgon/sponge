package proxykit

import (
	"net/http"
	"sync"

	"go.uber.org/zap"
)

// Middleware is a function that takes a http.Handler and returns a http.Handler,
// used to build a middleware chain.
type Middleware func(http.Handler) http.Handler

// Chain links multiple middlewares together to form a single http.Handler.
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	_ = "STUB: not implemented"
	// Start from the last middleware and wrap backwards,
	// so that the first middleware is the outermost layer.
	return *new(http.Handler)
}

type logger struct {
	*zap.Logger
}

func newLogger() *logger { _ = "STUB: not implemented"; return nil }

func (l *logger) Printf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Println(v ...interface{}) { _ = "STUB: not implemented"; return }

var log = newLogger()

var doOnce sync.Once

func SetLogger(l *zap.Logger) { _ = "STUB: not implemented"; return }
