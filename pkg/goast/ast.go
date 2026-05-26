// Package goast is a library for parsing Go code and extracting information from it.
package goast

import (
	"go/ast"
	"go/token"
)

const (
	// ast types

	PackageType = "package"
	ImportType  = "import"
	ConstType   = "const"
	VarType     = "var"
	FuncType    = "func"
	TypeType    = "type"

	// for TypeType

	StructType    = "struct"
	InterfaceType = "interface"
	ArrayType     = "array"
	MapType       = "map"
	ChanType      = "chan"
)

// AstInfo Go code block information
type AstInfo struct {
	// Type is the type of the code block, such as "func", "type", "const", "var", "import", "package".
	Type string

	// Names is the name of the code block, such as "func Name", "type Names", "const Names", "var Names", "import Paths".
	// If Type is "func", a standalone function without a receiver has a single name.
	// If the function is a method belonging to a struct, it has two names: the first
	// represents the function name, and the second represents the struct name.
	Names []string

	Comment string

	Body string
}

func (a *AstInfo) IsPackageType() bool { _ = "STUB: not implemented"; return false }

func (a *AstInfo) IsImportType() bool { _ = "STUB: not implemented"; return false }

func (a *AstInfo) IsConstType() bool { _ = "STUB: not implemented"; return false }

func (a *AstInfo) IsVarType() bool { _ = "STUB: not implemented"; return false }

func (a *AstInfo) IsTypeType() bool { _ = "STUB: not implemented"; return false }

func (a *AstInfo) IsFuncType() bool { _ = "STUB: not implemented"; return false }

func (a *AstInfo) GetName() string { _ = "STUB: not implemented"; return "" }

// ParseFile parses a go file and returns a list of AstInfo
func ParseFile(goFilePath string) ([]*AstInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// ParseGoCode parses a go code and returns a list of AstInfo
func ParseGoCode(filename string, data []byte) ([]*AstInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// traverse AST code blocks

//case *ast.BadDecl:
//	code := getBadDeclCode(fset, node, src)
//	println(code)

func getPackageCode(fset *token.FileSet, f *ast.File, src string) (names []string, comment string, body string) {
	_ = "STUB: not implemented"
	return nil, "", ""
}

func getFuncDeclCode(fset *token.FileSet, fn *ast.FuncDecl, src string) (receiverName string, comment string, body string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// the starting position of the func keyword
// end position of function body

func getCodeFromPos(fset *token.FileSet, start, end token.Pos, src string) string {
	_ = "STUB: not implemented"
	return ""
}

func getGenDeclCode(fset *token.FileSet, gen *ast.GenDecl, src string) (names []string, comment string, body string) {
	_ = "STUB: not implemented"
	return nil, "", ""
}

// keyword starting position

// end position of parentheses

// end position of the last Spec

// in the case of keywords only

func getGenName(gen *ast.GenDecl) []string { _ = "STUB: not implemented"; return nil }

// -----------------------------------------------------------------------------------

func adaptPackage(src string) string { _ = "STUB: not implemented"; return "" }

// nolint
func parseBody(body string) (*token.FileSet, *ast.File, string, error) {
	_ = "STUB: not implemented"
	return nil, nil, "", nil
}

type ImportInfo struct {
	Path    string
	Alias   string
	Comment string
	Body    string
}

// ParseImportGroup parse import group from source code
func ParseImportGroup(body string) ([]*ImportInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get path

// get alias

// get comment doc

// get source code of import path

type ConstInfo struct {
	Name    string
	Value   string
	Comment string
	Body    string
}

// ParseConstGroup parse const group from source code
func ParseConstGroup(body string) ([]*ConstInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// get line content

// get code content

// get value (if exists)

type VarInfo struct {
	Name    string
	Value   string
	Comment string
	Body    string
}

// ParseVarGroup parse var group from source code
func ParseVarGroup(body string) ([]*VarInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// get comment

// get code content

// get var value (if exists)

type TypeInfo struct {
	Type    string
	Name    string
	Comment string
	Body    string
	IsIdent bool
}

// ParseTypeGroup parse type group from source code
func ParseTypeGroup(body string) ([]*TypeInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// get comment

// get code content

// get type definition

type InterfaceInfo struct {
	Name        string
	Comment     string
	MethodInfos []*MethodInfo
}

// ParseInterface parse interface group from source code
func ParseInterface(body string) ([]*InterfaceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get interface comment

// get method name

// embedded interface

// get method comment

// get method line content

// MethodInfo method function info
type MethodInfo struct {
	Name         string
	Comment      string
	Body         string
	ReceiverName string
	IsIdent      bool
}

// ParseStructMethods parse struct methods from ast infos
func ParseStructMethods(astInfos []*AstInfo) map[string][]*MethodInfo {
	_ = "STUB: not implemented"
	return nil
}

// map[structName][]*MethodInfo

type StructInfo struct {
	Name    string
	Comment string
	Fields  []*StructFieldInfo
}

type StructFieldInfo struct {
	Name    string
	Type    string
	Comment string
	Body    string
}

// ParseStruct parse struct info from source code
func ParseStruct(body string) (map[string]*StructInfo, error) {
	_ = "STUB: not implemented" //nolint
	return nil, nil
}

// get struct comment

// 处理嵌入字段

// get field name

// get comment

// get source code of field

func getTypeString(expr ast.Expr) string { _ = "STUB: not implemented"; return "" }

func getSrcContent(srcLines []string, start, end int) string { _ = "STUB: not implemented"; return "" }
