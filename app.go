package main

import (
	"./engine"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
)

func main() {
	m := func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, engine.GetWebPage(ctx))
	}
	fasthttp.ListenAndServe(":3000", m)
}
