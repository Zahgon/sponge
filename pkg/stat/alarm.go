package stat

import (
	"time"
)

var (
	cpuThreshold            = 0.8 // 80% CPU usage
	memoryThreshold         = 0.8 // 80% memory usage
	triggerInterval float64 = 900 // unit(s)
)

// AlarmOption set the alarm options field.
type AlarmOption func(*alarmOptions)

type alarmOptions struct{}

func (o *alarmOptions) apply(opts ...AlarmOption) { _ = "STUB: not implemented"; return }

// WithCPUThreshold set cpu threshold, range 0 to 1
func WithCPUThreshold(threshold float64) AlarmOption {
	_ = "STUB: not implemented"
	return *new(AlarmOption)
}

// WithMemoryThreshold set memory threshold, range 0 to 1
func WithMemoryThreshold(threshold float64) AlarmOption {
	_ = "STUB: not implemented"
	return *new(AlarmOption)
}

type statGroup struct {
	data    [3]*StatData
	alarmAt time.Time
}

func newStatGroup() *statGroup { _ = "STUB: not implemented"; return nil }

func (g *statGroup) check(sd *StatData) bool { _ = "STUB: not implemented"; return false }

func (g *statGroup) checkCPU(threshold float64) bool { _ = "STUB: not implemented"; return false }

// average cpu usage exceeds cpuCores*threshold

func (g *statGroup) checkMemory(threshold float64) bool { _ = "STUB: not implemented"; return false }

// processes occupying more than threshold of system memory
