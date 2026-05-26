package template

import (
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var (
	printSQLOnce    sync.Once
	printSQLContent *strings.Builder
)

// SQLCommand generate code based on sql and custom template
func SQLCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// template directory
// fields defined in json
// table names
// output directory
// only print template code and all fields

//nolint

type sqlGenerator struct {
	tplDir    string
	fields    map[string]interface{}
	onlyPrint bool
	outPath   string
}

func (g *sqlGenerator) generateCode() (string, error) { _ = "STUB: not implemented"; return "", nil }
