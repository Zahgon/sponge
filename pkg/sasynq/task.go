package sasynq

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
)

// TaskHandler is a generic interface for handling a task with a specific payload type.
type TaskHandler[T any] interface {
	Handle(ctx context.Context, payload T) error
}

// TaskHandleFunc is a function adapter for TaskHandler.
type TaskHandleFunc[T any] func(ctx context.Context, payload T) error

// Handle calls the wrapped function.
func (f TaskHandleFunc[T]) Handle(ctx context.Context, payload T) error {
	_ = "STUB: not implemented"
	return nil

	// NewTask creates a new asynq.Task with a typed payload.
	// It automatically marshals the payload into JSON.
}

func NewTask[P any](typeName string, payload P, opts ...asynq.Option) (*asynq.Task, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HandleFunc creates a TaskHandler from a function.
func HandleFunc[T any](f func(ctx context.Context, payloadType T) error) TaskHandler[T] {
	_ = "STUB: not implemented"
	return nil
}

// RegisterTaskHandler registers a generic, type-safe task handler with the server's mux.
// It automatically unmarshals the JSON payload into the specified type.
func RegisterTaskHandler[T any](mux *asynq.ServeMux, typeName string, handler TaskHandler[T]) {
	_ = "STUB: not implemented"
	return
}

// --- Task Option Helpers ---

// WithMaxRetry specifies the max number of times the task will be retried.
func WithMaxRetry(maxRetry int) asynq.Option { _ = "STUB: not implemented"; return *new(asynq.Option) }

// WithTimeout specifies the timeout duration for the task.
func WithTimeout(timeout time.Duration) asynq.Option {
	_ = "STUB: not implemented"
	return *new(asynq.Option)
}

// WithDeadline specifies the deadline for the task.
func WithDeadline(t time.Time) asynq.Option {
	_ = "STUB: not implemented"
	return *

	// WithTaskID specifies the ID for the task, if another task with the same ID already exists, it will be rejeceted.
	new(asynq.Option)
}

func WithTaskID(id string) asynq.Option {
	_ = "STUB: not implemented"
	return *

	// WithQueue specifies which queue the task should be sent to.
	new(asynq.Option)
}

func WithQueue(name string) asynq.Option { _ = "STUB: not implemented"; return *new(asynq.Option) }
