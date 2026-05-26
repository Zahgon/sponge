package kafka

import (
	"crypto/tls"
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

// -------------------------------------- sync producer ------------------------------------

// SyncProducerOption set options.
type SyncProducerOption func(*syncProducerOptions)

type syncProducerOptions struct {
	version         sarama.KafkaVersion           // default V2_1_0_0
	requiredAcks    sarama.RequiredAcks           // default WaitForAll
	partitioner     sarama.PartitionerConstructor // default NewHashPartitioner
	returnSuccesses bool                          // default true
	clientID        string                        // default "sarama"
	tlsConfig       *tls.Config                   // default nil

	// custom config, if not nil, it will override the default config, the above parameters are invalid
	config *sarama.Config // default nil
}

func (o *syncProducerOptions) apply(opts ...SyncProducerOption) { _ = "STUB: not implemented"; return }

func defaultSyncProducerOptions() *syncProducerOptions { _ = "STUB: not implemented"; return nil }

// SyncProducerWithVersion set kafka version.
func SyncProducerWithVersion(version sarama.KafkaVersion) SyncProducerOption {
	_ = "STUB: not implemented"
	return *new(SyncProducerOption)
}

// SyncProducerWithRequiredAcks set requiredAcks.
func SyncProducerWithRequiredAcks(requiredAcks sarama.RequiredAcks) SyncProducerOption {
	_ = "STUB: not implemented"
	return *new(SyncProducerOption)
}

// SyncProducerWithPartitioner set partitioner.
func SyncProducerWithPartitioner(partitioner sarama.PartitionerConstructor) SyncProducerOption {
	_ = "STUB: not implemented"
	return *new(SyncProducerOption)
}

// SyncProducerWithReturnSuccesses set returnSuccesses.
func SyncProducerWithReturnSuccesses(returnSuccesses bool) SyncProducerOption {
	_ = "STUB: not implemented"
	return *new(SyncProducerOption)
}

// SyncProducerWithClientID set clientID.
func SyncProducerWithClientID(clientID string) SyncProducerOption {
	_ = "STUB: not implemented"
	return *new(SyncProducerOption)
}

// SyncProducerWithTLS set tlsConfig, if isSkipVerify is true, crypto/tls accepts any certificate presented by
// the server and any host name in that certificate.
func SyncProducerWithTLS(certFile, keyFile, caFile string, isSkipVerify bool) SyncProducerOption {
	_ = "STUB: not implemented"
	return *new(SyncProducerOption)
}

// SyncProducerWithConfig set custom config.
func SyncProducerWithConfig(config *sarama.Config) SyncProducerOption {
	_ = "STUB: not implemented"
	return *new(SyncProducerOption)
}

// -------------------------------------- async producer -----------------------------------

// AsyncSendFailedHandlerFn is a function that handles failed messages.
type AsyncSendFailedHandlerFn func(msg *sarama.ProducerMessage) error

// AsyncProducerOption set options.
type AsyncProducerOption func(*asyncProducerOptions)

type asyncProducerOptions struct {
	version         sarama.KafkaVersion           // default V2_1_0_0
	requiredAcks    sarama.RequiredAcks           // default WaitForLocal
	partitioner     sarama.PartitionerConstructor // default NewHashPartitioner
	returnSuccesses bool                          // default true
	clientID        string                        // default "sarama"
	flushMessages   int                           // default 20
	flushFrequency  time.Duration                 // default 2 second
	flushBytes      int                           // default 0
	tlsConfig       *tls.Config

	// custom config, if not nil, it will override the default config, the above parameters are invalid
	config *sarama.Config // default nil

	zapLogger      *zap.Logger              // default NewProduction
	handleFailedFn AsyncSendFailedHandlerFn // default nil
}

func (o *asyncProducerOptions) apply(opts ...AsyncProducerOption) {
	_ = "STUB: not implemented"
	return
}

func defaultAsyncProducerOptions() *asyncProducerOptions { _ = "STUB: not implemented"; return nil }

// AsyncProducerWithVersion set kafka version.
func AsyncProducerWithVersion(version sarama.KafkaVersion) AsyncProducerOption {
	_ = "STUB: not implemented"
	return *new(AsyncProducerOption)
}

// AsyncProducerWithRequiredAcks set requiredAcks.
func AsyncProducerWithRequiredAcks(requiredAcks sarama.RequiredAcks) AsyncProducerOption {
	_ = "STUB: not implemented"
	return *new(AsyncProducerOption)
}

// AsyncProducerWithPartitioner set partitioner.
func AsyncProducerWithPartitioner(partitioner sarama.PartitionerConstructor) AsyncProducerOption {
	_ = "STUB: not implemented"
	return *new(AsyncProducerOption)
}

// AsyncProducerWithReturnSuccesses set returnSuccesses.
func AsyncProducerWithReturnSuccesses(returnSuccesses bool) AsyncProducerOption {
	_ = "STUB: not implemented"
	return *new(AsyncProducerOption)
}

// AsyncProducerWithClientID set clientID.
func AsyncProducerWithClientID(clientID string) AsyncProducerOption {
	_ = "STUB: not implemented"
	return *new(AsyncProducerOption)
}

// AsyncProducerWithFlushMessages set flushMessages.
func AsyncProducerWithFlushMessages(flushMessages int) AsyncProducerOption {
	_ = "STUB: not implemented"
	return *new(AsyncProducerOption)
}

// AsyncProducerWithFlushFrequency set flushFrequency.
func AsyncProducerWithFlushFrequency(flushFrequency time.Duration) AsyncProducerOption {
	_ = "STUB: not implemented"
	return *new(AsyncProducerOption)
}

// AsyncProducerWithFlushBytes set flushBytes.
func AsyncProducerWithFlushBytes(flushBytes int) AsyncProducerOption {
	_ = "STUB: not implemented"
	return *new(AsyncProducerOption)
}

// AsyncProducerWithTLS set tlsConfig, if isSkipVerify is true, crypto/tls accepts any certificate presented by
// the server and any host name in that certificate.
func AsyncProducerWithTLS(certFile, keyFile, caFile string, isSkipVerify bool) AsyncProducerOption {
	_ = "STUB: not implemented"
	return *new(AsyncProducerOption)
}

// AsyncProducerWithZapLogger set zapLogger.
func AsyncProducerWithZapLogger(zapLogger *zap.Logger) AsyncProducerOption {
	_ = "STUB: not implemented"
	return *new(AsyncProducerOption)
}

// AsyncProducerWithHandleFailed set handleFailedFn.
func AsyncProducerWithHandleFailed(handleFailedFn AsyncSendFailedHandlerFn) AsyncProducerOption {
	_ = "STUB: not implemented"
	return *new(AsyncProducerOption)
}

// AsyncProducerWithConfig set custom config.
func AsyncProducerWithConfig(config *sarama.Config) AsyncProducerOption {
	_ = "STUB: not implemented"
	return *new(AsyncProducerOption)
}

func getTLSConfig(certFile, keyFile, caFile string, isSkipVerify bool) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
