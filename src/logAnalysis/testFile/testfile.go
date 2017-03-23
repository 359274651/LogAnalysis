package main

import (
	////"bufio"
	//"fmt"
	////"github.com/hpcloud/tail"
	"log"
	////"os"
	//"io"
	"net/http"
	//"time"
	_ "net/http/pprof"
)

//import "fmt"

//func HelloServer(w http.ResponseWriter, req *http.Request) {
//	io.WriteString(w, "hello, world!\n")
//}
func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	//log.Println("start")
	//
	//go func() {
	//	http.HandleFunc("/hello", HelloServer)
	//	log.Println("zhixingalalll")
	//	err := http.ListenAndServe("127.0.0.1:8090", nil)
	//
	//	//log.Fatal(http.ListenAndServe("127.0.0.1:8090", nil))
	//	log.Fatal(err)
	//	log.Println("zhixingalalll..............")
	//}()
	//
	//log.Println("main runing .......")
	exit := make(chan bool)
	<-exit
	//fmt.Println(time.Now().String())
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//t, _ := tail.TailFile("/Users/zhangyachuan/Documents/proxy_access.log", tail.Config{Poll: true, Follow: true})
	//for line := range t.Lines {
	//	log.Println(line.Text)
	//	//fmt.Println(line.Text)
	//}
	//ReadLogfile("/Users/zhangyachuan/Documents/proxy_access.log")
}

//
//func ReadLogfile(logpath string) {
//	var debugLine = ""
//	// 读取配置文件
//	logfile, fileerr := os.Open(logpath)
//	if fileerr != nil {
//		log.Fatalln(fileerr)
//	}
//	defer logfile.Close()
//
//	defer func() {
//		if err := recover(); err != nil {
//			log.Println(string(debugLine), err)
//
//			panic(err)
//		}
//	}()
//	rblog := bufio.NewReader(logfile)
//	for {
//		time.Sleep(1)
//		readline, _, nerr := rblog.ReadLine()
//		if nerr != nil {
//			log.Fatalln("----------------------" + nerr.Error())
//		}
//		debugLine = string(readline)
//		log.Println(debugLine)
//	}
//
//}
