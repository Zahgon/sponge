// Package parse is parsed proto file to struct
package parse

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// ServiceMethod method fields
type ServiceMethod struct {
	MethodName    string   // e.g. Create
	Request       string   // e.g. CreateRequest
	RequestFields []*Field // request fields
	Reply         string   // e.g. CreateReply
	ReplyFields   []*Field
	Comment       string // e.g. Create a record
	InvokeType    int    // 0:unary, 1: client-side streaming, 2: server-side streaming, 3: bidirectional streaming

	ServiceName      string // Greeter
	LowerServiceName string // greeter first character to lower

	RequestImportPkgName string // e.g. userV1
	ReplyImportPkgName   string // e.g. userV1
	ProtoPkgName         string // e.g. userV1
}

// Field request fields
type Field struct {
	Name      string
	FieldType string
	Comment   string
}

// GoTypeZero default zero value for type
func (r Field) GoTypeZero() string { _ = "STUB: not implemented"; return "" }

// AddOne counter
func (t *ServiceMethod) AddOne(i int) int {
	_ = "STUB: not implemented"

	// PbService service fields
	return 0
}

type PbService struct {
	Name      string           // Greeter
	LowerName string           // greeter first character to lower
	ProtoName string           // proto file name greeter.proto
	Methods   []*ServiceMethod // service methods

	ImportPkgMap map[string]string // e.g. [userV1]:[userV1 "user/api/user/v1"]

	ProtoFileDir string // e.g. api/user/v1
	ProtoPkgName string // e.g. userV1
	ModuleName   string
}

// RandNumber rand number 1~100
func (s *PbService) RandNumber() int { _ = "STUB: not implemented"; return 0 }

func parsePbService(s *protogen.Service, protoFileDir string, moduleName string) *PbService {
	_ = "STUB: not implemented"
	return nil
}

func getFields(m *protogen.Message) []*Field { _ = "STUB: not implemented"; return nil }

func getMethodComment(m *protogen.Method) string { _ = "STUB: not implemented"; return "" }

func getComment(commentSet protogen.CommentSet) string { _ = "STUB: not implemented"; return "" }

func getCommentStr(comment string) string { _ = "STUB: not implemented"; return "" }

// GetServices parse protobuf services
func GetServices(file *protogen.File, moduleName string) []*PbService {
	_ = "STUB: not implemented"
	return nil
}

func getInvokeType(isStreamingClient bool, isStreamingServer bool) int {
	_ = "STUB: not implemented"
	return 0
}

// bidirectional streaming

// client-side streaming

// server-side streaming

// unary

func getProtoFileDir(protoPath string) string { _ = "STUB: not implemented"; return "" }

func convertToPkgName(importPath string) string { _ = "STUB: not implemented"; return "" }

func isVersionNum(pkgName string) bool { _ = "STUB: not implemented"; return false }

func removeMiddleLine(str string) string { _ = "STUB: not implemented"; return "" }

func getProtoFilename(filenamePrefix string) string { _ = "STUB: not implemented"; return "" }

func getPathDelimiter() string { _ = "STUB: not implemented"; return "" }

// GetImportPkg get import package
func GetImportPkg(services []*PbService) []byte { _ = "STUB: not implemented"; return nil }

// GetSourceImportPkg get source import package
func GetSourceImportPkg(services []*PbService) []byte { _ = "STUB: not implemented"; return nil }
