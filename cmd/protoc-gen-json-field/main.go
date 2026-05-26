// Package main is a library that parses protobuf files into json
package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const helpInfo = `
Generate json file based on protobuf files.

Usage example:
    protoc --proto_path=. --json-field_out=. --json-field_opt=paths=source_relative demo.proto
`

func main() {
	var h bool
	flag.BoolVar(&h, "h", false, "help information")
	flag.Parse()
	if h {
		fmt.Printf("%s", helpInfo)
		return
	}

	var flags flag.FlagSet

	options := protogen.Options{
		ParamFunc: flags.Set,
	}

	options.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			// save json file
			if err := saveJSONFiles(f); err != nil {
				return err
			}
		}
		return nil
	})
}

func saveJSONFiles(f *protogen.File) error { _ = "STUB: not implemented"; return nil }
