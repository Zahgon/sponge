package handler

import (
	"context"
	"time"

	serverNameExampleV1 "github.com/go-dev-frame/sponge/api/serverNameExample/v1"
	"github.com/go-dev-frame/sponge/internal/dao"
	"github.com/go-dev-frame/sponge/internal/model"
)

var _ serverNameExampleV1.UserExampleLogicer = (*userExamplePbHandler)(nil)
var _ time.Time

type userExamplePbHandler struct {
	userExampleDao dao.UserExampleDao
}

// NewUserExamplePbHandler create a handler
func NewUserExamplePbHandler() serverNameExampleV1.UserExampleLogicer {
	_ = "STUB: not implemented"
	return *new(serverNameExampleV1.UserExampleLogicer)
}

// todo show db driver name here

// Create a new userExample
func (h *userExamplePbHandler) Create(ctx context.Context, req *serverNameExampleV1.CreateUserExampleRequest) (*serverNameExampleV1.CreateUserExampleReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: if copier.Copy cannot assign a value to a field, add it here

// DeleteByID delete a userExample by id
func (h *userExamplePbHandler) DeleteByID(ctx context.Context, req *serverNameExampleV1.DeleteUserExampleByIDRequest) (*serverNameExampleV1.DeleteUserExampleByIDReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateByID update a userExample by id
func (h *userExamplePbHandler) UpdateByID(ctx context.Context, req *serverNameExampleV1.UpdateUserExampleByIDRequest) (*serverNameExampleV1.UpdateUserExampleByIDReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: if copier.Copy cannot assign a value to a field, add it here

// GetByID get a userExample by id
func (h *userExamplePbHandler) GetByID(ctx context.Context, req *serverNameExampleV1.GetUserExampleByIDRequest) (*serverNameExampleV1.GetUserExampleByIDReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List get a paginated list of userExamples by custom conditions
func (h *userExamplePbHandler) List(ctx context.Context, req *serverNameExampleV1.ListUserExampleRequest) (*serverNameExampleV1.ListUserExampleReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: if copier.Copy cannot assign a value to a field, add it here

func convertUserExamplePb(record *model.UserExample) (*serverNameExampleV1.UserExample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: if copier.Copy cannot assign a value to a field, add it here
