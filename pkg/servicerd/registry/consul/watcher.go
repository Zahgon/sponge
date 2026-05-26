package consul

import (
	"context"

	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"
)

type watcher struct {
	event chan struct{}
	set   *serviceSet

	// for cancel
	ctx    context.Context
	cancel context.CancelFunc
}

func (w *watcher) Next() (services []*registry.ServiceInstance, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint

func (w *watcher) Stop() error { _ = "STUB: not implemented"; return nil }
