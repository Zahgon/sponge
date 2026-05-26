package cpu

import (
	"errors"
)

type cgroupCPU struct {
	frequency uint64
	quota     float64
	cores     uint64

	preSystem uint64
	preTotal  uint64
}

func newCgroupCPU() (cpu *cgroupCPU, err error) { _ = "STUB: not implemented"; return nil, nil }

func (cpu *cgroupCPU) Usage() (u uint64, err error) { _ = "STUB: not implemented"; return 0, nil }

func (cpu *cgroupCPU) Info() Info { _ = "STUB: not implemented"; return *new(Info) }

const nanoSecondsPerSecond = 1e9

// ErrNoCFSLimit is no quota limit
var ErrNoCFSLimit = errors.New("no quota limit")

var clockTicksPerSecond = uint64(getClockTicks())

// systemCPUUsage returns the host system's cpu usage in
// nanoseconds. An error is returned if the format of the underlying
// file does not match.
//
// Uses /proc/stat defined by POSIX. Looks for the cpu
// statistics line and then sums up the first seven fields
// provided. See man 5 proc for details on specific field
// information.
func systemCPUUsage() (usage uint64, err error) { _ = "STUB: not implemented"; return 0, nil }

func totalCPUUsage() (usage uint64, err error) { _ = "STUB: not implemented"; return 0, nil }

func perCPUUsage() (usage []uint64, err error) { _ = "STUB: not implemented"; return nil, nil }

func cpuSets() (sets []uint64, err error) { _ = "STUB: not implemented"; return nil, nil }

func cpuQuota() (quota int64, err error) { _ = "STUB: not implemented"; return 0, nil }

func cpuPeriod() (peroid uint64, err error) { _ = "STUB: not implemented"; return 0, nil }

func cpuFreq() uint64 { _ = "STUB: not implemented"; return 0 }

// treat this as the fallback value, thus we ignore error

func cpuMaxFreq() uint64 { _ = "STUB: not implemented"; return 0 }

// override the max freq from /proc/cpuinfo

// GetClockTicks get the OS's ticks per second
func getClockTicks() int {
	_ = "STUB: not implemented"
	// TODO figure out a better alternative for platforms where we're missing cgo
	//
	// TODO Windows. This could be implemented using Win32 QueryPerformanceFrequency().
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms644905(v=vs.85).aspx
	//
	// An example of its usage can be found here.
	// https://msdn.microsoft.com/en-us/library/windows/desktop/dn553408(v=vs.85).aspx
	return 0
}
