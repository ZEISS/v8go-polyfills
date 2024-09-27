package listeners

import (
	"fmt"
	"sync"

	"github.com/zeiss/pkg/conv"
	v8 "github.com/zeiss/v8go"
	"github.com/zeiss/v8go-polyfills/runtime"
)

var _ runtime.Polyfill = (*Listener)(nil)

// Opt is a functional option for configuring the listener.
type Opt func(*Listener)

// Listener is a polyfill for the addEventListener method.
type Listener struct {
	in  sync.Map
	out sync.Map
}

// GetMethodName ...
func (l *Listener) GetMethodName() string {
	return "addEventListener"
}

// New ...
func New(opt ...Opt) *Listener {
	c := new(Listener)
	c.in = sync.Map{}
	c.out = sync.Map{}

	for _, o := range opt {
		o(c)
	}

	return c
}

// WithEvents ...
func WithEvents(name string, in chan *v8.Object, out chan *v8.Value) Opt {
	return func(l *Listener) {
		l.in.Store(name, in)
		l.out.Store(name, out)
	}
}

// Add ...
func Add(iso *v8.Isolate, global *v8.ObjectTemplate, opts ...Opt) error {
	if iso == nil {
		return fmt.Errorf("v8-polyfills/listeners: isolate is required")
	}

	l := New(opts...)

	ctxFn := v8.NewFunctionTemplate(iso, l.GetFunctionCallback())

	if err := global.Set(l.GetMethodName(), ctxFn, v8.ReadOnly); err != nil {
		return fmt.Errorf("v8-polyfills/listener: %w", err)
	}

	return nil
}

// GetFunctionCallback ...
func (l *Listener) GetFunctionCallback() v8.FunctionCallback {
	return func(info *v8.FunctionCallbackInfo) *v8.Value {
		ctx := info.Context()
		args := info.Args()

		if len(args) <= 1 {
			err := fmt.Errorf("listeners: expected 2 arguments, got %d", len(args))

			return newErrorValue(ctx, err)
		}

		fn, err := args[1].AsFunction()
		if err != nil {
			err := fmt.Errorf("%w", err)

			return newErrorValue(ctx, err)
		}

		chn, ok := l.in.Load(conv.String(args[0]))
		if !ok {
			err := fmt.Errorf("listeners: event %s not found", args[0].String())

			return newErrorValue(ctx, err)
		}

		go func(chn chan *v8.Object, fn *v8.Function) {
			for e := range chn {
				v, err := fn.Call(ctx.Global(), e)
				if err != nil {
					fmt.Printf("listeners: %v", err)
				}

				out, ok := l.out.Load(conv.String(args[0]))
				if !ok {
					fmt.Println("listeners: out channel not found")
				}

				out.(chan *v8.Value) <- v
			}
		}(chn.(chan *v8.Object), fn)

		return v8.Undefined(ctx.Isolate())
	}
}

func newErrorValue(ctx *v8.Context, err error) *v8.Value {
	iso := ctx.Isolate()
	e, _ := v8.NewValue(iso, fmt.Sprintf("addListener: %v", err))

	return e
}
