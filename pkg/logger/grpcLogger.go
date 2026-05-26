package logger

import (
	"go.uber.org/zap"
)

type grpcLogger struct {
	zLog      *zap.Logger
	verbosity int
}

// ReplaceGRPCLoggerV2 replace grpc logger v2
func ReplaceGRPCLoggerV2(l *zap.Logger) { _ = "STUB: not implemented"; return }

func (l *grpcLogger) Info(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *grpcLogger) Infoln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *grpcLogger) Infof(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *grpcLogger) Warning(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *grpcLogger) Warningln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *grpcLogger) Warningf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *grpcLogger) Error(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *grpcLogger) Errorln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *grpcLogger) Errorf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *grpcLogger) Fatal(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *grpcLogger) Fatalln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *grpcLogger) Fatalf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *grpcLogger) V(level int) bool { _ = "STUB: not implemented"; return false }
