// Package parser is a library that parses to go structures based on sql
// and generates the code needed based on the template.
package parser

import (
	"text/template"

	"github.com/zhufuyi/sqlparser/ast"
	"github.com/zhufuyi/sqlparser/dependency/types"
)

const (
	// TableName table name
	TableName = "__table_name__"
	// CodeTypeModel model code
	CodeTypeModel = "model"
	// CodeTypeJSON json code
	CodeTypeJSON = "json"
	// CodeTypeDAO update fields code
	CodeTypeDAO = "dao"
	// CodeTypeHandler handler request and respond code
	CodeTypeHandler = "handler"
	// CodeTypeProto proto file code
	CodeTypeProto = "proto"
	// CodeTypeService grpc service code
	CodeTypeService = "service"
	// CodeTypeCrudInfo crud info json data
	CodeTypeCrudInfo = "crud_info"
	// CodeTypeTableInfo table info json data
	CodeTypeTableInfo = "table_info"

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

	jsonTypeName     = "datatypes.JSON"
	jsonPkgPath      = "gorm.io/datatypes"
	boolTypeName     = "sgorm.Bool"
	boolTypeTinyName = "sgorm.TinyBool"
	boolPkgPath      = "github.com/go-dev-frame/sponge/pkg/sgorm"
	decimalTypeName  = "decimal.Decimal"
	decimalPkgPath   = "github.com/shopspring/decimal"

	unknownCustomType = "UnknownCustomType"
)

// Codes content
type Codes struct {
	Model         []string // model code
	UpdateFields  []string // update fields code
	ModelJSON     []string // model json code
	HandlerStruct []string // handler request and respond code
}

// modelCodes model code
type modelCodes struct {
	Package    string
	ImportPath []string
	StructCode []string
}

// ParseSQL generate different usage codes based on sql
func ParseSQL(sql string, options ...Option) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type tmplData struct {
	TableNamePrefix string

	RawTableName    string // raw table name, example: foo_bar
	TableName       string // table name in camel case, example: FooBar
	TName           string // table name first letter in lower case, example: fooBar
	NameFunc        bool
	Fields          []tmplField
	Comment         string
	SubStructs      string // sub structs for model
	ProtoSubStructs string // sub structs for protobuf
	DBDriver        string

	CrudInfo *CrudInfo
}

type tmplField struct {
	IsPrimaryKey bool   // is primary key
	ColName      string // table column name
	Name         string // convert to camel case
	GoType       string // convert to go type
	Tag          string
	Comment      string
	JSONName     string
	DBDriver     string

	rewriterField *rewriterField
}

type rewriterField struct {
	goType string
	path   string
}

func (d tmplData) isCommonStyle(isEmbed bool) bool { _ = "STUB: not implemented"; return false }

// ConditionZero type of condition 0, used in dao template code
func (t tmplField) ConditionZero() string { _ = "STUB: not implemented"; return "" }

//nolint

//nolint
//nolint

//nolint

//nolint

//nolint
//nolint
//nolint

//nolint

//nolint

// GoZero type of 0, used in model to json template code
func (t tmplField) GoZero() string { _ = "STUB: not implemented"; return "" }

//nolint
//nolint
//nolint

// GoTypeZero type of 0, used in service template code, corresponding protobuf type
func (t tmplField) GoTypeZero() string { _ = "STUB: not implemented"; return "" }

//nolint

//nolint
//nolint
//nolint

//nolint

//nolint

// AddOne counter
func (t tmplField) AddOne(i int) int {
	_ = "STUB: not implemented"

	// AddOneWithTag counter and add id tag
	return 0
}

func (t tmplField) AddOneWithTag(i int) string { _ = "STUB: not implemented"; return "" }

func (t tmplField) AddOneWithTag2(i int) string { _ = "STUB: not implemented"; return "" }

func getProtoFieldName(fields []tmplField) string { _ = "STUB: not implemented"; return "" }

