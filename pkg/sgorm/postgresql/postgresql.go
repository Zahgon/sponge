// Package postgresql provides a gorm driver for postgresql.
package postgresql

import (
	"gorm.io/gorm"
)

// Init postgresql
func Init(dsn string, opts ...Option) (*gorm.DB, error) { _ = "STUB: not implemented"; return nil, nil }

// register trace plugin

// register plugins

// gorm setting
func gormConfig(o *options) *gorm.Config { _ = "STUB: not implemented"; return nil }

// disable foreign key constraints, not recommended for production environments

// removing the plural of an epithet

// print SQL

// print only slow queries

// use the standard output asWriter

// set the logging level, only above the specified level will output the slow query log

// Close close gorm db
func Close(db *gorm.DB) error { _ = "STUB: not implemented"; return nil }
