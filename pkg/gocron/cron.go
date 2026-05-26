// Package gocron is scheduled task library.
package gocron

import (
	"sync"

	"github.com/robfig/cron/v3"
)

var (
	c *cron.Cron
	// task name and id mapping, used to add, delete, modify and query tasks
	nameID = sync.Map{}
	// id and task name mapping, used in log printing
	idName = sync.Map{}
)

// Task scheduled task
type Task struct {
	// seconds (0-59) minutes (0- 59) hours (0-23) days (1-31) months (1-12) weeks (0-6)
	// "*/5 * * * * *"  means every five seconds.
	// "0 15,45 9-12 * * * "  indicates execution at the 15th and 45th minutes from 9 a.m. to 12 a.m. each day
	TimeSpec string

	Name      string // task name
	Fn        func() // task function
	IsRunOnce bool   // if the task is only run once
}

// Init initialize and start timed tasks
func Init(opts ...Option) error { _ = "STUB: not implemented"; return nil }

// second-level granularity, default is minute-level granularity

// Run the tasks
func Run(tasks ...*Task) error { _ = "STUB: not implemented"; return nil }

func checkRunOnce(task *Task) error { _ = "STUB: not implemented"; return nil }

// IsRunningTask determine if the task is running
func IsRunningTask(name string) bool { _ = "STUB: not implemented"; return false }

// GetRunningTasks gets a list of running task names
func GetRunningTasks() []string { _ = "STUB: not implemented"; return nil }

// DeleteTask stop and delete the specified task
func DeleteTask(name string) { _ = "STUB: not implemented"; return }

// Stop all scheduled tasks
func Stop() { _ = "STUB: not implemented"; return }

// EverySecond every second size (1~59)
func EverySecond(size int) string { _ = "STUB: not implemented"; return "" }

// EveryMinute every minute size (1~59)
func EveryMinute(size int) string { _ = "STUB: not implemented"; return "" }

// EveryHour every hour size (1~23)
func EveryHour(size int) string { _ = "STUB: not implemented"; return "" }

// Everyday size (1~31)
func Everyday(size int) string { _ = "STUB: not implemented"; return "" }
