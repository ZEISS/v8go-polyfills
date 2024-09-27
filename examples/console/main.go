package main

import (
	"os"

	"github.com/zeiss/v8go-polyfills/console"

	"github.com/zeiss/pkg/errorx"
	v8 "github.com/zeiss/v8go"
)

func main() {
	iso := v8.NewIsolate()
	ctx := v8.NewContext(iso)

	defer iso.Dispose()
	defer ctx.Close()

	errorx.Panic(console.Add(ctx, console.WithOutput(os.Stdout)))
	errorx.Must(ctx.RunScript("console.log('Hello, World!')", "main.js"))
}
