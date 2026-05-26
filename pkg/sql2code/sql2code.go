// Package sql2code is a code generation engine that generates CRUD code for model,
// dao, handler, service, protobuf based on sql and supports database types mysql,
// mongodb, postgresql, sqlite3.
package sql2code

import (
	"github.com/go-dev-frame/sponge/pkg/sql2code/parser"
)

// Args generate code arguments
type Args struct {
	SQL string // DDL sql

	DDLFile string // DDL file

	DBDriver   string            // db driver name, such as mysql, mongodb, postgresql, sqlite, default is mysql
	DBDsn      string            // connecting to mysql's dsn, if DBDriver is sqlite, DBDsn is local db file
	DBTable    string            // table name
	fieldTypes map[string]string // field name:type

	Package        string // specify the package name (only valid for model types)
	GormType       bool   // whether to display the gorm type name (only valid for model type codes)
	JSONTag        bool   // does it include a json tag
	JSONNamedType  int    // json field naming type, 0: snake case such as my_field_name, 1: camel sase, such as myFieldName
	IsEmbed        bool   // is gorm.Model embedded
	IsWebProto     bool   // proto file type, true: include router path and swagger info, false: normal proto file without router and swagger
	CodeType       string // specify the different types of code to be generated, namely model (default), json, dao, handler, proto
	ForceTableName bool
	Charset        string
	Collation      string
	TablePrefix    string
	ColumnPrefix   string
	NoNullType     bool
	NullStyle      string
	IsExtendedAPI  bool // true: generate extended api (9 api), false: generate basic api (5 api)

	IsCustomTemplate bool // whether to use custom template, default is false
}

func (a *Args) checkValid() error { _ = "STUB: not implemented"; return nil }

func getSQL(args *Args) (string, map[string]string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func setOptions(args *Args) []parser.Option { _ = "STUB: not implemented"; return nil }

// GenerateOne generate gorm code from sql, which can be obtained from parameters, files and db, with priority from highest to lowest
func GenerateOne(args *Args) (string, error) { _ = "STUB: not implemented"; return "", nil }

// default is model code

// Generate model, json, dao, handler, proto codes
func Generate(args *Args) (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }
