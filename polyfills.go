package polyfills

import (
	v8 "github.com/zeiss/v8go"
)

// Polyfill is the interface that wraps the Inject method.
type Polyfill interface {
	// Inject injects the polyfill into the given context.
	Inject(ctx *v8.Context) error
}
