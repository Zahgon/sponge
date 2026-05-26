package sasynq

import (
	"github.com/hibiken/asynq"
)

// Server is a wrapper around asynq.Server providing integrated features.
type Server struct {
	srv *asynq.Server
	mux *asynq.ServeMux
	cfg ServerConfig
}

// NewServer creates a new consumer server.
func NewServer(redisCfg RedisConfig, serverCfg ServerConfig) *Server {
	_ = "STUB: not implemented"
	return nil
}

// Mux returns the underlying ServeMux to register handlers.
func (s *Server) Mux() *asynq.ServeMux {
	_ = "STUB: not implemented"

	// Use adds middleware to the server's handler chain.
	return nil
}

func (s *Server) Use(middlewares ...asynq.MiddlewareFunc) { _ = "STUB: not implemented"; return }

// Register a task processor
func (s *Server) Register(typeName string, handler asynq.Handler) {
	_ = "STUB: not implemented"
	return
}

// RegisterFunc a task handler function
func (s *Server) RegisterFunc(typeName string, handlerFunc asynq.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

// Run runs the asynq server in a separate goroutine
func (s *Server) Run() { _ = "STUB: not implemented"; return }

// Shutdown the server.
func (s *Server) Shutdown() { _ = "STUB: not implemented"; return }

// WaitShutdown for interrupt signals for graceful shutdown the server.
func (s *Server) WaitShutdown() { _ = "STUB: not implemented"; return }
