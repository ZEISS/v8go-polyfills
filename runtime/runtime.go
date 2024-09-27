package runtime

import (
	v8 "github.com/zeiss/v8go"
)

// Polyfill is an interface that represents a polyfill.
type Polyfill interface {
	// GetFunctionCallback returns the function callback.
	GetFunctionCallback() v8.FunctionCallback
	// GetMethodName returns the method name.
	GetMethodName() string
}

// CompatibilityDate is a string that represents the date of the last update.
type CompatibilityDate string

// CompatibilityFlag is a string that represents the name of a compatibility flag.
type CompatibilityFlag string

// CompatibilityFlags is a map of compatibility flags.
type CompatibilityFlags map[CompatibilityFlag]bool

// CompatibilityMatrix is a map of compatibility dates.
type CompatibilityMatrix map[CompatibilityDate]CompatibilityFlags

// Compatibility is a map of compatibility matrices.
var Compatibility = CompatibilityMatrix{
	"2024-10-01": {
		"console":          true,
		"addEventListener": true,
	},
}
