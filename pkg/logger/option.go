package logger

import (
	"go.uber.org/zap/zapcore"
)

var (
	defaultLevel    = "debug" // output log levels debug, info, warn, error, default is debug
	defaultEncoding = formatConsole
	defaultIsSave   = false // false:output to terminal, true:output to file, default is false

	defaultFilename      = "out.log" // file name
	defaultMaxSize       = 10        // maximum file size (MB)
	defaultMaxBackups    = 100       // maximum number of old files
	defaultMaxAge        = 30        // maximum number of days for old documents
	defaultIsCompression = false     // whether to compress and archive old files
	defaultIsLocalTime   = true      // whether to use local time
)

type options struct {
	level    string
	encoding string
	isSave   bool

	fileConfig *fileOptions

	hooks []func(zapcore.Entry) error
}

func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// Option set the logger options.
type Option func(*options)

// WithLevel setting the log level
func WithLevel(levelName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithFormat set the output log format, console or json
func WithFormat(format string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSave save log to file
func WithSave(isSave bool, opts ...FileOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithHooks set the log hooks
func WithHooks(hooks ...func(zapcore.Entry) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// ------------------------------------------------------------------------------------------

type fileOptions struct {
	filename      string
	maxSize       int
	maxBackups    int
	maxAge        int
	isCompression bool
	isLocalTime   bool
}

func defaultFileOptions() *fileOptions { _ = "STUB: not implemented"; return nil }

func (o *fileOptions) apply(opts ...FileOption) { _ = "STUB: not implemented"; return }

// FileOption set the file options.
type FileOption func(*fileOptions)

// WithFileName set log filename
func WithFileName(filename string) FileOption { _ = "STUB: not implemented"; return *new(FileOption) }

// WithFileMaxSize set maximum file size (MB)
func WithFileMaxSize(maxSize int) FileOption { _ = "STUB: not implemented"; return *new(FileOption) }

// WithFileMaxBackups set maximum number of old files
func WithFileMaxBackups(maxBackups int) FileOption {
	_ = "STUB: not implemented"
	return *new(FileOption)
}

// WithFileMaxAge set maximum number of days for old documents
func WithFileMaxAge(maxAge int) FileOption { _ = "STUB: not implemented"; return *new(FileOption) }

// WithFileIsCompression set whether to compress log files
func WithFileIsCompression(isCompression bool) FileOption {
	_ = "STUB: not implemented"
	return *new(FileOption)
}

// WithLocalTime set whether to use local time
func WithLocalTime(isLocalTime bool) FileOption { _ = "STUB: not implemented"; return *new(FileOption) }
