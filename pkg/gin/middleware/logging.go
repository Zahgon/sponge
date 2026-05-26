package middleware

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	// Print body max length
	defaultMaxLength = 300

	// default zap log
	defaultLogger, _ = zap.NewProduction()

	// Ignore route list
	defaultIgnoreRoutes = map[string]struct{}{
		"/ping":   {},
		"/pong":   {},
		"/health": {},
	}

	// Print error by specified codes
	printErrorBySpecifiedCodes = map[int]bool{
		http.StatusInternalServerError: true,
		http.StatusBadGateway:          true,
		http.StatusServiceUnavailable:  true,
	}

	emptyBody   = []byte("")
	contentMark = []byte(" ...... ")
)

// Option set the gin logger options.
type Option func(*options)

func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

type options struct {
	maxLength     int
	log           *zap.Logger
	ignoreRoutes  map[string]struct{}
	requestIDFrom int // 0: ignore, 1: from context, 2: from header
}

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// WithMaxLen logger content max length
func WithMaxLen(maxLen int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLog set log
func WithLog(log *zap.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithIgnoreRoutes no logger content routes
func WithIgnoreRoutes(routes ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPrintErrorByCodes set print error by specified codes
func WithPrintErrorByCodes(code ...int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRequestIDFromContext name is field in context, default value is request_id
func WithRequestIDFromContext() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRequestIDFromHeader name is field in header, default value is X-Request-Id
func WithRequestIDFromHeader() Option { _ = "STUB: not implemented"; return *new(Option) }

// ------------------------------------------------------------------------------------------

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// If there is sensitive information in the body, you can use WithIgnoreRoutes set the route to ignore logging
func getResponseBody(buf *bytes.Buffer, maxLen int) []byte { _ = "STUB: not implemented"; return nil }

// If there is sensitive information in the body, you can use WithIgnoreRoutes set the route to ignore logging
func getRequestBody(buf *bytes.Buffer, maxLen int) []byte { _ = "STUB: not implemented"; return nil }

// Logging print request and response info
func Logging(opts ...Option) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

// ignore printing of the specified route

// print input information before processing

// replace writer

// processing requests

// print response message after processing

// SimpleLog print response info
func SimpleLog(opts ...Option) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

// ignore printing of the specified route

// processing requests

// print return message after processing
