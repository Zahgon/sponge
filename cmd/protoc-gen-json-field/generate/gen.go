// Package generate is generate json field code.
package generate

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/go-dev-frame/sponge/cmd/protoc-gen-json-field/parser"
)

// ProtoInfo is the info for parsing proto file
type ProtoInfo struct {
	FileName            string    // proto file name, example: foo_bar.proto
	FileNamePrefix      string    // proto file name, not include suffix, example: foo_bar
	FileNamePrefixCamel string    // proto file name, not include suffix, example: FooBar
	Services            []Service // services in proto file

	FileDir string // proto file directory, example: api/user/v1
	Package string // proto package name, example: api.user.v1

	// go related fields
	GoPkgName         string            // go package name, example: userV1
	GoPackage         string            // go package name, example: "moduleName/api/user/v1"
	ImportPkgMap      map[string]string // rpc params import packages, example: userV1 -> userV1 "moduleName/api/user/v1"
	FieldImportPkgMap map[string]string // message import packages, example: userV1 -> userV1 "moduleName/api/user/v1"
}

type Service struct {
	ServiceName               string // service name, example: foobar or Foobar
	ServiceNameCamel          string // service name camel case, example: FooBar
	ServiceNameCamelFCL       string // service name camel case first character to lower, example: fooBar
	ServiceNamePluralCamel    string // service name plural, camel case, example: FooBars
	ServiceNamePluralCamelFCL string // service name plural, camel case and first character lower, example: fooBars

	GoPkgName string // go package name, example: userV1

	Methods []RPCMethod // rpc methods
}

type RPCMethod struct {
	MethodName string // method name, example: Create
	Comment    string // method comment, example: // Create a record
	InvokeType string // rpc invoke type: unary_call, client_side_streaming, server_side_streaming, bidirectional_streaming

	RequestName          string  // method request message, example: CreateRequest
	RequestFields        []Field // request fields
	RequestImportPkgName string  // request import package name, example: emptypb

	ReplyName          string  // method reply message, example: CreateReply
	ReplyFields        []Field // reply fields
	ReplyImportPkgName string  // reply import package name, example: emptypb

	// google.api.http options fields
	HTTPRouter        string // http router, example: /api/user/v1/create
	HTTPRequestMethod string // http request method, example: POST
	HTTPRequestBody   string // http request body, example: CreateRequest

	// google.api.http selector custom options, only for gin
	IsPassGinContext bool
	IsIgnoreGinBind  bool
}

type Field struct {
	Name           string // field name
	GoType         string // field go type
	GoTypeCrossPkg string // field go type cross package
	Comment        string // field comment
	FieldType      string // field type, if field type is message, it will be used as import package name
	ImportPkgName  string // import package name, example: anypb
	ImportPkgPath  string // import path e.g. google.golang.org/protobuf/types/known/anypb
}

// GenerateFiles generate service logic, router, error code files.
func GenerateFiles(file *protogen.File) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func newProtoInfo(pss []*parser.PbService, goPkgName string) *ProtoInfo {
	_ = "STUB: not implemented"
	return nil
}

func firstLetterToLower(str string) string { _ = "STUB: not implemented"; return "" }

func customEndOfLetterToLower(srcStr string, str string) string {
	_ = "STUB: not implemented"
	return ""
}

func getInvokeType(t int) string { _ = "STUB: not implemented"; return "" }

func getMessageFields(fs []*parser.Field) []Field { _ = "STUB: not implemented"; return nil }

func mergeMap(m1, m2 map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }
