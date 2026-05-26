// Package process provides functions to manage processes.
package process

// Kill terminates a process by its PID.
func Kill(pid int) error { _ = "STUB: not implemented"; return nil }

// 1. First, attempt a graceful shutdown.

// Wait for confirmation of exit: check every 0.25s for up to 10 times (5s total).

// If the graceful signal failed, check if the process is already gone.

// 2. If graceful shutdown failed or timed out, force kill the process.

// If force kill failed, do one last check to see if it's running.
