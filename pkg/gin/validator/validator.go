// Package validator is gin request parameter check library.
package validator

import (
	"sync"

	valid "github.com/go-playground/validator/v10"
)

// Init validator instance, used to gin request parameter check
func Init() *CustomValidator { _ = "STUB: not implemented"; return nil }

// CustomValidator Custom valid objects
type CustomValidator struct {
	once     sync.Once
	Validate *valid.Validate
}

// NewCustomValidator Instantiate
func NewCustomValidator() *CustomValidator { _ = "STUB: not implemented"; return nil }

// ValidateStruct validates a struct or slice/array
func (v *CustomValidator) ValidateStruct(obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// pointer type: if nil, no validation required; otherwise recursive validation after dereference

// slice or array type: iterates over each element, recursively validating one by one

// Engine set tag name "binding", which is implementing the validator interface of the gin framework
func (v *CustomValidator) Engine() interface{} { _ = "STUB: not implemented"; return nil }

func (v *CustomValidator) lazyInit() { _ = "STUB: not implemented"; return }
