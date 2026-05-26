package rabbitmq

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

// ProducerOption producer option.
type ProducerOption func(*producerOptions)

type producerOptions struct {
	exchangeDeclare *exchangeDeclareOptions
	queueDeclare    *queueDeclareOptions
	queueBind       *queueBindOptions
	deadLetter      *deadLetterOptions

	isPersistent bool // is it persistent

	// If true, the message will be returned to the sender if the queue cannot be
	// found according to its own exchange type and routeKey rules.
	mandatory bool

	// only for publish-subscribe mode
	isPublisherConfirm bool
}

func (o *producerOptions) apply(opts ...ProducerOption) { _ = "STUB: not implemented"; return }

// default producer settings
func defaultProducerOptions() *producerOptions { _ = "STUB: not implemented"; return nil }

// WithProducerExchangeDeclareOptions set exchange declare option.
func WithProducerExchangeDeclareOptions(opts ...ExchangeDeclareOption) ProducerOption {
	_ = "STUB: not implemented"
	return *new(ProducerOption)
}

// WithProducerQueueDeclareOptions set queue declare option.
func WithProducerQueueDeclareOptions(opts ...QueueDeclareOption) ProducerOption {
	_ = "STUB: not implemented"
	return *new(ProducerOption)
}

// WithProducerQueueBindOptions set queue bind option.
func WithProducerQueueBindOptions(opts ...QueueBindOption) ProducerOption {
	_ = "STUB: not implemented"
	return *new(ProducerOption)
}

// WithDeadLetterOptions set dead letter options.
func WithDeadLetterOptions(opts ...DeadLetterOption) ProducerOption {
	_ = "STUB: not implemented"
	return *new(ProducerOption)
}

// WithProducerPersistent set producer persistent option.
func WithProducerPersistent(enable bool) ProducerOption {
	_ = "STUB: not implemented"
	return *new(ProducerOption)
}

// WithProducerMandatory set producer mandatory option.
func WithProducerMandatory(enable bool) ProducerOption {
	_ = "STUB: not implemented"
	return *new(ProducerOption)
}

// WithPublisherConfirm enables publisher confirm.
func WithPublisherConfirm() ProducerOption { _ = "STUB: not implemented"; return *new(ProducerOption) }

// -------------------------------------------------------------------------------------------

// Producer session
type Producer struct {
	Exchange  *Exchange        // exchange
	QueueName string           // queue name
	conn      *amqp.Connection // rabbitmq connection
	ch        *amqp.Channel    // rabbitmq channel

	// persistent or not
	isPersistent bool
	deliveryMode uint8 // amqp.Persistent or amqp.Transient

	// If true, the message will be returned to the sender if the queue cannot be
	// found according to its own exchange type and routeKey rules.
	mandatory bool

	zapLog *zap.Logger

	exchangeArgs  amqp.Table
	queueArgs     amqp.Table
	queueBindArgs amqp.Table

	// only for publish-subscribe mode
	isPublisherConfirm bool
}

// NewProducer create a producer
func NewProducer(exchange *Exchange, queueName string, connection *Connection, opts ...ProducerOption) (*Producer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// crate a new channel

// declare the exchange type

// declare a queue and create it automatically if it doesn't exist, or skip creation if it does.

// binding queue and exchange

// create dead letter exchange and queue if enabled

// PublishDirect send direct type message
func (p *Producer) PublishDirect(ctx context.Context, body []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishFanout send fanout type message
func (p *Producer) PublishFanout(ctx context.Context, body []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishTopic send topic type message
func (p *Producer) PublishTopic(ctx context.Context, topicKey string, body []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishHeaders send headers type message
func (p *Producer) PublishHeaders(ctx context.Context, headersKeys map[string]interface{}, body []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishDelayedMessage send delayed type message
func (p *Producer) PublishDelayedMessage(ctx context.Context, delayTime time.Duration, body []byte, opts ...DelayedMessagePublishOption) error {
	_ = "STUB: not implemented"
	return nil
}

// delay time: milliseconds

// Close the consumer
func (p *Producer) Close() { _ = "STUB: not implemented"; return }

// ExchangeArgs returns the exchange declare args.
func (p *Producer) ExchangeArgs() amqp.Table {
	_ = "STUB: not implemented"
	return *

	// QueueArgs returns the queue declare args.
	new(amqp.Table)
}

func (p *Producer) QueueArgs() amqp.Table {
	_ = "STUB: not implemented"

	// QueueBindArgs returns the queue bind args.
	return *new(amqp.Table)
}

func (p *Producer) QueueBindArgs() amqp.Table { _ = "STUB: not implemented"; return *new(amqp.Table) }

func logFields(queueName string, exchange *Exchange) []zap.Field {
	_ = "STUB: not implemented"
	return nil
}

// -------------------------------------------------------------------------------------------

func createDeadLetter(ch *amqp.Channel, o *deadLetterOptions) error {
	_ = "STUB: not implemented"
	// declare the exchange type
	return nil
}

// declare a queue and create it automatically if it doesn't exist, or skip creation if it does.

// binding queue and exchange
