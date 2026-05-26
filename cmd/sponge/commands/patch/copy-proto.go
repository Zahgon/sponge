// Package patch is command set for patching service code.
package patch

import (
	"github.com/spf13/cobra"
)

var copyCount = 0

// CopyProtoCommand copy proto file from the grpc service directory
func CopyProtoCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// server dir
// proto file version folder
// output directory
// proto file names
// target module

func getModuleAndServerName(dir string) (moduleName string, serverName string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

type protoCopier struct {
	moduleName string
	outPath    string

	srcModuleName    string
	srcServerName    string
	srcDir           string
	srcVersionFolder string

	selectProtoFiles []string
	copiedFiles      map[string]struct{}
}

func (c *protoCopier) copyProtoFiles() error { _ = "STUB: not implemented"; return nil }

// match proto files

func (c *protoCopier) copyDependencyProtoFile(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *protoCopier) copyProtoFile(srcProtoFile string, targetProtoFile string, isDependency bool) error {
	_ = "STUB: not implemented"
	return nil
}

// replace go_package

func (c *protoCopier) replacePackage(data []byte, isDependency bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (c *protoCopier) isCopied(targetProtoFile string) bool {
	_ = "STUB: not implemented"
	return false
}

func backupProtoFiles(outPath string) error { _ = "STUB: not implemented"; return nil }
