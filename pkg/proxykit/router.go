package proxykit

import (
	"net/http"
	"sync"
)

// ManagementRequest is for the management API.
type ManagementRequest struct {
	PrefixPath  string            `json:"prefixPath"`
	Targets     []string          `json:"targets"`
	HealthCheck HealthCheckConfig `json:"healthCheck"`
}

// Route holds all components for a specific routing rule.
type Route struct {
	PrefixPath string
	Backends   []*Backend
	Balancer   Balancer
	Proxy      *Proxy
	mu         sync.RWMutex
}

// RouteManager manages all routing rules.
type RouteManager struct {
	routes map[string]*Route
	mu     sync.RWMutex
}

// NewRouteManager creates a new manager.
func NewRouteManager() *RouteManager { _ = "STUB: not implemented"; return nil }

// AddRoute adds a new routing rule and configures its proxy to strip the given prefix.
func (m *RouteManager) AddRoute(prefixPath string, balancer Balancer) (*Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRoute safely retrieves a route.
func (m *RouteManager) GetRoute(prefixPath string) (*Route, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// HandleAddBackends handles the HTTP request to add new backends to a route.
func (m *RouteManager) HandleAddBackends(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// HandleRemoveBackends handles the HTTP request to remove backends from a route.
func (m *RouteManager) HandleRemoveBackends(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// HandleGetBackend handles the HTTP request to get a backend in a route.
func (m *RouteManager) HandleGetBackend(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// HandleListBackends handles the HTTP request to list backends in a route.
func (m *RouteManager) HandleListBackends(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func containsTarget(backends []*Backend, targetStr string) bool {
	_ = "STUB: not implemented"
	return false
}

func containsString(slice []string, str string) bool { _ = "STUB: not implemented"; return false }

func AnyRelativePath(prefixPath string) string { _ = "STUB: not implemented"; return "" }
