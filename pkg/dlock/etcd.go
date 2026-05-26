package dlock

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var defaultTTL = 15 // seconds

type EtcdLock struct {
	session *concurrency.Session
	mutex   *concurrency.Mutex
}

// NewEtcd creates a new etcd locker with the given key and ttl.
func NewEtcd(client *clientv3.Client, key string, ttl int) (Locker, error) {
	_ = "STUB: not implemented"
	return *new(Locker), nil
}

//nolint

// Lock blocks until the lock is acquired or the context is canceled.
func (l *EtcdLock) Lock(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Unlock releases the lock.
func (l *EtcdLock) Unlock(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// TryLock tries to acquire the lock without blocking.
func (l *EtcdLock) TryLock(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Close releases the lock and the etcd session.
func (l *EtcdLock) Close() error { _ = "STUB: not implemented"; return nil }
