package assistant

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/aicli"
	"github.com/go-dev-frame/sponge/pkg/goast"
)

// GenerateCommand  command
func GenerateCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// for test only

// specified Go files

//nolint

type assistantGenerator struct {
	maxAssistantNum int
	asst            *assistantParams

	fileCodeMap map[string]*codeInfo

	dir   string
	files []string

	isOnlyPrintPrompt bool
}

func (g *assistantGenerator) generateCode() error { _ = "STUB: not implemented"; return nil }

// initialize worker pool

// submit tasks to worker pool

// wait for all tasks to complete

// handle results from worker pool

// stop worker pool

type assistantTask struct {
	jobID         int
	file          string
	dependentFile string
	funcNames     []string
	prompt        string
	client        aicli.Assistanter
	Type          string

	isOnlyPrintPrompt bool
}

// Reply execute assistant task result
type Reply struct {
	JobID     int      `json:"jobID"`
	ErrMsg    string   `json:"errMsg"`
	SrcFile   string   `json:"srcFile"`
	Functions []string `json:"functions"`
	Prompt    string   `json:"prompt"`

	Contents map[string]string `json:"contents"` // reply contents
}

// Execute execute assistant task
func (t *assistantTask) Execute(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFiles(dir string, specifiedFiles []string) (bool, map[string]*codeInfo, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

//nolint

func saveAssistantCode(file string, code string, asstType string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type codeInfo struct {
	funcInfos []goast.FuncInfo
	code      []byte
}

func extractFuncCodeBlock(file string) *codeInfo { _ = "STUB: not implemented"; return nil }

func isAllHaveExampleCode(info *codeInfo) bool { _ = "STUB: not implemented"; return false }

type promptParams struct {
	TargetFilePath    string
	FunctionNamesList string
	TargetFileCode    string

	TargetDirName          string
	DaoFilePath            string
	DaoStructName          string
	ExistingDaoMethodsList string
	DaoInterfaceName       string
	DaoFileCode            string
}

func newDefaultPromptParams(file string, info *codeInfo) *promptParams {
	_ = "STUB: not implemented"
	return nil
}

func getDefaultPrompt(file string, info *codeInfo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// nolint
func getPrompt(file string, info *codeInfo) (dependentFileFullPath string, prompt string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// end with filepath.Separator

// check is need to depend on dao

func (c *codeInfo) isUseChinesePrompt() bool { _ = "STUB: not implemented"; return false }

func (c *codeInfo) getFuncNames() []string { _ = "STUB: not implemented"; return nil }

func containsChinese(s string) bool { _ = "STUB: not implemented"; return false }

func getLastDirName(path string) string { _ = "STUB: not implemented"; return "" }

func isMongoOrmType(dbFile string) bool { _ = "STUB: not implemented"; return false }

func checkDirAndFile(dir string, files []string) error { _ = "STUB: not implemented"; return nil }
