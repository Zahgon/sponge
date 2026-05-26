// Package module provides the functions to merge the code of two files.
package module

import (
	"github.com/go-dev-frame/sponge/pkg/goast"
)

// CodeAst is the struct for code
type CodeAst struct {
	FilePath string
	SrcCode  string
	FileSize int
	AstInfos []*goast.AstInfo

	replaceCodeMap         map[string]string   // src code -> new code
	excludeReceiverNameMap map[string]struct{} // receiver name -> exclude or not
	appendCodes            []string
}

func NewCodeAst(filePath string) (*CodeAst, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *CodeAst) compareExistedImportCode(genAst *CodeAst) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *CodeAst) compareExistedStructMethodsCode(genAst *CodeAst) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *CodeAst) findNonExistedCode(genAsts []*goast.AstInfo) { _ = "STUB: not implemented"; return }

type serviceMethods struct {
	methodNames    []string
	nameCodeBlocks map[string]string
}

func (s *serviceMethods) lastMethodCode() string { _ = "STUB: not implemented"; return "" }

// nolint
func (a *CodeAst) compareExistedGRPCMethodsTestCode(genAst *CodeAst) error {
	_ = "STUB: not implemented"
	return nil
}

// compare method names

func lastImportPath(astInfos []*goast.AstInfo) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (a *CodeAst) replaceCode() { _ = "STUB: not implemented"; return }

func parseImportCode(astInfos []*goast.AstInfo) ([]*goast.ImportInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseMethodFuncCode(astInfos []*goast.AstInfo) map[string][]*goast.MethodInfo {
	_ = "STUB: not implemented"
	return nil
}

// parse grpc methods test code block from source code
func parseGRPCMethodsTestCode(body string) ([]string, map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// find assignments to tests variables

// traversing the elements of the tests slice

// ParseHandlerAndServiceCode parses the source code and generated code of the handler and service file,
func ParseHandlerAndServiceCode(srcFile string, genFile string) (*CodeAst, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// compare existing import path

// compare existing struct methods

// get nonexistent variables

// replace source code

// ParseGRPCMethodsTestAndBenchmarkCode parses the source code and generated code of the service file
func ParseGRPCMethodsTestAndBenchmarkCode(srcFile string, genFile string) (*CodeAst, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// compare existing import path

// compare existing struct methods

// get nonexistent variables

// replace source code
