package parser

// GetSqliteTableInfo get table info from sqlite
func GetSqliteTableInfo(dbFile string, tableName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

//nolint

// SqliteField sqlite field struct
type SqliteField struct {
	Cid          int    `gorm:"column:cid" json:"cid"`
	Name         string `gorm:"column:name" json:"name"`
	Type         string `gorm:"column:type" json:"type"`
	Notnull      int    `gorm:"column:notnull" json:"notnull"`
	DefaultValue string `gorm:"column:dflt_value" json:"dflt_value"`
	Pk           int    `gorm:"column:pk" json:"pk"`
}

var sqliteToMysqlType = map[string]string{
	"integer":       "INT",
	"text":          "TEXT",
	"real":          "FLOAT",
	"datetime":      "DATETIME",
	"blob":          "BLOB",
	"boolean":       "TINYINT",
	"numeric":       " VARCHAR(255)",
	"autoincrement": "auto_increment",
}

func (field *SqliteField) getMysqlType() string { _ = "STUB: not implemented"; return "" }

// SqliteFields sqlite fields
type SqliteFields []*SqliteField

func (fields SqliteFields) getPrimaryField() *SqliteField { _ = "STUB: not implemented"; return nil }

func convertToSQLBySqliteFields(tableName string, fields SqliteFields) string {
	_ = "STUB: not implemented"
	return ""
}
