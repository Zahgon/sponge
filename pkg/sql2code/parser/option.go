package parser

// NullStyle null type
type NullStyle int

// nolint
const (
	NullDisable NullStyle = iota
	NullInSql
	NullInPointer
)

// Option function
type Option func(*options)

type options struct {
	DBDriver       string
	FieldTypes     map[string]string // name:type
	Charset        string
	Collation      string
	JSONTag        bool
	JSONNamedType  int
	TablePrefix    string
	ColumnPrefix   string
	NoNullType     bool
	NullStyle      NullStyle
	Package        string
	GormType       bool
	ForceTableName bool
	IsEmbed        bool // is gorm.Model embedded
	IsWebProto     bool // true: proto file include router path and swagger info, false: normal proto file without router and swagger
	IsExtendedAPI  bool // true: extended api (9 api), false: basic api (5 api)

	IsCustomTemplate bool // true: custom extend template, false: sponge template
}

var defaultOptions = options{
	DBDriver:   "mysql",
	FieldTypes: map[string]string{},
	NullStyle:  NullInSql,
	Package:    "model",
}

// WithDBDriver set db driver
func WithDBDriver(driver string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithFieldTypes set field types
func WithFieldTypes(fieldTypes map[string]string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCharset set charset
func WithCharset(charset string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCollation set collation
func WithCollation(collation string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTablePrefix set table prefix
func WithTablePrefix(p string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithColumnPrefix set column prefix
func WithColumnPrefix(p string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithJSONTag set json tag, 0 for underscore, other values for hump
func WithJSONTag(namedType int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNoNullType set NoNullType
func WithNoNullType() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNullStyle set NullType
func WithNullStyle(s NullStyle) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPackage set package name
func WithPackage(pkg string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithGormType will write type in gorm tag
func WithGormType() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithForceTableName set forceFloats
func WithForceTableName() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEmbed is embed gorm.Model
func WithEmbed() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithWebProto set proto file type
func WithWebProto() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithExtendedAPI set extended api
func WithExtendedAPI() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCustomTemplate set custom template
func WithCustomTemplate() Option { _ = "STUB: not implemented"; return *new(Option) }

func parseOption(options []Option) options { _ = "STUB: not implemented"; return *new(options) }
