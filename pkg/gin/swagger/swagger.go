// Package swagger is gin swagger library.
package swagger

import (
	"github.com/gin-gonic/gin"
)

// DefaultRouter default swagger router, request url is http://<ip:port>/swagger/index.html
func DefaultRouter(r *gin.Engine, jsonContent []byte) { _ = "STUB: not implemented"; return }

// DefaultRouterByFile  default swagger router from file, request url is http://<ip:port>/swagger/index.html
func DefaultRouterByFile(r *gin.Engine, jsonFile string) { _ = "STUB: not implemented"; return }

// CustomRouter custom swagger routing, request url is http://<ip:port>/<name>/swagger/index.html
func CustomRouter(r *gin.Engine, name string, jsonContent []byte) {
	_ = "STUB: not implemented"
	return
}

// CustomRouterByFile custom swagger router from file, request url is http://<ip:port>/<filename prefix>/swagger/index.html
func CustomRouterByFile(r *gin.Engine, jsonFile string) { _ = "STUB: not implemented"; return }

func registerSwagger(infoInstanceName string, jsonContent []byte) {
	_ = "STUB: not implemented"
	return
}
