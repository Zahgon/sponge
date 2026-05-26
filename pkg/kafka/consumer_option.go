package kafka

import (
	"crypto/tls"
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

// HandleMessageFn is a function that handles a message from a partition consumer
type HandleMessageFn func(msg *sarama.ConsumerMessage) error

// ConsumerOption set options.
type ConsumerOption func(*consumerOptions)

type consumerOptions struct {
	version   sarama.KafkaVersion // default V2_1_0_0
	clientID  string              // default "sarama"
	tlsConfig *tls.Config         // default nil

	// consumer group options
	groupStrategies           []sarama.BalanceStrategy // default NewBalanceStrategyRange
	offsetsInitial            int64                    // default OffsetOldest
	offsetsAutoCommitEnable   bool                     // default true
	offsetsAutoCommitInterval time.Duration            // default 1s, when offsetsAutoCommitEnable is true

	// custom config, if not nil, it will override the default config, the above parameters are invalid
	config *sarama.Config // default nil

	zapLogger *zap.Logger // default NewProduction
}

func (o *consumerOptions) apply(opts ...ConsumerOption) { _ = "STUB: not implemented"; return }

func defaultConsumerOptions() *consumerOptions { _ = "STUB: not implemented"; return nil }

// ConsumerWithVersion set kafka version.
func ConsumerWithVersion(version sarama.KafkaVersion) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}

// ConsumerWithGroupStrategies set groupStrategies.
func ConsumerWithGroupStrategies(groupStrategies ...sarama.BalanceStrategy) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}

// ConsumerWithOffsetsInitial set offsetsInitial.
func ConsumerWithOffsetsInitial(offsetsInitial int64) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}

// ConsumerWithOffsetsAutoCommitEnable set offsetsAutoCommitEnable.
func ConsumerWithOffsetsAutoCommitEnable(offsetsAutoCommitEnable bool) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}

// ConsumerWithOffsetsAutoCommitInterval set offsetsAutoCommitInterval.
func ConsumerWithOffsetsAutoCommitInterval(offsetsAutoCommitInterval time.Duration) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}

// ConsumerWithClientID set clientID.
func ConsumerWithClientID(clientID string) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}

// ConsumerWithTLS set tlsConfig, if isSkipVerify is true, crypto/tls accepts any certificate presented by
// the server and any host name in that certificate.
func ConsumerWithTLS(certFile, keyFile, caFile string, isSkipVerify bool) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}

// ConsumerWithZapLogger set zapLogger.
func ConsumerWithZapLogger(zapLogger *zap.Logger) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}

// ConsumerWithConfig set custom config.
func ConsumerWithConfig(config *sarama.Config) ConsumerOption {
	_ = "STUB: not implemented"
	return *new(ConsumerOption)
}
