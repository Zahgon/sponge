package generate

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/replacer"
)

// ModelCommand generate model code
func ModelCommand(parentName string) *cobra.Command { _ = "STUB: not implemented"; return nil }

// output directory
// table names

//nolint

type modelGenerator struct {
	codes   map[string]string
	outPath string
}

func (g *modelGenerator) generateCode() (string, error) { _ = "STUB: not implemented"; return "", nil }

// specify the subdirectory and files

func (g *modelGenerator) addFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

// replace the contents of the model/userExample.go file
