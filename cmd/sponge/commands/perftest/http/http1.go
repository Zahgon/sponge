package http

import (
	"net/http"

	"github.com/spf13/cobra"
)

// PerfTestHTTPCMD creates a new cobra.Command for HTTP/1.1 performance test.
func PerfTestHTTPCMD() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Cluster mode parameters

//nolint:lll

// force push to collector host

// wait for all goroutines to exit

// Cluster mode parameters

func newHTTPClient(worker int) *http.Client { _ = "STUB: not implemented"; return nil }

// default 1 second
// skip certificate validation
