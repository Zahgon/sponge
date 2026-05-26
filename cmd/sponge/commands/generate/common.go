// Package generate is to generate code, including model, cache, dao, handler, http, service, grpc, grpc-gw, grpc-cli code.
package generate

import (
	"fmt"

	"github.com/go-dev-frame/sponge/pkg/replacer"
	"github.com/go-dev-frame/sponge/pkg/sql2code/parser"
)

const (
	defaultGoModVersion      = "go 1.23.0"
	defaultImageGoModVersion = "golang:1.23-alpine"

	// TplNameSponge name of the template
	TplNameSponge = "sponge"

	// DBDriverMysql mysql driver
	DBDriverMysql = "mysql"
	// DBDriverPostgresql postgresql driver
	DBDriverPostgresql = "postgresql"
	// DBDriverTidb tidb driver
	DBDriverTidb = "tidb"
	// DBDriverSqlite sqlite driver
	DBDriverSqlite = "sqlite"
	// DBDriverMongodb mongodb driver
	DBDriverMongodb = "mongodb"

	// code name
	codeNameHTTP        = "http"
	codeNameGRPC        = "grpc"
	codeNameHTTPPb      = "http-pb"
	codeNameGRPCPb      = "grpc-pb"
	codeNameGRPCGW      = "grpc-gw-pb"
	codeNameGRPCHTTP    = "grpc-http"
	codeNameGRPCHTTPPb  = "grpc-http-pb"
	codeNameHandler     = "handler"
	codeNameHandlerPb   = "handler-pb"
	codeNameService     = "service"
	codeNameServiceHTTP = "service-handler"
	codeNameDao         = "dao"
	codeNameProtobuf    = "protobuf"
	codeNameModel       = "model"
	codeNameGRPCConn    = "grpc-conn"
	codeNameCache       = "cache"

	wellPrefix    = "## "
	mgoSuffix     = ".mgo"
	pkgPathSuffix = "/pkg"
	expSuffix     = ".exp"
	tplSuffix     = ".tpl"
	apiDocsSuffix = " api docs"
)

var (
	undeterminedDBDriver = "undetermined" // used in services created based on protobuf.

	modelFile     = "model/userExample.go"
	modelFileMark = "// todo generate model code to here"

	databaseInitDBFile     = "database/init.go"
	databaseInitDBFileMark = "// todo generate initialisation database code here"

	showDbNameMark  = "// todo show db driver name here"
	CurrentDbDriver = func(dbDriver string) string { return "// db driver is " + dbDriver }

	cacheFile = "cache/cacheNameExample.go"

	daoFile     = "dao/userExample.go"
	daoMgoFile  = "dao/userExample.go.mgo"
	daoFileMark = "// todo generate the update fields code to here"
	daoTestFile = "dao/userExample_test.go"

	typesFile         = "types/userExample_types.go"
	typesMgoFile      = "types/userExample_types.go.mgo"
	handlerFileMark   = "// todo generate the request and response struct to here"
	handlerTestFile   = "handler/userExample_test.go"
	handlerPbFile     = "handler/userExample_logic.go"
	handlerPbTestFile = "handler/userExample_logic_test.go"

	handlerLogicFile = "handler/userExample_logic.go"
	serviceLogicFile = "service/userExample.go"
	embedTimeMark    = "// todo generate the conversion createdAt and updatedAt code here"

	httpFile = "server/http.go"

	protoFile     = "v1/userExample.proto"
	protoFileMark = "// todo generate the protobuf code here"

	serviceTestFile      = "service/userExample_test.go"
	serviceClientFile    = "service/userExample_client_test.go"
	serviceClientMgoFile = "service/userExample_client_test.go.mgo"
	serviceFile          = "service/userExample.go"
	serviceFileMark      = "// todo generate the service struct code here"

	dockerFile     = "scripts/build/Dockerfile"
	dockerFileMark = "# todo generate dockerfile code for http or grpc here"

	dockerFileBuild     = "scripts/build/Dockerfile_build"
	dockerFileBuildMark = "# todo generate dockerfile_build code for http or grpc here"

	imageBuildFile     = "scripts/image-build.sh"
	imageBuildFileMark = "# todo generate image-build code for http or grpc here"

	imageBuildLocalFile     = "scripts/image-build-local.sh"
	imageBuildLocalFileMark = "# todo generate image-build-local code for http or grpc here"

	dockerComposeFile     = "deployments/docker-compose/docker-compose.yml"
	dockerComposeFileMark = "# todo generate docker-compose.yml code for http or grpc here"

	k8sDeploymentFile     = "deployments/kubernetes/serverNameExample-deployment.yml"
	k8sDeploymentFileMark = "# todo generate k8s-deployment.yml code for http or grpc here"

	k8sServiceFile     = "deployments/kubernetes/serverNameExample-svc.yml"
	k8sServiceFileMark = "# todo generate k8s-svc.yml code for http or grpc here"

	protoShellFile         = "scripts/protoc.sh"
	protoShellFileGRPCMark = "# todo generate grpc files here"
	protoShellFileMark     = "# todo generate api template code command here"

	appConfigFile      = "configs/serverNameExample.yml"
	appConfigFileMark  = "# todo generate http or rpc server configuration here"
	appConfigFileMark2 = "# todo generate the database configuration here"
	appConfigFileMark3 = "# todo generate the registry and discovery configuration here"

	expectedSQLForDeletion = "expectedSQLForDeletion := \"UPDATE .*\""

	//deploymentConfigFile     = "kubernetes/serverNameExample-configmap.yml"
	//deploymentConfigFileMark = "# todo generate the database configuration for deployment here"

	spongeTemplateVersionMark = "// todo generate the local sponge template code version here"

	configmapFileMark = "# todo generate server configuration code here"

	readmeFile    = "sponge/README.md"
	makeFile      = "sponge/Makefile"
	gitIgnoreFile = "sponge/.gitignore"

	startMarkStr  = "// delete the templates code start"
	endMarkStr    = "// delete the templates code end"
	startMark     = []byte(startMarkStr)
	endMark       = []byte(endMarkStr)
	wellStartMark = symbolConvert(startMarkStr)
	wellEndMark   = symbolConvert(endMarkStr)

	// embed FS template file when using
	selfPackageName = "github.com/go-dev-frame/sponge"
)

