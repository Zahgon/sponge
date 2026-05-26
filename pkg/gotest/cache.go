package gotest

import (
	"context"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

// Cache redis cache
type Cache struct {
	Ctx           context.Context
	TestDataSlice []interface{}
	TestDataMap   map[string]interface{}
	RedisClient   *redis.Client
	redisServer   *miniredis.Miniredis
	ICache        interface{}
}

// NewCache instantiated redis cache
func NewCache(testDataMap map[string]interface{}) *Cache { _ = "STUB: not implemented"; return nil }

// Close redis server
func (c *Cache) Close() { _ = "STUB: not implemented"; return }

// GetIDs get test data ids
func (c *Cache) GetIDs() []uint64 { _ = "STUB: not implemented"; return nil }

// GetFields get test data fields
func (c *Cache) GetFields() []string { _ = "STUB: not implemented"; return nil }

// GetTestData get test data
func (c *Cache) GetTestData() map[string]interface{} { _ = "STUB: not implemented"; return nil }

// -------------------------------------------------------------------------------------------

// RCCache redis cluster cache
type RCCache struct {
	Ctx           context.Context
	TestDataSlice []interface{}
	TestDataMap   map[string]interface{}
	RedisClient   *redis.ClusterClient
	redisServer   *miniredis.Miniredis
	ICache        interface{}
}

// NewRCCache instantiated redis cluster cache
func NewRCCache(testDataMap map[string]interface{}) *RCCache { _ = "STUB: not implemented"; return nil }

// Close redis server
func (c *RCCache) Close() { _ = "STUB: not implemented"; return }

// GetIDs get test data ids
func (c *RCCache) GetIDs() []uint64 { _ = "STUB: not implemented"; return nil }

// GetTestData get test data
func (c *RCCache) GetTestData() map[string]interface{} { _ = "STUB: not implemented"; return nil }
