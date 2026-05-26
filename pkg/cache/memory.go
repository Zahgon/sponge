package cache

import (
	"context"
	"sync"
	"time"

	"github.com/dgraph-io/ristretto"

	"github.com/go-dev-frame/sponge/pkg/encoding"
)

type options struct {
	numCounters int64
	maxCost     int64
	bufferItems int64
}

func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

// number of keys to track frequency of (10M).
// maximum cost of cache (1GB).
// number of keys per Get buffer.

// Option set the jwt options.
type Option func(*options)

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// WithNumCounters set number of keys.
func WithNumCounters(numCounters int64) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxCost set maximum cost of cache.
func WithMaxCost(maxCost int64) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBufferItems set number of keys per Get buffer.
func WithBufferItems(bufferItems int64) Option { _ = "STUB: not implemented"; return *new(Option) }

// InitMemory create a memory cache
func InitMemory(opts ...Option) *ristretto.Cache { _ = "STUB: not implemented"; return nil }

// see: https://dgraph.io/blog/post/introducing-ristretto-high-perf-go-cache/
//		https://www.start.io/blog/we-chose-ristretto-cache-for-go-heres-why/

// ----------------------------------------------------------------------------

// global memory cache client
var (
	memoryCli *ristretto.Cache
	once      sync.Once
)

// InitGlobalMemory init global memory cache
func InitGlobalMemory(opts ...Option) { _ = "STUB: not implemented"; return }

// GetGlobalMemoryCli get memory cache client
func GetGlobalMemoryCli() *ristretto.Cache { _ = "STUB: not implemented"; return nil }

// default options

// CloseGlobalMemory close memory cache
func CloseGlobalMemory() error { _ = "STUB: not implemented"; return nil }

// ----------------------------------------------------------------------------

type memoryCache struct {
	client            *ristretto.Cache
	KeyPrefix         string
	encoding          encoding.Encoding
	DefaultExpireTime time.Duration
	newObject         func() interface{}
}

// NewMemoryCache create a memory cache
func NewMemoryCache(keyPrefix string, encode encoding.Encoding, newObject func() interface{}) Cache {
	_ = "STUB: not implemented"
	return *new(Cache)
}

// Set data
func (m *memoryCache) Set(_ context.Context, key string, val interface{}, expiration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Get data
func (m *memoryCache) Get(_ context.Context, key string, val interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// not found, convert to redis nil error

// Del delete data
func (m *memoryCache) Del(_ context.Context, keys ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// MultiSet multiple set data
func (m *memoryCache) MultiSet(ctx context.Context, valueMap map[string]interface{}, expiration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// MultiGet multiple get data
func (m *memoryCache) MultiGet(ctx context.Context, keys []string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// SetCacheWithNotFound set not found
func (m *memoryCache) SetCacheWithNotFound(_ context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}
