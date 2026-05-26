package circuitbreaker

import (
	"math/rand"
	"sync"
	"time"

	"github.com/go-dev-frame/sponge/pkg/shield/window"
)

// Option is sre breaker option function.
type Option func(*options)

const (
	// StateOpen when circuit breaker open, request not allowed, after sleep
	// some duration, allow one single request for testing the health, if ok
	// then state reset to closed, if not continue the step.
	StateOpen int32 = iota
	// StateClosed when circuit breaker closed, request allowed, the breaker
	// calc the succeed ratio, if request num greater request setting and
	// ratio lower than the setting ratio, then reset state to open.
	StateClosed
)

var (
	_ CircuitBreaker = &Breaker{}
)

// options is a breaker options.
type options struct {
	success float64
	request int64
	bucket  int
	window  time.Duration
}

// WithSuccess with the K = 1 / Success value of sre breaker, default success is 0.5
// Reducing the K will make adaptive throttling behave more aggressively,
// Increasing the K will make adaptive throttling behave less aggressively.
func WithSuccess(s float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRequest with the minimum number of requests allowed.
func WithRequest(r int64) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithWindow with the duration size of the statistical window.
func WithWindow(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBucket set the bucket number in a window duration.
func WithBucket(b int) Option { _ = "STUB: not implemented"; return *new(Option) }

// Breaker is a sre CircuitBreaker pattern.
type Breaker struct {
	stat window.RollingCounter
	r    *rand.Rand
	// rand.New(...) returns a non thread safe object
	randLock sync.Mutex

	// Reducing the k will make adaptive throttling behave more aggressively,
	// Increasing the k will make adaptive throttling behave less aggressively.
	k       float64
	request int64

	state int32
}

// NewBreaker return a sreBresker with options
func NewBreaker(opts ...Option) CircuitBreaker {
	_ = "STUB: not implemented"
	return *new(CircuitBreaker)
}

func (b *Breaker) summary() (success int64, total int64) { _ = "STUB: not implemented"; return 0, 0 }

//nolint

// Allow request if error returns nil.
func (b *Breaker) Allow() error {
	_ = "STUB: not implemented"
	// The number of requests accepted by the backend
	return nil
}

// The number of requests attempted by the application layer(at the client, on top of the adaptive throttling system)

// check overflow requests = K * accepts

// MarkSuccess mark request is success.
func (b *Breaker) MarkSuccess() {
	_ = "STUB: not implemented"

	// MarkFailed mark request is failed.
	return
}

func (b *Breaker) MarkFailed() {
	_ = "STUB: not implemented"
	// NOTE: when client reject request locally, keep adding counter let the drop ratio higher.
	return
}

func (b *Breaker) trueOnProba(proba float64) (truth bool) { _ = "STUB: not implemented"; return false }
