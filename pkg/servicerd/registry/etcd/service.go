// Package etcd is registered as a service using etcd.
package etcd

import (
	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"
)

func marshal(si *registry.ServiceInstance) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// nolint
func unmarshal(data []byte) (si *registry.ServiceInstance, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
