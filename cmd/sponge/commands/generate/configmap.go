package generate

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/replacer"
)

// ConfigmapCommand generate k8s configmap command
func ConfigmapCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

type copyConfigGenerator struct {
	serverName  string
	projectName string
	content     string
	outPath     string
}

func (g *copyConfigGenerator) generateCode() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// setting up template information
// only the specified subdirectory is processed, if empty or no subdirectory is specified, it means all files

// specify the directory in the subdirectory where processing is ignored
// specify the files in the subdirectory to be ignored for processing

func (g *copyConfigGenerator) addFields() []replacer.Field { _ = "STUB: not implemented"; return nil }

// snake_case to kebab_case
