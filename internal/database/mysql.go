package database

import (
	"github.com/go-dev-frame/sponge/pkg/sgorm"
)

// InitMysql connect mysql
func InitMysql() *sgorm.DB { _ = "STUB: not implemented"; return nil }

// setting mysql slave and master dsn addresses
//opts = append(opts, mysql.WithRWSeparation(
//	mysqlCfg.SlavesDsn,
//	mysqlCfg.MastersDsn...,
//))

// add custom gorm plugin
//opts = append(opts, mysql.WithGormPlugin(yourPlugin))
