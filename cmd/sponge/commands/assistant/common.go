package assistant

import (
	"fmt"
	"time"

	"github.com/fatih/color"

	"github.com/go-dev-frame/sponge/pkg/aicli"
	"github.com/go-dev-frame/sponge/pkg/aicli/chatgpt"
	"github.com/go-dev-frame/sponge/pkg/aicli/deepseek"
	"github.com/go-dev-frame/sponge/pkg/aicli/gemini"
	"github.com/go-dev-frame/sponge/pkg/utils"
)

const (
	typeChatGPT  = "chatgpt"
	typeDeepSeek = "deepseek"
	typeGemini   = "gemini"

	gopherRoleDescCN = aicli.GopherRoleDescCN
	gopherRoleDescEN = aicli.GopherRoleDescEN

	successSymbol = "✓"
)

var assistantTypeMap = map[string]string{
	typeChatGPT:  "ChatGPT",
	typeDeepSeek: "DeepSeek",
	typeGemini:   "Gemini",
}

var defaultModelMap = map[string]string{
	typeChatGPT:  chatgpt.DefaultModel,
	typeDeepSeek: deepseek.DefaultModel,
	typeGemini:   gemini.DefaultModel,
}

type assistantParams struct {
	Type string

	apiKey        string
	model         string
	enableContext bool

	// only for chatgpt and deepseek
	roleDesc    string
	maxToken    int
	temperature float32
}

func (a *assistantParams) newClient() (aicli.Assistanter, error) {
	_ = "STUB: not implemented"
	return *new(aicli.Assistanter), nil
}

// --------------------------------------------------------------------------

const (
	gormDao = `package dao

import (
	"context"
	"errors"

	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"github.com/go-dev-frame/sponge/pkg/logger"
	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
	"github.com/go-dev-frame/sponge/pkg/utils"

	"github.com/go-dev-frame/sponge/internal/cache"
	"github.com/go-dev-frame/sponge/internal/database"
	"github.com/go-dev-frame/sponge/internal/model"
)

var _ UserExampleDao = (*userExampleDao)(nil)

// UserExampleDao defining the dao interface
type UserExampleDao interface {
	Create(ctx context.Context, table *model.UserExample) error
	UpdateByID(ctx context.Context, table *model.UserExample) error
	GetByID(ctx context.Context, id uint64) (*model.UserExample, error)

	CreateByTx(ctx context.Context, tx *gorm.DB, table *model.UserExample) (uint64, error)
	DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.UserExample) error
}

type userExampleDao struct {
	db    *gorm.DB
	cache cache.UserExampleCache // if nil, the cache is not used.
	sfg   *singleflight.Group    // if cache is nil, the sfg is not used.
}

// NewUserExampleDao creating the dao interface
func NewUserExampleDao(db *gorm.DB, xCache cache.UserExampleCache) UserExampleDao {
	if xCache == nil {
		return &userExampleDao{db: db}
	}
	return &userExampleDao{
		db:    db,
		cache: xCache,
		sfg:   new(singleflight.Group),
	}
}

`

	mongoDao = `package dao

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/singleflight"

	"github.com/go-dev-frame/sponge/pkg/logger"
	"github.com/go-dev-frame/sponge/pkg/mgo"
	"github.com/go-dev-frame/sponge/pkg/mgo/query"

	"github.com/go-dev-frame/sponge/internal/cache"
	"github.com/go-dev-frame/sponge/internal/database"
	"github.com/go-dev-frame/sponge/internal/model"
)

var _ UserExampleDao = (*userExampleDao)(nil)

// UserExampleDao defining the dao interface
type UserExampleDao interface {
	Create(ctx context.Context, record *model.UserExample) error
	UpdateByID(ctx context.Context, record *model.UserExample) error
	GetByID(ctx context.Context, id string) (*model.UserExample, error)
}

type userExampleDao struct {
	collection *mongo.Collection
	cache      cache.UserExampleCache // if nil, the cache is not used.
	sfg        *singleflight.Group    // if cache is nil, the sfg is not used.
}

// NewUserExampleDao creating the dao interface
func NewUserExampleDao(collection *mongo.Collection, xCache cache.UserExampleCache) UserExampleDao {
	if xCache == nil {
		return &userExampleDao{collection: collection}
	}
	return &userExampleDao{
		collection: collection,
		cache:      xCache,
		sfg:        new(singleflight.Group),
	}
}

`

	codeDelimiterMarker = "/**code-delimiter**/"
)

// nolint
var ErrnoAssistantMarker = fmt.Errorf("\n%s%s\n\n",
	`No Go code requiring AI assistant generation was detected. To trigger the AI assistant, ensure the following conditions are met:
    1. Define a function in Go code.
    2. Add detailed function comments (used as AI prompts).
    3. Add panic("implement me") inside the function body.

Example:`, fmt.Sprintf(`

    %s
    func FunctionName() {
        %s
    }`, color.HiCyanString("// Describe the specific functionality of the function"), color.HiCyanString(`panic("implement me")`)))

// get dao file path
func getDaoFilePath(path string) string { _ = "STUB: not implemented"; return "" }

type daoCodeInfo struct {
	structName    string
	interfaceName string
	methodNames   string
	code          string
}

func parseDaoCode(filePath string) (*daoCodeInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type daoInfo struct {
	code          string
	structName    string
	methodNames   string
	interfaceName string
}

func newDaoInfo(daoFile string, isMongo bool, objName string, isChinese bool) *daoInfo {
	_ = "STUB: not implemented"
	return nil
}

func getDaoDefaultMethodNames(isMongo bool, isChinese bool) string {
	_ = "STUB: not implemented"
	return ""
}

func getModelCode(file string, dirName string, fileName string, isChinese bool) string {
	_ = "STUB: not implemented"
	return ""
}

func capitalize(s string) string { _ = "STUB: not implemented"; return "" }

// extractGoCode extracts the Go code blocks from the given markdown string.
func extractGoCode(markdown string) []string { _ = "STUB: not implemented"; return nil }

// dealing with new ```go

// if it is already in the code block, it means that the previous code block is missing the close identifier and stores it first.

// processing ``` close code block

// record code content

// prevents the last code block from not closing

func parseCode(code string) []string { _ = "STUB: not implemented"; return nil }

func reassembleGoMarkdown(codes []string) []string { _ = "STUB: not implemented"; return nil }

func cutFilePath(fullPath string) string { _ = "STUB: not implemented"; return "" }

func newPrintLog(t ...time.Duration) *utils.WaitPrinter { _ = "STUB: not implemented"; return nil }

func getAssistantSuffixed(assistantType string) string { _ = "STUB: not implemented"; return "" }

func deleteGenFiles(files []string, assistantType string) { _ = "STUB: not implemented"; return }

func getRelativeDirAndFile(srcFile string) (dir string, file string) {
	_ = "STUB: not implemented"
	return "", ""
}

func getBackupDir() string { _ = "STUB: not implemented"; return "" }

func backupFile(file string, backupDir string) { _ = "STUB: not implemented"; return }
