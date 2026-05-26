package window

import (
	"time"
)

// Metric is a sample interface.
// Implementations of Metrics in metric package are Counter, Gauge,
// PointGauge, RollingCounter and RollingGauge.
type Metric interface {
	// Add adds the given value to the counter.
	Add(int64)
	// Value gets the current value.
	// If the metric's type is PointGauge, RollingCounter, RollingGauge,
	// it returns the sum value within the window.
	Value() int64
}

// Aggregation contains some common aggregation function.
// Each aggregation can compute summary statistics of window.
type Aggregation interface {
	// Min finds the min value within the window.
	Min() float64
	// Max finds the max value within the window.
	Max() float64
	// Avg computes average value within the window.
	Avg() float64
	// Sum computes sum value within the window.
	Sum() float64
}

// RollingCounter represents a ring window based on time duration.
// e.g. [[1], [3], [5]]
type RollingCounter interface {
	Metric
	Aggregation

	Timespan() int
	// Reduce applies the reduction function to all buckets within the window.
	Reduce(func(Iterator) float64) float64
}

// RollingCounterOpts contains the arguments for creating RollingCounter.
type RollingCounterOpts struct {
	Size           int
	BucketDuration time.Duration
}

type rollingCounter struct {
	policy *RollingPolicy
}

// NewRollingCounter creates a new RollingCounter bases on RollingCounterOpts.
func NewRollingCounter(opts RollingCounterOpts) RollingCounter {
	_ = "STUB: not implemented"
	return *new(RollingCounter)
}

func (r *rollingCounter) Add(val int64) { _ = "STUB: not implemented"; return }

func (r *rollingCounter) Reduce(f func(Iterator) float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (r *rollingCounter) Avg() float64 { _ = "STUB: not implemented"; return 0 }

func (r *rollingCounter) Min() float64 { _ = "STUB: not implemented"; return 0 }

func (r *rollingCounter) Max() float64 { _ = "STUB: not implemented"; return 0 }

func (r *rollingCounter) Sum() float64 { _ = "STUB: not implemented"; return 0 }

func (r *rollingCounter) Value() int64 { _ = "STUB: not implemented"; return 0 }

func (r *rollingCounter) Timespan() int { _ = "STUB: not implemented"; return 0 }
