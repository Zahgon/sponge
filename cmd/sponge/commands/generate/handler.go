package generate

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/replacer"
)

// HandlerCommand generate handler code
func HandlerCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// module name for go.mod
// output directory
// table names

// server name
// whether the generated code is suitable for mono-repo

//_ = cmd.MarkFlagRequired("module-name")

//nolint

type handlerGenerator struct {
	moduleName     string
	dbDriver       string
	codes          map[string]string
	outPath        string
	serverName     string
	isEmbed        bool
	isExtendedAPI  bool
	suitedMonoRepo bool

	fields        []replacer.Field
	isCommonStyle bool
}

func (g *handlerGenerator) generateCode() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// specify the subdirectory and files

func (g *handlerGenerator) addFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

// replace the contents of the model/userExample.go file

// replace the contents of the dao/userExample.go file

// replace the contents of the handler/userExample.go file

func handlerExtendedAPI(r replacer.Replacer, codeName string) (map[string][]string, []replacer.Field) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handlerMongoDBExtendedAPI(r replacer.Replacer, codeName string) (map[string][]string, []replacer.Field) {
	_ = "STUB: not implemented"
	return nil, nil
}

func commonHandlerFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

func commonHandlerExtendedFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}
