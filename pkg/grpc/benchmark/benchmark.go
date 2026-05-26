// Package benchmark is compression testing of rpc methods and generation of reported results.
package benchmark

import (
	"github.com/bojand/ghz/runner"
	"google.golang.org/protobuf/proto"
)

type Option = runner.Option

// Runner interface
type Runner interface {
	Run() error
}

// bench pressing parameters
type bench struct {
	rpcServerHost string // rpc server address

	protoFile              string        // proto file
	packageName            string        // proto file package name
	serviceName            string        // proto file service name
	methodName             string        // name of pressure test method
	methodRequest          proto.Message // input parameters corresponding to the method
	dependentProtoFilePath []string      // dependent proto file path

	total uint // number of requests

	options []runner.Option
}

// New create a pressure test instance
//
// invalid parameter total if the option runner.WithRunDuration is set
func New(host string, protoFile string, methodName string, req proto.Message, dependentProtoFilePath []string, total int, options ...runner.Option) (Runner, error) {
	_ = "STUB: not implemented"
	return *new(Runner), nil
}

// Run operational performance benchmarking
func (b *bench) Run() error { _ = "STUB: not implemented"; return nil }

// more parameter settings https://github.com/bojand/ghz/blob/master/runner/options.go#L41
// example settings: https://github.com/bojand/ghz/blob/master/runner/options_test.go#L79

func (b *bench) saveReport(callMethod string, report *runner.Report) error {
	_ = "STUB: not implemented"
	// specify the output path
	return nil
}
