package httpsrv

import (
	"net/http"
	"time"
)

// TLSRemoteAPIOption set tlsRemoteAPIOptions.
type TLSRemoteAPIOption func(*tlsRemoteAPIOptions)

type tlsRemoteAPIOptions struct {
	headers  map[string]string
	timeout  time.Duration
	cacheDir string
}

func (o *tlsRemoteAPIOptions) apply(opts ...TLSRemoteAPIOption) { _ = "STUB: not implemented"; return }

func defaultTLSRemoteAPIOptions() *tlsRemoteAPIOptions { _ = "STUB: not implemented"; return nil }

// WithTLSRemoteAPIHeaders set headers for tlsRemoteAPI.
func WithTLSRemoteAPIHeaders(headers map[string]string) TLSRemoteAPIOption {
	_ = "STUB: not implemented"
	return *new(TLSRemoteAPIOption)
}

// WithTLSRemoteAPITimeout set timeout for tlsRemoteAPI.
func WithTLSRemoteAPITimeout(timeout time.Duration) TLSRemoteAPIOption {
	_ = "STUB: not implemented"
	return *new(TLSRemoteAPIOption)
}

// WithTLSRemoteAPICacheDir set cacheDir for tlsRemoteAPI.
func WithTLSRemoteAPICacheDir(cacheDir string) TLSRemoteAPIOption {
	_ = "STUB: not implemented"
	return *new(TLSRemoteAPIOption)
}

// -------------------------------------------------------------------------------------------

var _ TLSer = (*TLSRemoteAPIConfig)(nil)

// TLSRemoteAPIConfig implements certificate retrieval from other service API
type TLSRemoteAPIConfig struct {
	url      string            // Certificate PEM download URL
	headers  map[string]string // Optional: request headers
	timeout  time.Duration     // Optional: HTTP timeout
	cacheDir string            // Cache directory

	certFile string // Cached certificate file path
	keyFile  string // Cached private key file path

	httpClient *http.Client // Internal HTTP client
}

func NewTLSRemoteAPIConfig(url string, opts ...TLSRemoteAPIOption) *TLSRemoteAPIConfig {
	_ = "STUB: not implemented"
	return nil
}

func (c *TLSRemoteAPIConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *TLSRemoteAPIConfig) Run(server *http.Server) error { _ = "STUB: not implemented"; return nil }

// retry 3 times to download cert from API

// write cert and key to file

type TLSRemoteAPIResponse struct {
	CertFile []byte `json:"cert_file"`
	KeyFile  []byte `json:"key_file"`
}

func (c *TLSRemoteAPIConfig) downloadFile() (certData []byte, keyData []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
