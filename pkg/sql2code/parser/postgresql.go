package parser

import (
	"gorm.io/gorm"
)

// GetPostgresqlTableInfo get table info from postgres
func GetPostgresqlTableInfo(dsn string, tableName string) (PGFields, error) {
	_ = "STUB: not implemented"
	return *new(PGFields), nil
}

// ConvertToSQLByPgFields convert to mysql table ddl
func ConvertToSQLByPgFields(tableName string, fields PGFields) (string, map[string]string) {
	_ = "STUB: not implemented"
	return "", nil
}

// name:type

// PGField postgresql field
type PGField struct {
	Name         string `gorm:"column:name;" json:"name"`
	Type         string `gorm:"column:type;" json:"type"`
	Comment      string `gorm:"column:comment;" json:"comment"`
	Length       int    `gorm:"column:length;" json:"length"`
	Lengthvar    int    `gorm:"column:lengthvar;" json:"lengthvar"`
	Notnull      bool   `gorm:"column:notnull;" json:"notnull"`
	IsPrimaryKey bool   `gorm:"column:is_primary_key;" json:"is_primary_key"`
}

// nolint
func (field *PGField) getMysqlType() string { _ = "STUB: not implemented"; return "" }

//nolint
//nolint

// unknown type convert to varchar

type PGFields []*PGField

func (fields PGFields) getPrimaryField() *PGField { _ = "STUB: not implemented"; return nil }

func getPostgresqlTableFields(db *gorm.DB, tableName string) (PGFields, error) {
	_ = "STUB: not implemented"
	return *new(PGFields), nil
}

func getType(field *PGField) string { _ = "STUB: not implemented"; return "" }

func closeDB(db *gorm.DB) { _ = "STUB: not implemented"; return }
