// Package parser is parsed proto file to struct
package parser

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// PbService service fields
type PbService struct {
	Name      string           // Greeter
	LowerName string           // greeter first character to lower
	Methods   []*ServiceMethod // service methods

	CutServiceName      string // GreeterService --> Greeter
	LowerCutServiceName string // GreeterService --> greeter

	ImportPkgMap      map[string]string // e.g. `userV1`:`userV1 "user/api/user/v1"`
	FieldImportPkgMap map[string]string // e.g. `userV1`:`userV1 "user/api/user/v1"`

	ProtoFileDir string // e.g. api/user/v1
	ProtoPkgName string // e.g. userV1
}

// ServiceMethod RPCMethod fields
type ServiceMethod struct {
	MethodName    string // Create
	Request       string // CreateRequest
	RequestFields []*Field
	Reply         string // CreateReply
	ReplyFields   []*Field
	Comment       string // e.g. Create a record
	InvokeType    int    // 0:unary, 1: client-side streaming, 2: server-side streaming, 3: bidirectional streaming

	ServiceName         string // Greeter
	LowerServiceName    string // greeter first character to lower
	LowerCutServiceName string // GreeterService --> greeter

	// http_rule
	Path   string // rule
	Method string // HTTP Method
	Body   string

	IsPassGinContext   bool
	IsIgnoreShouldBind bool

	RequestImportPkgName string // e.g. userV1
	ReplyImportPkgName   string // e.g. userV1
	ProtoPkgName         string // e.g. userV1
}

// Field message field
type Field struct {
	Name           string // field name
	GoType         string // field go type
	GoTypeCrossPkg string // field go type cross package
	Comment        string // field comment
	FieldType      string // field type
	IsList         bool   // is list
	IsMap          bool   // is map
	ImportPkgName  string // e.g. anypb
	ImportPkgPath  string // import path e.g. google.golang.org/protobuf/types/known/anypb
}

// GoTypeZero default zero value for type
func (r Field) GoTypeZero() string { _ = "STUB: not implemented"; return "" }

//nolint

// AddOne counter
func (t *ServiceMethod) AddOne(i int) int {
	_ = "STUB: not implemented"

	// RandNumber rand number 1~100
	return 0
}

func (s *PbService) RandNumber() int { _ = "STUB: not implemented"; return 0 }

func parsePbService(s *protogen.Service, protoFileDir string) *PbService {
	_ = "STUB: not implemented"
	return nil
}

//nolint

// GetServices parse protobuf services
func GetServices(file *protogen.File) []*PbService { _ = "STUB: not implemented"; return nil }

func getCutServiceName(name string) string { _ = "STUB: not implemented"; return "" }

type fieldPkgInfo struct {
	fieldType      string
	importPkgPath  string
	importPkgName  string
	goType         string
	goTypeCrossPkg string
}

func newFieldPkgInfo(ident protogen.GoIdent, importPkgMap map[string]string) *fieldPkgInfo {
	_ = "STUB: not implemented"
	return nil
}

func getFields(m *protogen.Message, fieldImportPkgMap map[string]string) []*Field {
	_ = "STUB: not implemented"
	return nil
}

// map value is message

// map value is not message

// field is message

// field is list of message

// field is not message

// field is list of not message

// toGoType convert protobuf type to go type, not support message type
func toGoType(protoKind protoreflect.Kind) string { _ = "STUB: not implemented"; return "" }

func getMethodComment(m *protogen.Method) string { _ = "STUB: not implemented"; return "" }

func getFieldComment(commentSet protogen.CommentSet) string { _ = "STUB: not implemented"; return "" }

func getFieldCommentStr(comment string) string { _ = "STUB: not implemented"; return "" }

func getInvokeType(isStreamingClient bool, isStreamingServer bool) int {
	_ = "STUB: not implemented"
	return 0
}

// bidirectional streaming

// client-side streaming

// server-side streaming

// unary

func GetProtoFileDir(protoPath string) string { _ = "STUB: not implemented"; return "" }

func GetProtoPkgName(importPath string) string { _ = "STUB: not implemented"; return "" }

func convertToPkgName(importPath string) string { _ = "STUB: not implemented"; return "" }

func isVersionNum(pkgName string) bool { _ = "STUB: not implemented"; return false }

func removeMiddleLine(str string) string { _ = "STUB: not implemented"; return "" }
