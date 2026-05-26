package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

// ErrClosed closed
var ErrClosed = amqp.ErrClosed

const (
	exchangeTypeDirect         = "direct"
	exchangeTypeTopic          = "topic"
	exchangeTypeFanout         = "fanout"
	exchangeTypeHeaders        = "headers"
	exchangeTypeDelayedMessage = "x-delayed-message"

	// HeadersTypeAll all
	HeadersTypeAll HeadersType = "all"
	// HeadersTypeAny any
	HeadersTypeAny HeadersType = "any"
)

// HeadersType headers type
type HeadersType = string

// Exchange rabbitmq minimum management unit
type Exchange struct {
	name               string                 // exchange name
	eType              string                 // exchange type: direct, topic, fanout, headers, x-delayed-message
	routingKey         string                 // route key
	headersKeys        map[string]interface{} // this field is required if eType=headers.
	delayedMessageType string                 // this field is required if eType=headers, support direct, topic, fanout, headers
}

// Name exchange name
func (e *Exchange) Name() string {
	_ = "STUB: not implemented"

	// Type exchange type
	return ""
}

func (e *Exchange) Type() string {
	_ = "STUB: not implemented"

	// RoutingKey exchange routing key
	return ""
}

func (e *Exchange) RoutingKey() string { _ = "STUB: not implemented"; return "" }

// HeadersKeys exchange headers keys
func (e *Exchange) HeadersKeys() map[string]interface{} { _ = "STUB: not implemented"; return nil }

// DelayedMessageType exchange delayed message type
func (e *Exchange) DelayedMessageType() string { _ = "STUB: not implemented"; return "" }

// NewDirectExchange create a direct exchange
func NewDirectExchange(exchangeName string, routingKey string) *Exchange {
	_ = "STUB: not implemented"
	return nil
}

// NewTopicExchange create a topic exchange
func NewTopicExchange(exchangeName string, routingKey string) *Exchange {
	_ = "STUB: not implemented"
	return nil
}

// NewFanoutExchange create a fanout exchange
func NewFanoutExchange(exchangeName string) *Exchange { _ = "STUB: not implemented"; return nil }

// NewHeadersExchange create a headers exchange, the headerType supports "all" and "any"
func NewHeadersExchange(exchangeName string, headersType HeadersType, keys map[string]interface{}) *Exchange {
	_ = "STUB: not implemented"
	return nil
}

// NewDelayedMessageExchange create a delayed message exchange
func NewDelayedMessageExchange(exchangeName string, e *Exchange) *Exchange {
	_ = "STUB: not implemented"
	return nil
}

// -------------------------------------------------------------------------------------------

// QueueDeclareOption declare queue option.
type QueueDeclareOption func(*queueDeclareOptions)

type queueDeclareOptions struct {
	autoDelete bool       // delete automatically
	exclusive  bool       // exclusive (only available to the program that created it)
	noWait     bool       // block processing
	args       amqp.Table // additional properties
}

func (o *queueDeclareOptions) apply(opts ...QueueDeclareOption) { _ = "STUB: not implemented"; return }

// default queue declare settings
func defaultQueueDeclareOptions() *queueDeclareOptions { _ = "STUB: not implemented"; return nil }

// WithQueueDeclareAutoDelete set queue declare auto delete option.
func WithQueueDeclareAutoDelete(enable bool) QueueDeclareOption {
	_ = "STUB: not implemented"
	return *new(QueueDeclareOption)
}

// WithQueueDeclareExclusive set queue declare exclusive option.
func WithQueueDeclareExclusive(enable bool) QueueDeclareOption {
	_ = "STUB: not implemented"
	return *new(QueueDeclareOption)
}

// WithQueueDeclareNoWait set queue declare no wait option.
func WithQueueDeclareNoWait(enable bool) QueueDeclareOption {
	_ = "STUB: not implemented"
	return *new(QueueDeclareOption)
}

// WithQueueDeclareArgs set queue declare args option.
func WithQueueDeclareArgs(args map[string]interface{}) QueueDeclareOption {
	_ = "STUB: not implemented"
	return *new(QueueDeclareOption)
}

// -------------------------------------------------------------------------------------------

// ExchangeDeclareOption declare exchange option.
type ExchangeDeclareOption func(*exchangeDeclareOptions)

type exchangeDeclareOptions struct {
	autoDelete bool       // delete automatically
	internal   bool       // public or not, false means public
	noWait     bool       // block processing
	args       amqp.Table // additional properties
}

func (o *exchangeDeclareOptions) apply(opts ...ExchangeDeclareOption) {
	_ = "STUB: not implemented"
	return
}

// default exchange declare settings
func defaultExchangeDeclareOptions() *exchangeDeclareOptions { _ = "STUB: not implemented"; return nil }

//durable:    true,