const (
	__mysqlModel__ = "__mysqlModel__" //nolint
	__type__       = "__type__"       //nolint
)

var replaceFields = map[string]string{
	__mysqlModel__: "sgorm.Model",
	__type__:       "",
}

const (
	columnID         = "id"
	_columnID        = "_id"
	columnCreatedAt  = "created_at"
	columnUpdatedAt  = "updated_at"
	columnDeletedAt  = "deleted_at"
	columnMysqlModel = __mysqlModel__
)

var ignoreColumns = map[string]struct{}{
	columnID:         {},
	columnCreatedAt:  {},
	columnUpdatedAt:  {},
	columnDeletedAt:  {},
	columnMysqlModel: {},
}

func isIgnoreFields(colName string, falseColumn ...string) bool {
	_ = "STUB: not implemented"
	return false
}

var newlineIdentifier = []struct{ old, new string }{
	{"\r\n", "\n//"},
	{"\n", "\n//"},
	{"\r", "\n//"},
	{"\n////", "\n//"},
}

func replaceCommentNewline(comment string) string { _ = "STUB: not implemented"; return "" }

type codeText struct {
	importPaths   []string
	modelStruct   string
	modelJSON     string
	updateFields  string
	handlerStruct string
	protoFile     string
	serviceStruct string
	crudInfo      string
	tableInfo     []byte
}

