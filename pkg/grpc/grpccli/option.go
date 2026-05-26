package grpccli

import (
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/go-dev-frame/sponge/pkg/servicerd/registry"
)

var (
	secureOneWay = "one-way"
	secureTwoWay = "two-way"
)

// Option grpc dial options
type Option func(*options)

// options grpc dial options
type options struct {
	requestTimeout time.Duration // request timeout, valid only for unary

	// secure setting
	secureType string // secure type "","one-way","two-way"
	serverName string // server name
	caFile     string // ca file
	certFile   string // cert file
	keyFile    string // key file

	// token setting
	enableToken bool // whether to turn on token
	appID       string
	appKey      string

	// interceptor setting
	enableLog            bool // whether to turn on the log
	log                  *zap.Logger
	enableRequestID      bool               // whether to turn on the request id
	enableTrace          bool               // whether to turn on tracing
	enableMetrics        bool               // whether to turn on metrics
	enableRetry          bool               // whether to turn on retry
	enableLoadBalance    bool               // whether to turn on load balance
	enableCircuitBreaker bool               // whether to turn on circuit breaker
	discovery            registry.Discovery // if not nil means use service discovery

	discoveryInsecure bool

	// custom setting
	dialOptions        []grpc.DialOption              // custom options
	unaryInterceptors  []grpc.UnaryClientInterceptor  // custom unary interceptor
	streamInterceptors []grpc.StreamClientInterceptor // custom stream interceptor
}

func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// WithTimeout set dial timeout
func WithTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableRequestID enable request id
func WithEnableRequestID() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableLog enable log
func WithEnableLog(log *zap.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableTrace enable trace
func WithEnableTrace() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableMetrics enable metrics
func WithEnableMetrics() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableLoadBalance enable load balance
func WithEnableLoadBalance() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableRetry enable registry
func WithEnableRetry() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableCircuitBreaker enable circuit breaker
func WithEnableCircuitBreaker() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDiscoveryInsecure setting discovery insecure
func WithDiscoveryInsecure(b bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func (o *options) isSecure() bool { _ = "STUB: not implemented"; return false }

// WithSecure support setting one-way or two-way secure
func WithSecure(t string, serverName string, caFile string, certFile string, keyFile string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithOneWaySecure set one-way secure
func WithOneWaySecure(serverName string, certFile string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTwoWaySecure set two-way secure
func WithTwoWaySecure(serverName string, caFile string, certFile string, keyFile string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithToken set token
func WithToken(enable bool, appID string, appKey string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDialOptions set dial options
func WithDialOptions(dialOptions ...grpc.DialOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithUnaryInterceptors set dial unaryInterceptors
func WithUnaryInterceptors(unaryInterceptors ...grpc.UnaryClientInterceptor) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStreamInterceptors set dial streamInterceptors
func WithStreamInterceptors(streamInterceptors ...grpc.StreamClientInterceptor) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDiscovery set dial discovery
func WithDiscovery(discovery registry.Discovery) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
