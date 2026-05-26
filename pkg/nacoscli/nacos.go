// Package nacoscli provides for getting the configuration from the nacos configuration center and parse it into a structure.
package nacoscli

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
)

// Params nacos parameters
type Params struct {
	IPAddr      string // server address
	Port        uint64 // port
	Scheme      string // http or grpc
	ContextPath string // path
	// if you set this parameter, the above fields(IPAddr, Port, Scheme, ContextPath) are invalid
	serverConfigs []constant.ServerConfig

	NamespaceID string // namespace id
	// if you set this parameter, the above field(NamespaceID) is invalid
	clientConfig *constant.ClientConfig

	Group  string // group, example: dev, prod, test
	DataID string // config file id
	Format string // configuration file type: json,yaml,toml
}

func (p *Params) valid() error { _ = "STUB: not implemented"; return nil }

func setParams(params *Params, opts ...Option) { _ = "STUB: not implemented"; return }

// create clientConfig

// create serverConfig

// GetConfig get configuration from nacos
func GetConfig(params *Params, opts ...Option) (string, []byte, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// create a dynamic configuration client

// read config content

// Init get configuration from nacos and parse to struct, use for configuration center
//
// Deprecated: use GetConfig instead.
func Init(_ interface{}, _ *Params, _ ...Option) error { _ = "STUB: not implemented"; return nil }

// NewNamingClient create a service registration and discovery of nacos client.
// Note: If parameter WithClientConfig is set, nacosNamespaceID is invalid,
// if parameter WithServerConfigs is set, nacosIPAddr and nacosPort are invalid.
func NewNamingClient(nacosIPAddr string, nacosPort int, nacosNamespaceID string, opts ...Option) (naming_client.INamingClient, error) {
	_ = "STUB: not implemented"
	return *new(naming_client.INamingClient), nil
}
