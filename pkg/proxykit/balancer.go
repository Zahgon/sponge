package proxykit

import (
	"errors"
	"net/http"
	"sync"
)

var (
	ErrNoHealthyBackends = errors.New("no healthy backends available")
)

// Balancer is an interface that must be implemented for all load balancing strategies.
type Balancer interface {
	Next(r *http.Request) (*Backend, error)
	GetBackends() []*Backend
	AddBackend(b *Backend)
	RemoveBackend(b *Backend)
}

// --- RoundRobin ---

type RoundRobin struct {
	backends []*Backend
	next     uint32
	mu       sync.RWMutex
}

func NewRoundRobin(backends []*Backend) *RoundRobin { _ = "STUB: not implemented"; return nil }

func (r *RoundRobin) Next(_ *http.Request) (*Backend, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RoundRobin) getHealthy() []*Backend { _ = "STUB: not implemented"; return nil }

// GetBackends implement the Balancer interface
func (r *RoundRobin) GetBackends() []*Backend { _ = "STUB: not implemented"; return nil }

// AddBackend implement the Balancer interface
func (r *RoundRobin) AddBackend(b *Backend) { _ = "STUB: not implemented"; return }

// RemoveBackend implement the Balancer interface
func (r *RoundRobin) RemoveBackend(b *Backend) { _ = "STUB: not implemented"; return }

// --- LeastConnections ---

type LeastConnections struct {
	backends []*Backend
	mu       sync.RWMutex
}

func NewLeastConnections(backends []*Backend) *LeastConnections {
	_ = "STUB: not implemented"
	return nil
}

func (lc *LeastConnections) Next(_ *http.Request) (*Backend, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lc *LeastConnections) GetBackends() []*Backend { _ = "STUB: not implemented"; return nil }

func (lc *LeastConnections) AddBackend(b *Backend) { _ = "STUB: not implemented"; return }

func (lc *LeastConnections) RemoveBackend(b *Backend) { _ = "STUB: not implemented"; return }

// --- IPHash ---

type IPHash struct {
	backends []*Backend
	mu       sync.RWMutex
}

func NewIPHash(backends []*Backend) *IPHash { _ = "STUB: not implemented"; return nil }

func (h *IPHash) Next(r *http.Request) (*Backend, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *IPHash) getHealthy() []*Backend { _ = "STUB: not implemented"; return nil }

func (h *IPHash) GetBackends() []*Backend { _ = "STUB: not implemented"; return nil }

func (h *IPHash) AddBackend(b *Backend) { _ = "STUB: not implemented"; return }

func (h *IPHash) RemoveBackend(b *Backend) { _ = "STUB: not implemented"; return }

func getClientIP(r *http.Request) string { _ = "STUB: not implemented"; return "" }
