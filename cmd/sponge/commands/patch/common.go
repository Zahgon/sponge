package patch

import (
	"go/ast"
)

const (
	httpType = "http"
	httpMark = "errcode.NewError"
	grpcType = "grpc"
	grpcMark = "errcode.NewRPCStatus"
)

// get moduleName and serverName from directory
func getNamesFromOutDir(dir string) (moduleName string, serverName string, suitedMonoRepo bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func cutPath(srcProtoFile string) string { _ = "STUB: not implemented"; return "" }

func cutPathPrefix(srcProtoFile string) string { _ = "STUB: not implemented"; return "" }

func listErrorCodeFiles(dir string) (map[string][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getSubFiles(selectedFiles map[string][]string) []string { _ = "STUB: not implemented"; return nil }

// ------------------------------------------------------------------------------------------

func getServiceName(varNames map[string]struct{}, body string) string {
	_ = "STUB: not implemented"
	return ""
}

func findVarLineContent(varName string, src string) (srcLineCode string, valueStr string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

type Result struct {
	VarName        string
	BaseCodeOffset int
}

// Visitor is a visitor that walks the AST and extracts error code offsets.
type Visitor struct {
	PackageName  string
	CallFuncName string

	stack   []ast.Node
	results []Result
}

func (v *Visitor) parseParts(src string) ([]Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Visit implements the ast.Visitor interface.
func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}

func (v *Visitor) isNewErrorCall(callExpr *ast.CallExpr) bool {
	_ = "STUB: not implemented"
	return false
}

// nolint
func (v *Visitor) processNode(node ast.Node) {
	_ = "STUB: not implemented"
	// handling calls
	return
}

// get variable name
