package console_test

import (
	"os"
	"testing"

	"github.com/zeiss/v8go-polyfills/console"

	"github.com/stretchr/testify/require"
	v8 "github.com/zeiss/v8go"
)

func TestAdd(t *testing.T) {
	iso := v8.NewIsolate()
	global := v8.NewObjectTemplate(iso)

	ctx := v8.NewContext(iso, global)

	err := console.Add(ctx, console.WithOutput(os.Stdout))
	require.NoError(t, err)

	defer ctx.Close()

	_, err = ctx.RunScript("console.log('hello world')", "console.js")
	require.NoError(t, err)
}
