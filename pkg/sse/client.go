package sse

import (
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"
)

type ClientOption func(*clientOptions)

type clientOptions struct {
	headers               map[string]string
	zapLogger             *zap.Logger
	reconnectTimeInterval time.Duration
}

func defaultClientOptions() *clientOptions { _ = "STUB: not implemented"; return nil }

func (o *clientOptions) apply(opts ...ClientOption) { _ = "STUB: not implemented"; return }

// WithClientHeaders set HTTP headers
func WithClientHeaders(headers map[string]string) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// WithClientLogger set logger
func WithClientLogger(logger *zap.Logger) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// WithClientReconnectTimeInterval set reconnect time interval
func WithClientReconnectTimeInterval(d time.Duration) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// -------------------------------------------------------------------------------------------

// EventCallback event callback
type EventCallback func(event *Event)

// SSEClient sse client
type SSEClient struct {
	url         string
	client      *http.Client
	callbacks   map[string]EventCallback
	lastEventID string
	mu          sync.RWMutex
	connected   bool
	stopCh      chan struct{}

	headers   map[string]string
	zapLogger *zap.Logger
	// default is 2 seconds, backoff time interval will double after each retry
	reconnectTimeInterval time.Duration
}

// NewClient create a new sse client
func NewClient(url string, opts ...ClientOption) *SSEClient { _ = "STUB: not implemented"; return nil }

// no timeout

// OnEvent register event callback
func (c *SSEClient) OnEvent(eventType string, callback EventCallback) {
	_ = "STUB: not implemented"
	return
}

// Connect to server
func (c *SSEClient) Connect() error { _ = "STUB: not implemented"; return nil }

// Disconnect to server
func (c *SSEClient) Disconnect() { _ = "STUB: not implemented"; return }

// Wait returns a channel that will be closed when the client is disconnected.
func (c *SSEClient) Wait() <-chan struct{} {
	_ = "STUB: not implemented"

	// GetConnectStatus get connect status
	return nil
}

func (c *SSEClient) GetConnectStatus() bool { _ = "STUB: not implemented"; return false }

func (c *SSEClient) setConnectStatus(st bool) { _ = "STUB: not implemented"; return }

// support auto reconnect, retry strategy, exponential backoff algorithm, max retries is 5 times, max backoff is 30 seconds
func (c *SSEClient) connectServer() { _ = "STUB: not implemented"; return }

// wait for next retry

func (c *SSEClient) stream(s *retryStrategy) error { _ = "STUB: not implemented"; return nil }

// add request headers

// send request

// blank lines indicate the end of the event, process the event

// parse event fields

func (c *SSEClient) processEvent(eventType, eventID, data string) error {
	_ = "STUB: not implemented"
	return nil
}

// default event type

// -------------------------------------------------------------------------------------------

var noConnectionErrText = "No connection could be made because the target machine actively refused it"

func SetNoConnectionErrText(errText string) { _ = "STUB: not implemented"; return }

type retryStrategy struct {
	backoff    time.Duration // reconnect time
	maxBackoff time.Duration // max reconnect time

	retryCount int // retry count
	maxRetries int // max retry count, -1 means unlimited retries
}

// retry strategy, exponential backoff algorithm, returns true when the number of retries reaches the maximum
func (s *retryStrategy) retry() bool {
	_ = "STUB: not implemented"

	// calculate the next reconnect time (exponential backoff)
	return false
}

// check if maximum number of retries is reached

func (s *retryStrategy) reset(d time.Duration) { _ = "STUB: not implemented"; return }
