package module

import (
	"go/ast"
	"go/token"

	"github.com/go-dev-frame/sponge/pkg/goast"
)

// define the extracted information structure
type middlewareFuncInfo struct {
	funcLineSrcCode string
	singlePaths     []*singlePath
}

func (i *middlewareFuncInfo) String() string { _ = "STUB: not implemented"; return "" }

type singlePath struct {
	name        string // method->path
	lineContent string // whole line content
}

// extract information about setSinglePath calls from the function body
func extractSinglePaths(src string, fset *token.FileSet, fn *ast.FuncDecl) *middlewareFuncInfo {
	_ = "STUB: not implemented"
	return nil
}

// gets the starting and ending line numbers of the function body

// regular expression matches c. setSinglePath ("Method "," Path ",...)

// first parameter: HTTP method
// second parameter: Path

// whole line content

func findNonExistedSinglePaths(srcMiddlewareFunc *middlewareFuncInfo,
	genMiddlewareFunc *middlewareFuncInfo) (srcCode string, targetCode string) {
	_ = "STUB: not implemented"
	return "", ""
}

// RouterCodeAst is the struct for router code
type RouterCodeAst struct {
	FilePath string
	SrcCode  string
	FileSize int
	AstInfos []*goast.AstInfo

	moduleCodeMap map[string]string // src code -> new code
	appendCodes   []string
}

func NewRouterCodeAst(filePath string) (*RouterCodeAst, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *RouterCodeAst) parseMiddlewareFunc() (map[string]*middlewareFuncInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// traversing declarations in AST

// CompareExistedMiddlewareFunc compare the existed middlewareFunc in the source code
// and generated code, and find the non-existed singlePaths
func (a *RouterCodeAst) CompareExistedMiddlewareFunc(genAst *RouterCodeAst) error {
	_ = "STUB: not implemented"
	return nil
}

// FindNonExistedName find non-existed names in the src code
func (a *RouterCodeAst) FindNonExistedName(genAsts []*goast.AstInfo) {
	_ = "STUB: not implemented"
	return
}

// ParseRouterCode parse router code from source and generated file, and merge them.
func ParseRouterCode(srcFile string, genFile string) (*RouterCodeAst, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// compare existing middlewareFunc

// get nonexistent variables

// replace source code
