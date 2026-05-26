package etcd

import (
	"context"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/go-dev-frame/sponge/pkg/etcdcli"
	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"
)

var (
	_ registry.Registry  = &Registry{}
	_ registry.Discovery = &Registry{}
)

// Option is etcd registry option.
type Option func(o *options)

type options struct {
	ctx       context.Context
	namespace string
	ttl       time.Duration
	maxRetry  int
}

func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

// WithContext with registry context.
func WithContext(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNamespace with registry namespace.
func WithNamespace(ns string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRegisterTTL with register ttl.
func WithRegisterTTL(ttl time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxRetry set max retry times.
func WithMaxRetry(num int) Option { _ = "STUB: not implemented"; return *new(Option) }

// NewRegistry instantiating the etcd registry
// Note: If the etcdcli.WithConfig(*clientv3.Config) parameter is set, the etcdEndpoints parameter is ignored!
func NewRegistry(etcdEndpoints []string, id string, instanceName string, instanceEndpoints []string, opts ...etcdcli.Option) (registry.Registry, *registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return *new(registry.Registry), nil, nil
}

// NewRegistryWithOptions instantiating the etcd registry, opts can be set by etcdcli.WithXXX and etcd.WithXXX functions.
// Note: If the etcdcli.WithConfig(*clientv3.Config) parameter is set, the etcdEndpoints parameter is ignored!
func NewRegistryWithOptions(etcdEndpoints []string, id string, instanceName string, instanceEndpoints []string, opts ...interface{}) (registry.Registry, *registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return *new(registry.Registry), nil, nil
}

// Registry is etcd registry.
type Registry struct {
	opts   *options
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}

// New create a etcd registry
func New(client *clientv3.Client, opts ...Option) (r *Registry) {
	_ = "STUB: not implemented"
	return nil
}

// Register the registration.
func (r *Registry) Register(ctx context.Context, service *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

// Deregister the registration.
func (r *Registry) Deregister(ctx context.Context, service *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

// GetService return the service instances in memory according to the service name.
func (r *Registry) GetService(ctx context.Context, name string) ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Watch creates a watcher according to the service name.
func (r *Registry) Watch(ctx context.Context, name string) (registry.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(registry.Watcher), nil
}

// registerWithKV create a new lease, return current leaseID
func (r *Registry) registerWithKV(ctx context.Context, key string, value string) (clientv3.LeaseID, error) {
	_ = "STUB: not implemented"
	return *new(clientv3.LeaseID), nil
}

func (r *Registry) heartBeat(ctx context.Context, leaseID clientv3.LeaseID, key string, value string) {
	_ = "STUB: not implemented"
	return
}

//nolint

// try to registerWithKV

// prevent infinite blocking

// retry failed

// channel closed due to context cancel

// need to retry registration
