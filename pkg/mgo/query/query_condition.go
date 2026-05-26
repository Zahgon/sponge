// Package query is a library of custom condition queries, support for complex conditional paging queries.
package query

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// Eq equal
	Eq       = "eq"
	eqSymbol = "="
	// Neq not equal
	Neq       = "neq"
	neqSymbol = "!="
	// Gt greater than
	Gt       = "gt"
	gtSymbol = ">"
	// Gte greater than or equal
	Gte       = "gte"
	gteSymbol = ">="
	// Lt less than
	Lt       = "lt"
	ltSymbol = "<"
	// Lte less than or equal
	Lte       = "lte"
	lteSymbol = "<="
	// Like fuzzy lookup
	Like = "like"
	// In include
	In = "in"
	// NotIn exclude
	NotIn = "nin"
	// IsNull is null
	IsNull = "isnull"
	// IsNotNull is not null
	IsNotNull = "isnotnull"

	// AND logic and
	AND        string = "and" //nolint
	andSymbol1        = "&"
	andSymbol2        = "&&"
	// OR logic or
	OR        string = "or" //nolint
	orSymbol1        = "|"
	orSymbol2        = "||"
)

var expMap = map[string]string{
	Eq:            eqSymbol,
	eqSymbol:      eqSymbol,
	Neq:           neqSymbol,
	neqSymbol:     neqSymbol,
	Gt:            gtSymbol,
	gtSymbol:      gtSymbol,
	Gte:           gteSymbol,
	gteSymbol:     gteSymbol,
	Lt:            ltSymbol,
	ltSymbol:      ltSymbol,
	Lte:           lteSymbol,
	lteSymbol:     lteSymbol,
	Like:          Like,
	In:            In,
	NotIn:         NotIn,
	"notin":       NotIn,
	"not in":      NotIn,
	IsNull:        IsNull,
	IsNotNull:     IsNotNull,
	"is null":     IsNull,
	"is not null": IsNotNull,
}

var logicMap = map[string]string{
	AND:        AND,
	"AND":      AND,
	andSymbol1: AND,
	andSymbol2: AND,

	OR:        OR,
	"OR":      OR,
	orSymbol1: OR,
	orSymbol2: OR,

	"and:(": AND,
	"and:)": AND,
	"or:(":  OR,
	"or:)":  OR,
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
	Exp   string      `json:"exp" form:"exp"`     // expressions, default value is "=", support =, !=, >, >=, <, <=, like, in
	Value interface{} `json:"value" form:"value"` // column value
	Logic string      `json:"logic" form:"logic"` // logical type, defaults to and when the value is null, with &(and), ||(or)
}

func (c *Column) checkName(whitelists map[string]bool) error { _ = "STUB: not implemented"; return nil }

func (c *Column) checkValid() error { _ = "STUB: not implemented"; return nil }

func (c *Column) convertLogic() error { _ = "STUB: not implemented"; return nil }

//nolint

func (c *Column) checkLogic() error { _ = "STUB: not implemented"; return nil }

//nolint

// converting ExpType to sql expressions and LogicType to sql using characters
func (c *Column) convert() error { _ = "STUB: not implemented"; return nil }

// nolint
func (c *Column) convertValue() error { _ = "STUB: not implemented"; return nil }

// force to "_id"

// case eqSymbol:

// ConvertToPage converted to page
func (p *Params) ConvertToPage() (sort bson.D, limit int, skip int) {
	_ = "STUB: not implemented" //nolint
	return *new(bson.D), 0, 0
}

//nolint

// ConvertToMongoFilter conversion to mongo-compliant parameters based on the Columns parameter
// ignore the logical type of the last column, whether it is a one-column or multi-column query
func (p *Params) ConvertToMongoFilter(opts ...RulerOption) (bson.M, error) {
	_ = "STUB: not implemented"
	return *new(bson.M), nil
}

// l == 1

// l == 2

// l >=3

func (p *Params) convertMultiColumns(whitelistNames map[string]bool) (bson.M, error) {
	_ = "STUB: not implemented"
	return *new(bson.M), nil
}

func isObjectID(v interface{}) (primitive.ObjectID, bool) {
	_ = "STUB: not implemented"
	return *new(primitive.ObjectID), false
}

type filterGroup struct {
	operator string   // "$and", "$or"
	filters  []bson.M // list of filters within this group
}

// use stack to handle explicit grouping
func buildFilterWithStack(columns []Column) (bson.M, error) {
	_ = "STUB: not implemented"
	return *new(bson.M), nil
}

// use precedence rules to handle flat lists (AND has higher precedence than OR)
func buildFilterWithPrecedence(columns []Column) (bson.M, error) {
	_ = "STUB: not implemented"
	return *new(bson.M), nil
}

// convert a single Column to a BSON condition (no change)
func (c *Column) createSingleCondition() (bson.M, error) {
	_ = "STUB: not implemented"
	return *new(bson.M), nil
}

// if the value is a string or an integer, if true means it is a string, otherwise it is an integer
func convertValue(v interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// try to parse as RFC3339

// support other formats

// -------------------------------------------------------------------------------------------

// Conditions query conditions
type Conditions struct {
	Columns []Column `json:"columns" form:"columns" binding:"min=1"` // columns info
}

// CheckValid check valid
func (c *Conditions) CheckValid() error { _ = "STUB: not implemented"; return nil }

// ConvertToMongo conversion to mongo-compliant parameters based on the Columns parameter
// ignore the logical type of the last column, whether it is a one-column or multi-column query
func (c *Conditions) ConvertToMongo(opts ...RulerOption) (bson.M, error) {
	_ = "STUB: not implemented"
	return *new(bson.M), nil
}
