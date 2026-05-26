// Package frontend is a library for serving static files in a Gin web application.
// It supports local static files and embedding static files in binary.
package frontend

import (
	"embed"

	"github.com/gin-gonic/gin"
)

type options struct {
	isUseEmbedFS    bool
	embedFS         embed.FS
	handleContentFn func(content []byte) []byte
	specifiedFile   map[string]struct{}
	is404ToHome     bool
}

func defaultOptions() *options {
	_ = "STUB: not implemented"

	// Option set the jwt options.
	return nil
}

type Option func(*options)

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

// WithEmbedFS set embedFS to use embed.FS static resources.
func WithEmbedFS(efs embed.FS) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHandleContent set function to handle content and specified files
func WithHandleContent(fn func(content []byte) []byte, files ...string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// With404ToHome set 404 to home page
func With404ToHome() Option { _ = "STUB: not implemented"; return *new(Option) }

// ------------------------------------------------------------------------------------

// FrontEnd is the frontend router configuration
type FrontEnd struct {
	//basePath string // custom prefix route path, default is empty.

	sourceDir string // directory where static files is located, e.g. dist, it's also gin's route path.

	isUseEmbedFS bool     // if true, use embed.FS, otherwise local static file.
	embedFS      embed.FS // embed.FS static resources.

	// only used for EmbedFS, e.g. config.js content, replace apiBaseUrl to backend api address.
	handleContentFn func(content []byte) []byte
	specifiedFile   map[string]struct{} // specified files to handle content, e.g. config.js

	// when request route notfound
	// true: redirect to index.html
	// false: returns 404, default is false.
	is404ToHome bool
}

// New create a new frontend, default use local static file, you can use WithEmbedFS to use embed.FS.
func New(sourceDir string, opts ...Option) *FrontEnd { _ = "STUB: not implemented"; return nil }

//basePath:        o.basePath,

// SetRouter set frontend router
func (f *FrontEnd) SetRouter(r *gin.Engine) error {
	_ = "STUB: not implemented"
	// use embed file
	return nil
}

// use local file

func (f *FrontEnd) setEmbedFSRouter(r *gin.Engine) { _ = "STUB: not implemented"; return }

// solve using history route 404 problem

func (f *FrontEnd) setLocalFileRouter(r *gin.Engine) { _ = "STUB: not implemented"; return }

// solve using history route 404 problem

func (f *FrontEnd) saveFSToLocal() error { _ = "STUB: not implemented"; return nil }

// Walk through the embedded filesystem

// Create the corresponding directory structure locally

// Read the file from the embedded filesystem

// handle file content

// Save the content to the local file

// solve vue using history route 404 problem, for embed.FS
func browserRefreshFS(efs embed.FS, path string) func(c *gin.Context) {
	_ = "STUB: not implemented"
	return nil
}

// solve vue using history route 404 problem, for system file
func browserRefresh(path string) func(c *gin.Context) { _ = "STUB: not implemented"; return nil }
