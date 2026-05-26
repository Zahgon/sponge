package patch

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/replacer"
)

// GenTypesPbCommand generate types.proto code
func GenTypesPbCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// go.mod module name
// output directory

func runTypesPbCommand(moduleName string, outPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// setting up template information
// only the specified subdirectory is processed, if empty or no subdirectory is specified, it means all files
// specify the directory in the subdirectory where processing is ignored
// specify the files in the subdirectory to be ignored for processing

func addTypePbFields(moduleName string) []replacer.Field { _ = "STUB: not implemented"; return nil }
