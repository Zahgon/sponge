package template

import (
	"strings"

	"github.com/spf13/cobra"
)

var (
	printCustomContent *strings.Builder
)

// FieldCommand generate code based on custom template and fields
func FieldCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// template directory
// json file
// only print template code and all fields, do not generate code
// output directory

type customGenerator struct {
	tplDir    string
	fields    map[string]interface{}
	onlyPrint bool
	outPath   string
}

func (g *customGenerator) generateCode() (string, error) { _ = "STUB: not implemented"; return "", nil }
