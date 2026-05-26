// Package service A grpc server-side or client-side package that handles business logic.
package service

import (
	"google.golang.org/grpc"
)

var (
	// registerFns collection of registration methods
	registerFns []func(server *grpc.Server)
)

// RegisterAllService register all services to the service
func RegisterAllService(server *grpc.Server) { _ = "STUB: not implemented"; return }

// Register for Health Screening
