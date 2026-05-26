package server

import (
	"net/http"

	"github.com/go-dev-frame/sponge/pkg/app"
	"github.com/go-dev-frame/sponge/pkg/httpsrv"
	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"

	"github.com/go-dev-frame/sponge/internal/config"
)

var _ app.IServer = (*httpServer)(nil)

type httpServer struct {
	addr   string
	server *httpsrv.Server

	instance  *registry.ServiceInstance
	iRegistry registry.Registry
}

// Start http service
func (s *httpServer) Start() error { _ = "STUB: not implemented"; return nil }

//nolint

// Stop http service
func (s *httpServer) Stop() error { _ = "STUB: not implemented"; return nil }

//nolint

// String comment
func (s *httpServer) String() string { _ = "STUB: not implemented"; return "" }

func newServer(server *http.Server, tls config.TLS) *httpsrv.Server {
	_ = "STUB: not implemented"
	return nil
}

// enable http redirect to https, port 80 to 443, default is false
//httpsrv.WithTLSEncryptEnableRedirect(),

// default is http, no tls

// NewHTTPServer creates a new http server
func NewHTTPServer(addr string, opts ...HTTPOption) app.IServer {
	_ = "STUB: not implemented"
	return *new(app.IServer)
}

//ReadTimeout:    time.Second*30,
//WriteTimeout:   time.Second*60,

// delete the templates code start

// NewHTTPServer_pbExample creates a new web server
func NewHTTPServer_pbExample(addr string, opts ...HTTPOption) app.IServer {
	_ = "STUB: not implemented" //nolint
	return *new(app.IServer)
}

//ReadTimeout:    time.Second*30,
//WriteTimeout:   time.Second*60,

// delete the templates code end
