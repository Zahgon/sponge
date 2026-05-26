package rabbitmq

import (
	"context"
)

// Subscriber session
type Subscriber struct {
	*Consumer
}

// NewSubscriber create a subscriber, channelName is exchange name, identifier is queue name
func NewSubscriber(channelName string, identifier string, connection *Connection, opts ...ConsumerOption) (*Subscriber, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Subscribe and handle message
func (s *Subscriber) Subscribe(ctx context.Context, handler Handler) {
	_ = "STUB: not implemented"
	return

	// Close subscriber
}

func (s *Subscriber) Close() { _ = "STUB: not implemented"; return }
