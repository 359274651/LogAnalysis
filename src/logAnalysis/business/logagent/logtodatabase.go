package logagent

import (
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/qiniu/log.v1"
	. "logAnalysis/CommonLibrary"
	"logAnalysis/database/mysql"
	"time"
)

type AgentLogInfo struct {
	AgentName string //agent 名称
	LogPath   string
	Seperate  string //分隔符

}

func (ali *AgentLogInfo) ReadLogToMysql(connsqluri, agentName, logPath, seperate, title string) {

	omysql := mysql.OperateMysql{Sqluri: connsqluri}
	defer omysql.CloseDB()
	omysql.Init()
	id, err := omysql.InsertData("insert agentInfo SET agentname=?,logpath=?,separate=?,title=?;", agentName, logPath, seperate, title)
	//id, err := omysql.InsertData("insert into agentInfo(agentname,logpath,separate,title) VALUES(?,?,?,?);", agentName, logPath, seperate, title)
	//开始读取日志文件并进行插入
	t, err := tail.TailFile(logPath, tail.Config{Poll: true, Follow: true})
	CheckError(err)

	//针对读取到的数据进行插入到数据库，采用持续监听读取的方式
	for line := range t.Lines {
		if line.Text != "" {

			offset, _ := t.Tell()
			log.Printf("文件名 %s,偏移量是 %d", t.Filename, offset)
			omysql.InsertData("insert logStore SET logtext=?,insertTime=?,agentlogid=?;", line.Text, time.Now().Format("2006-01-02 15:04:05"), id)

		}
	}

	fmt.Println("日志停止监控了")
}
