// Package stat provides for counting system and process cpu and memory information, alarm notification support.
package stat

import (
	"context"
	"time"

	"go.uber.org/zap"
)

var (
	printInfoInterval = time.Minute // minimum 1 second
	zapLog, _         = zap.NewProduction()

	notifyCh = make(chan struct{})
)

// Option set the stat options field.
type Option func(*options)

type options struct {
	enableAlarm   bool
	zapFields     []zap.Field
	customHandler func(ctx context.Context, sd *StatData) error
}

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// WithPrintInterval set print interval
func WithPrintInterval(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLog set zapLog
func WithLog(l *zap.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPrintField set print field
func WithPrintField(fields ...zap.Field) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAlarm enable alarm and notify, except windows
func WithAlarm(opts ...AlarmOption) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCustomHandler set custom handler and interval, will replace default print stat data handler
func WithCustomHandler(handler func(ctx context.Context, sd *StatData) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Init initialize statistical information
func Init(opts ...Option) { _ = "STUB: not implemented"; return }

//nolint

func handleAlarm(sg *statGroup, data *StatData, o *options) { _ = "STUB: not implemented"; return }

// Windows system does not support alarm

func handleCustom(data *StatData, o *options) { _ = "STUB: not implemented"; return }

//nolint

// nolint
func sendSystemSignForLinux() { _ = "STUB: not implemented"; return }

func getStatData() *StatData { _ = "STUB: not implemented"; return nil }

// rounding

func printUsageInfo(statData *StatData, fields ...zap.Field) { _ = "STUB: not implemented"; return }

// System information
type System struct {
	CPUUsage float64 `json:"cpu_usage"` // system cpu usage, unit(%)
	MemUsage float64 `json:"mem_usage"` // system memory usage, unit(%)
	CPUCores int32   `json:"cpu_cores"` // cpu cores, multiple cpu accumulation
	MemTotal uint64  `json:"mem_total"` // system total physical memory, unit(M)
	MemFree  uint64  `json:"mem_free"`  // system free physical memory, unit(M)
}

// Process information
type Process struct {
	CPUUsage   float64 `json:"cpu_usage"`   // process cpu usage, unit(%)
	RSS        uint64  `json:"rss"`         // use of physical memory, unit(M)
	VMS        uint64  `json:"vms"`         // use of virtual memory, unit(M)
	Alloc      uint64  `json:"alloc"`       // allocated memory capacity, unit(M)
	TotalAlloc uint64  `json:"total_alloc"` // cumulative allocated memory capacity, unit(M)
	Sys        uint64  `json:"sys"`         // requesting memory capacity from the system, unit(M)
	NumGc      uint32  `json:"num_gc"`      // number of GC cycles
	Goroutines int     `json:"goroutines"`  // number of goroutines
}

// StatData statistical data
type StatData struct {
	Sys  System
	Proc Process
}
