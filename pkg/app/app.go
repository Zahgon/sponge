// Package app is starting and stopping services gracefully, using golang.org/x/sync/errgroup to ensure that multiple services are started properly at the same time.
package app

import (
	"context"
)

// IServer server interface
type IServer interface {
	Start() error
	Stop() error
	String() string
}

// Close app close
type Close func() error

// App servers
type App struct {
	servers []IServer
	closes  []Close
}

// New create an app
func New(servers []IServer, closes []Close) *App { _ = "STUB: not implemented"; return nil }

// Run servers
func (a *App) Run() {
	_ = "STUB: not implemented"
	// ctx will be notified whenever an error occurs in one of the goroutines
	return
}

// start all servers

// watch and stop app

// watch the os signal and the ctx signal from the errgroup, and stop the service if either signal is triggered
func (a *App) watch(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// service error

// system notification signal

// start or stop sampling profile

// stopping services and releasing resources
func (a *App) stop() error { _ = "STUB: not implemented"; return nil }
