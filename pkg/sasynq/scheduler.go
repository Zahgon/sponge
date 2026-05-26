package sasynq

import (
	"github.com/hibiken/asynq"
)

// SchedulerOption set options.
type SchedulerOption func(*schedulerOptions)

type schedulerOptions struct {
	schedulerOptions *asynq.SchedulerOpts
	logger           asynq.Logger
	loggerLevel      asynq.LogLevel
}

func (o *schedulerOptions) apply(opts ...SchedulerOption) { _ = "STUB: not implemented"; return }

func defaultSchedulerOptions() *schedulerOptions { _ = "STUB: not implemented"; return nil }

// WithSchedulerOptions sets the options for the scheduler.
func WithSchedulerOptions(opts *asynq.SchedulerOpts) SchedulerOption {
	_ = "STUB: not implemented"
	return *new(SchedulerOption)
}

// WithSchedulerLogLevel sets the log level for the scheduler.
func WithSchedulerLogLevel(level asynq.LogLevel) SchedulerOption {
	_ = "STUB: not implemented"
	return *new(SchedulerOption)
}

// WithSchedulerLogger sets the logger for the scheduler.
func WithSchedulerLogger(opts ...LoggerOption) SchedulerOption {
	_ = "STUB: not implemented"
	return *new(SchedulerOption)
}

// --------------------------------------------------------------------

// Scheduler is a wrapper around asynq.Scheduler.
type Scheduler struct {
	*asynq.Scheduler
}

// NewScheduler creates a new periodic task scheduler.
func NewScheduler(cfg RedisConfig, opts ...SchedulerOption) *Scheduler {
	_ = "STUB: not implemented"
	return nil
}

// Register adds a new periodic task.
func (s *Scheduler) Register(cronSpec string, task *asynq.Task, opts ...asynq.Option) (entryID string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RegisterTask adds a new periodic task with a given type name.
func (s *Scheduler) RegisterTask(cronSpec string, typeName string, payload any, opts ...asynq.Option) (entryID string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Unregister removes a periodic task, cancel task execution.
func (s *Scheduler) Unregister(entryID string) error { _ = "STUB: not implemented"; return nil }

// Run runs the asynq Scheduler in a separate goroutine
func (s *Scheduler) Run() { _ = "STUB: not implemented"; return }

// Shutdown the Scheduler.
func (s *Scheduler) Shutdown() { _ = "STUB: not implemented"; return }
