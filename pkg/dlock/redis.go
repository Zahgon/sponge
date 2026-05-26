package dlock

import (
	"context"

	"github.com/go-redsync/redsync/v4"
	"github.com/redis/go-redis/v9"
)

// RedisLock implements Locker using Redis.
type RedisLock struct {
	mutex *redsync.Mutex
}

// NewRedisLock creates a new RedisLock.
func NewRedisLock(client *redis.Client, key string, options ...redsync.Option) (Locker, error) {
	_ = "STUB: not implemented"
	return *new(Locker), nil
}

// NewRedisClusterLock creates a new RedisClusterLock.
func NewRedisClusterLock(clusterClient *redis.ClusterClient, key string, options ...redsync.Option) (Locker, error) {
	_ = "STUB: not implemented"
	return *new(Locker), nil
}

func newLocker(delegate redis.UniversalClient, key string, options ...redsync.Option) Locker {
	_ = "STUB: not implemented"
	return *new(Locker)
}

// TryLock tries to acquire the lock without blocking.
func (l *RedisLock) TryLock(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Lock blocks until the lock is acquired or the context is canceled.
func (l *RedisLock) Lock(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Unlock releases the lock, if unlocking the key is successful, the key will be automatically deleted
func (l *RedisLock) Unlock(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Close no-op for RedisLock.
func (l *RedisLock) Close() error { _ = "STUB: not implemented"; return nil }
