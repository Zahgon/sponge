// Package router is to generate gin router code.
package router

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/go-dev-frame/sponge/cmd/protoc-gen-go-gin/internal/parse"
)

// GenerateFiles generate gin router code.
func GenerateFiles(file *protogen.File) []byte { _ = "STUB: not implemented"; return nil }

func genGinRouterFile(services parse.HTTPPbServices, goPackageName string) []byte {
	_ = "STUB: not implemented"
	return nil
}

type ginRouterFields struct {
	*parse.HTTPPbService
}

func (f *ginRouterFields) execute() []byte { _ = "STUB: not implemented"; return nil }

type importPkg struct {
	PackageName  string
	PackagePaths string
}

func (f *importPkg) execute() []byte { _ = "STUB: not implemented"; return nil }
