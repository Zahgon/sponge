package gotest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler info
type Handler struct {
	TestData interface{}
	MockDao  *Dao
	IHandler interface{}

	Engine      *gin.Engine
	HTTPServer  *http.Server
	httpAddr    string
	requestAddr string
	routers     map[string]RouterInfo
}

// RouterInfo router info
type RouterInfo struct {
	FuncName    string
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

// NewHandler instantiated handler
func NewHandler(dao *Dao, testData interface{}) *Handler { _ = "STUB: not implemented"; return nil }

// GoRunHTTPServer run http server
func (h *Handler) GoRunHTTPServer(fns []RouterInfo) { _ = "STUB: not implemented"; return }

// GetRequestURL get request url from name
func (h *Handler) GetRequestURL(funcName string, pathVal ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// Close handler
func (h *Handler) Close() { _ = "STUB: not implemented"; return }

//nolint
