package database

import (
	"github.com/go-dev-frame/sponge/pkg/sgorm"
)

// InitPostgresql connect postgresql
func InitPostgresql() *sgorm.DB { _ = "STUB: not implemented"; return nil }

// add custom gorm plugin
//opts = append(opts, postgresql.WithGormPlugin(yourPlugin))
