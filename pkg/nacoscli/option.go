package nacoscli

import (
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
)

type options struct {
	username string
	password string

	// if set the clientConfig, the above fields(username, password) are invalid
	clientConfig  *constant.ClientConfig
	serverConfigs []constant.ServerConfig
}

func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

// Option set the nacos client options.
type Option func(*options)

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// WithAuth set authentication
func WithAuth(username string, password string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithClientConfig set nacos client config
func WithClientConfig(clientConfig *constant.ClientConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithServerConfigs set nacos server config
func WithServerConfigs(serverConfigs []constant.ServerConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
