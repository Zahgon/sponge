// Package handlerfunc is used for public http request handler.
package handlerfunc

import (
	"embed"

	"github.com/gin-gonic/gin"
)

// CheckHealthReply check health result
type CheckHealthReply struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}

// CheckHealth check healthy.
// @Summary check system health status
// @Description Returns system health information including status and hostname
// @Tags system
// @Accept  json
// @Produce  json
// @Success 200 {object} CheckHealthReply "Returns health status information"
// @Router /health [get]
func CheckHealth(c *gin.Context) { _ = "STUB: not implemented"; return }

type PingReply struct{}

// Ping the server
// @Summary ping the server
// @Description Simple ping endpoint to check if server is responsive
// @Tags system
// @Accept  json
// @Produce  json
// @Success 200 {object} PingReply "Returns empty JSON object"
// @Router /ping [get]
func Ping(c *gin.Context) { _ = "STUB: not implemented"; return }

// ListCodes list error codes info
// @Summary list all error codes
// @Description Returns a list of all defined HTTP error codes and their descriptions
// @Tags system
// @Accept  json
// @Produce  json
// @Success 200 {array} errcode.ErrInfo "List of error codes"
// @Router /codes [get]
func ListCodes(c *gin.Context) { _ = "STUB: not implemented"; return }

// BrowserRefresh solve vue using history route 404 problem, for system file
func BrowserRefresh(path string) func(c *gin.Context) { _ = "STUB: not implemented"; return nil }

// BrowserRefreshFS solve vue using history route 404 problem, for embed.FS
func BrowserRefreshFS(fs embed.FS, path string) func(c *gin.Context) {
	_ = "STUB: not implemented"
	return nil
}
