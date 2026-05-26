package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/go-dev-frame/sponge/pkg/encoding"
)

// CacheNotFound no hit cache
var CacheNotFound = redis.Nil

// redisCache redis cache object
type redisCache struct {
	client            *redis.Client
	KeyPrefix         string
	encoding          encoding.Encoding
	DefaultExpireTime time.Duration
	newObject         func() interface{}
}

// NewRedisCache new a cache, client parameter can be passed in for unit testing
func NewRedisCache(client *redis.Client, keyPrefix string, encode encoding.Encoding, newObject func() interface{}) Cache {
	_ = "STUB: not implemented"
	return *new(Cache)
}

// Set one value
func (c *redisCache) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

//if expiration == 0 {
//	expiration = DefaultExpireTime
//}

// Get one value
func (c *redisCache) Get(ctx context.Context, key string, val interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: don't handle the case where redis value is nil
// but leave it to the upstream for processing

// prevent Unmarshal from reporting an error if data is empty

// MultiSet set multiple values
func (c *redisCache) MultiSet(ctx context.Context, valueMap map[string]interface{}, expiration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

//if expiration == 0 {
//	expiration = DefaultExpireTime
//}

// the key-value is paired and has twice the capacity of a map

// MultiGet get multiple values
func (c *redisCache) MultiGet(ctx context.Context, keys []string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Injection into map via reflection

// Del delete multiple values
func (c *redisCache) Del(ctx context.Context, keys ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// SetCacheWithNotFound set value for notfound
func (c *redisCache) SetCacheWithNotFound(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// BuildCacheKey construct a cache key with a prefix
func BuildCacheKey(keyPrefix string, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// -------------------------------------------------------------------------------------------

// redisClusterCache redis cluster cache object
type redisClusterCache struct {
	client            *redis.ClusterClient
	KeyPrefix         string
	encoding          encoding.Encoding
	DefaultExpireTime time.Duration
	newObject         func() interface{}
}

// NewRedisClusterCache new a cache
func NewRedisClusterCache(client *redis.ClusterClient, keyPrefix string, encode encoding.Encoding, newObject func() interface{}) Cache {
	_ = "STUB: not implemented"
	return *new(Cache)
}

// Set one value
func (c *redisClusterCache) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

//if expiration == 0 {
//	expiration = DefaultExpireTime
//}

// Get one value
func (c *redisClusterCache) Get(ctx context.Context, key string, val interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: don't handle the case where redis value is nil
// but leave it to the upstream for processing

// prevent Unmarshal from reporting an error if data is empty

// MultiSet set multiple values
func (c *redisClusterCache) MultiSet(ctx context.Context, valueMap map[string]interface{}, expiration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// the key-value is paired and has twice the capacity of a map

// MultiGet get multiple values
func (c *redisClusterCache) MultiGet(ctx context.Context, keys []string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Injection into map via reflection

// Del delete multiple values
func (c *redisClusterCache) Del(ctx context.Context, keys ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// SetCacheWithNotFound set value for notfound
func (c *redisClusterCache) SetCacheWithNotFound(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}
