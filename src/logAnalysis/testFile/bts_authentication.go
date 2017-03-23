package main

import (
	//"container/list"
	"encoding/json"
	"fmt"
	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type AuthData struct {
	Path   string      `json:"path"`
	Params ArrayString `json:"params"`
	Header ArrayString `json:"header"`
}

type ArrayString map[string]string

//POST:
//{
//"path":"/a/test.mp4",
//"params":[
//["sig","123123123"],
//["token","123"]
//],
//"header":[
//["Host","vcdn.chengzivr.com"],
//["Pragma","no-cache"],
//["Cache-Control","no-cache"],
//["Accept-Encoding","gzip, deflate, sdch"],
//["cookie","name1=value1; name2=value2"],
//["cookie","name3=value3; name4=value4"]
//]
//}

func main() {

	iris.Post("/auth", authSource)

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

	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.Writef("status %d , zyc not fuond", iris.StatusNotFound)
		//ctx.HTML(iris.StatusNotFound, "<h1> You are not allowed here 404 </h1>")
	})
	if len(os.Args) > 1 {
		iris.Listen(fmt.Sprintf("127.0.0.1:%s", os.Args[1]))
	} else {
		iris.Listen(fmt.Sprintf("127.0.0.1:%s", "8080"))
	}

}

func authSource(ctx *iris.Context) {

	// 回复 403 是失败 ，其他状态吗都是成功
	//code, _ := ctx.URLParamInt("param")
	//ctx.SetStatusCode(code)
	ad := AuthData{}
	//ctx.SetStatusCode(200)
	rawData, _ := ioutil.ReadAll(ctx.Request.Body)

	log.Println(string(rawData))
	errs := json.Unmarshal(rawData, &ad)
	//errs := ctx.ReadJSON(&ad)
	if errs != nil {
		log.Println("转换字符串出错")
		ctx.SetStatusCode(403)
	}
	byt, _ := json.MarshalIndent(ad, "", "\t")
	log.Println(string(byt))
	for key, pal := range ad.Params {
		//for key, val := range pal {
		if key == "param" {
			code, _ := strconv.ParseInt(pal, 10, 0)
			ctx.SetStatusCode(int(code))

		}
		if key == "pause" {
			pause, _ := strconv.ParseInt(pal, 10, 0)
			time.Sleep(time.Second * time.Duration(pause))
			log.Println("暂停结束了，哈哈哈")
		}
		//}
	}
	//for key, val := range values {
	//	ctx.Writef("%s ----%x", key, val)
	//
	//}

	ctx.WriteString("testzyc")
}
