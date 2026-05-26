package generate

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/replacer"
)

// ServiceCommand generate service code
func ServiceCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// module name for go.mod
// server name
// output directory
// table names

// whether the generated code is suitable for mono-repo

//_ = cmd.MarkFlagRequired("module-name")

//_ = cmd.MarkFlagRequired("server-name")

//nolint

type serviceGenerator struct {
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
func (g *serviceGenerator) generateCode() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// setting up template information

/*"userExample_client_test.go",*/

/*"userExample_client_test.go.mgo",*/

func (g *serviceGenerator) addFields(r replacer.Replacer) []replacer.Field {
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

func serviceExtendedAPI(r replacer.Replacer, codeName string) (map[string][]string, []replacer.Field) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*"userExample_client_test.go.exp",*/

/*"userExample_client_test.go.exp"*/

func serviceMongoDBExtendedAPI(r replacer.Replacer, codeName string) (map[string][]string, []replacer.Field) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*"userExample_client_test.go.mgo.exp",*/

/*"userExample_client_test.go.mgo.exp"*/

func commonServiceFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

func commonServiceExtendedFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}
