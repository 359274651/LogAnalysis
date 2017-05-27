package controll

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gopkg.in/kataras/iris.v6"

	"logAnalysis/CommonLibrary"
	"logAnalysis/handle/logserver/server"
)

func CountStatusArea(ctx *iris.Context) {

	if ctx.Method() != http.MethodGet {
		ctx.SetStatusCode(403)
		return
	}
	ctx.Log(iris.DevMode, ""+ctx.RequestPath(true))
	ctx.Log(iris.DevMode, "i'm coming here csa")
	//"select count(status) from %s where time_local > %s and time > now - %s group by status time(%s) fill(0)"
	timelocal := ctx.Param("reqtime")
	clresult, err := server.CountStatusArea(timelocal)
	ctx.Log(iris.DevMode, "i'm coming here csa")
	CommonLibrary.CheckHtmlError(err, ctx)
	result, _ := json.Marshal(clresult)

	fmt.Println(string(result))
	ctx.WriteString(string(result))
}

func ListMaxRespTime(ctx *iris.Context) {

	if ctx.Method() != http.MethodGet {
		ctx.SetStatusCode(403)
		return
	}

	//"select count(status) from %s where time_local > %s and time > now - %s group by status time(%s) fill(0)"
	starttime := ctx.Param("starttime")
	resptime := ctx.Param("resptime")
	tt, _ := strconv.ParseFloat(resptime, 32)
	clresult, err := server.ListMaxRespTime(float32(tt), starttime)

	CommonLibrary.CheckHtmlError(err, ctx)
	result, _ := json.Marshal(clresult)

	fmt.Println(string(result))
	ctx.WriteString(string(result))
}

func ListMaxBodySent(ctx *iris.Context) {
	ctx.Log(iris.DevMode, "%s%s", ctx.Path(), ctx.Method())

	ctx.Render("hi.html", map[string]interface{}{"Name": "iris"}, iris.RenderOptions{"gzip": true})
}

func ListMaxRespTimeBodySent(ctx *iris.Context) {
	ctx.Log(iris.DevMode, "%s%s", ctx.Path(), ctx.Method())

	ctx.Render("hi.html", map[string]interface{}{"Name": "iris"}, iris.RenderOptions{"gzip": true})
}
