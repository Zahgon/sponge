package server

import (
	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"
)

// GrpcOption grpc settings
type GrpcOption func(*grpcOptions)

type grpcOptions struct {
	instance  *registry.ServiceInstance
	iRegistry registry.Registry
}

func defaultGrpcOptions() *grpcOptions { _ = "STUB: not implemented"; return nil }

func (o *grpcOptions) apply(opts ...GrpcOption) { _ = "STUB: not implemented"; return }

// WithGrpcRegistry registration services
func WithGrpcRegistry(iRegistry registry.Registry, instance *registry.ServiceInstance) GrpcOption {
	_ = "STUB: not implemented"
	return *new(GrpcOption)
}
