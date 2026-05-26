package patch

import (
	"github.com/spf13/cobra"
)

// ModifyDuplicateErrorCodeNumCommand Command modify duplicate error code numbers
func ModifyDuplicateErrorCodeNumCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// distinguish between _http.go and _rpc.go files, check and correct duplicate error code NO
func checkAndModifyGoFileErrorCodeNO(files []string) error { _ = "STUB: not implemented"; return nil }

// ErrorCodeNOAst is the struct for error code NO
type ErrorCodeNOAst struct {
	FilePath string
	SrcCode  string
	fileSize int

	VarName    string // example: userExampleNO
	VarValue   int    // example: 1
	VarSrcCode string // example: "userExampleNO = 1"
}

// SaveToFile save the modified source code to file
func (e *ErrorCodeNOAst) SaveToFile(maxErrorCodeNO int) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckAndModifyDuplicateErrorCodeNO check and modify duplicate error code NO
func CheckAndModifyDuplicateErrorCodeNO(errorCodeNOs []*ErrorCodeNOAst) error {
	_ = "STUB: not implemented"
	return nil
}

// NewErrorCodeNOAst create a ErrorCodeNOAst object
func NewErrorCodeNOAst(filePath string) ([]*ErrorCodeNOAst, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
