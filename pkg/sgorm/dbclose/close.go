// Package dbclose provides a function to close gorm db.
package dbclose

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// Close close gorm db
func Close(db *gorm.DB) error { _ = "STUB: not implemented"; return nil }

func checkInUse(sqlDB *sql.DB, duration time.Duration) { _ = "STUB: not implemented"; return }

//nolint
