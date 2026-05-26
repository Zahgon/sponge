package window

// Iterator iterates the buckets within the window.
type Iterator struct {
	count         int
	iteratedCount int
	cur           *Bucket
}

// Next returns true util all of the buckets has been iterated.
func (i *Iterator) Next() bool { _ = "STUB: not implemented"; return false }

// Bucket gets current bucket.
func (i *Iterator) Bucket() Bucket { _ = "STUB: not implemented"; return *new(Bucket) }
