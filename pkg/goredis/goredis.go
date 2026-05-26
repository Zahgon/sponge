// Package goredis is a library wrapped on top of github.com/go-redis/redis.
package goredis

import (
	"github.com/redis/go-redis/v9"
)

// Client is a redis client
type Client = redis.Client

const (
	// ErrRedisNotFound not exist in redis
	ErrRedisNotFound = redis.Nil
	// DefaultRedisName default redis name
	DefaultRedisName = "default"
)

// Init connecting to redis
// dsn supported formats.
// (1) no password, no db: localhost:6379
// (2) with password and db: <user>:<pass>@localhost:6379/2
// (3) redis://default:123456@localhost:6379/0?max_retries=3
// for more parameters see the redis source code for the setupConnParams function
func Init(dsn string, opts ...Option) (*redis.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// replace single options if provided

//nolint

// InitSingle connecting to single redis instance
func InitSingle(addr string, password string, db int, opts ...Option) (*redis.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// replace single options if provided

//nolint

// InitSentinel connecting to redis for sentinel, all redis username and password are the same
func InitSentinel(masterName string, addrs []string, username string, password string, opts ...Option) (*redis.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// replace sentinel options if provided

//nolint

// InitCluster connecting to redis for cluster, all redis username and password are the same
func InitCluster(addrs []string, username string, password string, opts ...Option) (*redis.ClusterClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// replace cluster options if provided

//nolint

func getRedisOpt(dsn string, opts *options) (*redis.Options, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use db 0 by default

// Close redis client
func Close(rdb *redis.Client) error { _ = "STUB: not implemented"; return nil }

// CloseCluster redis cluster client
func CloseCluster(clusterRdb *redis.ClusterClient) error { _ = "STUB: not implemented"; return nil }
