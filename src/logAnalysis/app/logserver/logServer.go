package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
	"gopkg.in/kataras/iris.v6/middleware/logger"

	"logAnalysis/config/serverConf"
	. "logAnalysis/routes"
	"logAnalysis/runglobal"
)

var confpath *string
var defaultconf = "/Users/zhangyachuan/Documents/workCode/log-analysis/src/logAnalysis/app/logserver/logServer.conf"

func init() {
	const (
		filepathusage = "配置文件路径"
	)
	confpath = flag.String("f", "", filepathusage)

}

func main() {

	flag.Parse()
	var agentConfig *serverConf.Config
	if *confpath != "" {
		agentConfig = serverConf.ReadConfig(*confpath)

	} else {
		file, _ := exec.LookPath(os.Args[0])
		path, _ := filepath.Abs(file)
		fmt.Println(path)
		agentConfig = serverConf.ReadConfig(defaultconf)

	}
	// 为全局变量扶植
	runglobal.GlobalConf = agentConfig
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

	app := iris.New(iris.Configuration{Gzip: false, Charset: "UTF-8"})
	app.Use(customLogger)
	app.Adapt(httprouter.New())
	app.Adapt(iris.DevLogger())
	errorLogger := logger.New()
	app.Favicon("./templates/icon/Taget_Icon_128px_1183257_easyicon.net.ico")
	//app.Adapt(html.New()).Directory("../../templates", ".html")
	tmpl := view.HTML("./templates/pages", ".html")
	tmpl.Reload(true)
	tmpl.Layout("layout/layout.html") //这里等到定义好文件之后就可以填写
	app.Adapt(tmpl)
	//app.Adapt(view.HTML("./templates/pages", ".html").Reload(true))

	app.StaticWeb("/vendor", "./templates/vendor")
	app.StaticWeb("/dist", "./templates/dist")
	app.StaticWeb("/data", "./templates/data")
	app.OnError(iris.StatusForbidden, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.HTML(iris.StatusForbidden, "<h1> You are not allowed here 403 </h1>")
	})

	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.HTML(iris.StatusNotFound, "<h1> You are not allowed here 404 </h1>")
	})

	InitRoute(app)
	app.Listen(agentConfig.ServerAddr)
}
