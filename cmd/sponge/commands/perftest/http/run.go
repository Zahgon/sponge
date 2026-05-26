package http

import (
	"context"
	"net/http"
	"time"
)

// PerfTestHTTP performance test parameters for HTTP
type PerfTestHTTP struct {
	ID string // performance test ID

	Client *http.Client
	Params *HTTPReqParams

	Worker        int
	TotalRequests uint64
	Duration      time.Duration

	PushURL           string
	PrometheusJobName string
	pushInterval      time.Duration

	agentID            string
	clusterEnable      bool
	pushToCollectorURL string
}

func (p *PerfTestHTTP) checkParams() error { _ = "STUB: not implemented"; return nil }

// Run the performance test with fixed number of requests or fixed duration.
func (p *PerfTestHTTP) Run(ctx context.Context, duration time.Duration, out string) error {
	_ = "STUB: not implemented"
	return nil
}

// RunWithFixedRequestsNum implements performance with a fixed number of requests.
func (p *PerfTestHTTP) RunWithFixedRequestsNum(globalCtx context.Context) (*Statistics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint

// The collector counts the results of each request, closes the statsDone
// channel, and notifies the main thread of the end

// Distribute tasks and listen for context cancellation events

// RunWithFixedDuration implements performance with a fixed duration.
func (p *PerfTestHTTP) RunWithFixedDuration(globalCtx context.Context) (*Statistics, error) {
	_ = "STUB: not implemented"
	// Create a context that will be canceled when the duration is over
	return nil, nil
}

//nolint

// Since the total number of requests is unknown, we initialize with a reasonable capacity to reduce reallocation's.

// Start workers

// Keep sending requests until the context is canceled

// Exit goroutine when context is canceled

// Wait for the timeout or a signal

// Wait for all workers to finish their current request

// Close the result channel to signal the collector that no more results will be sent

// Wait for the collector to process all the results in the channel

// The total number of requests is the count of collected results

func pushStatistics(spc *statsPrometheusCollector, p *PerfTestHTTP, totalTime time.Duration, status AgentStatus) {
	_ = "STUB: not implemented"
	return
}

//nolint

// -------------------------------------------------------------------------------------------

// Result record the results of the request
type Result struct {
	Duration   time.Duration
	ReqSize    int64
	RespSize   int64
	StatusCode int
	Err        error
}

type HTTPReqParams struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    []byte

	version string
}

func buildRequest(params *HTTPReqParams) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func requestOnce(client *http.Client, params *HTTPReqParams, ch chan<- Result) {
	_ = "STUB: not implemented"
	return
}

// Check for request-level errors (e.g. timeout, DNS resolution failure)

// Check if the response status code is not 2xx

func captureSignal() context.Context { _ = "STUB: not implemented"; return *new(context.Context) }

// handle manual interruption (Ctrl+C)
