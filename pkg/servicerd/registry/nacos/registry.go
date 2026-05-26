// Package nacos is registered as a service using nacos.
package nacos

import (
	"context"

	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"

	"github.com/go-dev-frame/sponge/pkg/nacoscli"
	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"
)

var (
	_ registry.Registry  = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

type options struct {
	prefix  string
	weight  float64
	cluster string
	group   string
	kind    string
}

// Option is nacos option.
type Option func(o *options)

// WithPrefix with prefix path.
func WithPrefix(prefix string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithWeight with weight option.
func WithWeight(weight float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCluster with cluster option.
func WithCluster(cluster string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithGroup with group option.
func WithGroup(group string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDefaultKind with default kind option.
func WithDefaultKind(kind string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Registry is nacos registry.
type Registry struct {
	opts options
	cli  naming_client.INamingClient
}

// NewRegistry instantiating the nacos registry
func NewRegistry(nacosIPAddr string, nacosPort int, nacosNamespaceID string,
	id string, instanceName string, instanceEndpoints []string,
	opts ...nacoscli.Option) (registry.Registry, *registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return *new(registry.Registry), nil, nil
}

// NewRegistryWithOptions instantiating the nacos registry, opts can be set by nacoscli.WithXXX and nacos.WithXXX functions.
func NewRegistryWithOptions(nacosIPAddr string, nacosPort int, nacosNamespaceID string,
	id string, instanceName string, instanceEndpoints []string,
	opts ...interface{}) (registry.Registry, *registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return *new(registry.Registry), nil, nil
}

// New new a nacos registry.
func New(cli naming_client.INamingClient, opts ...Option) (r *Registry) {
	_ = "STUB: not implemented"
	return nil
}

// Register the registration.
func (r *Registry) Register(_ context.Context, si *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

// Deregister the registration.
func (r *Registry) Deregister(_ context.Context, service *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

// Watch creates a watcher according to the service name.
func (r *Registry) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(registry.Watcher), nil
}

// GetService return the service instances in memory according to the service name.
func (r *Registry) GetService(_ context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
