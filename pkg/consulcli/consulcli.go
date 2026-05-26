// Package consulcli is connecting to the consul service client.
package consulcli

import (
	"github.com/hashicorp/consul/api"
)

// Init connecting to the consul service
// Note: If the WithConfig(*api.Config) parameter is set, the addr parameter is ignored!
func Init(addr string, opts ...Option) (*api.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
