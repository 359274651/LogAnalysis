package handle

import "github.com/kataras/iris"

func Hi(ctx *iris.Context) {
	ctx.Render("hi.html", map[string]interface{}{"Name": "iris"}, iris.RenderOptions{"gzip": true})
}

//func Index(ctx *iris.Context) {
//	ctx.Render("Dashboard.html", nil, iris.RenderOptions{"gzip": true})
//	//ctx.render
//}

func Index(ctx *iris.Context) {
	ctx.Render("index.html", nil, iris.RenderOptions{"gzip": true})
	//ctx.render
}
