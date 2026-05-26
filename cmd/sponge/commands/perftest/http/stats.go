package http

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// collector of statistical results
type statsCollector struct {
	durations      []float64
	totalReqBytes  int64
	totalRespBytes int64
	successCount   uint64
	errorCount     uint64
	errSet         map[string]struct{}
	statusCodeSet  map[int]int64
}

func (c *statsCollector) collect(results <-chan Result, done chan<- struct{}) {
	_ = "STUB: not implemented"
	return
}

// nolint
func (c *statsCollector) collectAndPush(ctx context.Context, results <-chan Result, done chan<- struct{},
	spc *statsPrometheusCollector, p *PerfTestHTTP, start time.Time) {
	_ = "STUB: not implemented"
	return
}

func (c *statsCollector) toStatistics(totalTime time.Duration, totalRequests uint64, params *HTTPReqParams) *Statistics {
	_ = "STUB: not implemented"
	return nil
}

// convert float64 to string with specified precision, automatically process the last 0
func float64ToString(f float64, precision int) string { _ = "STUB: not implemented"; return "" }

func float64ToStringNoRound(f float64) string { _ = "STUB: not implemented"; return "" }

func (c *statsCollector) printReport(totalDuration time.Duration, totalRequests uint64, params *HTTPReqParams, id string) (*Statistics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func printStatusCodeSet(builder *Builder, statusCodeSet map[int]int64) {
	_ = "STUB: not implemented"
	return
}

func printErrorSet(builder *Builder, errSet map[string]struct{}) { _ = "STUB: not implemented"; return }

// --------------------------------------------------------------------------------

// Statistics performance test statistical data
type Statistics struct {
	ID string `json:"id"` // Performance Test ID

	URL    string `json:"url"`    // performed request URL
	Method string `json:"method"` // request method
	Body   string `json:"body"`   // request body (JSON)

	TotalRequests uint64   `json:"total_requests"` // total requests
	TotalDuration float64  `json:"total_duration"` //  total duration (seconds)
	SuccessCount  uint64   `json:"success_count"`  // successful requests (status code 2xx)
	ErrorCount    uint64   `json:"error_count"`    // failed requests (status code not 2xx)
	Errors        []string `json:"errors"`         // error details

	QPS        float64 `json:"qps"`         // requests per second (Throughput)
	AvgLatency float64 `json:"avg_latency"` // average latency (ms)
	P25Latency float64 `json:"p25_latency"` // 25th percentile latency (ms)
	P50Latency float64 `json:"p50_latency"` // 50th percentile latency (ms)
	P95Latency float64 `json:"p95_latency"` // 95th percentile latency (ms)
	P99Latency float64 `json:"p99_latency"` // 95th percentile latency (ms)
	MinLatency float64 `json:"min_latency"` // minimum latency (ms)
	MaxLatency float64 `json:"max_latency"` // maximum latency (ms)

	TotalSent     int64 `json:"total_sent"`     // total sent (bytes)
	TotalReceived int64 `json:"total_received"` // total received (bytes)

	StatusCodes map[int]int64 `json:"status_codes"` // status code distribution (count)

	CreatedAt time.Time `json:"created_at"` // created time

	Status  string `json:"status"`   // running, finished
	AgentID string `json:"agent_id"` // identify agent
}

// Save saves the statistics data to a JSON file.
func (s *Statistics) Save(filePath string) error { _ = "STUB: not implemented"; return nil }

func ensureFileExists(filePath string) error { _ = "STUB: not implemented"; return nil }

func convertToMilliseconds(f float64) float64 { _ = "STUB: not implemented"; return 0 }

// --------------------------------------------------------------------------

type statsPrometheusCollector struct {
	statsCollector *statsCollector

	// prometheus metrics
	totalRequestsGauge prometheus.Gauge
	successGauge       prometheus.Gauge
	errorGauge         prometheus.Gauge
	totalTimeGauge     prometheus.Gauge
	qpsGauge           prometheus.Gauge
	avgLatencyGauge    prometheus.Gauge
	p25LatencyGauge    prometheus.Gauge
	p50LatencyGauge    prometheus.Gauge
	p95LatencyGauge    prometheus.Gauge
	minLatencyGauge    prometheus.Gauge
	maxLatencyGauge    prometheus.Gauge
	totalSentGauge     prometheus.Gauge
	totalRecvGauge     prometheus.Gauge
	statusCodeGaugeVec *prometheus.GaugeVec
}

func newStatsPrometheusCollector() *statsPrometheusCollector { _ = "STUB: not implemented"; return nil }

func (spc *statsPrometheusCollector) copyStatsCollector(s *statsCollector) {
	_ = "STUB: not implemented"
	return
}

func percentile(sorted []float64, p float64) float64 { _ = "STUB: not implemented"; return 0 }

// PushToPrometheus pushes the statistics to a Prometheus.
func (spc *statsPrometheusCollector) PushToPrometheus(ctx context.Context, pushGatewayURL, jobName string, elapsed time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// set gauges

// PushToPrometheusAsync pushes the statistics to a Prometheus asynchronously
func (spc *statsPrometheusCollector) PushToPrometheusAsync(ctx context.Context, pushGatewayURL, jobName string, elapsed time.Duration) {
	_ = "STUB: not implemented"
	return
}

// PushToServer pushes the statistics data to a custom server
// body is the JSON data of Statistics struct
func (spc *statsPrometheusCollector) PushToServer(ctx context.Context, pushURL string, elapsed time.Duration, httpReqParams *HTTPReqParams, id string, agentID string, status AgentStatus) error {
	_ = "STUB: not implemented"
	return nil
}

// PushToServerAsync pushes the statistics data to a custom server asynchronously
func (spc *statsPrometheusCollector) PushToServerAsync(ctx context.Context, pushURL string, elapsed time.Duration, httpReqParams *HTTPReqParams, id string, agentID string, status AgentStatus) {
	_ = "STUB: not implemented"
	return
}

func postWithContext(ctx context.Context, url string, data *Statistics) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip certificate validation
