package main

import (
	"context"
	"os"

	"goContext/cancellingCtx"
)

func main() {
	ss := cancellingCtx.SlowServer()
	defer ss.Close()
	fs := cancellingCtx.FastServer()
	defer fs.Close()

	ctx := context.Background()
	cancellingCtx.CallBoth(ctx, os.Args[1], ss.URL, fs.URL)
}
