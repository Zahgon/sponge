// Package handler is to generate template code, router code, and error code.
package handler

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/go-dev-frame/sponge/cmd/protoc-gen-go-gin/internal/parse"
)

// GenerateFiles generate handler logic, router, error code files.
func GenerateFiles(file *protogen.File, isMixType bool, moduleName string) (logicContent []byte, routerFileContent []byte, errCodeFileContent []byte) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func genHandlerLogicFile(fields []*parse.PbService) []byte { _ = "STUB: not implemented"; return nil }

func genRouterFile(fields []*parse.PbService) []byte { _ = "STUB: not implemented"; return nil }

func genErrCodeFile(fields []*parse.PbService) []byte { _ = "STUB: not implemented"; return nil }

func genMixLogicFile(fields []*parse.PbService) []byte { _ = "STUB: not implemented"; return nil }

func genMixRouterFile(fields []*parse.PbService) []byte { _ = "STUB: not implemented"; return nil }

type handlerLogicFields struct {
	PbServices []*parse.PbService
}

func (f *handlerLogicFields) execute() []byte { _ = "STUB: not implemented"; return nil }

type routerFields struct {
	PbServices []*parse.PbService
}

func (f *routerFields) execute() []byte { _ = "STUB: not implemented"; return nil }

type errCodeFields struct {
	PbServices []*parse.PbService
}

func (f *errCodeFields) execute() []byte { _ = "STUB: not implemented"; return nil }

type mixLogicFields struct {
	PbServices []*parse.PbService
}

func (f *mixLogicFields) execute() []byte { _ = "STUB: not implemented"; return nil }

type mixRouterFields struct {
	PbServices []*parse.PbService
}

func (f *mixRouterFields) execute() []byte { _ = "STUB: not implemented"; return nil }

const importPkgPathMark = "// import api service package here"
