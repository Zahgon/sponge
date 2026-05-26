package window

import (
	"sync"
	"time"
)

// RollingPolicy is a policy for ring window based on time duration.
// RollingPolicy moves bucket offset with time duration.
// e.g. If the last point is appended one bucket duration ago,
// RollingPolicy will increment current offset.
type RollingPolicy struct {
	mu     sync.RWMutex
	size   int
	window *Window
	offset int

	bucketDuration time.Duration
	lastAppendTime time.Time
}

// RollingPolicyOpts contains the arguments for creating RollingPolicy.
type RollingPolicyOpts struct {
	BucketDuration time.Duration
}

// NewRollingPolicy creates a new RollingPolicy based on the given window and RollingPolicyOpts.
func NewRollingPolicy(window *Window, opts RollingPolicyOpts) *RollingPolicy {
	_ = "STUB: not implemented"
	return nil
}

// timespan returns passed bucket number since lastAppendTime,
// if it is one bucket duration earlier than the last recorded
// time, it will return the size.
func (r *RollingPolicy) timespan() int { _ = "STUB: not implemented"; return 0 }

// maybe time backwards

// apply applies function f with value val on
// current offset bucket, expired bucket will be reset
func (r *RollingPolicy) apply(f func(offset int, val float64), val float64) {
	_ = "STUB: not implemented"
	return
}

// calculate current offset

// reset the expired buckets

// Append appends the given points to the window.
func (r *RollingPolicy) Append(val float64) { _ = "STUB: not implemented"; return }

// Add adds the given value to the latest point within bucket.
func (r *RollingPolicy) Add(val float64) { _ = "STUB: not implemented"; return }

// Reduce applies the reduction function to all buckets within the window.
func (r *RollingPolicy) Reduce(f func(Iterator) float64) (val float64) {
	_ = "STUB: not implemented"
	return 0
}
