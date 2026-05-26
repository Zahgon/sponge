package template

import (
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var (
	printProtoOnce    sync.Once
	printProtoContent *strings.Builder
)

// ProtobufCommand generate code based on protobuf and custom template
func ProtobufCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// protobuf file, support * matching
// dependency protobuf files directory

// template directory
// fields defined in json

// output directory
// only print template code and all fields

type protoGenerator struct {
	tplDir    string
	fields    map[string]interface{}
	onlyPrint bool
	outPath   string
}

func (g *protoGenerator) generateCode() (string, error) { _ = "STUB: not implemented"; return "", nil }

func copyThirdPartyProtoFiles(depProtoDir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// out dir is third_party

func convertProtoToJSON(protoFile string, thirdPartyDir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getProtoDataFromJSON(jsonFile string) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deleteFileOrDir(path string) { _ = "STUB: not implemented"; return }
