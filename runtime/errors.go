package runtime

import (
	"fmt"

	"github.com/zeiss/pkg/errorx"
	v8 "github.com/zeiss/v8go"
)

// Value is a v8 value.
type Value interface {
	Value(ctx *v8.Context) *v8.Value
}

// PolyfillError is an error that occurs when a polyfill fails.
type PolyfillError struct {
	Err  error
	Name string
}

// NewPolyfillError creates a new PolyfillError.
func NewPolyfillError(name string, err error) *PolyfillError {
	return &PolyfillError{err, name}
}

// Error returns the error message.
func (e *PolyfillError) Error() string {
	return fmt.Sprintf("v8-polyfills/%s: %v", e.Name, e.Err)
}

// Unwrap implements the errors.Wrapper interface.
func (e *PolyfillError) Unwrap() error { return e.Err }

// Value returns the error value.
func (e *PolyfillError) Value(ctx *v8.Context) *v8.Value {
	iso := v8.NewIsolate()

	return errorx.Ignore(v8.NewValue(iso, fmt.Sprintf("%s: %v", e.Name, e.Err)))
}
