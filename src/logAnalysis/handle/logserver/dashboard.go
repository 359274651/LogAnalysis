package logserver

import (
	"gopkg.in/kataras/iris.v6"
	"net/http"
)

func CountStatusArea(ctx *iris.Context) {

	if ctx.Method() != http.MethodGet {
		ctx.SetStatusCode(403)
		return
	}

	ctx.Render("hi.html", map[string]interface{}{"Name": "iris"}, iris.RenderOptions{"gzip": true})
}

func ListMaxRespTime(ctx *iris.Context) {
	ctx.Log(iris.DevMode, "%s%s", ctx.Path(), ctx.Method())

	ctx.Render("hi.html", map[string]interface{}{"Name": "iris"}, iris.RenderOptions{"gzip": true})
}

func ListMaxBodySent(ctx *iris.Context) {
	ctx.Log(iris.DevMode, "%s%s", ctx.Path(), ctx.Method())

	ctx.Render("hi.html", map[string]interface{}{"Name": "iris"}, iris.RenderOptions{"gzip": true})
}

func ListMaxRespTimeBodySent(ctx *iris.Context) {
	ctx.Log(iris.DevMode, "%s%s", ctx.Path(), ctx.Method())

	ctx.Render("hi.html", map[string]interface{}{"Name": "iris"}, iris.RenderOptions{"gzip": true})
}
