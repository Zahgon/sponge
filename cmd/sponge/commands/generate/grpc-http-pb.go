package generate

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/replacer"
)

// GRPCAndHTTPPbCommand generate grpc+http servers code based on protobuf file
func GRPCAndHTTPPbCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// module name for go.mod
// server name
// project name for deployment name
// image repo address
// output directory
// protobuf file, support * matching

// whether the generated code is suitable for mono-repo

type httpAndGRPCPbGenerator struct {
	moduleName        string
	serverName        string
	projectName       string
	protobufFile      string
	repoAddr          string
	outPath           string
	suitedMonoRepo    bool
	isHandleProtoFile bool

	// grpc+http servers code generation related
	isAddDBInitCode    bool
	dbDriver           string
	extraReplaceFields []replacer.Field
}

func (g *httpAndGRPCPbGenerator) generateCode() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// specify the subdirectory and files

// ignore some directories and files

func (g *httpAndGRPCPbGenerator) addFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

// replace the configuration of the *.yml file

// replace the configuration of the *.yml file

// replace the contents of the Dockerfile file

// replace the contents of the Dockerfile_build file

// replace the contents of the image-build.sh file

// replace the contents of the image-build-local.sh file

// replace the contents of the docker-compose.yml file

// replace the contents of the *-deployment.yml file

// replace the contents of the *-svc.yml file

// replace the contents of the proto.sh file

// replace the contents of the proto.sh file

// replace the sponge version of the go.mod file

// docker image and k8s deployment script replacement

// snake_case to kebab_case

// docker image and k8s deployment script replacement

// force to use grpc type when using mono-repo

func adaptGRPCHTTPType(isAddDBInitCode bool) string { _ = "STUB: not implemented"; return "" }
