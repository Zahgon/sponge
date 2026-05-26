package consul

import (
	"context"

	"github.com/hashicorp/consul/api"

	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"
)

// Client is consul client config
type Client struct {
	client *api.Client
	ctx    context.Context
	cancel context.CancelFunc
}

// NewClient creates consul client
func NewClient(cli *api.Client) *Client { _ = "STUB: not implemented"; return nil }

// Service get services from consul
func (d *Client) Service(ctx context.Context, service string, index uint64, passingOnly bool) ([]*registry.ServiceInstance, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Register register service instance to consul
func (d *Client) Register(_ context.Context, svc *registry.ServiceInstance, enableHealthCheck bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Deregister deregister service by service ID
func (d *Client) Deregister(_ context.Context, serviceID string) error {
	_ = "STUB: not implemented"
	return nil
}
