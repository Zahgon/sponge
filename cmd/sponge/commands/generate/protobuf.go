package generate

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/replacer"
)

// ProtobufCommand generate protobuf code
func ProtobufCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// module name for go.mod
// server name
// output directory
// table names

//_ = cmd.MarkFlagRequired("module-name")

//_ = cmd.MarkFlagRequired("server-name")

//nolint

type protobufGenerator struct {
	moduleName string
	serverName string
	codes      map[string]string
	outPath    string
}

func (g *protobufGenerator) generateCode() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// specify the subdirectory and files

func (g *protobufGenerator) addFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

// replace the contents of the v1/userExample.proto file

// replace directory name

// Note: protobuf package no "-" signs allowed
