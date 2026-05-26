// Package mysql provides a gorm driver for mysql.
package mysql

import (
	"gorm.io/gorm"
)

// Init mysql
func Init(dsn string, opts ...Option) (*gorm.DB, error) { _ = "STUB: not implemented"; return nil, nil }

// set the maximum number of connections in the idle connection pool
// set the maximum number of open database connections
// set the maximum time a connection can be reused

// automatic appending of table suffixes when creating tables

// register trace plugin

// register read-write separation plugin

// register plugins

// InitTidb init tidb
func InitTidb(dsn string, opts ...Option) (*gorm.DB, error) {
	_ = "STUB: not implemented"
	return nil,

		// gorm setting
		nil
}

func gormConfig(o *options) *gorm.Config { _ = "STUB: not implemented"; return nil }

// disable foreign key constraints, not recommended for production environments

// removing the plural of an epithet

// print SQL

// print only slow queries

// use the standard output asWriter

// set the logging level, only above the specified level will output the slow query log

func rwSeparationPlugin(o *options) gorm.Plugin {
	_ = "STUB: not implemented"
	return *new(gorm.Plugin)
}

// Close close gorm db
func Close(db *gorm.DB) error { _ = "STUB: not implemented"; return nil }
