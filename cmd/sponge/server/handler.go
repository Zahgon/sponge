package server

import (
	"archive/zip"
	"fmt"
	"os"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/go-dev-frame/sponge/pkg/errcode"
	"github.com/go-dev-frame/sponge/pkg/sgorm"
)

var (
	recordDirName = "sponge_record"
	saveDir       = fmt.Sprintf("%s/.%s", getSpongeDir(), recordDirName)
)

type dbInfoForm struct {
	Dsn      string `json:"dsn" binding:"required"`
	DbDriver string `json:"dbDriver"`
}

type kv struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// ListDbDrivers list db drivers
func ListDbDrivers(c *gin.Context) { _ = "STUB: not implemented"; return }

// ListLLM list llm info
func ListLLM(c *gin.Context) { _ = "STUB: not implemented"; return }

// ListTables list tables
func ListTables(c *gin.Context) { _ = "STUB: not implemented"; return }

// GenerateCodeForm generate code form
type GenerateCodeForm struct {
	Arg  string `json:"arg" binding:"required"`
	Path string `json:"path" binding:"required"`
}

// GenerateCode generate code
func GenerateCode(c *gin.Context) {
	_ = "STUB: not implemented"
	// Allow getting the value of the request header when crossing domains
	return
}

// GetTemplateInfo get template info
func GetTemplateInfo(c *gin.Context) { _ = "STUB: not implemented"; return }

// nolint
func handleGenerateCode(c *gin.Context, outPath string, arg string) {
	_ = "STUB: not implemented"
	return
}

// nolint

// first line is the command

// HandleAssistant handle assistant generate and merge code
func HandleAssistant(c *gin.Context) { _ = "STUB: not implemented"; return }

// nolint

// first line is the command

var processMap = sync.Map{}

func getCommand(args []string) string { _ = "STUB: not implemented"; return "" }

func addProcess(key string, pid int) { _ = "STUB: not implemented"; return }

func getPid(key string) (int, bool) { _ = "STUB: not implemented"; return 0, false }

func removeProcess(key string) { _ = "STUB: not implemented"; return }

// HandlePerformanceTest handle performance test
func HandlePerformanceTest(c *gin.Context) { _ = "STUB: not implemented"; return }

// nolint

// first line is the command and pid value

// HandleStopPerformanceTest handle stop performance test
func HandleStopPerformanceTest(c *gin.Context) { _ = "STUB: not implemented"; return }

func splitString(str string, sep string) (lineContent string, out string) {
	_ = "STUB: not implemented"
	return "", ""
}

// GetRecord generate run command record
func GetRecord(c *gin.Context) { _ = "STUB: not implemented"; return }

func responseErr(c *gin.Context, err error, ec *errcode.Error) { _ = "STUB: not implemented"; return }

// UploadFiles batch files upload
func UploadFiles(c *gin.Context) { _ = "STUB: not implemented"; return }

//spongeArg, err := getFormValue(form.Value, "spongeArg")
//if err != nil {
//	response.Error(c, errcode.InvalidParams.RewriteMsg("the field 'spongeArg' cannot be empty"))
//	return
//}

//if !checkFileType(fileType) {
//	response.Error(c, errcode.InvalidParams.RewriteMsg("only .proto or yaml files are allowed to be uploaded"))
//	return
//}

//func getFormValue(valueMap map[string][]string, key string) (string, error) {
//	valueSlice := valueMap[key]
//	if len(valueSlice) == 0 {
//		return "", fmt.Errorf("form '%s' is empty", key)
//	}
//
//	return valueSlice[0], nil
//}

//func checkFileType(typeName string) bool {
//	switch typeName {
//	case ".proto", ".yml", ".yaml", "json":
//		return true
//	}
//
//	return false
//}

func checkSameFile(files []string, file string) bool { _ = "STUB: not implemented"; return false }

func getSavePath() string { _ = "STUB: not implemented"; return "" }

// CompressPathToZip compressed directory to zip file
func CompressPathToZip(outPath, targetFile string) error { _ = "STUB: not implemented"; return nil }

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func AdaptToWindowsZip(outPath, targetFile string) error { _ = "STUB: not implemented"; return nil }

// convert to slash path

func getSpongeDir() string { _ = "STUB: not implemented"; return "" }

func getMysqlTables(dsn string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint

func getPostgresqlTables(dsn string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint

type pgSchema struct {
	SchemaName string
}

type pgTable struct {
	TableName string
}

func getSchemas(db *sgorm.DB, dsn string) ([]pgSchema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getSchemaTables(db *sgorm.DB, schemas []pgSchema) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getSqliteTables(dbFile string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint

func getMongodbTables(dsn string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint
