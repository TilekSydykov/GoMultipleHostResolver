package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			index(ctx)
		case "/{id}":
			g(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}

	fasthttp.ListenAndServe(":3000", m)
}

func index(ctx *fasthttp.RequestCtx)  {
	start := time.Now()
	db, _ := sql.Open("mysql", "icamp:itcamp@/db_nodejs_1")
	db.Prepare("select * from users")
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q %q", ctx.RequestURI(), ctx)
	elapsed := time.Since(start)

	fmt.Printf("execution time %q  || host: %q \n", elapsed, ctx.Host())
}

func g(ctx *fasthttp.RequestCtx){
	fmt.Fprintf(ctx, "another uri")
}


