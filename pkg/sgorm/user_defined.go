package sgorm

import (
	"database/sql/driver"
)

// Bool is a custom type for MySQL bit(1) type and PostgreSQL boolean type.
type Bool bool
type BitBool = Bool

// Scan implement Scanner interface to read values from database
func (b *Bool) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

// Value implement Valuer interface to write values to database
func (b Bool) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// default MySQL processing

var currentDriver string

// SetDriver sets the name of the current database driver, such as "postgres"
// if you use postgres, you need to call SetDriver("postgres") after initializing gorm
func SetDriver(driverName string) { _ = "STUB: not implemented"; return }

// --------------------------------------------------------------------------------------

// TinyBool is a custom type for MySQL tinyint(1) type
type TinyBool bool

// Scan implement Scanner interface to read values from database
func (b *TinyBool) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

// Value implement Valuer interface to write values to database
func (b TinyBool) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}
