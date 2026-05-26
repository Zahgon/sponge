// Package merge is merge the generated code into the template file, you don't worry about it affecting
// the logic code you have already written, in case of accidents, you can find the
// pre-merge code in the directory /tmp/sponge_merge_backup_code
package merge

const defaultFuzzyFilename = "*.go.gen20*"

// path end with "/" or "\" will be removed
func adaptDir(dir string) string { _ = "STUB: not implemented"; return "" }

func getRelativeFilePath(srcFile string) string { _ = "STUB: not implemented"; return "" }

func getRelativeDirAndFile(srcFile string) (dir string, file string) {
	_ = "STUB: not implemented"
	return "", ""
}

func getBackupDir() string { _ = "STUB: not implemented"; return "" }

func backupFile(file string, backupDir string) { _ = "STUB: not implemented"; return }

func deleteGenFiles(files []string) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------------

type mergeType int

const (
	errCodeType           mergeType = 1
	routersType           mergeType = 2
	handlerType           mergeType = 3
	serviceGRPCTmplType   mergeType = 4
	serviceGRPCClientType mergeType = 5
)

type mergeParams struct {
	serverDir     string    // specify the server directory
	Type          mergeType // type of code to be merged
	genCodeDir    string    // directory where code is generated
	fuzzyFilename string    // fuzzy matching file name
	backupDir     string    // backup Code Catalog
}

func newMergeParams(dir string, dirType mergeType) *mergeParams {
	_ = "STUB: not implemented"
	return nil
}

func (m *mergeParams) runMerge() error { _ = "STUB: not implemented"; return nil }

// "*.go.gen20*"

// delete generated files

func (m *mergeParams) mergeErrCodeFile(groupFiles map[string]string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *mergeParams) mergeRoutersFile(groupFiles map[string]string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *mergeParams) mergeHandlerAndServiceFile(groupFiles map[string]string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *mergeParams) mergeServiceGRPCClientFile(groupFiles map[string]string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
