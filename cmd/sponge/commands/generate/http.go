package generate

import (
	"github.com/spf13/cobra"

	"github.com/go-dev-frame/sponge/pkg/replacer"
)

// HTTPCommand generate web server code
func HTTPCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// module name for go.mod
// server name
// project name for deployment name
// image repo address
// output directory
// table names

// whether the generated code is suitable for mono-repo

//nolint

//nolint

type httpGenerator struct {
	moduleName     string
	serverName     string
	projectName    string
	repoAddr       string
	dbDSN          string
	dbDriver       string
	codes          map[string]string
	outPath        string
	isEmbed        bool
	isExtendedAPI  bool
	suitedMonoRepo bool

	fields        []replacer.Field
	isCommonStyle bool
}

func (g *httpGenerator) generateCode() (string, error) { _ = "STUB: not implemented"; return "", nil }

// specify the subdirectory and files

// ignore some directories and files

func (g *httpGenerator) addFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

//fields = append(fields, deleteAllFieldsMark(r, makeFile, wellStartMark, wellEndMark)...)

//fields = append(fields, deleteFieldsMark(r, deploymentConfigFile, wellStartMark, wellEndMark)...)

// replace the configuration of the *.yml file

// replace the configuration of the *.yml file

// replace the contents of the model/userExample.go file

// replace the contents of the database/init.go file

// replace the contents of the dao/userExample.go file

// replace the contents of the handler/userExample.go file

// replace the contents of the Dockerfile file

// replace the contents of the Dockerfile_build file

// replace the contents of the image-build.sh file

// replace the contents of the image-build-local.sh file

// replace the contents of the docker-compose.yml file

//{ // replace the contents of the *-configmap.yml file
//	Old: deploymentConfigFileMark,
//	New: getDBConfigCode(g.dbDriver, true),
//},
// replace the contents of the *-deployment.yml file

// replace the contents of the *-svc.yml file

// replace github.com/go-dev-frame/sponge/templates/sponge

// replace the sponge version of the go.mod file

// docker image and k8s deployment script replacement

// snake_case to kebab_case

// docker image and k8s deployment script replacement

func commonHTTPFields(r replacer.Replacer) []replacer.Field { _ = "STUB: not implemented"; return nil }

func commonHTTPExtendedFields(r replacer.Replacer) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}
