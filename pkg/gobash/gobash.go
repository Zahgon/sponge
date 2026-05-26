// Package gobash provides the ability to execute commands, scripts, executables in the go environment with live log output.
package gobash

import (
	"context"
	"io"
	"os/exec"
)

// Exec suitable for executing a single non-blocking command, outputting standard and error logs,
// but the log output is not real time, no execution, command name must be in system path,
// Note: If the execution of a command blocks permanently, it can cause a concurrent leak.
func Exec(name string, args ...string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// cmdName is absolute path

// Result of the execution of the command
type Result struct {
	StdOut chan string
	Err    error // If nil after the command is executed, the command is executed successfully
	Pid    int   // Process ID of the command
}

// Run execute the command, no execution, command name must be in system path,
// you can actively end the command, the execution results are returned in real time in Result.StdOut
func Run(ctx context.Context, name string, args ...string) *Result {
	_ = "STUB: not implemented"
	return nil
}

// execution complete, channel closed
// cmdName is absolute path

func handleExec(ctx context.Context, cmd *exec.Cmd, result *Result) {
	_ = "STUB: not implemented"
	return
}

// command name and pid

// reads each line in real time

// determine if it has been read

// capture error logs

func getCmdReader(cmd *exec.Cmd) (stdout io.ReadCloser, stderr io.ReadCloser, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), *new(io.ReadCloser), nil
}

func getResult(cmd *exec.Cmd) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ParsePid extracts the process ID from the command output string.
func ParsePid(s string) int { _ = "STUB: not implemented"; return 0 }
