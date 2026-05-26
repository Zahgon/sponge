package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap/zapcore"
)

// Field type
type Field = zapcore.Field

// Int type
func Int(key string, val int) Field {
	_ = "STUB: not implemented"
	return *

	// Int32 type
	new(Field)
}

func Int32(key string, val int32) Field { _ = "STUB: not implemented"; return *new(Field) }

// Int64 type
func Int64(key string, val int64) Field { _ = "STUB: not implemented"; return *new(Field) }

// Uint type
func Uint(key string, val uint) Field {
	_ = "STUB: not implemented"
	return *

	// Uint32 type
	new(Field)
}

func Uint32(key string, val uint32) Field { _ = "STUB: not implemented"; return *new(Field) }

// Uint64 type
func Uint64(key string, val uint64) Field { _ = "STUB: not implemented"; return *new(Field) }

// Uintptr type
func Uintptr(key string, val uintptr) Field { _ = "STUB: not implemented"; return *new(Field) }

// Float64 type
func Float64(key string, val float64) Field { _ = "STUB: not implemented"; return *new(Field) }

// Bool type
func Bool(key string, val bool) Field {
	_ = "STUB: not implemented"
	return *

	// String type
	new(Field)
}

func String(key string, val string) Field { _ = "STUB: not implemented"; return *new(Field) }

// ByteString type
func ByteString(key string, val []byte) Field { _ = "STUB: not implemented"; return *new(Field) }

// Stringer type
func Stringer(key string, val fmt.Stringer) Field { _ = "STUB: not implemented"; return *new(Field) }

// Time type
func Time(key string, val time.Time) Field {
	_ = "STUB: not implemented"
	return *

	// Duration type
	new(Field)
}

func Duration(key string, val time.Duration) Field { _ = "STUB: not implemented"; return *new(Field) }

// Err type
func Err(err error) Field {
	_ = "STUB: not implemented"
	return *

	// Any type, if it is a composite type such as object, slice, map, etc., use Any
	new(Field)
}

func Any(key string, val interface{}) Field { _ = "STUB: not implemented"; return *new(Field) }
