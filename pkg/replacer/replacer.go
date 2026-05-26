// Package replacer is a library of replacement file content, supports replacement of
// files in local directories and embedded directory files via embed.
package replacer

import (
	"embed"
)

var _ Replacer = (*replacerInfo)(nil)

// Replacer interface
type Replacer interface {
	SetReplacementFields(fields []Field)
	SetSubDirsAndFiles(subDirs []string, subFiles ...string)
	SetIgnoreSubDirs(dirs ...string)
	SetIgnoreSubFiles(filenames ...string)
	SetOutputDir(absDir string, name ...string) error
	GetOutputDir() string
	GetSourcePath() string
	SaveFiles() error
	ReadFile(filename string) ([]byte, error)
	GetFiles() []string
	SaveTemplateFiles(m map[string]interface{}, parentDir ...string) error
}

// replacerInfo replacer information
type replacerInfo struct {
	path              string   // template directory or file
	fs                embed.FS // Template directory corresponding to binary objects
	isActual          bool     // true: use os to manipulate files, false: use fs to manipulate files
	files             []string // list of template files
	ignoreFiles       []string // ignore the list of replaced files, e.g. ignore.txt or myDir/ignore.txt
	ignoreDirs        []string // ignore processed subdirectories
	replacementFields []Field  // characters to be replaced when converting from a template file to a new file
	outPath           string   // the directory where the file is saved after replacement
}

// New create replacer with local directory
func New(path string) (Replacer, error) { _ = "STUB: not implemented"; return *new(Replacer), nil }

// NewFS create replacer with embed.FS
func NewFS(path string, fs embed.FS) (Replacer, error) {
	_ = "STUB: not implemented"
	return *new(Replacer), nil
}

// Field replace field information
type Field struct {
	Old             string // old field
	New             string // new field
	IsCaseSensitive bool   // whether the first letter is case-sensitive
}

// SetReplacementFields set the replacement field, note: old characters should not be included in the relationship,
// if they exist, pay attention to the order of precedence when setting the Field
func (r *replacerInfo) SetReplacementFields(fields []Field) { _ = "STUB: not implemented"; return }

// splitting the initial case field

// convert the first letter to upper case

// convert the first letter to lower case

// GetFiles get files
func (r *replacerInfo) GetFiles() []string {
	_ = "STUB: not implemented"

	// SetSubDirsAndFiles set up processing of specified subdirectories, files in other directories are ignored
	return nil
}

func (r *replacerInfo) SetSubDirsAndFiles(subDirs []string, subFiles ...string) {
	_ = "STUB: not implemented"
	return
}

// use map to avoid duplicate files

// SetIgnoreSubFiles specify files to be ignored
func (r *replacerInfo) SetIgnoreSubFiles(filenames ...string) { _ = "STUB: not implemented"; return }

// SetIgnoreSubDirs specify subdirectories to be ignored
func (r *replacerInfo) SetIgnoreSubDirs(dirs ...string) { _ = "STUB: not implemented"; return }

// SetOutputDir specify the output directory, preferably using absPath, if absPath is empty,
// the output directory is automatically generated in the current directory according to the name of the parameter
func (r *replacerInfo) SetOutputDir(absPath string, name ...string) error {
	_ = "STUB: not implemented"
	// output to the specified directory
	return nil
}

// output to the current directory

// GetOutputDir get output directory
func (r *replacerInfo) GetOutputDir() string {
	_ = "STUB: not implemented"

	// GetSourcePath get source directory
	return ""
}

func (r *replacerInfo) GetSourcePath() string {
	_ = "STUB: not implemented"

	// ReadFile read file content
	return ""
}

func (r *replacerInfo) ReadFile(filename string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SaveFiles save file with setting
func (r *replacerInfo) SaveFiles() error { _ = "STUB: not implemented"; return nil }

// read from local files

// read from local embed.FS

// replace text content

// get new file path

// replace file names and directory names

//nolint

// SaveTemplateFiles save file with setting
func (r *replacerInfo) SaveTemplateFiles(m map[string]interface{}, parentDir ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func replaceTemplateData(file string, m map[string]interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func replaceTemplateFilePath(file string, m map[string]interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func trimExt(file string) string { _ = "STUB: not implemented"; return "" }

func (r *replacerInfo) isIgnoreFile(file string) bool { _ = "STUB: not implemented"; return false }

func (r *replacerInfo) isInIgnoreDir(file string) bool { _ = "STUB: not implemented"; return false }

func isForbiddenFile(file string, path string) bool { _ = "STUB: not implemented"; return false }

func (r *replacerInfo) getNewFilePath(file string) string {
	_ = "STUB: not implemented"
	// var newFilePath string
	//
	//	if r.isActual {
	//		newFilePath = r.outPath + strings.Replace(file, r.path, "", 1)
	//	} else {
	//
	//		newFilePath = r.outPath + strings.Replace(file, r.path, "", 1)
	//	}
	return ""
}

func (r *replacerInfo) getNewFilePath2(file string, refDir string) string {
	_ = "STUB: not implemented"
	return ""
}

// if windows, convert the path splitter
func (r *replacerInfo) convertPathDelimiter(filePath string) string {
	_ = "STUB: not implemented"
	return ""
}

// if windows, batch convert path splitters
func (r *replacerInfo) convertPathsDelimiter(filePaths ...string) []string {
	_ = "STUB: not implemented"
	return nil
}

func saveToNewFile(filePath string, data []byte) error {
	_ = "STUB: not implemented"
	// create directory
	return nil
}

// save file

// iterates over all files in the embedded directory, returning the absolute path to the file
func listFiles(path string, fs embed.FS) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// iterating through the embedded catalog
func walkDir(dirPath string, allFiles *[]string, fs embed.FS) error {
	_ = "STUB: not implemented"
	return nil
}

// determine if the first character of a string is a letter
func isFirstAlphabet(str string) bool { _ = "STUB: not implemented"; return false }

func isSubPath(filePath string, subPath string) bool { _ = "STUB: not implemented"; return false }

func isMatchFile(filePath string, sf string) bool { _ = "STUB: not implemented"; return false }
