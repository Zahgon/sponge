package consul

import (
	"sync"
	"sync/atomic"

	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"
)

type serviceSet struct {
	serviceName string
	watcher     map[*watcher]struct{}
	services    *atomic.Value
	lock        sync.RWMutex
}

func (s *serviceSet) broadcast(ss []*registry.ServiceInstance) { _ = "STUB: not implemented"; return }
