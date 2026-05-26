package patch

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/replacer"
)

// GenerateDBInitCommand generate database initialization code
func GenerateDBInitCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// go.mod module name
// database driver e.g. mysql, mongodb, postgresql, sqlite
// output directory

// check handler and service directory db driver mark

type dbInitGenerator struct {
	moduleName string
	dbDriver   string
	outPath    string

	serverName     string
	suitedMonoRepo bool
}

func (g *dbInitGenerator) generateCode() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (g *dbInitGenerator) addFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

// replace the contents of the model/init.go file

func getContentMark(dbDriver string) []byte { _ = "STUB: not implemented"; return nil }

func checkDbDriver(files []string) string { _ = "STUB: not implemented"; return "" }

func detectDbDriverName() string { _ = "STUB: not implemented"; return "" }
