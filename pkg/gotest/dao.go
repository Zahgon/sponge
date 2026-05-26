package gotest

import (
	"context"
	"database/sql/driver"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

// Dao dao info
type Dao struct {
	Ctx      context.Context
	TestData interface{}
	SQLMock  sqlmock.Sqlmock
	Cache    *Cache
	DB       *gorm.DB
	IDao     interface{}
	AnyTime  *anyTime
	closeFns []func()
}

// NewDao instantiated dao
func NewDao(c *Cache, testData interface{}) *Dao { _ = "STUB: not implemented"; return nil }

// mock mysql

// Close dao
func (d *Dao) Close() { _ = "STUB: not implemented"; return }

// GetAnyArgs Dynamic generation of parameter types based on structures
func (d *Dao) GetAnyArgs(obj interface{}) []driver.Value { _ = "STUB: not implemented"; return nil }

//nolint

type anyTime struct{}

// Match satisfies sqlmock.Argument interface,
// if the table has fields of type time.Time, this method must be implemented
func (a *anyTime) Match(v driver.Value) bool { _ = "STUB: not implemented"; return false }
