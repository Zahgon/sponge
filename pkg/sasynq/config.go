package sasynq

import (
	"crypto/tls"

	"github.com/hibiken/asynq"
)

// RedisMode defines the Redis connection mode.
type RedisMode string

const (
	// RedisModeSingle uses a single Redis instance.
	RedisModeSingle RedisMode = "single"
	// RedisModeSentinel uses Redis Sentinel for high availability.
	RedisModeSentinel RedisMode = "sentinel"
	// RedisModeCluster uses a Redis Cluster for horizontal scaling.
	RedisModeCluster RedisMode = "cluster"
)

// RedisConfig holds all configurations for connecting to Redis.
type RedisConfig struct {
	Mode RedisMode `yaml:"mode"`

	// For Single Mode
	Addr string `yaml:"addr"`

	// For Sentinel Mode
	SentinelAddrs []string `yaml:"sentinelAddrs"`
	MasterName    string   `yaml:"masterName"`

	// For Cluster Mode
	ClusterAddrs []string `yaml:"clusterAddrs"`

	// Common options
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`

	TLSConfig *tls.Config `yaml:"tlsConfig"`
}

// GetAsynqRedisConnOpt converts RedisConfig to asynq's RedisConnOpt interface.
// This is the core of the high-availability switching logic.
func (c RedisConfig) GetAsynqRedisConnOpt() asynq.RedisConnOpt {
	_ = "STUB: not implemented"
	return *new(asynq.RedisConnOpt)
}

// ServerConfig holds configurations for the asynq server.
type ServerConfig struct {
	*asynq.Config
}

// DefaultServerConfig returns a default server configuration.
func DefaultServerConfig(opts ...LoggerOption) ServerConfig {
	_ = "STUB: not implemented"
	return *new(ServerConfig)
}
