package goast

import (
	"go/ast"
	"go/token"
)

// FuncInfo represents function information
type FuncInfo struct {
	Name    string
	Comment string
}

// ExtractComment extracts function comments in Go code
func (f FuncInfo) ExtractComment() string { _ = "STUB: not implemented"; return "" }

// regular matching `//` or `/* */` comments

// remove the `//` or `/* */` tags

// remove the space at the beginning of the line and split the line

// output the comment string

// containsPanicCall determine if there is a panic("implement me"), or customized flag, e.g. panic("ai to do")
func containsPanicCall(fn *ast.FuncDecl, customFlag ...string) bool {
	_ = "STUB: not implemented"
	return false
}

// stop traversing if you find it.

// stop traversing if you find it.

// interval indicates an area to delete
type interval struct {
	start token.Pos
	end   token.Pos
}

// FilterFuncCodeByFile filters out the code of functions that contain panic("implement me") or customized flag, e.g. panic("ai to do")
func FilterFuncCodeByFile(goFilePath string, customFlag ...string) ([]byte, []FuncInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// FilterFuncCode filters out the code of functions that contain panic("implement me") or customized flag, e.g. panic("ai to do")
func FilterFuncCode(filename string, data []byte, customFlag ...string) ([]byte, []FuncInfo, error) {
	_ = "STUB: not implemented"
	return nil,

		// parse source code for comments
		nil, nil
}

// used to record the code interval corresponding to the function to be deleted (including its Doc comment)

// used to collect function names and comment information that contain panic ("implementation me")

// traverse declarations in the file, keeping only qualified function declarations

// preserve if function name starts with New

// if the function body is nil, it remains

// if matches call panic("implement me") and has function comment, the function is retained

// delete other cases: record deletion interval

// filter comment groups to remove comments that fall within the deletion interval
