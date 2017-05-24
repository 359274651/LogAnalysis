package CommonLibrary

import (
	"fmt"
	"log"
	//"time"

	"gopkg.in/kataras/iris.v6"
)

//import "log"

func CheckError(err error) {
	if err != nil {
		log.Fatalln("error: %s", err.Error())
	}
}

func CheckHtmlError(err error, ctx *iris.Context) {
	if err != nil {
		ctx.Log(iris.DevMode, err.Error())
	}

}

func CheckErrorPrintln(filename string, method string, err error) bool {
	if err != nil {
		return true
		fmt.Println(" 文件名：", filename, " 方法：", method, " 异常：", err)
	}
	return false
}
