package parse

import (
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
)

var methodSets = make(map[string]int)

// GetMethods get rpc method descriptions
func GetMethods(m *protogen.Method, protoSelfPkgPath string) []*RPCMethod {
	_ = "STUB: not implemented"
	return nil

	// http rule config
}

func buildHTTPRule(m *protogen.Method, rule *annotations.HttpRule, protoSelfPkgPath string) *RPCMethod {
	_ = "STUB: not implemented"
	return nil
}

// default

func buildMethodDesc(m *protogen.Method, httpMethod, path string, customKind string, selector string, protoSelfPkgPath string) *RPCMethod {
	_ = "STUB: not implemented"
	return nil
}

// RPCMethod describes a rpc method
type RPCMethod struct {
	Name       string // SayHello
	Num        int    // one rpc RPCMethod can correspond to multiple http requests
	Request    string // SayHelloReq
	Reply      string // SayHelloResp
	InvokeType int    // 0:unary, 1: client-side streaming, 2: server-side streaming, 3: bidirectional streaming

	// http_rule
	Path         string // rule
	Method       string // HTTP Method
	Body         string
	ResponseBody string

	CustomKind string
	Selector   string
	// if Selector is [ctx], and IsPassGinContext is true
	// if true, pass gin.Context to the rpc method
	IsPassGinContext bool
	// if Selector is [no_bind], IsPassGinContext and IsPassGinContext are both true
	// if true, ignore c.ShouldBindXXX for this method, you must use c.ShouldBindXXX() in rpc method
	IsIgnoreShouldBind bool

	RequestImportPkgName string // e.g. empty or userV1
	ReplyImportPkgName   string // e.g. empty or userV1

	ProtoSelfPkgPath string              // e.g. "module/api/user/v1"
	ImportPkgPaths   map[string]struct{} // exclude ProtoSelfPkgPath
}

// HandlerName for gin handler name
func (m *RPCMethod) HandlerName() string { _ = "STUB: not implemented"; return "" }

// HasPathParams whether to include routing parameters
func (m *RPCMethod) HasPathParams() bool { _ = "STUB: not implemented"; return false }

// parse selector and set custom control variables
func parseVariable(str string) (prefixStr string, isPassGinContext bool, isIgnoreShouldBind bool) {
	_ = "STUB: not implemented"
	return "", false, false
}

// pass gin.Context

func (m *RPCMethod) checkCustomKind() { _ = "STUB: not implemented"; return }

func (m *RPCMethod) checkSelector() { _ = "STUB: not implemented"; return }

// InitPathParams conversion parameter routing {xx} --> :xx
func (m *RPCMethod) InitPathParams() { _ = "STUB: not implemented"; return }
