package handle

//import "github.com/kataras/iris"
import (
	"fmt"
	"gopkg.in/kataras/iris.v6"
)

func Hi(ctx *iris.Context) {
	ctx.Log(iris.DevMode, "%s%s", ctx.Path(), ctx.Method())

	ctx.Render("index.html", map[string]interface{}{"Name": "iris"}, iris.RenderOptions{"gzip": true})
}

//func Index(ctx *iris.Context) {
//	ctx.Render("Dashboard.html", nil, iris.RenderOptions{"gzip": true})
//	//ctx.render
//}

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
