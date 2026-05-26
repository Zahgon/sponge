// Package service is to generate template code, router code, and error code.
package service

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/go-dev-frame/sponge/cmd/protoc-gen-go-gin/internal/parse"
)

// GenerateFiles generate service logic, router, error code files.
func GenerateFiles(file *protogen.File, moduleName string) (logicContent []byte,
	routerFileContent []byte, errCodeFileContent []byte) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func genServiceLogicFile(fields []*parse.PbService) []byte { _ = "STUB: not implemented"; return nil }

func genRouterFile(fields []*parse.PbService) []byte { _ = "STUB: not implemented"; return nil }

func genErrCodeFile(fields []*parse.PbService) []byte { _ = "STUB: not implemented"; return nil }

type serviceLogicFields struct {
	PbServices []*parse.PbService
}

func (f *serviceLogicFields) execute() []byte { _ = "STUB: not implemented"; return nil }

type routerFields struct {
	PbServices []*parse.PbService
}

func (f *routerFields) execute() []byte { _ = "STUB: not implemented"; return nil }

type errCodeFields struct {
	PbServices []*parse.PbService
}

func (f *errCodeFields) execute() []byte { _ = "STUB: not implemented"; return nil }

const importPkgPathMark = "// import api service package here"
