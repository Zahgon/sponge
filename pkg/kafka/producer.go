// Package kafka is a kafka client package.
package kafka

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

// ProducerMessage is sarama ProducerMessage
type ProducerMessage = sarama.ProducerMessage

// ---------------------------------- sync producer ---------------------------------------

// SyncProducer is a sync producer.
type SyncProducer struct {
	Producer sarama.SyncProducer
}

// InitSyncProducer init sync producer.
func InitSyncProducer(addrs []string, opts ...SyncProducerOption) (*SyncProducer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SendMessage sends a message to a topic.
func (p *SyncProducer) SendMessage(msg *sarama.ProducerMessage) (int32, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// SendData sends a message to a topic with multiple types of data.
func (p *SyncProducer) SendData(topic string, data interface{}) (int32, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// Close closes the producer.
func (p *SyncProducer) Close() error { _ = "STUB: not implemented"; return nil }

// Message is a message to be sent to a topic.
type Message struct {
	Topic string `json:"topic"`
	Data  []byte `json:"data"`
	Key   []byte `json:"key"`
}

// ---------------------------------- async producer ---------------------------------------

// AsyncProducer is async producer.
type AsyncProducer struct {
	Producer  sarama.AsyncProducer
	zapLogger *zap.Logger
	exit      chan struct{}
}

// InitAsyncProducer init async producer.
func InitAsyncProducer(addrs []string, opts ...AsyncProducerOption) (*AsyncProducer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SendMessage sends messages to a topic.
func (p *AsyncProducer) SendMessage(messages ...*sarama.ProducerMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// SendData sends messages to a topic with multiple types of data.
func (p *AsyncProducer) SendData(topic string, multiData ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// handleResponse handles the response of async producer, if producer message failed, you can handle it, e.g. add to other queue to handle later.
func (p *AsyncProducer) handleResponse(handleFn AsyncSendFailedHandlerFn) {
	_ = "STUB: not implemented"
	return
}

// Close closes the producer.
func (p *AsyncProducer) Close() error { _ = "STUB: not implemented"; return nil }

// ignore error
