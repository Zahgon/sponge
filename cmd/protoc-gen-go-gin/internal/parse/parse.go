// Package parse is parsed proto file to struct
package parse

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Field message field
type Field struct {
	Name      string // field name
	FieldType string // field type
	Comment   string // field comment
}

// GoTypeZero default zero value for type
func (r Field) GoTypeZero() string { _ = "STUB: not implemented"; return "" }

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

// AddOne counter
func (t *ServiceMethod) AddOne(i int) int {
	_ = "STUB: not implemented"

	// PbService service fields
	return 0
}

type PbService struct {
	Name      string           // Greeter
	LowerName string           // greeter first character to lower
	Methods   []*ServiceMethod // service methods

	CutServiceName      string // GreeterService --> Greeter
	LowerCutServiceName string // GreeterService --> greeter

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

//nolint

/*else {
	// if the http method and path is not set, set default value.
	//rpcMethod = defaultMethod(m)
}*/

// GetServices parse protobuf services
func GetServices(file *protogen.File, moduleName string) []*PbService {
	_ = "STUB: not implemented"
	return nil
}

func getCutServiceName(name string) string { _ = "STUB: not implemented"; return "" }

func getFields(m *protogen.Message) []*Field { _ = "STUB: not implemented"; return nil }

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

func getProtoFileDir(protoPath string) string { _ = "STUB: not implemented"; return "" }

func convertToPkgName(importPath string) string { _ = "STUB: not implemented"; return "" }

func isVersionNum(pkgName string) bool { _ = "STUB: not implemented"; return false }

func removeMiddleLine(str string) string { _ = "STUB: not implemented"; return "" }

// GetImportPkg get import package
func GetImportPkg(services []*PbService) []byte { _ = "STUB: not implemented"; return nil }

//pkgName := convertToPkgName(protoFileDir)
//if _, ok := pkgMap[pkgName]; !ok {
//	pkgMap[pkgName] = fmt.Sprintf(`%s "%s"`, pkgName, moduleName+"/"+protoFileDir)
//}

// real package path priority

// GetSourceImportPkg get source import package
func GetSourceImportPkg(services []*PbService) []byte { _ = "STUB: not implemented"; return nil }

// -------------------------------------------------------------------------------------------

// HTTPPbService http service fields
type HTTPPbService struct {
	Name      string // Greeter
	LowerName string // greeter first character to lower

	Methods       []*RPCMethod // service methods
	UniqueMethods []*RPCMethod

	ImportPkgMap map[string]string // [userV1]:[userV1 "user/api/user/v1"]
}

type HTTPPbServices []*HTTPPbService

// ParseHTTPPbServices parse protobuf services
func ParseHTTPPbServices(file *protogen.File) []*HTTPPbService {
	_ = "STUB: not implemented"
	return nil
}

// MergeImportPkgPath merge import package path
func (services HTTPPbServices) MergeImportPkgPath() string { _ = "STUB: not implemented"; return "" }

func removeDuplicates(methods []*RPCMethod) []*RPCMethod { _ = "STUB: not implemented"; return nil }
