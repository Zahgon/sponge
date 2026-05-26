package parser

var peculiarNouns = map[string]string{
	"ID":    "Id",
	"UID":   "Uid",
	"UUID":  "Uuid",
	"GUID":  "Guid",
	"URI":   "Uri",
	"URL":   "Url",
	"IP":    "Ip",
	"QPS":   "Qps",
	"API":   "Api",
	"ASCII": "Ascii",
	"CPU":   "Cpu",
	"CSS":   "Css",
	"DNS":   "Dns",
	"EOF":   "Eof",
	"HTML":  "Html",
	"HTTP":  "Http",
	"HTTPS": "Https",
	"JSON":  "Json",
	"LHS":   "Lhs",
	"RAM":   "Ram",
	"RHS":   "Rhs",
	"RPC":   "Rpc",
	"SLA":   "Sla",
	"SMTP":  "Smtp",
	"SSH":   "Ssh",
	"TLS":   "Tls",
	"TTL":   "Ttl",
	"UI":    "Ui",
	"UTF8":  "Utf8",
	"VM":    "Vm",
	"XML":   "Xml",
	"XSRF":  "Xsrf",
	"XSS":   "Xss",
}

func toCamel(s string) string { _ = "STUB: not implemented"; return "" }

// special case for table column ID

func firstLetterToLower(str string) string { _ = "STUB: not implemented"; return "" }

func customToCamel(str string) string { _ = "STUB: not implemented"; return "" }

func customToSnake(str string) string { _ = "STUB: not implemented"; return "" }
