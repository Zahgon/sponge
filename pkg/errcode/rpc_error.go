package errcode

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var grpcErrCodes = map[int]string{}

// RPCStatus rpc status
type RPCStatus struct {
	status *status.Status
}

var statusCodes = map[codes.Code]string{}

// NewRPCStatus create a new rpc status
func NewRPCStatus(code codes.Code, msg string) *RPCStatus { _ = "STUB: not implemented"; return nil }

// Detail error details
type Detail struct {
	key string
	val interface{}
}

// String detail key-value
func (d *Detail) String() string { _ = "STUB: not implemented"; return "" }

// Any type key value
func Any(key string, val interface{}) Detail { _ = "STUB: not implemented"; return *new(Detail) }

// Code get code
func (s *RPCStatus) Code() codes.Code {
	_ = "STUB: not implemented"
	return *

	// Msg get message
	new(codes.Code)
}

func (s *RPCStatus) Msg() string { _ = "STUB: not implemented"; return "" }

// Err return error
// if there is a parameter 'desc', it will replace the original message
func (s *RPCStatus) Err(desc ...string) error { _ = "STUB: not implemented"; return nil }

// ErrToHTTP convert to standard error add ToHTTPCodeLabel to error message,
// usually used when HTTP calls the GRPC API,
// if there is a parameter 'desc', it will replace the original message.
func (s *RPCStatus) ErrToHTTP(desc ...string) error { _ = "STUB: not implemented"; return nil }

// ToRPCErr converted to standard RPC error,
// use it if you need to convert to standard RPC errors,
// if there is a parameter 'desc', it will replace the original message.
func (s *RPCStatus) ToRPCErr(desc ...string) error { _ = "STUB: not implemented"; return nil }

func toRPCErr(code codes.Code, descs ...string) error { _ = "STUB: not implemented"; return nil }

// ToRPCCode converted to standard RPC error code
func (s *RPCStatus) ToRPCCode() codes.Code { _ = "STUB: not implemented"; return *new(codes.Code) }

// converted grpc code to http code
func convertToHTTPCode(code codes.Code) int { _ = "STUB: not implemented"; return 0 }

// GetStatusCode get status code from error returned by RPC invoke
func GetStatusCode(err error) codes.Code { _ = "STUB: not implemented"; return *new(codes.Code) }

// IsStatus check if error is a specific status
func IsStatus(err error, e *RPCStatus) bool { _ = "STUB: not implemented"; return false }

// ErrInfo error info
type ErrInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func getErrorInfo(codeInfo map[int]string) []ErrInfo { _ = "STUB: not implemented"; return nil }

// ListGRPCErrCodes list grpc error codes, http handle func
func ListGRPCErrCodes(w http.ResponseWriter, _ *http.Request) { _ = "STUB: not implemented"; return }

// ShowConfig show config info
// @Summary get system configuration
// @Description Returns the current system configuration in JSON format. This includes all runtime configuration parameters.
// @Tags system
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "Returns the complete system configuration"
// @Router /config [get]
func ShowConfig(jsonData []byte) func(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return nil
}

//w.Header().Set("Content-Type", "application/json")