var (
	ModelInitDBFile     = databaseInitDBFile
	ModelInitDBFileMark = databaseInitDBFileMark
	//AppConfigFileDBMark = appConfigFileMark2
	StartMark = startMark
	EndMark   = endMark
)

func symbolConvert(str string, additionalChar ...string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func convertServerName(serverName string) string { _ = "STUB: not implemented"; return "" }

func convertProjectAndServerName(projectName, serverName string) (pn string, sn string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func adjustmentOfIDType(handlerCodes string, dbDriver string, isCommonStyle bool) string {
	_ = "STUB: not implemented"
	return ""
}

func idTypeFixToUint64(handlerCodes string) string { _ = "STUB: not implemented"; return "" }

func idTypeToUint64(handlerCodes string) string { _ = "STUB: not implemented"; return "" }

func idTypeToStr(handlerCodes string) string { _ = "STUB: not implemented"; return "" }

func deleteFieldsMark(r replacer.Replacer, filename string, startMark []byte, endMark []byte) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

//fmt.Printf("readFile error: %v\n", err)

// clear marked template code

// DeleteCodeMark delete code mark fragment
func DeleteCodeMark(r replacer.Replacer, filename string, startMark []byte, endMark []byte) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

func deleteAllFieldsMark(r replacer.Replacer, filename string, startMark []byte, endMark []byte) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

//fmt.Printf("readFile error: %v\n", err)

// clear marked template code

func replaceFileContentMark(r replacer.Replacer, filename string, newContent string) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

// resolving mirror repository host and name
func parseImageRepoAddr(addr string) (host string, name string) {
	_ = "STUB: not implemented"
	return "", ""
}

// default docker hub official repo address

// unofficial repo address

// ------------------------------------------------------------------------------------------

func parseProtobufFiles(protobufFile string) ([]string, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// ParseFuzzyProtobufFiles parse fuzzy protobuf files
func ParseFuzzyProtobufFiles(protobufFile string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// save the moduleName and serverName to the specified file for external use
func saveGenInfo(moduleName string, serverName string, suitedMonoRepo bool, outputDir string) error {
	_ = "STUB: not implemented"
	return nil
}

func saveEmptySwaggerJSON(outputDir string) error { _ = "STUB: not implemented"; return nil }

// get moduleName and serverName from directory
func getNamesFromOutDir(dir string) (moduleName string, serverName string, suitedMonoRepo bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func saveProtobufFiles(moduleName string, serverName string, suitedMonoRepo bool, outputDir string, protobufFiles []string) error {
	_ = "STUB: not implemented"
	return nil
}

func isExistServiceName(data []byte) bool { _ = "STUB: not implemented"; return false }

func isDependImport(protoData []byte, pkgName string) bool { _ = "STUB: not implemented"; return false }

func replacePackage(data []byte, moduleName string, serverName string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func getDBConfigCode(dbDriver string) string { _ = "STUB: not implemented"; return "" }

// GetDBConfigurationCode get db config code
func GetDBConfigurationCode(dbDriver string) string { _ = "STUB: not implemented"; return "" }

func getInitDBCode(dbDriver string) string { _ = "STUB: not implemented"; return "" }

// do nothing

// GetInitDataBaseCode get init db code
func GetInitDataBaseCode(dbDriver string) string { _ = "STUB: not implemented"; return "" }

func getLocalSpongeTemplateVersion() string { _ = "STUB: not implemented"; return "" }

func getEmbedTimeCode(isEmbed bool) string { _ = "STUB: not implemented"; return "" }

func getExpectedSQLForDeletion(isEmbed bool) string { _ = "STUB: not implemented"; return "" }

func getExpectedSQLForDeletionField(isEmbed bool) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

func convertYamlConfig(configFile string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

//nolint

func generateConfigmap(serverName string, outPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func sqliteDSNAdaptation(dbDriver string, dsn string) string { _ = "STUB: not implemented"; return "" }

func removeElements(slice []string, elements ...string) []string {
	_ = "STUB: not implemented"
	return nil
}

func moveProtoFileToAPIDir(moduleName string, serverName string, suitedMonoRepo bool, outputDir string) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	// for protoc.sh and protoc-doc.sh
	monoRepoAPIPath = `bash scripts/patch-mono.sh
cd ..

protoBasePath="api"`

	// for patch-mono.sh
	monoRepoHTTPPatch = `bash scripts/patch-mono.sh

HOST_ADDR=$1`

	// for patch.sh
	typePbShellCode = `
    if [ ! -d "../api/types" ]; then
        sponge patch gen-types-pb --out=./
        checkResult $?
        mv -f api/types ../api
        rmdir api
    fi`

	dupCodeMark = "--dir=internal/ecode"

	adaptDupCode = func(serverType string, serverName string) string {
		if serverType == codeNameHTTP {
			return dupCodeMark
		}
		return fmt.Sprintf("--dir=%s/internal/ecode", serverName)
	}
)

func serverCodeFields(serverType string, moduleName string, serverName string) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

// SubServerCodeFields sub server code fields
func SubServerCodeFields(moduleName string, serverName string) []replacer.Field {
	_ = "STUB: not implemented"
	return nil
}

func changeOutPath(outPath string, serverName string) string { _ = "STUB: not implemented"; return "" }

func getSubFiles(selectFiles map[string][]string, replaceFiles map[string][]string) []string {
	_ = "STUB: not implemented"
	return nil
}

type Version struct {
	major     string
	minor     string
	patch     string
	goVersion string
}

func getLocalGoVersion() string { _ = "STUB: not implemented"; return "" }

//  descending sort by major, minor, patch

func extractImageGoVersion() string { _ = "STUB: not implemented"; return "" }

func dbDriverErr(driver string) error { _ = "STUB: not implemented"; return nil }

func flagTip(name ...string) string { _ = "STUB: not implemented"; return "" }

func cutPath(srcFilePath string) string { _ = "STUB: not implemented"; return "" }

func getReadmeContent(moduleName, serverName, serverType, dbDriver string, suitedMonoRepo bool) string {
	_ = "STUB: not implemented"
	return ""
}

// GetGoModFields get go mod fields
func GetGoModFields(moduleName string) []replacer.Field { _ = "STUB: not implemented"; return nil }

func adaptPgDsn(dsn string) string { _ = "STUB: not implemented"; return "" }

func unmarshalCrudInfo(str string) (*parser.CrudInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getTemplateFiles(files map[string][]string) []string { _ = "STUB: not implemented"; return nil }

func replaceFilesContent(r replacer.Replacer, files []string, crudInfo *parser.CrudInfo) ([]replacer.Field, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func replaceTemplateFileContent(r replacer.Replacer, file string, crudInfo *parser.CrudInfo) (field replacer.Field, err error) {
	_ = "STUB: not implemented"
	return *new(replacer.Field), nil
}

// nolint
func SetSelectFiles(dbDriver string, selectFiles map[string][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func getHTTPServiceFields() []replacer.Field { _ = "STUB: not implemented"; return nil }

func getGRPCServiceFields() []replacer.Field { _ = "STUB: not implemented"; return nil }
