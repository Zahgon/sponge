// Package consul is registered as a service using consul.
package consul

import (
	"context"
	"sync"

	"github.com/hashicorp/consul/api"

	"github.com/go-dev-frame/sponge/pkg/consulcli"
	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"
)

var (
	_ registry.Registry  = &Registry{}
	_ registry.Discovery = &Registry{}
)

// Option is consul registry option.
type Option func(*Registry)

// WithHealthCheck with registry health check option.
func WithHealthCheck(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// Config is consul registry config
type Config struct {
	*api.Config
}

// Registry is consul registry
type Registry struct {
	cli               *Client
	enableHealthCheck bool
	registry          map[string]*serviceSet
	lock              sync.RWMutex
}

// NewRegistry instantiating the consul registry
// Note: If the consulcli.WithConfig(*api.Config) parameter is set, the consulAddr parameter is ignored!
func NewRegistry(consulAddr string, id string, instanceName string, instanceEndpoints []string, opts ...consulcli.Option) (registry.Registry, *registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return *new(registry.Registry), nil, nil
}

// NewRegistryWithOptions instantiating the consul registry, opts can be set by etcdcli.WithXXX and etcd.WithXXX functions.
// Note: If the consulcli.WithConfig(*api.Config) parameter is set, the consulAddr parameter is ignored!
func NewRegistryWithOptions(consulAddr string, id string, instanceName string, instanceEndpoints []string, opts ...interface{}) (registry.Registry, *registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return *new(registry.Registry), nil, nil
}

// New create a consul registry
func New(apiClient *api.Client, opts ...Option) *Registry { _ = "STUB: not implemented"; return nil }

// Register register service
func (r *Registry) Register(ctx context.Context, svc *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

// Deregister deregister service
func (r *Registry) Deregister(ctx context.Context, svc *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	// NOTE: invoke the func Deregister will block when err is not nil
	return nil
}

// GetService return service by name
func (r *Registry) GetService(_ context.Context, name string) (services []*registry.ServiceInstance, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint

// ListServices return service list.
func (r *Registry) ListServices() (allServices map[string][]*registry.ServiceInstance, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint

// Watch resolve service by name
func (r *Registry) Watch(_ context.Context, name string) (registry.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(registry.Watcher), nil
}

// If the service has a value, it needs to be pushed to the watcher,
// otherwise the initial data may be blocked forever during the watch.

func (r *Registry) resolve(ss *serviceSet) { _ = "STUB: not implemented"; return }
