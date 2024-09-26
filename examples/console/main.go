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

	errorx.Panic(console.Add(ctx, console.WithOutput(os.Stdout)))
	errorx.Must(ctx.RunScript("console.log('Hello, World!')", "main.js"))
}

func newContextObject(ctx *v8.Context) (*v8.Object, error) {
	iso := ctx.Isolate()
	obj := v8.NewObjectTemplate(iso)

	resObj, err := obj.NewInstance(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range []struct {
		Key string
		Val interface{}
	}{
		{Key: "sourceIP", Val: "127.0.0.1"},
	} {
		if err := resObj.Set(v.Key, v.Val); err != nil {
			return nil, err
		}
	}

	return resObj, nil
}
