package service

import (
	"context"
	"time"

	"google.golang.org/grpc"

	serverNameExampleV1 "github.com/go-dev-frame/sponge/api/serverNameExample/v1"
	"github.com/go-dev-frame/sponge/internal/dao"
	"github.com/go-dev-frame/sponge/internal/model"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		serverNameExampleV1.RegisterUserExampleServer(server, NewUserExampleServer()) // register service to the rpc service
	})
}

var _ serverNameExampleV1.UserExampleServer = (*userExample)(nil)
var _ time.Time

type userExample struct {
	serverNameExampleV1.UnimplementedUserExampleServer

	iDao dao.UserExampleDao
}

// NewUserExampleServer create a new service
func NewUserExampleServer() serverNameExampleV1.UserExampleServer {
	_ = "STUB: not implemented"
	return *new(serverNameExampleV1.UserExampleServer)
}

// todo show db driver name here

// Create a new userExample
func (s *userExample) Create(ctx context.Context, req *serverNameExampleV1.CreateUserExampleRequest) (*serverNameExampleV1.CreateUserExampleReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: if copier.Copy cannot assign a value to a field, add it here

// DeleteByID delete a userExample by id
func (s *userExample) DeleteByID(ctx context.Context, req *serverNameExampleV1.DeleteUserExampleByIDRequest) (*serverNameExampleV1.DeleteUserExampleByIDReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateByID update a userExample by id
func (s *userExample) UpdateByID(ctx context.Context, req *serverNameExampleV1.UpdateUserExampleByIDRequest) (*serverNameExampleV1.UpdateUserExampleByIDReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: if copier.Copy cannot assign a value to a field, add it here

// GetByID get a userExample by id
func (s *userExample) GetByID(ctx context.Context, req *serverNameExampleV1.GetUserExampleByIDRequest) (*serverNameExampleV1.GetUserExampleByIDReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List get a paginated list of userExamples by custom conditions
func (s *userExample) List(ctx context.Context, req *serverNameExampleV1.ListUserExampleRequest) (*serverNameExampleV1.ListUserExampleReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: if copier.Copy cannot assign a value to a field, add it here

func convertUserExample(record *model.UserExample) (*serverNameExampleV1.UserExample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: if copier.Copy cannot assign a value to a field, add it here
