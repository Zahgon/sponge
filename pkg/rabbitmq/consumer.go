package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

// ConsumerOption consumer option.
type ConsumerOption func(*consumerOptions)

type consumerOptions struct {
	exchangeDeclare *exchangeDeclareOptions
	queueDeclare    *queueDeclareOptions
	queueBind       *queueBindOptions
	qos             *qosOptions
	consume         *consumeOptions

	isPersistent bool // persistent or not
	isAutoAck    bool // auto-answer or not, if false, manual ACK required
}

func (o *consumerOptions) apply(opts ...ConsumerOption) { _ = "STUB: not implemented"; return }

// default consumer settings
func defaultConsumerOptions() *consumerOptions { _ = "STUB: not implemented"; return nil }

// WithConsumerExchangeDeclareOptions set exchange declare option.
func WithConsumerExchangeDeclareOptions(opts ...ExchangeDeclareOption) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}

// WithConsumerQueueDeclareOptions set queue declare option.
func WithConsumerQueueDeclareOptions(opts ...QueueDeclareOption) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}

// WithConsumerQueueBindOptions set queue bind option.
func WithConsumerQueueBindOptions(opts ...QueueBindOption) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}

// WithConsumerQosOptions set consume qos option.
func WithConsumerQosOptions(opts ...QosOption) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}

// WithConsumerConsumeOptions set consumer consume option.
func WithConsumerConsumeOptions(opts ...ConsumeOption) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}

// WithConsumerAutoAck set consumer auto ack option, if false, manual ACK required.
func WithConsumerAutoAck(enable bool) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}

// WithConsumerPersistent set consumer persistent option.
func WithConsumerPersistent(enable bool) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}

// -------------------------------------------------------------------------------------------

// ConsumeOption consume option.
type ConsumeOption func(*consumeOptions)

type consumeOptions struct {
	consumer  string     // used to distinguish between multiple consumers
	exclusive bool       // only available to the program that created it
	noLocal   bool       // if set to true, a message sent by a producer in the same Connection cannot be passed to a consumer in this Connection.
	noWait    bool       // block processing
	args      amqp.Table // additional properties
}

func (o *consumeOptions) apply(opts ...ConsumeOption) { _ = "STUB: not implemented"; return }

// default consume settings
func defaultConsumeOptions() *consumeOptions { _ = "STUB: not implemented"; return nil }

// WithConsumeConsumer set consume consumer option.
func WithConsumeConsumer(consumer string) ConsumeOption {
	_ = "STUB: not implemented"
	return *new(ConsumeOption)
}

// WithConsumeExclusive set consume exclusive option.
func WithConsumeExclusive(enable bool) ConsumeOption {
	_ = "STUB: not implemented"
	return *new(ConsumeOption)
}

// WithConsumeNoLocal set consume noLocal option.
func WithConsumeNoLocal(enable bool) ConsumeOption {
	_ = "STUB: not implemented"
	return *new(ConsumeOption)
}

// WithConsumeNoWait set consume no wait option.
func WithConsumeNoWait(enable bool) ConsumeOption {
	_ = "STUB: not implemented"
	return *new(ConsumeOption)
}

// WithConsumeArgs set consume args option.
func WithConsumeArgs(args map[string]interface{}) ConsumeOption {
	_ = "STUB: not implemented"
	return *new(ConsumeOption)
}

// -------------------------------------------------------------------------------------------

// QosOption qos option.
type QosOption func(*qosOptions)

type qosOptions struct {
	enable        bool
	prefetchCount int
	prefetchSize  int
	global        bool
}

func (o *qosOptions) apply(opts ...QosOption) { _ = "STUB: not implemented"; return }

// default qos settings
func defaultQosOptions() *qosOptions { _ = "STUB: not implemented"; return nil }

// WithQosEnable set qos enable option.
func WithQosEnable() QosOption { _ = "STUB: not implemented"; return *new(QosOption) }

// WithQosPrefetchCount set qos prefetch count option.
func WithQosPrefetchCount(count int) QosOption { _ = "STUB: not implemented"; return *new(QosOption) }

// WithQosPrefetchSize set qos prefetch size option.
func WithQosPrefetchSize(size int) QosOption { _ = "STUB: not implemented"; return *new(QosOption) }

// WithQosPrefetchGlobal set qos global option.
func WithQosPrefetchGlobal(enable bool) QosOption {
	_ = "STUB: not implemented"
	return *new(QosOption)
}

// -------------------------------------------------------------------------------------------

// Consumer session
type Consumer struct {
	Exchange   *Exchange
	QueueName  string
	connection *Connection
	ch         *amqp.Channel

	exchangeDeclareOption *exchangeDeclareOptions
	queueDeclareOption    *queueDeclareOptions
	queueBindOption       *queueBindOptions
	qosOption             *qosOptions
	consumeOption         *consumeOptions

	isPersistent bool // persistent or not
	isAutoAck    bool // auto ack or not

	zapLog *zap.Logger

	count int64 // consumer success message number
}

// Handler message
type Handler func(ctx context.Context, data []byte, tagID string) error

//type Handler func(ctx context.Context, d *amqp.Delivery, isAutoAck bool) error

// NewConsumer create a consumer
func NewConsumer(exchange *Exchange, queueName string, connection *Connection, opts ...ConsumerOption) (*Consumer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// initialize a consumer session
func (c *Consumer) initialize() error { _ = "STUB: not implemented"; return nil }

// crate a new channel

// declare the exchange type

// declare a queue and create it automatically if it doesn't exist, or skip creation if it does.

// binding queue and exchange

// setting the prefetch value, set channel.Qos on the consumer side to limit the number of messages consumed at a time,
// balancing message throughput and fairness, and prevent consumers from being hit by sudden bursts of information traffic.

func (c *Consumer) consumeWithContext(ctx context.Context) (<-chan amqp.Delivery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Consume messages for loop in goroutine
func (c *Consumer) Consume(ctx context.Context, handler Handler) { _ = "STUB: not implemented"; return }

// check connection for loop

// Close consumer
func (c *Consumer) Close() { _ = "STUB: not implemented"; return }

// Count consumer success message number
func (c *Consumer) Count() int64 { _ = "STUB: not implemented"; return 0 }
