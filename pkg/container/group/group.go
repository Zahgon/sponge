// Package group provides a sample lazy load container.
// The group only creating a new object not until the object is needed by user.
// And it will cache all the objects to reduce the creation of object.
package group

import "sync"

// Group is a lazy load container.
type Group struct {
	new  func() interface{}
	vals map[string]interface{}
	sync.RWMutex
}

// NewGroup news a group container.
func NewGroup(fn func() interface{}) *Group { _ = "STUB: not implemented"; return nil }

// Get gets the object by the given key.
func (g *Group) Get(key string) interface{} { _ = "STUB: not implemented"; return nil }

// slow path for group don`t have specified key value

// Reset resets the new function and deletes all existing objects.
func (g *Group) Reset(fn func() interface{}) { _ = "STUB: not implemented"; return }

// Clear deletes all objects.
func (g *Group) Clear() { _ = "STUB: not implemented"; return }
