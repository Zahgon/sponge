package httpsrv

import (
	"net/http"
)

var _ TLSer = (*TLSExternalConfig)(nil)

type TLSExternalConfig struct {
	certFile string
	keyFile  string
}

func NewTLSExternalConfig(certFile, keyFile string) *TLSExternalConfig {
	_ = "STUB: not implemented"
	return nil
}

func (c *TLSExternalConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *TLSExternalConfig) Run(server *http.Server) error { _ = "STUB: not implemented"; return nil }
