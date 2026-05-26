package sasynq

import (
	"time"

	"github.com/hibiken/asynq"
)

// Client is a wrapper around asynq.Client providing more convenient APIs.
type Client struct {
	*asynq.Client
}

// NewClient creates a new producer client.
func NewClient(cfg RedisConfig) *Client { _ = "STUB: not implemented"; return nil }

// NewFromClient creates a new producer client from an existing asynq.Client.
func NewFromClient(c *asynq.Client) *Client { _ = "STUB: not implemented"; return nil }

// Enqueue enqueues the given task to a queue.
func (c *Client) Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EnqueueNow enqueues a task for immediate processing, parameter payload should be supported json.Marshal
func (c *Client) EnqueueNow(typeName string, payload any, opts ...asynq.Option) (*asynq.Task, *asynq.TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// EnqueueIn enqueues a task to be processed after a specified delay.
func (c *Client) EnqueueIn(delay time.Duration, typeName string, payload any, opts ...asynq.Option) (*asynq.Task, *asynq.TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// EnqueueAt enqueues a task to be processed at a specific time.
func (c *Client) EnqueueAt(t time.Time, typeName string, payload any, opts ...asynq.Option) (*asynq.Task, *asynq.TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// EnqueueUnique enqueues a task with unique in the queue for a specified duration.
func (c *Client) EnqueueUnique(keepTime time.Duration, typeName string, payload any, opts ...asynq.Option) (*asynq.Task, *asynq.TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
