package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

func main() {
	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			index(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}

	fasthttp.ListenAndServe(":3000", m)
}

func index(ctx *fasthttp.RequestCtx)  {
	start := time.Now()
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q %q", ctx.RequestURI(), ctx)
	elapsed := time.Since(start)

	fmt.Printf("execution time %q  || host: %q \n", elapsed, ctx.Host())
}