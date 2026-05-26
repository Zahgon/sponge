package cache

import (
	"context"
	"time"

	"github.com/go-dev-frame/sponge/pkg/cache"

	"github.com/go-dev-frame/sponge/internal/database"
	"github.com/go-dev-frame/sponge/internal/model"
)

const (
	// cache prefix key, must end with a colon
	userExampleCachePrefixKey = "userExample:"
	// UserExampleExpireTime expire time
	UserExampleExpireTime = 5 * time.Minute
)

var _ UserExampleCache = (*userExampleCache)(nil)

// UserExampleCache cache interface
type UserExampleCache interface {
	Set(ctx context.Context, id uint64, data *model.UserExample, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.UserExample, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserExample, error)
	MultiSet(ctx context.Context, data []*model.UserExample, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// userExampleCache define a cache struct
type userExampleCache struct {
	cache cache.Cache
}

// NewUserExampleCache new a cache
func NewUserExampleCache(cacheType *database.CacheType) UserExampleCache {
	_ = "STUB: not implemented"
	return *new(UserExampleCache)
}

// no cache

// GetUserExampleCacheKey cache key
func (c *userExampleCache) GetUserExampleCacheKey(id uint64) string {
	_ = "STUB: not implemented"
	return ""
}

// Set write to cache
func (c *userExampleCache) Set(ctx context.Context, id uint64, data *model.UserExample, duration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Get cache value
func (c *userExampleCache) Get(ctx context.Context, id uint64) (*model.UserExample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MultiSet multiple set cache
func (c *userExampleCache) MultiSet(ctx context.Context, data []*model.UserExample, duration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *userExampleCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserExample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Del delete cache
func (c *userExampleCache) Del(ctx context.Context, id uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *userExampleCache) SetPlaceholder(ctx context.Context, id uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// IsPlaceholderErr check if cache is placeholder error
func (c *userExampleCache) IsPlaceholderErr(err error) bool {
	_ = "STUB: not implemented"
	return false
}
