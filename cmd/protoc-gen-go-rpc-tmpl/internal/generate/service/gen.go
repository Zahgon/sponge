// Package service is to generate template code, test code, and error code.
package service

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/go-dev-frame/sponge/cmd/protoc-gen-go-rpc-tmpl/internal/parse"
)

// GenerateFiles generate service template code and error codes
func GenerateFiles(file *protogen.File, moduleName string) (serviceTmplContent []byte,
	serviceTestTmplContent []byte, errCodeFileContent []byte) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func genServiceTmplFile(fields []*parse.PbService) []byte { _ = "STUB: not implemented"; return nil }

func genServiceTestTmplFile(pbs []*parse.PbService) []byte { _ = "STUB: not implemented"; return nil }

func genErrCodeFile(fields []*parse.PbService) []byte { _ = "STUB: not implemented"; return nil }

type serviceTmplFields struct {
	PbServices []*parse.PbService
}

func (f *serviceTmplFields) execute() []byte { _ = "STUB: not implemented"; return nil }

type serviceTestTmplFields struct {
	PbServices []*parse.PbService
}

func (f *serviceTestTmplFields) execute() []byte { _ = "STUB: not implemented"; return nil }

type errCodeFields struct {
	PbServices []*parse.PbService
}

func (f *errCodeFields) execute() []byte { _ = "STUB: not implemented"; return nil }

const importPkgPathMark = "// import api service package here"
