// Package mem is a library that counts system and process memory usage.
package mem

// System memory information
type System struct {
	Total        uint64  `json:"total"`         // total physical memory capacity, unit(M)
	Free         uint64  `json:"free"`          // free physical memory capacity, unit(M)
	UsagePercent float64 `json:"usage_percent"` // memory usage, unit(%)
}

// Process memory information
type Process struct {
	Alloc      uint64 `json:"alloc"`       // allocated memory capacity, unit(M)
	TotalAlloc uint64 `json:"total_alloc"` // cumulative allocated memory capacity, unit(M)
	Sys        uint64 `json:"sys"`         // requesting memory capacity from the system, unit(M)
	NumGc      uint32 `json:"num_gc"`      // number of GC cycles
}

// GetSystemMemory get system memory
func GetSystemMemory() *System { _ = "STUB: not implemented"; return nil }

// GetProcessMemory get process memory
func GetProcessMemory() *Process { _ = "STUB: not implemented"; return nil }
