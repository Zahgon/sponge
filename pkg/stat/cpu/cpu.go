// Package cpu is a library that counts system and process cpu usage.
package cpu

// System cpu information
type System struct {
	UsagePercent float64   `json:"usage_percent"` // cpu usage, unit(%), current logical CPU usage, total usage is cores*UsagePercent
	CPUInfo      []CPUInfo `json:"cpu_info"`
}

// CPUInfo cpu information
type CPUInfo struct {
	ModelName string  `json:"modelName"`
	Cores     int32   `json:"cores"`
	Frequency float64 `json:"frequency"` // cpu frequency, unit(Mhz)
}

// Process information
type Process struct {
	UsagePercent float64 `json:"usage_percent"` // cpu usage, unit(%), current process occupies current logical CPU, total usage is cores*UsagePercent

	RSS uint64 `json:"rss"` // use of physical memory, unit(M)
	VMS uint64 `json:"vms"` // use of virtual memory, unit(M)
}

// GetSystemCPU get system cpu info
func GetSystemCPU() *System { _ = "STUB: not implemented"; return nil }

// total cpu Percent

// GetProcess get current process info
func GetProcess() *Process { _ = "STUB: not implemented"; return nil }

func floatRound(f float64, n int) float64 { _ = "STUB: not implemented"; return 0 }
