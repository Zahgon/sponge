// Package certfile is used to locate the certificate file.
package certfile

import (
	"path/filepath"
	"runtime"
)

var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0) //nolint
	basepath = filepath.Dir(currentFile)
}

// Path return absolute path
func Path(rel string) string { _ = "STUB: not implemented"; return "" }
