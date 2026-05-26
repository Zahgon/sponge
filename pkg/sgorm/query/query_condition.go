// Package query is a library of custom condition queries, support for complex conditional paging queries.
package query

const (
	// Eq equal
	Eq = "eq"
	// Neq not equal
	Neq = "neq"
	// Gt greater than
	Gt = "gt"
	// Gte greater than or equal
	Gte = "gte"
	// Lt less than
	Lt = "lt"
	// Lte less than or equal
	Lte = "lte"
	// Like fuzzy lookup
	Like = "like"
	// In include
	In = "in"
	// NotIN not include
	NotIN = "notin"
	// IsNull is null
	IsNull = "isnull"
	// IsNotNull is not null
	IsNotNull = "isnotnull"

	// AND logic and
	AND string = "and"
	// OR logic or
	OR string = "or"
)

var expMap = map[string]string{
	Eq:        " = ",
	Neq:       " <> ",
	Gt:        " > ",
	Gte:       " >= ",
	Lt:        " < ",
	Lte:       " <= ",
	Like:      " LIKE ",
	In:        " IN ",
	NotIN:     " NOT IN ",
	IsNull:    " IS NULL ",
	IsNotNull: " IS NOT NULL ",

	"=":           " = ",
	"!=":          " <> ",
	">":           " > ",
	">=":          " >= ",
	"<":           " < ",
	"<=":          " <= ",
	"not in":      " NOT IN ",
	"is null":     " IS NULL ",
	"is not null": " IS NOT NULL ",
}

var logicMap = map[string]string{
	AND: " AND ",
	OR:  " OR ",

	"&":   " AND ",
	"&&":  " AND ",
	"|":   " OR ",
	"||":  " OR ",
	"AND": " AND ",
	"OR":  " OR ",

	"and:(": " AND ",
	"and:)": " AND ",
	"or:(":  " OR ",
	"or:)":  " OR ",
}

// ---------------------------------------------------------------------------

type rulerOptions struct {
	whitelistNames map[string]bool
	validateFn     func(columns []Column) error
}

// RulerOption set the parameters of ruler options
type RulerOption func(*rulerOptions)

func (o *rulerOptions) apply(opts ...RulerOption) { _ = "STUB: not implemented"; return }

// WithWhitelistNames set white list names of columns
func WithWhitelistNames(whitelistNames map[string]bool) RulerOption {
	_ = "STUB: not implemented"
	return *new(RulerOption)
}

// WithValidateFn set validate function of columns
func WithValidateFn(fn func(columns []Column) error) RulerOption {
	_ = "STUB: not implemented"
	return *new(RulerOption)
}

// -----------------------------------------------------------------------------

// Params query parameters
type Params struct {
	Page  int    `json:"page" form:"page" binding:"gte=0"`
	Limit int    `json:"limit" form:"limit" binding:"gte=1"`
	Sort  string `json:"sort,omitempty" form:"sort" binding:""`

	Columns []Column `json:"columns,omitempty" form:"columns"` // not required

	// Deprecated: use Limit instead in sponge version v1.8.6, will remove in the future
	Size int `json:"size" form:"size"`
}

// Column query info
type Column struct {
	Name  string      `json:"name" form:"name"`   // column name
	Exp   string      `json:"exp" form:"exp"`     // expressions, default value is "=", support =, !=, >, >=, <, <=, like, in, notin, isnull, isnotnull
	Value interface{} `json:"value" form:"value"` // column value
	Logic string      `json:"logic" form:"logic"` // logical type, defaults to and when the value is null, with &(and), ||(or)
}

// converting ExpType to sql expressions and LogicType to sql using characters
func (c *Column) checkExp() (string, error) { _ = "STUB: not implemented"; return "", nil }

//nolint

// Use rune-safe slicing to preserve multi-byte characters

//nolint

// ConvertToPage converted to page
func (p *Params) ConvertToPage() (order string, limit int, offset int) {
	_ = "STUB: not implemented" //nolint
	return "", 0, 0
}

//nolint

// ConvertToGormConditions conversion to gorm-compliant parameters based on the Columns parameter
// ignore the logical type of the last column, whether it is a one-column or multi-column query
func (p *Params) ConvertToGormConditions(opts ...RulerOption) (string, []interface{}, error) {
	_ = "STUB: not implemented" //nolint
	return "", nil, nil
}

// check name

// check value

// check exp

// ignore the logical type of the last column

// when multiple columns are the same, determine whether the use of IN

// if the value is a string or an integer, if true means it is a string, otherwise it is an integer
func convertValue(v interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// try to parse as RFC3339

// support other formats

// -------------------------------------------------------------------------------------------

// Conditions query conditions
type Conditions struct {
	Columns []Column `json:"columns" form:"columns" binding:"min=1"` // columns info
}

// ConvertToGorm conversion to gorm-compliant parameters based on the Columns parameter
// ignore the logical type of the last column, whether it is a one-column or multi-column query
func (c *Conditions) ConvertToGorm(opts ...RulerOption) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// CheckValid check valid
func (c *Conditions) CheckValid() error { _ = "STUB: not implemented"; return nil }
