package main

import (
	"flag"
	"fmt"
	"log"
	"logAnalysis/business/logagent"
	"logAnalysis/config/agentConf"
	"os"
	"os/exec"
	"path/filepath"
)

var block chan int
var confpath *string //用来存储agent 配置文件路径的变量
var defaultconf = "/Users/zhangyachuan/Documents/workCode/log-analysis/src/logAnalysis/app/logagent/logAgent.conf"

func init() {
	const (
		filepathusage = "配置文件路径"
	)
	confpath = flag.String("f", "", filepathusage)

}
func main() {
	flag.Parse()
	var agentConfig *agentConf.Config
	if *confpath != "" {
		agentConfig = agentConf.ReadConfig(*confpath)
		fmt.Println(agentConfig.Agentname)
	} else {
		file, _ := exec.LookPath(os.Args[0])
		path, _ := filepath.Abs(file)
		fmt.Println(path)
		agentConfig = agentConf.ReadConfig(defaultconf)
		fmt.Println(agentConfig.Agentname)
	}

	if agentConfig.Agentname == "" {
		log.Fatalln("agentname 为空停止运行")
	}

	go func() {
		ali := logagent.AgentLogInfo{}
		ali.ReadLogToMysql(agentConfig.MysqlConf.ToString(), agentConfig.Agentname, agentConfig.NLog.NginxAcessLogPath, agentConfig.NLog.Separator, agentConfig.NLog.Title)
	}()

	go func() {
		ali := logagent.AgentLogInfo{}
		ali.ReadLogToMysql(agentConfig.MysqlConf.ToString(), agentConfig.Agentname, agentConfig.NLog.NginxErrorLogPath, "", "")
	}()
	//
	//go func() {
	//	logagent.ReadLogToMysql(agentConfig.MysqlConf.ToString(), agentConfig.Agentname, agentConfig.AtsLog.AtsAcessLogPath, agentConfig.AtsLog.Separator, agentConfig.AtsLog.Title)
	//}()
	//
	//go func() {
	//	logagent.ReadLogToMysql(agentConfig.MysqlConf.ToString(), agentConfig.Agentname, agentConfig.AtsLog.AtsErrorLogPath, "", "")
	//}()

	//因为没有赋值所以按照chan的逻辑，空chan 取数据会造成阻塞
	<-block

}
