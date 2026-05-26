// Package resolve is setting grpc client-side load balancing policy.
package resolve

import (
	"sync"

	"google.golang.org/grpc/resolver"
)

var mutex = &sync.Mutex{}

// Register address and serviceName
func Register(scheme string, serviceName string, address []string) string {
	_ = "STUB: not implemented"
	return ""
}

// ResolverBuilder resolver struct
type ResolverBuilder struct {
	scheme      string
	serviceName string
	addrs       []string
	path        string
}

// Build resolver
func (r *ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	_ = "STUB: not implemented"
	return *new(resolver.Resolver), nil
}

// Scheme get scheme
func (r *ResolverBuilder) Scheme() string { _ = "STUB: not implemented"; return "" }

type blResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (b *blResolver) start() { _ = "STUB: not implemented"; return }

// ResolveNow Resolve now
func (*blResolver) ResolveNow(_ resolver.ResolveNowOptions) {
	_ = "STUB: not implemented"

	// Close resolver
	return
}

func (*blResolver) Close() { _ = "STUB: not implemented"; return }
