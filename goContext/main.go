package main

import (
	"context"
	"fmt"
	"os"

	"goContext/cancellingCtx"
	"goContext/timeoutCtx"
)

func main() {
	switch os.Args[1] {
	case "cancelling":
		ss := cancellingCtx.SlowServer()
		defer ss.Close()
		fs := cancellingCtx.FastServer()
		defer fs.Close()
		ctx := context.Background()
		cancellingCtx.CallBoth(ctx, os.Args[2], ss.URL, fs.URL)
	case "timeout":
		timeoutCtx.SimpleTimeout()
	case "timeout2":
		ctx := context.Background()
		res, _ := timeoutCtx.LongRunningThingManager(ctx, os.Args[2])
		fmt.Println(res)
	default:
		return
	}
}
