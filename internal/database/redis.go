package database

import (
	"sync"

	"github.com/go-dev-frame/sponge/pkg/goredis"
)

var (
	// ErrCacheNotFound No hit cache
	ErrCacheNotFound = goredis.ErrRedisNotFound
)

var (
	redisCli     *goredis.Client
	redisCliOnce sync.Once

	cacheType     *CacheType
	cacheTypeOnce sync.Once
)

// CacheType cache type
type CacheType struct {
	CType string          // cache type  memory or redis
	Rdb   *goredis.Client // if CType=redis, Rdb cannot be empty
}

// InitCache initial cache
func InitCache(cType string) { _ = "STUB: not implemented"; return }

// GetCacheType get cacheType
func GetCacheType() *CacheType { _ = "STUB: not implemented"; return nil }

// InitRedis connect redis
func InitRedis() { _ = "STUB: not implemented"; return }

// GetRedisCli get redis client
func GetRedisCli() *goredis.Client { _ = "STUB: not implemented"; return nil }

// CloseRedis close redis
func CloseRedis() error { _ = "STUB: not implemented"; return nil }
