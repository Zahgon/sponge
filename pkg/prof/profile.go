// Package prof is wrap the official `net/http/pprof` route and add the profile io wait time route.
package prof

import (
	"syscall"
)

var (
	durationSecond  uint32 = 60
	isSamplingTrace        = false

	serverName = getServerName()
	pid        = syscall.Getpid()
	timeFormat = "20060102T150405"

	status      uint32
	statusStart uint32 = 1 // status=1
	statusStop  uint32     // status=0
)

// WaitSign wait system notification signals
//func WaitSign() {
//	p := NewProfile()
//
//	signals := make(chan os.Signal, 1)
//	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGTRAP)
//
//	for {
//		v := <-signals
//		switch v {
//		case syscall.SIGTRAP:
//			p.StartOrStop()
//
//		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP:
//			gracefully exit
//			os.Exit(0)
//		}
//	}
//}

type Profile struct {
	files    []string
	closeFns []func()

	//ctx    context.Context
	stopCh chan struct{}
}

// NewProfile create a new profile
func NewProfile() *Profile { _ = "STUB: not implemented"; return nil }

// StartOrStop start and stop sampling profile, the first call to start sampling data, the default maximum is 60 seconds,
// in less than 60s, if the second execution will actively stop sampling profile
func (p *Profile) StartOrStop() { _ = "STUB: not implemented"; return }

func (p *Profile) startProfile() { _ = "STUB: not implemented"; return }

func (p *Profile) stopProfile() { _ = "STUB: not implemented"; return }

// reset profile
//nolint

func (p *Profile) checkTimeout() { _ = "STUB: not implemented"; return }

//nolint

func (p *Profile) cpu() error { _ = "STUB: not implemented"; return nil }

func (p *Profile) mem() error { _ = "STUB: not implemented"; return nil }

func (p *Profile) goroutine() error { _ = "STUB: not implemented"; return nil }

func (p *Profile) block() error { _ = "STUB: not implemented"; return nil }

func (p *Profile) mutex() error { _ = "STUB: not implemented"; return nil }

func (p *Profile) threadCreate() error { _ = "STUB: not implemented"; return nil }

func (p *Profile) tracing() error { _ = "STUB: not implemented"; return nil }

// SetDurationSecond set sampling profile duration
func SetDurationSecond(d uint32) { _ = "STUB: not implemented"; return }

// EnableTrace enable sampling trace profile
func EnableTrace() { _ = "STUB: not implemented"; return }

func isStart() bool { _ = "STUB: not implemented"; return false }

func isStop() bool { _ = "STUB: not implemented"; return false }

func getFilePath(profileName string) string { _ = "STUB: not implemented"; return "" }

func getServerName() string { _ = "STUB: not implemented"; return "" }

func joinPath(elem ...string) string { _ = "STUB: not implemented"; return "" }
