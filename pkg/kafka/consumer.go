package kafka

import (
	"context"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

// ---------------------------------- consume group---------------------------------------

// ConsumerGroup consume group
type ConsumerGroup struct {
	Group            sarama.ConsumerGroup
	groupID          string
	zapLogger        *zap.Logger
	autoCommitEnable bool
}

// InitConsumerGroup init consumer group
func InitConsumerGroup(addrs []string, groupID string, opts ...ConsumerOption) (*ConsumerGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConsumeLoop consume messages in a loop, with rebalanced handling
func (c *ConsumerGroup) ConsumeLoop(ctx context.Context, topics []string, handleMessageFn HandleMessageFn) {
	_ = "STUB: not implemented"
	return
}

// Normal exit (mostly due to rebalanced), should reset the backoff and wait for the next time to join the consumption group

// Consume consume messages
func (c *ConsumerGroup) Consume(ctx context.Context, topics []string, handleMessageFn HandleMessageFn) error {
	_ = "STUB: not implemented"
	return nil
}

// ConsumeCustomLoop consume messages for custom handler in a loop, with rebalanced handling
func (c *ConsumerGroup) ConsumeCustomLoop(ctx context.Context, topics []string, handler sarama.ConsumerGroupHandler) {
	_ = "STUB: not implemented"
	return
}

// Normal exit (mostly due to rebalanced), should reset the backoff and wait for the next time to join the consumption group

// ConsumeCustom consume messages for custom handler, you need to implement the sarama.ConsumerGroupHandler interface
func (c *ConsumerGroup) ConsumeCustom(ctx context.Context, topics []string, handler sarama.ConsumerGroupHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConsumerGroup) Close() error { _ = "STUB: not implemented"; return nil }

type defaultConsumerHandler struct {
	ctx              context.Context
	handleMessageFn  HandleMessageFn
	zapLogger        *zap.Logger
	autoCommitEnable bool
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (h *defaultConsumerHandler) Setup(sess sarama.ConsumerGroupSession) error {
	_ = "STUB: not implemented"
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (h *defaultConsumerHandler) Cleanup(sess sarama.ConsumerGroupSession) error {
	_ = "STUB: not implemented"
	return nil
}

// ConsumeClaim consumes messages
func (h *defaultConsumerHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------- consume partition------------------------------------

// Consumer consume partition
type Consumer struct {
	C         sarama.Consumer
	zapLogger *zap.Logger
}

// InitConsumer init consumer
func InitConsumer(addrs []string, opts ...ConsumerOption) (*Consumer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConsumePartition consumer one partition, blocking
func (c *Consumer) ConsumePartition(ctx context.Context, topic string, partition int32, offset int64, handleFn HandleMessageFn) {
	_ = "STUB: not implemented"
	return
}

// ConsumeAllPartition consumer all partitions, no blocking
func (c *Consumer) ConsumeAllPartition(ctx context.Context, topic string, offset int64, handleFn HandleMessageFn) {
	_ = "STUB: not implemented"
	return
}

// Close the consumer
func (c *Consumer) Close() error { _ = "STUB: not implemented"; return nil }
