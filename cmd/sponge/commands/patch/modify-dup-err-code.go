package patch

import (
	"github.com/spf13/cobra"
)

// ModifyDuplicateErrorCodeOffsetCommand modify duplicate error code offset command
func ModifyDuplicateErrorCodeOffsetCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func checkAndModifyDuplicateErrorCodeOffset(file string) error {
	_ = "STUB: not implemented"
	return nil
}

// ErrorCodeOffsetAst is the struct for error code offset
type ErrorCodeOffsetAst struct {
	FilePath string
	SrcCode  string
	FileSize int

	ErrorCodeNameInfoMap map[string]*ErrorCodeOffset
}

// ErrorCodeOffset is the struct for error code offset
type ErrorCodeOffset struct {
	Name string

	OffsetInfos []*Deconstruction
	offsetIndex int

	VarNameCodeSrcMap map[string]string // varName -> varErrCodeSrc
	Body              string
}

// Deconstruction is the struct for error code offset deconstruction
type Deconstruction struct {
	VarName    string
	VarValue   int
	VarSrcCode string
}

// NewErrorCodeOffsetAst create a ErrorCodeOffsetAst object
func NewErrorCodeOffsetAst(filePath string) (*ErrorCodeOffsetAst, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//varNOName := serviceName + "NO"

//varNOName = "_" + varNOName

// CheckDuplicateErrorCodeOffset check if the error code offset has duplicate value
func (c *ErrorCodeOffset) CheckDuplicateErrorCodeOffset(srcCode string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (c *ErrorCodeOffset) replaceErrorCodeOffset(srcCode string, duplicateNOMap map[int][]string, maxVarValue int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// CheckMergedItems check if the merged items have duplicate error code offset
func (c *ErrorCodeOffset) CheckMergedItems(srcCode string, codeInfo *ErrorCodeOffset) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
