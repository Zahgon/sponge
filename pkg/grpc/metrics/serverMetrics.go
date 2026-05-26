// Package metrics is grpc's server-side and client-side metrics can continue to be captured using prometheus.
package metrics

import (
	"net/http"
	"sync"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"google.golang.org/grpc"
)

// https://github.com/grpc-ecosystem/go-grpc-prometheus/tree/master/examples/grpc-server-with-prometheus

var (
	// server side default router
	serverPattern = "/metrics"

	// create a Registry
	srvReg = prometheus.NewRegistry()

	// initialize server-side default metrics
	grpcServerMetrics = grpc_prometheus.NewServerMetrics()

	// go metrics
	goMetrics = collectors.NewGoCollector()

	// user-defined metrics https://prometheus.io/docs/concepts/metric_types/#histogram
	customizedCounterMetrics   = []*prometheus.CounterVec{}
	customizedSummaryMetrics   = []*prometheus.SummaryVec{}
	customizedGaugeMetrics     = []*prometheus.GaugeVec{}
	customizedHistogramMetrics = []*prometheus.HistogramVec{}
	grpcConnectionGauge        = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "grpc_server_active_connections",
			Help: "Current number of active gRPC client connections",
		},
	)

	srvOnce sync.Once
)

// Option set metrics
type Option func(*options)

type options struct{}

func defaultMetricsOptions() *options { _ = "STUB: not implemented"; return nil }

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// WithCounterMetrics add Counter type indicator
func WithCounterMetrics(metrics ...*prometheus.CounterVec) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSummaryMetrics add Summary type indicator
func WithSummaryMetrics(metrics ...*prometheus.SummaryVec) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGaugeMetrics add Gauge type indicator
func WithGaugeMetrics(metrics ...*prometheus.GaugeVec) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithHistogramMetrics adding Histogram type indicators
func WithHistogramMetrics(metrics ...*prometheus.HistogramVec) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func srvRegisterMetrics() {
	_ = "STUB: not implemented"

	// enable time record
	return
}

// register go metrics

// register metrics to capture, custom metrics also need to be registered

// register custom Counter metrics

// SetServerPattern set the server pattern
func SetServerPattern(pattern string) { _ = "STUB: not implemented"; return }

// Register for http routing and grpc methods
func Register(mux *http.ServeMux, grpcServer *grpc.Server) {
	_ = "STUB: not implemented"
	// register for http routing
	return
}

// register all gRPC methods to metrics

// ServerHTTPService initialize the prometheus exporter service on the server side and fetch data using http://ip:port/metrics
func ServerHTTPService(addr string, grpcServer *grpc.Server) *http.Server {
	_ = "STUB: not implemented"
	return nil
}

// run http server

// initialize gRPC methods Metrics

// ---------------------------------- server interceptor ----------------------------------

// UnaryServerMetrics metrics unary interceptor
func UnaryServerMetrics(opts ...Option) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// StreamServerMetrics metrics stream interceptor
func StreamServerMetrics(opts ...Option) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}
