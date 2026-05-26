package generate

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/replacer"
)

// DaoCommand generate dao code
func DaoCommand(parentName string) *cobra.Command { _ = "STUB: not implemented"; return nil }

// go.mod module name
// output directory
// table names

// server name
// whether the generated code is suitable for mono-repo

// control to generate the initialization db code only once

//_ = cmd.MarkFlagRequired("module-name")

//nolint

type daoGenerator struct {
	moduleName      string
	dbDriver        string
	isIncludeInitDB bool
	codes           map[string]string
	outPath         string
	isEmbed         bool
	isExtendedAPI   bool
	serverName      string
	suitedMonoRepo  bool

	fields []replacer.Field
}

func (g *daoGenerator) generateCode() (string, error) { _ = "STUB: not implemented"; return "", nil }

// specify the subdirectory and files

// set fields
func (g *daoGenerator) addFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

// replace the contents of the model/userExample.go file

func daoExtendedAPI(r replacer.Replacer) (map[string][]string, []replacer.Field) {
	_ = "STUB: not implemented"
	return nil, nil
}

func daoMongoDBExtendedAPI(r replacer.Replacer) (map[string][]string, []replacer.Field) {
	_ = "STUB: not implemented"
	return nil, nil
}

func commonDaoFields(r replacer.Replacer) []replacer.Field { _ = "STUB: not implemented"; return nil }

func commonDaoExtendedFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}
