package handle

//import "github.com/kataras/iris"
import (
	"fmt"
	"gopkg.in/kataras/iris.v6"
	"logAnalysis/CommonLibrary"
	"logAnalysis/handle/logserver"
	"logAnalysis/handle/logserver/server"
	"net/http"
)

func Hi(ctx *iris.Context) {
	ctx.Log(iris.DevMode, "%s%s", ctx.Path(), ctx.Method())

	ctx.Render("index.html", map[string]interface{}{"Name": "iris"}, iris.RenderOptions{"gzip": true})
}

func InitMenu(ctx *iris.Context) {
	ctx.Log(iris.DevMode, "%s%s", ctx.Path(), ctx.Method())
	res, err := server.InitMenu()
	CommonLibrary.CheckHtmlError(err, ctx)
	ctx.JSON(http.StatusOK, res)
}

func InitDocumentKey(ctx *iris.Context) {
	ctx.Log(iris.DevMode, "%s%s", ctx.Path(), ctx.Method())
	ql := logserver.QueryKey{}
	ctx.ReadJSON(&ql)
	fmt.Println("json", ql.String())
	res := server.InitDocumentKey(&ql)
	ctx.JSON(http.StatusOK, res)

}

func Index(ctx *iris.Context) {
	err := ctx.Render("index.html", nil, nil)
	//ctx.WriteString("我尽力；爱了")
	if err != nil {
		fmt.Println(err)
	}
	//ctx.render
}

func Pages(ctx *iris.Context) {
	err := ctx.Render("index.html", nil)
	//ctx.WriteString("我尽力；爱了")
	if err != nil {
		fmt.Println(err)
	}

	//ctx.render
}
func Flot(ctx *iris.Context) {
	err := ctx.Render("flot.html", nil, nil)
	//ctx.WriteString("我尽力；爱了")
	if err != nil {
		fmt.Println(err)
	}
}

func Morris(ctx *iris.Context) {
	err := ctx.Render("morris.html", nil, nil)
	//ctx.WriteString("我尽力；爱了")
	if err != nil {
		fmt.Println(err)
	}
}

func Login(ctx *iris.Context) {
	err := ctx.Render("login.html", nil, nil)
	//ctx.WriteString("我尽力；爱了")
	if err != nil {
		fmt.Println(err)
	}
}
