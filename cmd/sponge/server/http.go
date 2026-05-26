// Package server is a sponge UI service that contains the front-end pages.
package server

import (
	"embed"

	"github.com/gin-gonic/gin"
)

//go:embed static
var staticFS embed.FS // index.html in the static directory

var defaultAddr = "http://localhost:24631"
var frontendDir = "frontend"
var ConfigJsFile = "static/appConfig.js"

// NewRouter create a router
func NewRouter(spongeAddr string, isLog bool) *gin.Engine { _ = "STUB: not implemented"; return nil }

// solve vue using history route 404 problem

// determine whether you need to use Embed.FS static resources based on the default configured address,
// if it is not the default address, copy the read-only Embed.FS static resources locally and then modify the default
// configured address, so dynamically configure the service address based on the parameter.

// RunHTTPServer run http server
func RunHTTPServer(spongeAddr string, port int, isLog bool) { _ = "STUB: not implemented"; return }

func checkIsUseEmbedFS(targetDir string, spongeAddr string) bool {
	_ = "STUB: not implemented"
	return false
}

func saveFSToLocal(targetDir string, spongeAddr string) error {
	_ = "STUB: not implemented"
	return nil
}

// Walk through the embedded filesystem

// Create the corresponding directory structure locally

// Read the file from the embedded filesystem

// replace config address

// Save the content to the local file
