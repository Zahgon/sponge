package cache

import (
	"context"
	"time"

	"github.com/go-dev-frame/sponge/pkg/cache"

	"github.com/go-dev-frame/sponge/internal/database"
)

// delete the templates code start
type keyTypeExample = string
type valueTypeExample = string

// delete the templates code end

const (
	// cache prefix key, must end with a colon
	cacheNameExampleCachePrefixKey = "prefixKeyExample:"
	// CacheNameExampleExpireTime expire time
	CacheNameExampleExpireTime = 5 * time.Minute
)

var _ CacheNameExampleCache = (*cacheNameExampleCache)(nil)

// CacheNameExampleCache cache interface
type CacheNameExampleCache interface {
	Set(ctx context.Context, keyNameExample keyTypeExample, valueNameExample valueTypeExample, duration time.Duration) error
	Get(ctx context.Context, keyNameExample keyTypeExample) (valueTypeExample, error)
	Del(ctx context.Context, keyNameExample keyTypeExample) error
}

type cacheNameExampleCache struct {
	cache cache.Cache
}

// NewCacheNameExampleCache create a new cache
func NewCacheNameExampleCache(cacheType *database.CacheType) CacheNameExampleCache {
	_ = "STUB: not implemented"
	return *new(CacheNameExampleCache)
}

// cache key
func (c *cacheNameExampleCache) getCacheKey(keyNameExample keyTypeExample) string {
	_ = "STUB: not implemented"
	return ""
}

// Set cache
func (c *cacheNameExampleCache) Set(ctx context.Context, keyNameExample keyTypeExample, valueNameExample valueTypeExample, duration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Get cache
func (c *cacheNameExampleCache) Get(ctx context.Context, keyNameExample keyTypeExample) (valueTypeExample, error) {
	_ = "STUB: not implemented"
	return *new(valueTypeExample), nil
}

// Del delete cache
func (c *cacheNameExampleCache) Del(ctx context.Context, keyNameExample keyTypeExample) error {
	_ = "STUB: not implemented"
	return nil
}
