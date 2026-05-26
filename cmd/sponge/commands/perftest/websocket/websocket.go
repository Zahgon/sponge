package websocket

import (
	"time"

	"github.com/spf13/cobra"
)

// PerfTestWebsocketCMD creates a cobra command for websocket performance test
func PerfTestWebsocketCMD() *cobra.Command { _ = "STUB: not implemented"; return nil }

// default payload data random(10)

type perfTestParams struct {
	targetURL    string
	worker       int
	duration     time.Duration
	sendInterval time.Duration
	rampUp       time.Duration

	payloadData []byte
	isJSON      bool

	out string
}

func (p *perfTestParams) run() error { _ = "STUB: not implemented"; return nil }

// Test duration timer

// OS interrupt signal

// Calculate ramp-up delay

// Start all client workers with ramp-up

// Wait for all workers to finish
