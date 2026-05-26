// Package cache is memory and redis cache libraries.
package cache

import (
	"context"
	"errors"
	"time"
)

var (
	// DefaultExpireTime default expiry time
	DefaultExpireTime = time.Hour * 24
	// DefaultNotFoundExpireTime expiry time when result is empty 1 minute,
	// often used for cache time when data is empty (cache pass-through)
	DefaultNotFoundExpireTime = time.Minute * 10

	// NotFoundPlaceholder placeholder
	NotFoundPlaceholder      = "*"
	NotFoundPlaceholderBytes = []byte(NotFoundPlaceholder)
	ErrPlaceholder           = errors.New("cache: placeholder")

	// DefaultClient generate a cache client, where keyPrefix is generally the business prefix
	DefaultClient Cache
)

// Cache driver interface
type Cache interface {
	Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, val interface{}) error
	MultiSet(ctx context.Context, valMap map[string]interface{}, expiration time.Duration) error
	MultiGet(ctx context.Context, keys []string, valueMap interface{}) error
	Del(ctx context.Context, keys ...string) error
	SetCacheWithNotFound(ctx context.Context, key string) error
}

// Set data
func Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Get data
func Get(ctx context.Context, key string, val interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// MultiSet multiple set data
func MultiSet(ctx context.Context, valMap map[string]interface{}, expiration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// MultiGet multiple get data
func MultiGet(ctx context.Context, keys []string, valueMap interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Del multiple delete data
func Del(ctx context.Context, keys ...string) error { _ = "STUB: not implemented"; return nil }

// SetCacheWithNotFound .
func SetCacheWithNotFound(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}
