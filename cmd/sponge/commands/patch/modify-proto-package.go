package patch

import (
	"github.com/spf13/cobra"
)

// ModifyProtoPackageCommand modifies the package and go_package names of proto files.
func ModifyProtoPackageCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getPackageName(ss []string, moduleName string) (packageName string, goPackageName string) {
	_ = "STUB: not implemented"
	return "", ""
}

func splitProtoFilePath(protoFilePath string) []string { _ = "STUB: not implemented"; return nil }

func replaceProtoPackages(protoFilePath, packageName, goPackage string) error {
	_ = "STUB: not implemented"
	return nil
}
