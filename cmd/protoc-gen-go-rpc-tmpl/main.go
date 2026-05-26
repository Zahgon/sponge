// Package main is to generate *.go(tmpl), *_client_test.go, *_rpc.go files.
package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	helpInfo = `
# generate *.go file
protoc --proto_path=. --proto_path=./third_party --go-rpc-tmpl_out=. --go-rpc-tmpl_opt=paths=source_relative \
  --go-rpc-tmpl_opt=moduleName=yourModuleName --go-rpc-tmpl_opt=serverName=yourServerName *.proto

# if you want the generated code to suited to mono-repo, you need to set the parameter --go-gin_opt=suitedMonoRepo=true

Tip:
    If you want to merge the code, after generating the code, execute the command "sponge merge rpc-pb",
    you don't worry about it affecting the logic code you have already written, in case of accidents,
    you can find the pre-merge code in the directory /tmp/sponge_merge_backup_code.
`

	optErrFormat = `--go-rpc-tmpl_opt error, '%s' cannot be empty.

Usage example: 
    protoc --proto_path=. --proto_path=./third_party \
      --go-rpc-tmpl_out=. --go-rpc-tmpl_opt=paths=source_relative \
      --go-rpc-tmpl_opt=moduleName=yourModuleName --go-rpc-tmpl_opt=serverName=yourServerName \
      *.proto
`
)

func main() {
	var h bool
	flag.BoolVar(&h, "h", false, "help information")
	flag.Parse()
	if h {
		fmt.Printf("%s", helpInfo)
		return
	}

	var flags flag.FlagSet

	var moduleName, serverName, tmplDir, ecodeOut string
	var suitedMonoRepo bool
	flags.StringVar(&moduleName, "moduleName", "", "module name")
	flags.StringVar(&serverName, "serverName", "", "server name")
	flags.StringVar(&tmplDir, "tmplDir", "", "rpc template file directory, the default value is internal/service")
	flags.StringVar(&ecodeOut, "ecodeOut", "", "rpc error code file directory, the default value is internal/ecode")
	flags.BoolVar(&suitedMonoRepo, "suitedMonoRepo", false, "whether the generated code is suitable for mono-repo")

	options := protogen.Options{
		ParamFunc: flags.Set,
	}

	options.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			dirName := "internal"
			if suitedMonoRepo {
				dirName = serverName + "/internal"
			}
			if tmplDir == "" {
				tmplDir = dirName + "/service"
			}
			if ecodeOut == "" {
				ecodeOut = dirName + "/ecode"
			}

			err := saveRPCTmplFiles(f, moduleName, serverName, tmplDir, ecodeOut, suitedMonoRepo)
			if err != nil {
				continue // skip error, process the next protobuf file
			}
		}
		return nil
	})
}

func saveRPCTmplFiles(f *protogen.File, moduleName string, serverName string, tmplOut string, ecodeOut string, suitedMonoRepo bool) error {
	_ = "STUB: not implemented"
	return nil
}

func saveFile(moduleName string, serverName string, out string, filePath string, content []byte, isNeedCovered bool, suitedMonoRepo bool) error {
	_ = "STUB: not implemented"
	return nil
}

func saveFileSimple(out string, filePath string, content []byte, isNeedCovered bool) error {
	_ = "STUB: not implemented"
	return nil
}

func isExists(f string) bool { _ = "STUB: not implemented"; return false }

func removeOldGenFile(file string) { _ = "STUB: not implemented"; return }

func firstLetterToUpper(s string) []byte { _ = "STUB: not implemented"; return nil }

func adaptMonoRepo(moduleName string, serverName string, data []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}
