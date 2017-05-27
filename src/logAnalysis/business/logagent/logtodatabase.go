package logagent

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
	"github.com/qiniu/log.v1"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	. "logAnalysis/CommonLibrary"
	"logAnalysis/config/agentConf"
	"logAnalysis/database/mongo"
	"logAnalysis/database/mysql"
	"os"
	"strings"
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

func (ali *AgentLogInfo) ReadLogToMongo(mo *agentConf.Config) {
	mg := mongo.CreateMO(mo.MongoC)
	mg.SwitchDB(mo.MongoC.Nodedb)
	nc := NodeCondition{mo.Agentname}
	ncl := NodeCollection{mo.Agentname, mo.NLog.NginxAcessLogPath, mo.NLog.NginxErrorLogPath, mo.AtsLog.AtsAcessLogPath, mo.AtsLog.AtsErrorLogPath, mo.NLog.HttpsNLlog, mo.NLog.HttpsNLErrorlog}
	//ncdata, _ := bson.Marshal(nc)
	//ncldata, _ := bson.Marshal(ncl)
	cinfo, err := mg.UpdateResultAndInsert("nodecollection", &nc, &ncl)
	CheckError(err)
	log.Println("chagenodeinfo: ", cinfo.Removed, " ", cinfo.Updated, " ", cinfo.UpsertedId)

	go readLogToMongo(mo.Agentname, mo.NLog.NginxAcessLogPath, mg, mo.NLog.Separator, mo.NLog.Title)
	go readLogToMongo(mo.Agentname, mo.NLog.NginxErrorLogPath, mg, "", "")
	go readLogToMongo(mo.Agentname, mo.AtsLog.AtsAcessLogPath, mg, mo.AtsLog.Separator, mo.AtsLog.Title)
	go readLogToMongo(mo.Agentname, mo.AtsLog.AtsErrorLogPath, mg, "", "")
	go readLogToMongo(mo.Agentname, mo.NLog.HttpsNLlog, mg, mo.NLog.Separator, mo.NLog.Title)
	go readLogToMongo(mo.Agentname, mo.NLog.HttpsNLErrorlog, mg, "", "")
}

func readLogToMongo(agentname string, logPath string, op *mongo.MgoOp, separate string, title string) {
	if logPath == "" {
		return
	}
	CollectLog(title, op, agentname, logPath, separate)

	fmt.Println("日志停止监控了")
	for {
		if FileIsExist(logPath) {
			log.Println(" 文件存在，开始读取")
			CollectLog(title, op, agentname, logPath, separate)
		}
		//file, _ := os.OpenFile()
		//file.Close()
		time.Sleep(2 * time.Second)
		log.Println("暂停2秒 ， 文件不存在")
	}
}
func FileIsExist(logPath string) bool {
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		return false
	}
	return true
}
func CollectLog(title string, op *mongo.MgoOp, agentname string, logPath string, separate string) {
	defer CatchExecption(logPath)

	titles := strings.Split(title, ",")
	op.SwitchDB(agentname)
	logs := strings.Split(logPath, "/")
	collectionname := logs[len(logs)-1]
	lineindex := mgo.Index{Key: []string{"createAt"}, ExpireAfter: 72 * time.Hour}

	errindex := op.CreateIndex(collectionname, lineindex)
	CheckError(errindex)
	//开始读取日志文件并进行插入
	t, err := tail.TailFile(logPath, tail.Config{Poll: true, Follow: true})
	CheckError(err)
	//针对读取到的数据进行插入到数据库，采用持续监听读取的方式
	for line := range t.Lines {
		if line.Text != "" {
			lineData := line.Text
			//如果有分隔符的就需要经过处理
			if separate != "" {
				linedatas := strings.Split(lineData, separate)
				if linedatas[3] == "~.*" {
					log.Printf("~.*", lineData)
					continue
				}
				if len(linedatas) != len(titles) {
					if len(linedatas)-len(titles) == 1 {
						linedatas[4] = linedatas[4] + "|" + linedatas[5]
						//tempsli := []string{}
						//log.Printf("the line's length is not equal to titles's length ,line is ", lineData)
						linedatas = append(linedatas[:5], linedatas[6:]...)
						//log.Printf("the line's length is not equal to titles's length ,but convert to line is ", linedatas)
					} else if len(linedatas)-len(titles) > 1 {
						//log.Printf("the line's length is not equal to titles's length ,line is ", lineData)
						middle := 13
						end, start := len(linedatas)-(middle+1), 4
						linedatas[start] = strings.Join(linedatas[start:end], "|")
						linedatas = append(linedatas[:start+1], linedatas[len(linedatas)-middle:]...)
						//log.Printf("the line's length is not equal to titles's length ,but convert to line is ", linedatas)
					} else {
						log.Printf("the line's length is not equal to titles's length ,line is ", lineData)
						continue
					}
					//log.Printf("the line's length is not equal to titles's length ,line is ", lineData)
				}
				//var doc string
				var bsonm bson.M = bson.M{}
				for i, val := range titles {
					//temp := val + ":" + linedatas[i]
					//if i == len(titles)-1 {
					//	doc = doc + temp
					//} else {
					//	doc = doc + temp + ","
					//}
					bsonm[strings.Split(val, ":")[0]] = IsDataType(val, linedatas[i])

				}
				bsonm["createAt"] = bson.Now()
				//doc = "{" + doc + "}"
				inerr := op.InsertResult(collectionname, bsonm)
				CheckError(inerr)
			}
			//如果是没有分隔符的就直接插入到数据库中
			if separate == "" {
				//var doc string
				//doc = "{ doc:" + lineData + "}"

				inerr := op.InsertResult(collectionname, bson.M{"createAt": bson.Now(), "doc": lineData})
				CheckError(inerr)
			}

			//omysql.InsertData("insert logStore SET logtext=?,insertTime=?,agentlogid=?;", line.Text, time.Now().Format("2006-01-02 15:04:05"), id)

		}
	}
}
