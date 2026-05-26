package interceptor

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// ---------------------------------- client interceptor ----------------------------------

var (
	// default error code for triggering a retry
	defaultErrCodes = []codes.Code{codes.Internal}
)

// RetryOption set the retry retryOptions.
type RetryOption func(*retryOptions)

type retryOptions struct {
	times    uint
	interval time.Duration
	errCodes []codes.Code
}

func defaultRetryOptions() *retryOptions { _ = "STUB: not implemented"; return nil }

// default retry times
// default retry interval 100 ms
// default error code for triggering a retry

func (o *retryOptions) apply(opts ...RetryOption) { _ = "STUB: not implemented"; return }

// WithRetryTimes set number of retries, max 10
func WithRetryTimes(n uint) RetryOption { _ = "STUB: not implemented"; return *new(RetryOption) }

// WithRetryInterval set the retry interval from 1 ms to 10 seconds
func WithRetryInterval(t time.Duration) RetryOption {
	_ = "STUB: not implemented"
	return *new(RetryOption)
}

// WithRetryErrCodes set the trigger retry error code
func WithRetryErrCodes(errCodes ...codes.Code) RetryOption {
	_ = "STUB: not implemented"
	return *new(RetryOption)
}

// UnaryClientRetry client-side retry unary interceptor
func UnaryClientRetry(opts ...RetryOption) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

// set the number of retries
// set retry interval

// set retry error code

// StreamClientRetry client-side retry stream interceptor
func StreamClientRetry(opts ...RetryOption) grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}

// set the number of retries
// set retry interval

// set retry error code
