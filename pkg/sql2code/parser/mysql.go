package parser

import (
	_ "github.com/go-sql-driver/mysql" //nolint
)

// GetMysqlTableInfo get table info from mysql
func GetMysqlTableInfo(dsn, tableName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

//nolint

//nolint

// GetTableInfo get table info from mysql
// Deprecated: replaced by GetMysqlTableInfo
func GetTableInfo(dsn, tableName string) (string, error) { _ = "STUB: not implemented"; return "", nil }
