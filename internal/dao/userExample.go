package dao

import (
	"context"

	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"github.com/go-dev-frame/sponge/pkg/sgorm/query"

	"github.com/go-dev-frame/sponge/internal/cache"
	"github.com/go-dev-frame/sponge/internal/model"
)

var _ UserExampleDao = (*userExampleDao)(nil)

// UserExampleDao defining the dao interface
type UserExampleDao interface {
	Create(ctx context.Context, table *model.UserExample) error
	DeleteByID(ctx context.Context, id uint64) error
	UpdateByID(ctx context.Context, table *model.UserExample) error
	GetByID(ctx context.Context, id uint64) (*model.UserExample, error)
	GetByColumns(ctx context.Context, params *query.Params) ([]*model.UserExample, int64, error)

	CreateByTx(ctx context.Context, tx *gorm.DB, table *model.UserExample) (uint64, error)
	DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.UserExample) error
}

type userExampleDao struct {
	db    *gorm.DB
	cache cache.UserExampleCache // if nil, the cache is not used.
	sfg   *singleflight.Group    // if cache is nil, the sfg is not used.
}

// NewUserExampleDao creating the dao interface
func NewUserExampleDao(db *gorm.DB, xCache cache.UserExampleCache) UserExampleDao {
	_ = "STUB: not implemented"
	return *new(UserExampleDao)
}

func (d *userExampleDao) deleteCache(ctx context.Context, id uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Create a new userExample, insert the record and the id value is written back to the table
func (d *userExampleDao) Create(ctx context.Context, table *model.UserExample) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteByID delete a userExample by id
func (d *userExampleDao) DeleteByID(ctx context.Context, id uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// delete cache

// UpdateByID update a userExample by id, support partial update
func (d *userExampleDao) UpdateByID(ctx context.Context, table *model.UserExample) error {
	_ = "STUB: not implemented"
	return nil
}

// delete cache

func (d *userExampleDao) updateDataByID(ctx context.Context, db *gorm.DB, table *model.UserExample) error {
	_ = "STUB: not implemented"
	return nil
}

// todo generate the update fields code to here
// delete the templates code start

// delete the templates code end

// GetByID get a userExample by id
func (d *userExampleDao) GetByID(ctx context.Context, id uint64) (*model.UserExample, error) {
	_ = "STUB: not implemented"
	// no cache
	return nil, nil
}

// get from cache

// get from database

// for the same id, prevent high concurrent simultaneous access to database
//nolint

// set placeholder cache to prevent cache penetration, default expiration time 10 minutes

// set cache

// GetByColumns get a paginated list of userExamples by custom conditions.
// For more details, please refer to https://go-sponge.com/component/data/custom-page-query.html
func (d *userExampleDao) GetByColumns(ctx context.Context, params *query.Params) ([]*model.UserExample, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// determine if count is required

// CreateByTx create a record in the database using the provided transaction
func (d *userExampleDao) CreateByTx(ctx context.Context, tx *gorm.DB, table *model.UserExample) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DeleteByTx delete a record by id in the database using the provided transaction
func (d *userExampleDao) DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// delete cache

// UpdateByTx update a record by id in the database using the provided transaction
func (d *userExampleDao) UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.UserExample) error {
	_ = "STUB: not implemented"
	return nil
}

// delete cache
