package sse

import "sync"

type AsyncTaskPool struct {
	tasks chan func()
	wg    sync.WaitGroup
}

// NewAsyncTaskPool creates a task pool with a fixed capacity
func NewAsyncTaskPool(maxWorkers int) *AsyncTaskPool { _ = "STUB: not implemented"; return nil }

// default capacity is 10000

func (p *AsyncTaskPool) worker() { _ = "STUB: not implemented"; return }

func (p *AsyncTaskPool) Submit(task func()) { _ = "STUB: not implemented"; return }

func (p *AsyncTaskPool) Wait() { _ = "STUB: not implemented"; return }

func (p *AsyncTaskPool) Stop() { _ = "STUB: not implemented"; return }
