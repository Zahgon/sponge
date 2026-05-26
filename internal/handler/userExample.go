package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/go-dev-frame/sponge/internal/dao"
	"github.com/go-dev-frame/sponge/internal/model"
	"github.com/go-dev-frame/sponge/internal/types"
)

var _ UserExampleHandler = (*userExampleHandler)(nil)

// UserExampleHandler defining the handler interface
type UserExampleHandler interface {
	Create(c *gin.Context)
	DeleteByID(c *gin.Context)
	UpdateByID(c *gin.Context)
	GetByID(c *gin.Context)
	List(c *gin.Context)
}

type userExampleHandler struct {
	iDao dao.UserExampleDao
}

// NewUserExampleHandler creating the handler interface
func NewUserExampleHandler() UserExampleHandler {
	_ = "STUB: not implemented"
	return *new(UserExampleHandler)
}

// todo show db driver name here

// Create a new userExample
// @Summary Create a new userExample
// @Description Creates a new userExample entity using the provided data in the request body.
// @Tags userExample
// @Accept json
// @Produce json
// @Param data body types.CreateUserExampleRequest true "userExample information"
// @Success 200 {object} types.CreateUserExampleReply{}
// @Router /api/v1/userExample [post]
// @Security BearerAuth
func (h *userExampleHandler) Create(c *gin.Context) { _ = "STUB: not implemented"; return }

// Note: if copier.Copy cannot assign a value to a field, add it here

// DeleteByID delete a userExample by id
// @Summary Delete a userExample by id
// @Description Deletes a existing userExample identified by the given id in the path.
// @Tags userExample
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.DeleteUserExampleByIDReply{}
// @Router /api/v1/userExample/{id} [delete]
// @Security BearerAuth
func (h *userExampleHandler) DeleteByID(c *gin.Context) { _ = "STUB: not implemented"; return }

// UpdateByID update a userExample by id
// @Summary Update a userExample by id
// @Description Updates the specified userExample by given id in the path, support partial update.
// @Tags userExample
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdateUserExampleByIDRequest true "userExample information"
// @Success 200 {object} types.UpdateUserExampleByIDReply{}
// @Router /api/v1/userExample/{id} [put]
// @Security BearerAuth
func (h *userExampleHandler) UpdateByID(c *gin.Context) { _ = "STUB: not implemented"; return }

// Note: if copier.Copy cannot assign a value to a field, add it here

// GetByID get a userExample by id
// @Summary Get a userExample by id
// @Description Gets detailed information of a userExample specified by the given id in the path.
// @Tags userExample
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetUserExampleByIDReply{}
// @Router /api/v1/userExample/{id} [get]
// @Security BearerAuth
func (h *userExampleHandler) GetByID(c *gin.Context) { _ = "STUB: not implemented"; return }

// Note: if copier.Copy cannot assign a value to a field, add it here

// List get a paginated list of userExamples by custom conditions
// @Summary Get a paginated list of userExamples by custom conditions
// @Description Returns a paginated list of userExample based on query filters, including page number and size.
// @Tags userExample
// @Accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.ListUserExamplesReply{}
// @Router /api/v1/userExample/list [post]
// @Security BearerAuth
func (h *userExampleHandler) List(c *gin.Context) { _ = "STUB: not implemented"; return }

func getUserExampleIDFromPath(c *gin.Context) (string, uint64, bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}

func convertUserExample(userExample *model.UserExample) (*types.UserExampleObjDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: if copier.Copy cannot assign a value to a field, add it here

func convertUserExamples(fromValues []*model.UserExample) ([]*types.UserExampleObjDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
