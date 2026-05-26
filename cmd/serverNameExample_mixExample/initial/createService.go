package initial

import (
	"github.com/go-dev-frame/sponge/pkg/app"
	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"
)

// CreateServices create grpc or http service
func CreateServices() []app.IServer { _ = "STUB: not implemented"; return nil }

// create a http service

// create a grpc service

// register service with consul or etcd or nacos, select one of them to use
func registerService(scheme string, host string, port int) (registry.Registry, *registry.ServiceInstance) {
	_ = "STUB: not implemented"
	return *new(registry.Registry), nil
}
