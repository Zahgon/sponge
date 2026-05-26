package websocket

import (
	"context"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Client represents a single WebSocket client worker.
type Client struct {
	id    int
	conn  *websocket.Conn
	stats *statsCollector
	url   string

	isJSON bool
	//sendPayloadTemplate map[string]any // isJSON=true, for JSON data
	sendPayloadBytes []byte // if isJSON = true, sendPayloadBytes is the JSON data, otherwise, it is the binary data
	sendTicker       *time.Ticker
}

// NewClient creates a new WebSocket client worker.
func NewClient(id int, url string, stats *statsCollector, sendInterval time.Duration, payloadData []byte, isJSON bool) *Client {
	_ = "STUB: not implemented"
	return nil
}

// Dial establishes a WebSocket connection to the server.
func (c *Client) Dial(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Run starts the client worker.
func (c *Client) Run(ctx context.Context, wg *sync.WaitGroup) { _ = "STUB: not implemented"; return }

// Wait for loops to finish with a timeout

func (c *Client) sendData(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// writeLoop handles sending messages with sequence numbers.
func (c *Client) writeLoop(ctx context.Context, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

// readLoop handles receiving messages and checking for latency and order.
func (c *Client) readLoop(ctx context.Context, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return

	//_ = c.conn.SetReadDeadline(time.Now().Add(time.Second))
}
