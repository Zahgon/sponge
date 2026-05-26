// Package gotest is a library that simulates the testing of cache, dao and handler.
package gotest

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

// Service info
type Service struct {
	Ctx      context.Context
	TestData interface{}
	MockDao  *Dao

	Server *grpc.Server
	listen net.Listener

	clientAddr     string
	clientConn     *grpc.ClientConn
	IServiceClient interface{}
}

// NewService instantiated service
func NewService(dao *Dao, testData interface{}) *Service { _ = "STUB: not implemented"; return nil }

// GoGrpcServer run grpc server
func (s *Service) GoGrpcServer() { _ = "STUB: not implemented"; return }

// GetClientConn dial rpc server
func (s *Service) GetClientConn() *grpc.ClientConn { _ = "STUB: not implemented"; return nil }

// Close service
func (s *Service) Close() { _ = "STUB: not implemented"; return }
