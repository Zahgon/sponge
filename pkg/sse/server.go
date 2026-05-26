package sse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ServeOption is an option for Serve
type ServeOption func(*serveOptions)

type serveOptions struct {
	extraHeaders map[string]string
}

func defaultServeOptions() *serveOptions { _ = "STUB: not implemented"; return nil }

func (o *serveOptions) apply(opts ...ServeOption) { _ = "STUB: not implemented"; return }

// WithServeExtraHeaders sets extra headers to be sent with the response.
func WithServeExtraHeaders(headers map[string]string) ServeOption {
	_ = "STUB: not implemented"
	return *new(ServeOption)
}

// -------------------------------------------------------------------------------------------

// Serve serves a client connection
func (h *Hub) Serve(c *gin.Context, uid string, opts ...ServeOption) {
	_ = "STUB: not implemented"
	return
}

// client connected, response once means connection is ok

// ServeHandler gin handler for sse server
func (h *Hub) ServeHandler(opts ...ServeOption) func(c *gin.Context) {
	_ = "STUB: not implemented"
	return nil
}

// set uid in auth middleware

// PushRequest push request
type PushRequest struct {
	UIDs   []string `json:"uids"`
	Events []*Event `json:"events"`
}

// PushEventHandler gin handler for push event request
func (h *Hub) PushEventHandler() func(c *gin.Context) { _ = "STUB: not implemented"; return nil }

// UserClient information
type UserClient struct {
	UID     string // user id
	Send    chan *Event
	writer  http.ResponseWriter
	flusher http.Flusher

	isSendClosedEvent bool
}

func (c *UserClient) sendEvent(e *Event) error { _ = "STUB: not implemented"; return nil }

func (c *UserClient) write(data []byte) error { _ = "STUB: not implemented"; return nil }

func responseCode400(c *gin.Context, msg string) { _ = "STUB: not implemented"; return }

func responseCode200(c *gin.Context) { _ = "STUB: not implemented"; return }
