package generate

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/replacer"
)

// HTTPPbCommand generate web server code based on protobuf file
func HTTPPbCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// module name for go.mod
// server name
// project name for deployment name
// image repo address
// output directory
// protobuf file, support * matching

// whether the generated code is suitable for mono-repo

type httpPbGenerator struct {
	moduleName   string
	serverName   string
	projectName  string
	protobufFile string
	repoAddr     string
	outPath      string

	suitedMonoRepo bool
}

func (g *httpPbGenerator) generateCode() (string, error) { _ = "STUB: not implemented"; return "", nil }

// specify the subdirectory and files

// ignore some directories and files

func (g *httpPbGenerator) addFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

//fields = append(fields, deleteFieldsMark(r, deploymentConfigFile, wellStartMark, wellEndMark)...)

// replace the configuration of the *.yml file

// replace the configuration of the *.yml file

//{ // replace the contents of the model/init.go file
//	Old: modelInitDBFileMark,
//	New: getInitDBCode(DBDriverMysql), // default is mysql
//},
// replace the contents of the Dockerfile file

// replace the contents of the Dockerfile_build file

// replace the contents of the image-build.sh file

// replace the contents of the image-build-local.sh file

// replace the contents of the docker-compose.yml file

//{ // replace the contents of the *-configmap.yml file
//	Old: deploymentConfigFileMark,
//	New: getDBConfigCode(DBDriverMysql, true),
//},
// replace the contents of the *-deployment.yml file

// replace the contents of the *-svc.yml file

// replace the contents of the proto.sh file

// replace the contents of the proto.sh file

// replace the sponge version of the go.mod file

// docker image and k8s deployment script replacement

// convert to kebab-case format

// docker image and k8s deployment script replacement
