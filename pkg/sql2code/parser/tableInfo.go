package parser

// TableInfo is the struct for extend template
type TableInfo struct {
	TableNamePrefix string // table name prefix, example: t_

	TableName               string // original table name, example: foo_bar
	TableNameCamel          string // camel case, example: FooBar
	TableNameCamelFCL       string // camel case and first character lower, example: fooBar
	TableNamePluralCamel    string // plural, camel case, example: FooBars
	TableNamePluralCamelFCL string // plural, camel case and first character lower, example: fooBars
	TableNameSnake          string // snake case, example: foo_bar

	TableComment string // table comment

	Columns    []Field     // columns of the table
	PrimaryKey *PrimaryKey // primary key information

	DBDriver string // database driver, example: mysql, postgresql, sqlite3, mongodb

	ColumnSubStructure string // column sub structure for model
	ColumnSubMessage   string // sub message for protobuf
}

// Field is the struct for column information
type Field struct {
	ColumnName         string // original column name, example: foo_bar
	ColumnNameCamel    string // first character lower, example: FooBar
	ColumnNameCamelFCL string // first character lower, example: fooBar

	ColumnComment string // column comment
	IsPrimaryKey  bool   // is primary key

	GoType string // convert to go type
	Tag    string // tag for model struct field, default gorm tag
}

// PrimaryKey is the struct for primary key information, it used for generate CRUD code
type PrimaryKey struct {
	Name               string // primary key name, example: foo_bar
	NameCamel          string // primary key name, camel case, example: FooBar
	NameCamelFCL       string // primary key name, camel case and first character lower, example: fooBar
	NamePluralCamel    string // primary key name, plural, camel case, example: FooBars
	NamePluralCamelFCL string // primary key name, plural, camel case and first character lower, example: fooBars

	GoType    string // go type, example:  int, string
	GoTypeFCU string // go type, first character upper, example: Int64, String

	IsStringType bool // go type is string or not
}

func newTableInfo(data tmplData) TableInfo { _ = "STUB: not implemented"; return *new(TableInfo) }

func (table TableInfo) getCode() []byte { _ = "STUB: not implemented"; return nil }

func getColumns(dbDriver string, fields []tmplField) []Field { _ = "STUB: not implemented"; return nil }

func handleTag(dbDriver string, tag string) string { _ = "STUB: not implemented"; return "" }

func getPrimaryKeyInfo(info *CrudInfo) *PrimaryKey { _ = "STUB: not implemented"; return nil }

// UnMarshalTableInfo unmarshal the json data to TableInfo struct
func UnMarshalTableInfo(data string) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
