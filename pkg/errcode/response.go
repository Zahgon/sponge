package errcode

import (
	"errors"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

// SkipResponse skip response
var SkipResponse = errors.New("skip response") //nolint

// Responser response interface
type Responser interface {
	Success(ctx *gin.Context, data interface{})
	ParamError(ctx *gin.Context, err error)
	Error(ctx *gin.Context, err error) bool
}

// NewResponser creates a new responser, if isFromRPC=true, it means return from rpc, otherwise default return from http
func NewResponser(isFromRPC bool, httpErrors []*Error, rpcStatus []*RPCStatus) Responser {
	_ = "STUB: not implemented"
	return *new(Responser)
}

type defaultResponse struct {
	isFromRPC  bool // error comes from grpc, if not, default is from http
	httpErrors map[int]*Error
	rpcStatus  map[int]*RPCStatus
}

func (resp *defaultResponse) response(c *gin.Context, respStatus, code int, msg string, data interface{}) {
	_ = "STUB: not implemented"
	return
}

// Success response success information
func (resp *defaultResponse) Success(c *gin.Context, data interface{}) {
	_ = "STUB: not implemented"
	return
}

// ParamError response parameter error information, does not return an error message
func (resp *defaultResponse) ParamError(c *gin.Context, _ error) { _ = "STUB: not implemented"; return }

// Error response error information, if return true, means that the error code is converted to a standard http code,
// otherwise the return http code is always 200
func (resp *defaultResponse) Error(c *gin.Context, err error) bool {
	_ = "STUB: not implemented"
	return false

	// error from rpc and response the corresponding http code
}

// error from http and response http code

// error from grpc
func (resp *defaultResponse) handleRPCError(c *gin.Context, err error) bool {
	_ = "STUB: not implemented"
	return false
}

// user defined err, response 200

// non-conforming err

// err created using NewRPCStatus

// default error code to http

// check if you need to return the standard http code

// user defined error code to http

// response 200

// error from http
func (resp *defaultResponse) handleHTTPError(c *gin.Context, err error) bool {
	_ = "STUB: not implemented"
	return false

	// default error code to http
}

// user requests to return standard HTTP code, if e.ToHTTPCode() not match, will return of 500

// user defined error code to http

// response 200

func (resp *defaultResponse) isUserDefinedRPCErrorCode(c *gin.Context, errCode int) bool {
	_ = "STUB: not implemented"
	return false
}

func (resp *defaultResponse) isUserDefinedHTTPErrorCode(c *gin.Context, errCode int) bool {
	_ = "STUB: not implemented"
	return false
}

// ToHTTPErr converted to http error
func ToHTTPErr(st *status.Status) *Error {
	_ = "STUB: not implemented" //nolint
	return nil
}

func parseCodeAndMsg(errStr string) (int, string) { _ = "STUB: not implemented"; return 0, "" }
