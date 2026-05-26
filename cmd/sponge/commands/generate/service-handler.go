package generate

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/replacer"
)

// ServiceAndHandlerCRUDCommand generate both service and handler CRUD code
func ServiceAndHandlerCRUDCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// module name for go.mod
// server name
// output directory
// table names

// whether the generated code is suitable for mono-repo

//_ = cmd.MarkFlagRequired("module-name")

//_ = cmd.MarkFlagRequired("server-name")

//nolint

type serviceAndHandlerGenerator struct {
	moduleName     string
	serverName     string
	dbDriver       string
	isEmbed        bool
	isExtendedAPI  bool
	codes          map[string]string
	outPath        string
	suitedMonoRepo bool

	fields        []replacer.Field
	isCommonStyle bool
}

// nolint
func (g *serviceAndHandlerGenerator) generateCode() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// specify the subdirectory and files

/*"userExample_client_test.go",*/

/*"userExample_client_test.go.mgo",*/

func (g *serviceAndHandlerGenerator) addFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

// replace the contents of the model/userExample.go file

// replace the contents of the dao/userExample.go file

// replace the contents of the handler/userExample_logic.go file

// replace the contents of the v1/userExample.proto file

// replace the contents of the service/userExample_client_test.go file

// replace directory name

// Note: protobuf package no "-" signs allowed

func serviceHandlerExtendedAPI(r replacer.Replacer) (map[string][]string, []replacer.Field) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*"userExample_client_test.go.exp",*/

func serviceHandlerMongoDBExtendedAPI(r replacer.Replacer) (map[string][]string, []replacer.Field) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*"userExample_client_test.go.mgo.exp",*/

func commonServiceHandlerFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

func commonServiceHandlerExtendedFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}
