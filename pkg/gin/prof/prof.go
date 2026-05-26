// Package prof is used for gin profiling.
package prof

import (
	"github.com/gin-gonic/gin"
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

// Register pprof for gin router
func Register(r *gin.Engine, opts ...Option) { _ = "STUB: not implemented"; return }

// Similar to /profile, add IO wait time,  https://github.com/felixge/fgprof
