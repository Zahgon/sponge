package sse

import "sync"

// SafeMap goroutine security Map structure encapsulating sync.Map
type SafeMap struct {
	m sync.Map // userID -> []*Client
}

// NewSafeMap creates a new SafeMap
func NewSafeMap() *SafeMap { _ = "STUB: not implemented"; return nil }

// Set store key-value pairs
func (sm *SafeMap) Set(uid string, client *UserClient) { _ = "STUB: not implemented"; return }

// Get value by key
func (sm *SafeMap) Get(uid string) (*UserClient, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Delete key-value pair
func (sm *SafeMap) Delete(uid string) {
	_ = "STUB: not implemented"

	// Has checked if key exists
	return
}

func (sm *SafeMap) Has(uid string) bool { _ = "STUB: not implemented"; return false }

// Range all key-value pairs
func (sm *SafeMap) Range(f func(key, value interface{}) bool) {
	_ = "STUB: not implemented"

	// Keys get all keys
	return
}

func (sm *SafeMap) Keys() []string { _ = "STUB: not implemented"; return nil }

// Values get all values
func (sm *SafeMap) Values() []*UserClient { _ = "STUB: not implemented"; return nil }

// Len Gets the number of elements in Map
// Note: Due to the nature of sync.Map, this operation is O(n) complex
func (sm *SafeMap) Len() int { _ = "STUB: not implemented"; return 0 }

// Clear all key-value pairs
func (sm *SafeMap) Clear() { _ = "STUB: not implemented"; return }
