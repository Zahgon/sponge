package rabbitmq

import (
	"context"
)

// Publisher session
type Publisher struct {
	*Producer
}

// NewPublisher create a publisher, channelName is exchange name
func NewPublisher(channelName string, connection *Connection, opts ...ProducerOption) (*Publisher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// crate a new channel

// enable publisher confirm

// declare the exchange type

func (p *Publisher) Publish(ctx context.Context, body []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// wait for publisher confirm

// Close publisher
func (p *Publisher) Close() { _ = "STUB: not implemented"; return }
