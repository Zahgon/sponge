package assistant

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/goast"
)

// MergeAssistantCode merge AI assistant generated code into source Go file
func MergeAssistantCode() *cobra.Command { _ = "STUB: not implemented"; return nil }

// specified Go files

type mergeParams struct {
	assistantType  string
	dir            string
	specifiedFiles []string
	isClean        bool
}

func (m *mergeParams) parseAssistantFiles() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// srcFile -> genFile

func (m *mergeParams) getSourceFile(file string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (m *mergeParams) getGoAndMDFile(file string, assistantType string) (goFile string, mdFile string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func (m *mergeParams) mergeGoFile(srcFile string, genFile string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func checkPackageName(file string, data []byte) []goast.CodeAstOption {
	_ = "STUB: not implemented"
	return nil
}
