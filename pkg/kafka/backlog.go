package kafka

import (
	"github.com/IBM/sarama"
)

// ClientManager client manager
type ClientManager struct {
	client        sarama.Client
	offsetManager sarama.OffsetManager
}

// Backlog info
type Backlog struct {
	Partition         int32 `json:"partition"`  // partition id
	Backlog           int64 `json:"backlog"`    // data backlog
	NextConsumeOffset int64 `json:"nextOffset"` // offset for next consumption
}

// InitClientManager init client manager
func InitClientManager(addrs []string, groupID string) (*ClientManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBacklog get topic backlog
func (m *ClientManager) GetBacklog(topic string) (int64, []*Backlog, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// get offset from kafka

// create topic/partition manager

// call sarama The NextOffset method of PartitionOffsetManager. Return the offset for the next consumption
// if the consumer group has not consumed the data for this section, the return value will be -1

// Close topic backlog
func (m *ClientManager) Close() error { _ = "STUB: not implemented"; return nil }
