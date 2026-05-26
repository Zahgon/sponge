package gofile

// IsExists determine if a file or folder exists
func IsExists(f string) bool { _ = "STUB: not implemented"; return false }

// GetRunPath get the absolute path of the program execution
func GetRunPath() string { _ = "STUB: not implemented"; return "" }

// GetFilename get file name
func GetFilename(filePath string) string { _ = "STUB: not implemented"; return "" }

// GetFileSuffixName get file suffix name, example: ".txt"
func GetFileSuffixName(filePath string) string { _ = "STUB: not implemented"; return "" }

// GetDir get dir, not include the last separator
func GetDir(filePath string) string { _ = "STUB: not implemented"; return "" }

// GetSuffixDir get suffix dir, not include the last separator
func GetSuffixDir(filePath string) string { _ = "STUB: not implemented"; return "" }

// GetFileDir get dir, include the last separator
func GetFileDir(filePath string) string { _ = "STUB: not implemented"; return "" }

// CreateDir create dir
func CreateDir(dir string) error { _ = "STUB: not implemented"; return nil }

// GetFilenameWithoutSuffix get file name without suffix
func GetFilenameWithoutSuffix(filePath string) string { _ = "STUB: not implemented"; return "" }

// GetRelativeFilePath get relative file path, force to use "/" as separator
func GetRelativeFilePath(absFilePath string) string { _ = "STUB: not implemented"; return "" }

// Join joins any number of path elements into a single path
func Join(elem ...string) string { _ = "STUB: not implemented"; return "" }

// IsWindows determining whether a window environment
func IsWindows() bool { _ = "STUB: not implemented"; return false }

// GetPathDelimiter get separator by system type
func GetPathDelimiter() string { _ = "STUB: not implemented"; return "" }

// ListFiles iterates over all files in the specified directory, returning the absolute path to the file
func ListFiles(dirPath string, opts ...Option) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListDirsAndFiles iterates through all subdirectories of the specified directory, returning the absolute path to the file
func ListDirsAndFiles(dirPath string) (map[string][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FuzzyMatchFiles fuzzy matching of documents, * only
func FuzzyMatchFiles(f string) []string { _ = "STUB: not implemented"; return nil }

// ListDirs list all sub dirs, not including itself
func ListDirs(specifiedDir string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// FilterDirs filter directories that meet the criteria
func FilterDirs(dirs []string, opts Option) []string { _ = "STUB: not implemented"; return nil }

// iterative traversal of documents with filter conditions
func walkDirWithFilter(dirPath string, allFiles *[]string, filter filterFn) error {
	_ = "STUB: not implemented"
	return nil
}

func walkDir2(dirPath string, allDirs *[]string, allFiles *[]string) error {
	_ = "STUB: not implemented"
	return nil
}

type filterFn func(string) bool

// suffix matching
func matchSuffix(suffixName string) filterFn { _ = "STUB: not implemented"; return *new(filterFn) }

// prefix Matching
func matchPrefix(prefixName string) filterFn { _ = "STUB: not implemented"; return *new(filterFn) }

// contains the string
func matchContain(containName string) filterFn { _ = "STUB: not implemented"; return *new(filterFn) }

// traversing the document by iteration
func walkDir(dirPath string, allFiles *[]string) error { _ = "STUB: not implemented"; return nil }

// ListSubDirs list all sub dirs that have the specified sub dir, if sub dir is empty, return all sub dirs
func ListSubDirs(root string, subDir string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hasSubDir(dirPath string, subDir string) bool { _ = "STUB: not implemented"; return false }
