package engine

import "github.com/valyala/fasthttp"

func GetWebPage(ctx *fasthttp.RequestCtx) string {
	return string(ctx.Host()) + " " + string(ctx.Path())
}
