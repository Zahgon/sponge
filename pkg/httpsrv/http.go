package httpsrv

import (
	"context"
	"net/http"
)

// TLSer abstract different TLS operation schemes
type TLSer interface {
	Validate() error
	Run(server *http.Server) error
}

// Server TLS server.
type Server struct {
	scheme string       // e.g. http or https.
	server *http.Server // server is the HTTP server instance.

	tlser TLSer
}

// New returns a new Server with TLSer injected.
func New(server *http.Server, tlser ...TLSer) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) validate() error { _ = "STUB: not implemented"; return nil }

// Run starts the server according to the provided configuration.
func (s *Server) Run() error { _ = "STUB: not implemented"; return nil }

// no TLS mode specified, run in http mode.

// Shutdown gracefully shuts down the server and releases resources.
func (s *Server) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// runHTTP starts the server in http mode, without TLS.
func (s *Server) runHTTP() error { _ = "STUB: not implemented"; return nil }

// Scheme returns the scheme of the server, e.g. http or https.
func (s *Server) Scheme() string { _ = "STUB: not implemented"; return "" }
