package prof

import (
	"net/http"
)

var defaultPrefix = "/debug/pprof"

// Option set defaultPrefix func
type Option func(o *options)

type options struct {
	prefix           string
	enableIOWaitTime bool
}

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// WithPrefix set route defaultPrefix
func WithPrefix(prefix string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithIOWaitTime enable IO wait time
func WithIOWaitTime() Option { _ = "STUB: not implemented"; return *new(Option) }

// Register pprof server mux
func Register(mux *http.ServeMux, opts ...Option) { _ = "STUB: not implemented"; return }

// Similar to /profile, add IO wait time,  https://github.com/felixge/fgprof
