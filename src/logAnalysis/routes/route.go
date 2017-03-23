package routes

import (
	"github.com/kataras/iris"
	"logAnalysis/handle"
)

func init() {
	iris.Get("/index", handle.Index)
	iris.Get("/hi", handle.Hi)

}
