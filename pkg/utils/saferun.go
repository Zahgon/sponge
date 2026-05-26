package utils

import (
	"context"
	"time"
)

// SafeRun safe run
func SafeRun(ctx context.Context, fn func(ctx context.Context)) { _ = "STUB: not implemented"; return }

// SafeRunWithTimeout safe run with limit timeouts
func SafeRunWithTimeout(d time.Duration, fn func(cancel context.CancelFunc)) {
	_ = "STUB: not implemented"
	return
}
