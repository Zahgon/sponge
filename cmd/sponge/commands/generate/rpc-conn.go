package generate

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/replacer"
)

// GRPCConnectionCommand generate grpc connection code
func GRPCConnectionCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// module name for go.mod
// output directory
// grpc service names

// server name
// whether the generated code is suitable for mono-repo

type grpcConnectionGenerator struct {
	moduleName string
	grpcName   string
	outPath    string

	serverName     string
	suitedMonoRepo bool
}

func (g *grpcConnectionGenerator) generateCode() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// specify the subdirectory and files

func (g *grpcConnectionGenerator) addFields() []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}
