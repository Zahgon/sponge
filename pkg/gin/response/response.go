// Package response provides wrapper gin returns json data in the same format.
package response

import (
	"github.com/gin-gonic/gin"

	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// Result output data format
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func newResp(code int, msg string, data interface{}) *Result { _ = "STUB: not implemented"; return nil }

// ensure that the data field is not nil on return, note that it is not nil when resp.data=[]interface {}, it is serialized to null

func respJSONWithStatusCode(c *gin.Context, code int, msg string, data ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Output return standard HTTP status codes and message, parameter code is HTTP status code
func Output(c *gin.Context, code int, data ...interface{}) { _ = "STUB: not implemented"; return }

// Out returns the standard HTTP status code and message, parameter err is errcode.Error
func Out(c *gin.Context, err *errcode.Error, data ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// status code flat 200, custom error codes in data.code
func respJSONWith200(c *gin.Context, code int, msg string, data ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Success return success
func Success(c *gin.Context, data ...interface{}) { _ = "STUB: not implemented"; return }

// Error return error
func Error(c *gin.Context, err *errcode.Error, data ...interface{}) {
	_ = "STUB: not implemented"
	return
}
