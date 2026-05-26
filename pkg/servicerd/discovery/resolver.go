// Package discovery is service discovery library, supports etcd, consul and nacos.
package discovery

import (
	"context"
	"net/url"

	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"

	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"
)

type discoveryResolver struct {
	w  registry.Watcher
	cc resolver.ClientConn

	ctx    context.Context
	cancel context.CancelFunc

	insecure         bool
	debugLogDisabled bool
}

func (r *discoveryResolver) watch() { _ = "STUB: not implemented"; return }

func (r *discoveryResolver) update(ins []*registry.ServiceInstance) {
	_ = "STUB: not implemented"
	return
}

//fmt.Printf("[resolver] Failed to parse discovery endpoint: %v\n", err)

// filter redundant endpoints

//fmt.Printf("[resolver] Zero endpoint found,refused to write, instances: %v\n", ins)

func (r *discoveryResolver) Close() { _ = "STUB: not implemented"; return }

func (r *discoveryResolver) ResolveNow(_ resolver.ResolveNowOptions) {
	_ = "STUB: not implemented"
	return
}

func parseAttributes(md map[string]string) *attributes.Attributes {
	_ = "STUB: not implemented"
	return nil
}

// parseEndpoint parses an Endpoint URL.
func parseEndpoint(endpoints []string, scheme string, isSecure bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IsSecure parses isSecure for Endpoint URL.
func IsSecure(u *url.URL) bool { _ = "STUB: not implemented"; return false }