// nolint
func makeCode(stmt *ast.CreateTableStmt, opt options) (*codeText, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// find table comment

// TODO: foreign key support

// snake case

// camel case (default)

// make GORM's tag

//gormTag.WriteString(";NULL")

// For Timestamp and Datetime only.

//return "", nil, errors.Errorf(" unsupport option %d\n", o.Tp)

// mongodb

// gorm

// get type in golang

// rewritten type

// nolint
func getModelStructCode(data tmplData, importPaths []string, isEmbed bool, jsonNamedType int) (string, []string, error) {
	_ = "STUB: not implemented"
	// filter to ignore field fields
	return "", nil, nil
}

// filter time package name

//nolint

// force conversion of ID field to uint64 type

// restore the real embedded fields

// snake case
// sgorm.Model2

func getTableColumnsCode(data tmplData, isEmbed bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getModelCode(data modelCodes) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getUpdateFieldsCode(data tmplData, isEmbed bool) (string, error) {
	_ = "STUB: not implemented"

	// filter fields
	return "", nil
}

func getHandlerStructCodes(data tmplData, jsonNamedType int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// mongodb

// snake case

// camel case (default)

// customized filter fields
func tmplExecuteWithFilter(data tmplData, tmpl *template.Template, reservedColumns ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// mongodb

func getModelJSONCode(data tmplData) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getProtoFileCode(data tmplData, jsonNamedType int, isWebProto bool, isExtendedAPI bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

const (
	createTableReplyFieldCodeMark         = "// createTableReplyFieldCode"
	deleteTableByIDRequestFieldCodeMark   = "// deleteTableByIDRequestFieldCode"
	deleteTableByIDsRequestFieldCodeMark  = "// deleteTableByIDsRequestFieldCode"
	getTableByIDRequestFieldCodeMark      = "// getTableByIDRequestFieldCode"
	getTableByIDsRequestFieldCodeMark     = "// getTableByIDsRequestFieldCode"
	listTableByLastIDRequestFieldCodeMark = "// listTableByLastIDRequestFieldCode"
)

var grpcDefaultProtoMessageFieldCodes = map[string]string{
	createTableReplyFieldCodeMark:         "uint64 id = 1;",
	deleteTableByIDRequestFieldCodeMark:   "uint64 id = 1 [(validate.rules).uint64.gt = 0];",
	deleteTableByIDsRequestFieldCodeMark:  "repeated uint64 ids = 1 [(validate.rules).repeated.min_items = 1];",
	getTableByIDRequestFieldCodeMark:      "uint64 id = 1 [(validate.rules).uint64.gt = 0];",
	getTableByIDsRequestFieldCodeMark:     "repeated uint64 ids = 1 [(validate.rules).repeated.min_items = 1];",
	listTableByLastIDRequestFieldCodeMark: "uint64 lastID = 1; // last id",
}

var webDefaultProtoMessageFieldCodes = map[string]string{
	createTableReplyFieldCodeMark:         "uint64 id = 1;",
	deleteTableByIDRequestFieldCodeMark:   `uint64 id =1 [(validate.rules).uint64.gt = 0, (tagger.tags) = "uri:\"id\""];`,
	deleteTableByIDsRequestFieldCodeMark:  "repeated uint64 ids = 1 [(validate.rules).repeated.min_items = 1];",
	getTableByIDRequestFieldCodeMark:      `uint64 id =1 [(validate.rules).uint64.gt = 0, (tagger.tags) = "uri:\"id\"" ];`,
	getTableByIDsRequestFieldCodeMark:     "repeated uint64 ids = 1 [(validate.rules).repeated.min_items = 1];",
	listTableByLastIDRequestFieldCodeMark: `uint64 lastID = 1 [(tagger.tags) = "form:\"lastID\""]; // last id`,
}

var grpcProtoMessageFieldCodes = map[string]string{
	createTableReplyFieldCodeMark:         "string id = 1;",
	deleteTableByIDRequestFieldCodeMark:   "string id = 1 [(validate.rules).string.min_len = 6];",
	deleteTableByIDsRequestFieldCodeMark:  "repeated string ids = 1 [(validate.rules).repeated.min_items = 1];",
	getTableByIDRequestFieldCodeMark:      "string id = 1 [(validate.rules).string.min_len = 6];",
	getTableByIDsRequestFieldCodeMark:     "repeated string ids = 1 [(validate.rules).repeated.min_items = 1];",
	listTableByLastIDRequestFieldCodeMark: "string lastID = 1; // last id",
}

var webProtoMessageFieldCodes = map[string]string{
	createTableReplyFieldCodeMark:         "string id = 1;",
	deleteTableByIDRequestFieldCodeMark:   `string id =1 [(validate.rules).string.min_len = 6, (tagger.tags) = "uri:\"id\""];`,
	deleteTableByIDsRequestFieldCodeMark:  "repeated string ids = 1 [(validate.rules).repeated.min_items = 1];",
	getTableByIDRequestFieldCodeMark:      `string id =1 [(validate.rules).string.min_len = 6, (tagger.tags) = "uri:\"id\"" ];`,
	getTableByIDsRequestFieldCodeMark:     "repeated string ids = 1 [(validate.rules).repeated.min_items = 1];",
	listTableByLastIDRequestFieldCodeMark: `string lastID = 1 [(tagger.tags) = "form:\"lastID\""]; // last id`,
}

func adaptedDbType(data tmplData, isWebProto bool, code string) string {
	_ = "STUB: not implemented"
	return ""
}

// mongodb

func replaceProtoMessageFieldCode(code string, messageFields map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

func getServiceStructCode(data tmplData) (string, error) { _ = "STUB: not implemented"; return "", nil }

func addCommaToJSON(modelJSONCode string) string { _ = "STUB: not implemented"; return "" }

// nolint
func mysqlToGoType(colTp *types.FieldType, style NullStyle) (name string, path string, rrField *rewriterField) {
	_ = "STUB: not implemented"
	return "", "", nil
}

//nolint

// nolint
func goTypeToProto(fields []tmplField, jsonNameType int, isCommonStyle bool) []tmplField {
	_ = "STUB: not implemented"
	return nil
}

// snake case

// camel case (default)

func makeTagStr(tags []string) string { _ = "STUB: not implemented"; return "" }

func getDefaultValue(expr ast.ExprNode) (value string) { _ = "STUB: not implemented"; return "" }
