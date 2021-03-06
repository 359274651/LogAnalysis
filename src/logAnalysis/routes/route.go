package routes

import (
	//"github.com/kataras/iris"
	"gopkg.in/kataras/iris.v6"
	"logAnalysis/handle"
	"logAnalysis/handle/logserver"
)

//func init() {
//	iris.Get("/index", handle.Index)
//	iris.Get("/hi", handle.Hi)
//
//}

func InitRoute(app *iris.Framework) {
	app.Get("/index", handle.Pages)
	//app.Get("/index.html", handle.Index)
	app.Get("/hi", handle.Pages)
	app.Get("/", handle.Pages)
	app.Get("/flot", handle.Flot)
	app.Get("/morris", handle.Morris)
	app.Get("/login", handle.Login)

	app.Get("/countstatusarea/:reqtime/:startime/:durations", logserver.CountStatusArea)
	app.Get("/ListMaxBodySent", logserver.ListMaxBodySent)
	app.Get("/ListMaxRespTime", logserver.ListMaxRespTime)
	app.Get("/ListMaxRespTimeBodySent", logserver.ListMaxRespTimeBodySent)
}
