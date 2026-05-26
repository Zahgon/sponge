// Package cpu is a library that calculates cpu and memory usage.
package cpu

import (
	"time"
)

type psutilCPU struct {
	interval time.Duration
}

func newPsutilCPU(interval time.Duration) (*psutilCPU, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ps *psutilCPU) Usage() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// convert to 10/1000 of a percent

func (ps *psutilCPU) Info() Info { _ = "STUB: not implemented"; return *new(Info) }
