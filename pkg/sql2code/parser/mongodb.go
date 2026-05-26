package parser

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	goTypeOID            = "primitive.ObjectID"
	goTypeInt            = "int"
	goTypeInt64          = "int64"
	goTypeFloat64        = "float64"
	goTypeString         = "string"
	goTypeTime           = "time.Time"
	goTypeBool           = "bool"
	goTypeNil            = "nil"
	goTypeBytes          = "[]byte"
	goTypeStrings        = "[]string"
	goTypeInts           = "[]int"
	goTypeInterface      = "interface{}"
	goTypeSliceInterface = "[]interface{}"

	// SubStructKey sub struct key
	SubStructKey = "_sub_struct_"
	// ProtoSubStructKey proto sub struct key
	ProtoSubStructKey = "_proto_sub_struct_"

	oidName = "_id"
)

var mgoTypeToGo = map[bsontype.Type]string{
	bson.TypeObjectID:         goTypeOID,
	bson.TypeInt32:            goTypeInt,
	bson.TypeInt64:            goTypeInt64,
	bson.TypeDouble:           goTypeFloat64,
	bson.TypeString:           goTypeString,
	bson.TypeArray:            goTypeSliceInterface,
	bson.TypeEmbeddedDocument: goTypeInterface,
	bson.TypeTimestamp:        goTypeTime,
	bson.TypeDateTime:         goTypeTime,
	bson.TypeBoolean:          goTypeBool,
	bson.TypeNull:             goTypeNil,
	bson.TypeBinary:           goTypeBytes,
	bson.TypeUndefined:        goTypeInterface,
	bson.TypeCodeWithScope:    goTypeString,
	bson.TypeSymbol:           goTypeString,
	bson.TypeRegex:            goTypeString,
	bson.TypeDecimal128:       goTypeInterface,
	bson.TypeDBPointer:        goTypeInterface,
	bson.TypeMinKey:           goTypeInt,
	bson.TypeMaxKey:           goTypeInt,
	bson.TypeJavaScript:       goTypeString,
}

var jsonTagFormat int32 = 1 // 0: snake case, 1: camel case

// SetJSONTagSnakeCase set json tag format to snake case
func SetJSONTagSnakeCase() { _ = "STUB: not implemented"; return }

// SetJSONTagCamelCase set json tag format to camel case
func SetJSONTagCamelCase() { _ = "STUB: not implemented"; return }

// MgoField mongo field
type MgoField struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	Comment        string `json:"comment"`
	ObjectStr      string `json:"objectStr"`
	ProtoObjectStr string `json:"protoObjectStr"`
}

// GetMongodbTableInfo get table info from mongodb
func GetMongodbTableInfo(dsn string, tableName string) ([]*MgoField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getMongodbTableFields(db *mongo.Database, collectionName string) ([]*MgoField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filter deleted_at, used for soft delete

func getTypeFromMgo(name string, element bson.RawElement) (goTypeStr string, goObjectStr string, protoObjectStr string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

//nolint

func parseObject(name string, elements []bson.RawElement) (goTypeStr string, goObjectStr string, protoObjectStr string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

func parseArray(name string, element bson.RawElement) (goTypeStr string, goObjectStr string, protoObjectStr string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

func toLowerFirst(str string) string { _ = "STUB: not implemented"; return "" }

func embedTimeField(names []string, fields []*MgoField) []*MgoField {
	_ = "STUB: not implemented"
	return nil
}

// ConvertToSQLByMgoFields convert to mysql table ddl
func ConvertToSQLByMgoFields(tableName string, fields []*MgoField) (string, map[string]string) {
	_ = "STUB: not implemented"
	return "", nil
}

// name:type

// nolint
func convertMongoToMysqlType(goType string) string { _ = "STUB: not implemented"; return "" }

//nolint

// nolint
func convertToProtoFieldType(name string, goType string) string {
	_ = "STUB: not implemented"
	return ""
}

//nolint

// MgoFieldToGoStruct convert to go struct
func MgoFieldToGoStruct(name string, fs []*MgoField) string { _ = "STUB: not implemented"; return "" }

func toSingular(word string) string { _ = "STUB: not implemented"; return "" }

func nameToSingular(goTypeStr string, targetObjectStr string, markStr string) string {
	_ = "STUB: not implemented"
	return ""
}

func convertToSingular(goTypeStr string, objectStr string, protoObjectStr string) (tStr string, oStr string, pObjStr string) {
	_ = "STUB: not implemented"
	return "", "", ""
}
