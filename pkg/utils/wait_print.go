package utils

import (
	"context"
	"time"
)

// WaitPrinter is a waiting printer.
type WaitPrinter struct {
	ctx            context.Context
	cancel         context.CancelFunc
	printFrequency time.Duration
}

// NewWaitPrinter create a new WaitPrinter instance.
func NewWaitPrinter(interval time.Duration) *WaitPrinter { _ = "STUB: not implemented"; return nil }

// LoopPrint start the waiting loop and print the running tip message.
func (p *WaitPrinter) LoopPrint(runningTip string) { _ = "STUB: not implemented"; return }

// StopPrint stop the waiting loop and print the tip message.
func (p *WaitPrinter) StopPrint(tip string) { _ = "STUB: not implemented"; return }

func (p *WaitPrinter) clearCurrentLine() { _ = "STUB: not implemented"; return }
