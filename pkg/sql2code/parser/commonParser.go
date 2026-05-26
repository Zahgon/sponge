package parser

import (
	"text/template"
)

// CrudInfo crud info for cache, dao, handler, service, protobuf, error
type CrudInfo struct {
	TableNameCamel          string `json:"tableNameCamel"`          // camel case, example: FooBar
	TableNameCamelFCL       string `json:"tableNameCamelFCL"`       // camel case and first character lower, example: fooBar
	TableNamePluralCamel    string `json:"tableNamePluralCamel"`    // plural, camel case, example: FooBars
	TableNamePluralCamelFCL string `json:"tableNamePluralCamelFCL"` // plural, camel case, example: fooBars

	ColumnName               string `json:"columnName"`               // column name, example: first_name
	ColumnNameCamel          string `json:"columnNameCamel"`          // column name, camel case, example: FirstName
	ColumnNameCamelFCL       string `json:"columnNameCamelFCL"`       // column name, camel case and first character lower, example: firstName
	ColumnNamePluralCamel    string `json:"columnNamePluralCamel"`    // column name, plural, camel case, example: FirstNames
	ColumnNamePluralCamelFCL string `json:"columnNamePluralCamelFCL"` // column name, plural, camel case and first character lower, example: firstNames

	GoType       string `json:"goType"`       // go type, example: string, uint64
	GoTypeFCU    string `json:"goTypeFCU"`    // go type, first character upper, example: String, Uint64
	ProtoType    string `json:"protoType"`    // proto type, example: string, uint64
	IsStringType bool   `json:"isStringType"` // go type is string or not

	PrimaryKeyColumnName string `json:"PrimaryKeyColumnName"` // primary key, example: id
	IsCommonType         bool   `json:"isCommonType"`         // custom primary key name and type
	IsStandardPrimaryKey bool   `json:"isStandardPrimaryKey"` // standard primary key id
}

func isDesiredGoType(t string) bool { _ = "STUB: not implemented"; return false }

//nolint

func setCrudInfo(field tmplField) *CrudInfo { _ = "STUB: not implemented"; return nil }

// if singular and plural are the same, force the suffix 's' to distinguish them

// if singular and plural are the same, force the suffix 's' to distinguish them

func newCrudInfo(data tmplData) *CrudInfo { _ = "STUB: not implemented"; return nil }

// if not found primary key, find the first xxx_id column as primary key

// xxx_id

// if not found xxx_id field, use the first field of integer or string type

// use the first column as primary key

func (info *CrudInfo) getCode() string { _ = "STUB: not implemented"; return "" }

func (info *CrudInfo) CheckCommonType() bool { _ = "STUB: not implemented"; return false }

func (info *CrudInfo) isIDPrimaryKey() bool { _ = "STUB: not implemented"; return false }

func (info *CrudInfo) GetGRPCProtoValidation() string { _ = "STUB: not implemented"; return "" }

func (info *CrudInfo) GetWebProtoValidation() string { _ = "STUB: not implemented"; return "" }

func getCommonHandlerStructCodes(data tmplData, jsonNamedType int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// snake case

// camel case (default)

func getCommonServiceStructCode(data tmplData) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getCommonProtoFileCode(data tmplData, jsonNamedType int, isWebProto bool, isExtendedAPI bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func tmplExecuteWithFilter2(data tmplData, tmpl *template.Template, reservedColumns ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// nolint
func simpleGoTypeToProtoType(goType string) string { _ = "STUB: not implemented"; return "" }

func adaptedDbType2(data tmplData, isWebProto bool, code string) string {
	_ = "STUB: not implemented"
	return ""
}

func firstLetterToUpper(str string) string { _ = "STUB: not implemented"; return "" }

func customFirstLetterToLower(str string) string { _ = "STUB: not implemented"; return "" }

func customEndOfLetterToLower(srcStr string, str string) string {
	_ = "STUB: not implemented"
	return ""
}

func getHandlerGoType(field *tmplField) string { _ = "STUB: not implemented"; return "" }
