package main

import (
	"fmt"

	"github.com/zeiss/pkg/errorx"
	v8 "github.com/zeiss/v8go"
	"github.com/zeiss/v8go-polyfills/listeners"
)

func main() {
	iso := v8.NewIsolate()
	global := v8.NewObjectTemplate(iso)
	defer iso.Dispose()

	in := make(chan *v8.Object, 1)
	out := make(chan *v8.Value, 1)

	errorx.Panic(listeners.Add(iso, global, listeners.WithEvents("auth", in, out)))

	ctx := v8.NewContext(iso, global)
	defer ctx.Close()

	errorx.Must(ctx.RunScript("addEventListener('auth', event => { return event.sourceIP === '127.0.0.1' })", "listener.js"))

	obj := errorx.Must(newContextObject(ctx))

	in <- obj
	v := <-out
	fmt.Println(v)
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
