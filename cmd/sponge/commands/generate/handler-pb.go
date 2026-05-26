package generate

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/replacer"
)

// HandlerPbCommand generate handler and protobuf code
func HandlerPbCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// module name for go.mod
// server name
// output directory
// table names

// whether the generated code is suitable for mono-repo

//_ = cmd.MarkFlagRequired("module-name")

//_ = cmd.MarkFlagRequired("server-name")

//nolint

type handlerPbGenerator struct {
	moduleName     string
	serverName     string
	dbDriver       string
	isEmbed        bool
	isExtendedAPI  bool
	codes          map[string]string
	outPath        string
	suitedMonoRepo bool

	fields []replacer.Field
}

func (g *handlerPbGenerator) generateCode() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// specify the subdirectory and files

func (g *handlerPbGenerator) addFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

// replace the contents of the model/userExample.go file

// replace the contents of the dao/userExample.go file

// replace the contents of the handler/userExample_logic.go file

// replace the contents of the v1/userExample.proto file

// replace directory name

// Note: protobuf package no "-" signs allowed

func handlerPbExtendedAPI(r replacer.Replacer) (map[string][]string, []replacer.Field) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handlerPbMongoDBExtendedAPI(r replacer.Replacer) (map[string][]string, []replacer.Field) {
	_ = "STUB: not implemented"
	return nil, nil
}

func commonHandlerPbFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

func commonHandlerPbExtendedFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}
