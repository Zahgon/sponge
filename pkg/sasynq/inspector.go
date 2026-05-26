package sasynq

import (
	"github.com/hibiken/asynq"
)

// Inspector provides access to the Redis backend used by asynq.
type Inspector struct {
	*asynq.Inspector
}

// NewInspector creates a new Inspector instance.
// Note: A new Redis connection will be created, in actual use, only once
func NewInspector(cfg RedisConfig) *Inspector { _ = "STUB: not implemented"; return nil }

// CancelTask cancels the processing of a task.
func (i *Inspector) CancelTask(queue string, taskID string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetTaskInfo returns information about a task.
func (i *Inspector) GetTaskInfo(queue string, taskID string) (*asynq.TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes the inspector.
func (i *Inspector) Close() error { _ = "STUB: not implemented"; return nil }
