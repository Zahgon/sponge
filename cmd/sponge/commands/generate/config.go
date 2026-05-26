package generate

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/jy2struct"
)

// ConfigCommand convert yaml to struct command
func ConfigCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// output directory

// k8s configmap command

func runGenConfigCommand(files map[string]configType, ysArgs jy2struct.Args) error {
	_ = "STUB: not implemented"
	return nil
}

type configType struct {
	configFile     string
	isConfigCenter bool
}

// read all yaml file directories from the config directory, one is .yml and the other is cc.yml
func getYAMLFile(serverDir string) (map[string]configType, error) {
	_ = "STUB: not implemented"
	// generate target file:configuration file
	return nil, nil
}

func saveFile(inputFile string, outputFile string, code string) error {
	_ = "STUB: not implemented"
	return nil
}

func convertToGoFile(ysArgs jy2struct.Args, outPath string) error {
	_ = "STUB: not implemented"
	return nil
}
