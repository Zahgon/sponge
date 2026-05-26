// Package ratelimit is an adaptive rate limit library, support for use in gin middleware and grpc interceptors.
package ratelimit

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-dev-frame/sponge/pkg/shield/window"
)

var (
	gCPU    int64
	decay   = 0.95
	cpuOnce sync.Once
	_       Limiter = &BBR{}
)

type (
	cpuGetter func() int64

	// Option function for bbr limiter
	Option func(*options)
)

// cpu = cpuᵗ⁻¹ * decay + cpuᵗ * (1 - decay)
func cpuproc() { _ = "STUB: not implemented"; return }

// same to cpu sample rate

// EMA algorithm: https://blog.csdn.net/m0_38106113/article/details/81542863

func getMin(l, r uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// Stat contains the metrics snapshot of bbr.
type Stat struct {
	CPU         int64
	InFlight    int64
	MaxInFlight int64
	MinRt       int64
	MaxPass     int64
}

// counterCache is used to cache maxPASS and minRt result.
// Value of current bucket is not counted in real time.
// Cache time is equal to a bucket duration.
type counterCache struct {
	val  int64
	time time.Time
}

// options of bbr limiter.
type options struct {
	// WindowSize defines time duration per window
	Window time.Duration
	// BucketNum defines bucket number for each window
	Bucket int
	// CPUThreshold
	CPUThreshold int64
	// CPUQuota
	CPUQuota float64
}

// WithWindow with window size.
func WithWindow(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBucket with bucket ize.
func WithBucket(b int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCPUThreshold with cpu threshold;
func WithCPUThreshold(threshold int64) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCPUQuota with real cpu quota(if it can not collect from process correct);
func WithCPUQuota(quota float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// BBR implements bbr-like limiter.
// It is inspired by sentinel.
// https://github.com/alibaba/Sentinel/wiki/%E7%B3%BB%E7%BB%9F%E8%87%AA%E9%80%82%E5%BA%94%E9%99%90%E6%B5%81
type BBR struct {
	cpu             cpuGetter
	passStat        window.RollingCounter
	rtStat          window.RollingCounter
	inFlight        int64
	bucketPerSecond int64
	bucketDuration  time.Duration

	// prevDropTime defines previous start drop since initTime
	prevDropTime atomic.Value
	maxPASSCache atomic.Value
	minRtCache   atomic.Value

	opts options
}

// NewLimiter returns a bbr limiter
func NewLimiter(opts ...Option) *BBR { _ = "STUB: not implemented"; return nil }

// if cpuQuota is set, use new cpuGetter,Calculate the real CPU value based on the number of CPUs and Quota.

func (l *BBR) maxPASS() int64 { _ = "STUB: not implemented"; return 0 }

// timespan returns the passed bucket count
// since lastTime, if it is one bucket duration earlier than
// the last recorded time, it will return the BucketNum.
func (l *BBR) timespan(lastTime time.Time) int { _ = "STUB: not implemented"; return 0 }

func (l *BBR) minRT() int64 { _ = "STUB: not implemented"; return 0 }

func (l *BBR) maxInFlight() int64 { _ = "STUB: not implemented"; return 0 }

func (l *BBR) shouldDrop() bool { _ = "STUB: not implemented"; return false }

// current cpu payload below the threshold

// haven't start drop,
// accept current request

// just start drop one second ago, check current inflight count

// current cpu payload exceeds the threshold

// already started drop, return directly

// store start drop time

// Stat tasks a snapshot of the bbr limiter.
func (l *BBR) Stat() Stat { _ = "STUB: not implemented"; return *new(Stat) }

// Allow checks all inbound traffic.
// Once overload is detected, it raises limit.ErrLimitExceed error.
func (l *BBR) Allow() (DoneFunc, error) { _ = "STUB: not implemented"; return *new(DoneFunc), nil }
