package generate

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/replacer"
)

// CacheCommand generate cache code
func CacheCommand(parentName string) *cobra.Command { _ = "STUB: not implemented"; return nil }

// module name for go.mod
// output directory
// cache name
// prefix key
// key name
// key type
// value name
// value type

// server name
// whether the generated code is suitable for mono-repo

type stringCacheGenerator struct {
	moduleName string
	cacheName  string
	prefixKey  string
	keyName    string
	keyType    string
	valueName  string
	valueType  string
	outPath    string

	serverName     string
	suitedMonoRepo bool
}

func (g *stringCacheGenerator) generateCode() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// specify the subdirectory and files

func (g *stringCacheGenerator) addFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

// match the case where the value type is a pointer
