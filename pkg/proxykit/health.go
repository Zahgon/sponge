package proxykit

import (
	"time"
)

// HealthCheckConfig defined the configuration for health check.
type HealthCheckConfig struct {
	Interval time.Duration `json:"interval"`
	Timeout  time.Duration `json:"timeout"`
}

// StartHealthChecks initiate backend health check for the backend server pool.
func StartHealthChecks(backends []*Backend, config HealthCheckConfig) {
	_ = "STUB: not implemented"
	return
}

func runHealthCheck(backend *Backend, config HealthCheckConfig) { _ = "STUB: not implemented"; return }