// WithExchangeDeclareAutoDelete set exchange declare auto delete option.
func WithExchangeDeclareAutoDelete(enable bool) ExchangeDeclareOption {
	_ = "STUB: not implemented"
	return *new(ExchangeDeclareOption)
}

// WithExchangeDeclareInternal set exchange declare internal option.
func WithExchangeDeclareInternal(enable bool) ExchangeDeclareOption {
	_ = "STUB: not implemented"
	return *new(ExchangeDeclareOption)
}

// WithExchangeDeclareNoWait set exchange declare no wait option.
func WithExchangeDeclareNoWait(enable bool) ExchangeDeclareOption {
	_ = "STUB: not implemented"
	return *new(ExchangeDeclareOption)
}

// WithExchangeDeclareArgs set exchange declare args option.
func WithExchangeDeclareArgs(args map[string]interface{}) ExchangeDeclareOption {
	_ = "STUB: not implemented"
	return *new(ExchangeDeclareOption)
}

// -------------------------------------------------------------------------------------------

// QueueBindOption declare queue bind option.
type QueueBindOption func(*queueBindOptions)

type queueBindOptions struct {
	noWait bool       // block processing
	args   amqp.Table // this parameter is invalid if the type is headers.
}

func (o *queueBindOptions) apply(opts ...QueueBindOption) { _ = "STUB: not implemented"; return }

// default queue bind settings
func defaultQueueBindOptions() *queueBindOptions { _ = "STUB: not implemented"; return nil }

// WithQueueBindNoWait set queue bind no wait option.
func WithQueueBindNoWait(enable bool) QueueBindOption {
	_ = "STUB: not implemented"
	return *new(QueueBindOption)
}

// WithQueueBindArgs set queue bind args option.
func WithQueueBindArgs(args map[string]interface{}) QueueBindOption {
	_ = "STUB: not implemented"
	return *new(QueueBindOption)
}

// -------------------------------------------------------------------------------------------

// DelayedMessagePublishOption declare queue bind option.
type DelayedMessagePublishOption func(*delayedMessagePublishOptions)

type delayedMessagePublishOptions struct {
	topicKey    string                 // the topic message type must be required
	headersKeys map[string]interface{} // the headers message type must be required
}

func (o *delayedMessagePublishOptions) apply(opts ...DelayedMessagePublishOption) {
	_ = "STUB: not implemented"
	return
}

// default delayed message publish settings
func defaultDelayedMessagePublishOptions() *delayedMessagePublishOptions {
	_ = "STUB: not implemented"
	return nil
}

// WithDelayedMessagePublishTopicKey set delayed message publish topicKey option.
func WithDelayedMessagePublishTopicKey(topicKey string) DelayedMessagePublishOption {
	_ = "STUB: not implemented"
	return *new(DelayedMessagePublishOption)
}

// WithDelayedMessagePublishHeadersKeys set delayed message publish headersKeys option.
func WithDelayedMessagePublishHeadersKeys(headersKeys map[string]interface{}) DelayedMessagePublishOption {
	_ = "STUB: not implemented"
	return *new(DelayedMessagePublishOption)
}

// -------------------------------------------------------------------------------------------

// DeadLetterOption declare dead letter option.
type DeadLetterOption func(*deadLetterOptions)

type deadLetterOptions struct {
	exchangeName string
	queueName    string
	routingKey   string

	exchangeDeclare *exchangeDeclareOptions
	queueDeclare    *queueDeclareOptions
	queueBind       *queueBindOptions
}

func (o *deadLetterOptions) apply(opts ...DeadLetterOption) { _ = "STUB: not implemented"; return }

func (o *deadLetterOptions) isEnabled() bool { _ = "STUB: not implemented"; return false }

func defaultDeadLetterOptions() *deadLetterOptions { _ = "STUB: not implemented"; return nil }

// WithDeadLetterExchangeDeclareOptions set dead letter exchange declare option.
func WithDeadLetterExchangeDeclareOptions(opts ...ExchangeDeclareOption) DeadLetterOption {
	_ = "STUB: not implemented"
	return *new(DeadLetterOption)
}

// WithDeadLetterQueueDeclareOptions set dead letter queue declare option.
func WithDeadLetterQueueDeclareOptions(opts ...QueueDeclareOption) DeadLetterOption {
	_ = "STUB: not implemented"
	return *new(DeadLetterOption)
}

// WithDeadLetterQueueBindOptions set dead letter queue bind option.
func WithDeadLetterQueueBindOptions(opts ...QueueBindOption) DeadLetterOption {
	_ = "STUB: not implemented"
	return *new(DeadLetterOption)
}

// WithDeadLetter set dead letter exchange, queue, routing key.
func WithDeadLetter(exchangeName string, queueName string, routingKey string) DeadLetterOption {
	_ = "STUB: not implemented"
	return *new(DeadLetterOption)
}
