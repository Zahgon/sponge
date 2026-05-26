package proxykit

import (
	"net/http"
)

// Proxy is a reverse proxy that implements the http.Handler interface.
type Proxy struct {
	balancer Balancer
}

// NewProxy creates a new reverse proxy instance.
func NewProxy(balancer Balancer) (*Proxy, error) { _ = "STUB: not implemented"; return nil, nil }

// ServeHTTP handles incoming HTTP requests and forwards them to the backend
// selected by the load balancer.
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	// Select a healthy backend according to the load balancing strategy.
	return
}

// Increase the active connection count, and ensure it is decremented
// when the request completes.
