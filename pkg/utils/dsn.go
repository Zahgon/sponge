package utils

// AdaptiveMysqlDsn adaptation of various mysql format dsn address
func AdaptiveMysqlDsn(dsn string) string { _ = "STUB: not implemented"; return "" }

// AdaptivePostgresqlDsn convert postgres dsn to kv string
func AdaptivePostgresqlDsn(dsn string) string { _ = "STUB: not implemented"; return "" }

// AdaptiveSqlite adaptive sqlite
func AdaptiveSqlite(dbFile string) string {
	_ = "STUB: not implemented"
	// todo convert to absolute path
	return ""
}

// AdaptiveMongodbDsn adaptive mongodb dsn
func AdaptiveMongodbDsn(dsn string) string { _ = "STUB: not implemented"; return "" }

// default scheme

// DeleteBrackets delete brackets in dsn
func DeleteBrackets(str string) string { _ = "STUB: not implemented"; return "" }
