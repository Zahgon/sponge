// Package window is a library that calculates windows cpu and memory usage.
package window

// Bucket contains multiple float64 points.
type Bucket struct {
	Points []float64
	Count  int64
	next   *Bucket
}

// Append appends the given value to the bucket.
func (b *Bucket) Append(val float64) { _ = "STUB: not implemented"; return }

// Add adds the given value to the point.
func (b *Bucket) Add(offset int, val float64) { _ = "STUB: not implemented"; return }

// Reset empties the bucket.
func (b *Bucket) Reset() { _ = "STUB: not implemented"; return }

// Next returns the next bucket.
func (b *Bucket) Next() *Bucket {
	_ = "STUB: not implemented"

	// Window contains multiple buckets.
	return nil
}

type Window struct {
	buckets []Bucket
	size    int
}

// Options contains the arguments for creating Window.
type Options struct {
	Size int
}

// NewWindow creates a new Window based on WindowOpts.
func NewWindow(opts Options) *Window { _ = "STUB: not implemented"; return nil }

// ResetWindow empties all buckets within the window.
func (w *Window) ResetWindow() { _ = "STUB: not implemented"; return }

// ResetBucket empties the bucket based on the given offset.
func (w *Window) ResetBucket(offset int) { _ = "STUB: not implemented"; return }

// ResetBuckets empties the buckets based on the given offsets.
func (w *Window) ResetBuckets(offset int, count int) { _ = "STUB: not implemented"; return }

// Append appends the given value to the bucket where index equals the given offset.
func (w *Window) Append(offset int, val float64) { _ = "STUB: not implemented"; return }

// Add adds the given value to the latest point within bucket where index equals the given offset.
func (w *Window) Add(offset int, val float64) { _ = "STUB: not implemented"; return }

// Bucket returns the bucket where index equals the given offset.
func (w *Window) Bucket(offset int) Bucket { _ = "STUB: not implemented"; return *new(Bucket) }

// Size returns the size of the window.
func (w *Window) Size() int {
	_ = "STUB: not implemented"

	// Iterator returns the count number buckets iterator from offset.
	return 0
}

func (w *Window) Iterator(offset int, count int) Iterator {
	_ = "STUB: not implemented"
	return *new(Iterator)
}
