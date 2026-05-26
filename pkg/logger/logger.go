// Package logger is log library encapsulated in https://github.com/uber-go/zap
//
// Support for terminal printing and log saving.
// Support for automatic log file cutting.
// Support for json format and console log format output.
// Supports Debug, Info, Warn, Error, Panic, Fatal, also supports fmt.Printf-like log printing, Debugf, Infof, Warnf, Errorf, Panicf, Fatalf.
package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	formatConsole = "console"
	formatJSON    = "json"

	levelDebug = "DEBUG"
	levelInfo  = "INFO"
	levelWarn  = "WARN"
	levelError = "ERROR"
)

var defaultLogger *zap.Logger
var defaultSugaredLogger *zap.SugaredLogger

func getLogger() *zap.Logger { _ = "STUB: not implemented"; return nil }

func getSugaredLogger() *zap.SugaredLogger { _ = "STUB: not implemented"; return nil }

// Init initial log settings
// print the debug level log in the terminal, example: Init()
// print the info level log in the terminal, example: Init(WithLevel("info"))
// print the json format, debug level log in the terminal, example: Init(WithFormat("json"))
// log with hooks, example: Init(WithHooks(func(zapcore.Entry) error{return nil}))
// output the log to the file out.log, using the default cut log-related parameters, debug-level log, example: Init(WithSave())
// output the log to the specified file, custom set the log file cut log parameters, json format, debug level log, example:
// Init(
//
//	  WithFormat("json"),
//	  WithSave(true,
//
//			WithFileName("my.log"),
//			WithFileMaxSize(5),
//			WithFileMaxBackups(5),
//			WithFileMaxAge(10),
//			WithFileIsCompression(true),
//		))
func Init(opts ...Option) (*zap.Logger, error) { _ = "STUB: not implemented"; return nil, nil }

func log2Terminal(levelName string, encoding string) (*zap.Logger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// logging color

// logging levels in the log file using upper case letters

// default time format

func log2File(encoding string, levelName string, fo *fileOptions) *zap.Logger {
	_ = "STUB: not implemented"
	return nil
}

// modify Time Encoder
// logging levels in the log file using upper case letters

// console format

// json format

// file name
// maximum file size (MB)
// maximum number of old files
// maximum number of days for old documents
// whether to compress and archive old files

// add the function call information log to the log.

// DEBUG(default), INFO, WARN, ERROR
func getLevelSize(levelName string) zapcore.Level {
	_ = "STUB: not implemented"
	return *new(zapcore.Level)
}

func timeFormatter(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

// GetWithSkip get defaultLogger, set the skipped caller value, customize the number of lines of code displayed
func GetWithSkip(skip int) *zap.Logger { _ = "STUB: not implemented"; return nil }

// Get logger
func Get() *zap.Logger { _ = "STUB: not implemented"; return nil }

func checkNil() { _ = "STUB: not implemented"; return }

// default output to console
