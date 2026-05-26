// Package errcode is used for http and grpc error codes, include system-level error codes and business-level error codes
package errcode

// ToHTTPCodeLabel need to convert to standard http code label
const ToHTTPCodeLabel = "[standard http code]"

var errCodes = map[int]*Error{}
var httpErrCodes = map[int]string{}

// Error error
type Error struct {
	code    int
	msg     string
	details []string

	// if true, need to convert to standard http code
	// use ErrToHTTP and ParseError will set this to true
	needHTTPCode bool
}

// NewError create a new error message
func NewError(code int, msg string, details ...string) *Error {
	_ = "STUB: not implemented"
	return nil
}

// Err convert to standard error,
// if there is a parameter 'msg', it will replace the original message.
func (e *Error) Err(msg ...string) error { _ = "STUB: not implemented"; return nil }

// ErrToHTTP convert to standard error add ToHTTPCodeLabel to error message,
// use it if you need to convert to standard HTTP status code,
// if there is a parameter 'msg', it will replace the original message.
// Tips: you can call the GetErrorCode function to get the standard HTTP status code.
func (e *Error) ErrToHTTP(msg ...string) error { _ = "STUB: not implemented"; return nil }

// Code get error code
func (e *Error) Code() int {
	_ = "STUB: not implemented"

	// Msg get error code message
	return 0
}

func (e *Error) Msg() string {
	_ = "STUB: not implemented"

	// NeedHTTPCode need to convert to standard http code
	return ""
}

func (e *Error) NeedHTTPCode() bool { _ = "STUB: not implemented"; return false }

// Details get error code details
func (e *Error) Details() []string {
	_ = "STUB: not implemented"

	// WithDetails add error details
	return nil
}

func (e *Error) WithDetails(details ...string) *Error { _ = "STUB: not implemented"; return nil }

// RewriteMsg rewrite error message
func (e *Error) RewriteMsg(msg string) *Error { _ = "STUB: not implemented"; return nil }

// WithOutMsg out error message
// Deprecated: use RewriteMsg instead
func (e *Error) WithOutMsg(msg string) *Error { _ = "STUB: not implemented"; return nil }

// WithOutMsgI18n out error message i18n
// langMsg example:
//
//	map[int]map[string]string{
//			20010: {
//				"en-US": "login failed",
//				"zh-CN": "登录失败",
//			},
//		}
//
// lang BCP 47 code https://learn.microsoft.com/en-us/openspecs/office_standards/ms-oe376/6c085406-a698-4e12-9d4d-c3b0ee3dbc4a
func (e *Error) WithOutMsgI18n(langMsg map[int]map[string]string, lang string) *Error {
	_ = "STUB: not implemented"
	return nil
}

// ToHTTPCode convert to http error code
func (e *Error) ToHTTPCode() int { _ = "STUB: not implemented"; return 0 }

// ParseError parsing out error codes from error messages
func ParseError(err error) *Error { _ = "STUB: not implemented"; return nil }

// GetErrorCode get Error code from error returned by http invoke
func GetErrorCode(err error) int { _ = "STUB: not implemented"; return 0 }

// Is check if error is equal to target error
func Is(err error, e *Error) bool { _ = "STUB: not implemented"; return false }

// ListHTTPErrCodes list http error codes
func ListHTTPErrCodes() []ErrInfo { _ = "STUB: not implemented"; return nil }
