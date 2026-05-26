// Package database provides database client initialization.
package database

import (
	"sync"

	"github.com/go-dev-frame/sponge/pkg/sgorm"
)

var (
	gdb     *sgorm.DB
	gdbOnce sync.Once

	ErrRecordNotFound = sgorm.ErrRecordNotFound
)

// todo generate initialisation database code here
// delete the templates code start

// InitDB connect database
func InitDB() { _ = "STUB: not implemented"; return }

// delete the templates code end

// GetDB get db
func GetDB() *sgorm.DB { _ = "STUB: not implemented"; return nil }

// CloseDB close db
func CloseDB() error { _ = "STUB: not implemented"; return nil }
