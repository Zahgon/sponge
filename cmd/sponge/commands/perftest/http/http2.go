package http

import (
	"net/http"

	"github.com/spf13/cobra"
)

// PerfTestHTTP2CMD creates a new cobra.Command for HTTP/2 performance test.
func PerfTestHTTP2CMD() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Cluster mode parameters

//nolint:lll

// force push to collector host

// wait for all goroutines to exit

// Cluster mode parameters

func newHTTP2Client(worker int) *http.Client { _ = "STUB: not implemented"; return nil }

// default 1 second
// skip certificate validation
