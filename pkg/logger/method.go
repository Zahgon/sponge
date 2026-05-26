package logger

import (
	"go.uber.org/zap"
)

// Debug level information
func Debug(msg string, fields ...Field) { _ = "STUB: not implemented"; return }

// Info level information
func Info(msg string, fields ...Field) { _ = "STUB: not implemented"; return }

// Warn level information
func Warn(msg string, fields ...Field) { _ = "STUB: not implemented"; return }

// Error level information
func Error(msg string, fields ...Field) { _ = "STUB: not implemented"; return }

// Panic level information
func Panic(msg string, fields ...Field) { _ = "STUB: not implemented"; return }

// Fatal level information
func Fatal(msg string, fields ...Field) { _ = "STUB: not implemented"; return }

// Debugf format level information
func Debugf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// Infof format level information
func Infof(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// Warnf format level information
func Warnf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// Errorf format level information
func Errorf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// Fatalf format level information
func Fatalf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// Sync flushing any buffered log entries, applications should take care to call Sync before exiting.
func Sync() error { _ = "STUB: not implemented"; return nil }

// WithFields carrying field information
func WithFields(fields ...Field) *zap.Logger { _ = "STUB: not implemented"; return nil }
