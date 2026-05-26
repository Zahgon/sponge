package consulcli

import (
	"time"

	"github.com/hashicorp/consul/api"
)

// Option set the consul client options.
type Option func(*options)

type options struct {
	scheme     string
	waitTime   time.Duration
	datacenter string
	token      string

	// if you set this parameter, all fields above are invalid
	config *api.Config
}

func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// WithWaitTime set wait time
func WithWaitTime(waitTime time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithScheme set scheme
func WithScheme(scheme string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDatacenter set datacenter
func WithDatacenter(datacenter string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithToken set token
func WithToken(token string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithConfig set consul config
func WithConfig(c *api.Config) Option { _ = "STUB: not implemented"; return *new(Option) }
