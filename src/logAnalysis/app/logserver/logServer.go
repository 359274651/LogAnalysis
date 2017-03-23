package main

import (
	//"github.com/kataras/go-template/django"
	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
	_ "logAnalysis/routes"
)

func main() {

	//定义日志格式

	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
	})

	iris.Use(customLogger)

	errorLogger := logger.New()
	iris.UseTemplate(html.New()).Directory("../../templates", ".html")
	iris.OnError(iris.StatusForbidden, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.HTML(iris.StatusForbidden, "<h1> You are not allowed here 403 </h1>")
	})

	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.HTML(iris.StatusNotFound, "<h1> You are not allowed here 404 </h1>")
	})
	iris.StaticWeb("/css", "../../templates/css")
	iris.StaticWeb("/jquery", "../../templates/jquery")
	iris.StaticWeb("/bootstrap", "../../templates/bootstrap/")

	iris.Listen("127.0.0.1:8080")
}
