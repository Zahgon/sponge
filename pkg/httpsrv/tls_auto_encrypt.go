package httpsrv

import (
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

// TLSEncryptOption set tlsEncryptOptions.
type TLSEncryptOption func(*tlsEncryptOptions)

type tlsEncryptOptions struct {
	cacheDir       string
	httpAddr       string
	enableRedirect bool
}

func (o *tlsEncryptOptions) apply(opts ...TLSEncryptOption) { _ = "STUB: not implemented"; return }

func defaultTLSEncryptOptions() *tlsEncryptOptions { _ = "STUB: not implemented"; return nil }

// WithTLSEncryptCacheDir sets the directory to store Let's Encrypt certificates.
func WithTLSEncryptCacheDir(cacheDir string) TLSEncryptOption {
	_ = "STUB: not implemented"
	return *new(TLSEncryptOption)
}

// WithTLSEncryptEnableRedirect enables the HTTP-to-HTTPS redirect service.
// By default, it listens on ":80".
// An optional httpAddr can be provided to specify a different address.
func WithTLSEncryptEnableRedirect(httpAddr ...string) TLSEncryptOption {
	_ = "STUB: not implemented"
	return *new(TLSEncryptOption)
}

// ------------------------------------------------------------------------------------------

var _ TLSer = (*TLSAutoEncryptConfig)(nil)

type TLSAutoEncryptConfig struct {
	domain         string // The domain to request a certificate for in production mode.
	email          string // Used for Let's Encrypt account registration and important notices.
	cacheDir       string // Directory to store Let's Encrypt certificates.
	httpAddr       string // Listen address for the HTTP redirect service (defaults to :80).
	enableRedirect bool   // Enable HTTP-to-HTTPS redirect service (default: false).

	m              *autocert.Manager // Manages certificates automatically.
	redirectServer *http.Server      // The HTTP redirect server.
}

func NewTLSEAutoEncryptConfig(domain string, email string, opts ...TLSEncryptOption) *TLSAutoEncryptConfig {
	_ = "STUB: not implemented"
	return nil
}

func (c *TLSAutoEncryptConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *TLSAutoEncryptConfig) Run(server *http.Server) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *TLSAutoEncryptConfig) redirectHTTP() error { _ = "STUB: not implemented"; return nil }

// Handles ACME challenges and redirection.

func (c *TLSAutoEncryptConfig) shutDownRedirectHTTP() error { _ = "STUB: not implemented"; return nil }
