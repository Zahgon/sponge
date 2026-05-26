package metrics

import (
	"net/http"
	"sync"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

var (
	// client side default router
	clientPattern = "/rpc_client/metrics"

	// create a Registry
	cliReg = prometheus.NewRegistry()

	// initialize the client's default metrics
	grpcClientMetrics = grpc_prometheus.NewClientMetrics()

	cliOnce sync.Once
)

func cliRegisterMetrics() {
	_ = "STUB: not implemented"

	// register metrics, including custom metrics
	return
}

// SetClientPattern set the client pattern
func SetClientPattern(pattern string) { _ = "STUB: not implemented"; return }

// ClientRegister for http routing and grpc methods
func ClientRegister(mux *http.ServeMux) {
	_ = "STUB: not implemented"
	// register for http routing
	return
}

// ClientHTTPService initialize the client's prometheus exporter service and use http://ip:port/metrics to fetch data
func ClientHTTPService(addr string) *http.Server { _ = "STUB: not implemented"; return nil }

// run http server

// ---------------------------------- client interceptor ----------------------------------

// UnaryClientMetrics metrics unary interceptor
func UnaryClientMetrics() grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

// StreamClientMetrics metrics stream interceptor
func StreamClientMetrics() grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}
