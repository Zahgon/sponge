package query

import (
	"go.mongodb.org/mongo-driver/bson"
)

var defaultMaxSize = 1000

const oidName = "_id"

// SetMaxSize change the default maximum number of pages per page
func SetMaxSize(maxValue int) { _ = "STUB: not implemented"; return }

// Page info
type Page struct {
	page  int // page number, starting from page 0
	limit int // number per page

	// sort fields, default is id backwards, you can add - sign before the field to indicate
	// reverse order, no - sign to indicate ascending order, multiple fields separated by comma
	sort bson.D
}

// Page get page value
func (p *Page) Page() int {
	_ = "STUB: not implemented"

	// Limit number per page
	return 0
}

func (p *Page) Limit() int {
	_ = "STUB: not implemented"

	// Size number per page
	// Deprecated: use Limit instead, will delete it in the future
	return 0
}

func (p *Page) Size() int {
	_ = "STUB: not implemented"

	// Sort get sort field
	return 0
}

func (p *Page) Sort() bson.D {
	_ = "STUB: not implemented"

	// Skip get offset value
	return *new(bson.D)
}

func (p *Page) Skip() int { _ = "STUB: not implemented"; return 0 }

// DefaultPage default page, number 20 per page, sorted by id backwards
func DefaultPage(page int) *Page { _ = "STUB: not implemented"; return nil }

//nolint

// NewPage custom page, starting from page 0.
// the parameter columnNames indicates a sort field, if empty means id descending, if there are multiple column names, separated by a comma,
// a '-' sign in front of each column name indicates descending order, otherwise ascending order.
func NewPage(page int, limit int, columnNames string) *Page { _ = "STUB: not implemented"; return nil }

// convert to mysql sort, each column name preceded by a '-' sign, indicating descending order, otherwise ascending order, example:
//
//	columnNames="name" means sort by name in ascending order,
//	columnNames="-name" means sort by name descending,
//	columnNames="name,age" means sort by name in ascending order, otherwise sort by age in ascending order,
//	columnNames="-name,-age" means sort by name descending before sorting by age descending.
func getSort(columnNames string) bson.D { _ = "STUB: not implemented"; return *new(bson.D) }

//nolint

//nolint

//nolint
