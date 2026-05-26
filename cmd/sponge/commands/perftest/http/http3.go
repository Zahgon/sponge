package http

import (
	"net/http"

	"github.com/spf13/cobra"
)

// PerfTestHTTP3CMD creates a new cobra.Command for HTTP/3 performance test.
func PerfTestHTTP3CMD() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Cluster mode parameters

//nolint:lll

// force push to collector host

// wait for all goroutines to exit

// Cluster mode parameters

func newHTTP3Client(worker int) *http.Client { _ = "STUB: not implemented"; return nil }

// quic.Config provides fine control over the underlying QUIC connections

// Skip certificate validation
