package httpsrv

import (
	"net"
	"net/http"
)

// TLSSelfSignedOption set tlsSelfSignedOptions.
type TLSSelfSignedOption func(*tlsSelfSignedOptions)

type tlsSelfSignedOptions struct {
	cacheDir       string
	expirationDays int
	wanIPs         []string // IP addresses to include in the certificate.
}

func (o *tlsSelfSignedOptions) apply(opts ...TLSSelfSignedOption) {
	_ = "STUB: not implemented"
	return
}

func defaultTLSSelfSignedOptions() *tlsSelfSignedOptions { _ = "STUB: not implemented"; return nil }

// WithTLSSelfSignedCacheDir sets the cache directory for self-signed certificates.
func WithTLSSelfSignedCacheDir(cacheDir string) TLSSelfSignedOption {
	_ = "STUB: not implemented"
	return *new(TLSSelfSignedOption)
}

// WithTLSSelfSignedExpirationDays sets the expiration days for self-signed certificates.
func WithTLSSelfSignedExpirationDays(expirationDays int) TLSSelfSignedOption {
	_ = "STUB: not implemented"
	return *new(TLSSelfSignedOption)
}

// WithTLSSelfSignedWanIPs sets the IP addresses to include in the certificate.
func WithTLSSelfSignedWanIPs(wanIPs ...string) TLSSelfSignedOption {
	_ = "STUB: not implemented"
	return *new(TLSSelfSignedOption)
}

// ------------------------------------------------------------------------------------------

var _ TLSer = (*TLSSelfSignedConfig)(nil)

type TLSSelfSignedConfig struct {
	cacheDir       string
	certFile       string
	keyFile        string
	expirationDays int
	wanIPs         []string // IP addresses to include in the certificate.
}

func NewTLSSelfSignedConfig(opts ...TLSSelfSignedOption) *TLSSelfSignedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (c *TLSSelfSignedConfig) Validate() error { _ = "STUB: not implemented"; return nil }

// generateCert checks for and generates a certificate for local development.
// If the certificate doesn't exist or expires in less than 30 days, it generates a new one.
func (c *TLSSelfSignedConfig) generateCert() error {
	_ = "STUB: not implemented"
	// Check if the certificate file exists.
	return nil
}

// If the cert exists, decode it and check its expiration.

// Renewal threshold is 30 days.

// createCert is a helper function that creates the certificate and key files.
func (c *TLSSelfSignedConfig) createCert() error { _ = "STUB: not implemented"; return nil }

// Write certificate file.

// Write key file.

func (c *TLSSelfSignedConfig) Run(server *http.Server) error { _ = "STUB: not implemented"; return nil }

func getLANIP() string { _ = "STUB: not implemented"; return "" }

func isValidIP(ip net.IP) bool { _ = "STUB: not implemented"; return false }
