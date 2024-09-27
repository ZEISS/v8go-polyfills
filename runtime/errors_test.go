package runtime_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	v8 "github.com/zeiss/v8go"
	"github.com/zeiss/v8go-polyfills/runtime"
)

func TestNewPolyfillError(t *testing.T) {
	err := runtime.NewPolyfillError("foo", errors.New("test"))
	require.NotNil(t, err)
	require.Implements(t, (*error)(nil), err)
	require.Equal(t, "v8-polyfills/foo: test", err.Error())
	require.Equal(t, "test", err.Unwrap().Error())
}

func TestValue(t *testing.T) {
	err := runtime.NewPolyfillError("foo", errors.New("test"))
	require.NotNil(t, err)

	iso := v8.NewIsolate()
	defer iso.Dispose()

	global := v8.NewObjectTemplate(iso)
	ctx := v8.NewContext(iso, global)
	defer ctx.Close()

	val := err.Value(ctx)
	require.NotNil(t, val)
	require.True(t, val.IsString())
	require.Equal(t, "foo: test", val.String())
}
