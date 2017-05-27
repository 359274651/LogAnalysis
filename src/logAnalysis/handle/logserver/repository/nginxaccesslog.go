package Repository

import (
	//"log"

	"github.com/influxdata/influxdb/client/v2"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"logAnalysis/CommonLibrary"
	//"logAnalysis/business/logagent"
	"encoding/json"
	"fmt"
	"github.com/qiniu/log.v1"
	"logAnalysis/config/agentConf"
	"logAnalysis/database/mongo"
	"logAnalysis/runglobal"
	"time"
)

//公共的node集合名称
var nodecoll = "nodecollection"

func InitMenu() *mgo.Query {
	mongodb := ConvertMoEntity()
	defer runglobal.PublicLock.Unlock()
	runglobal.PublicLock.Lock()
	mgdb := GetmongoDb(mongodb)
	data, _ := json.Marshal(mongodb)
	fmt.Println("-----------", string(data))
	//time.Sleep(15 * time.Second)
	mgdb.SwitchDB(mongodb.Nodedb)
	return mgdb.FindResult(nodecoll, bson.M{})
}

//统计状态码  时间 now－1d  1m durations=refresh time
func CountStatusArea(reqtime string) ([]client.Result, error) {
	//var nodersl []logagent.NodeCollection
	//mongodb := ConvertMoEntity()
	//defer runglobal.PublicLock.Unlock()
	//runglobal.PublicLock.Lock()
	//mgdb := GetmongoDb(mongodb)
	//
	////defer DeferCloseconn(nginxlog)
	//mgdb.SwitchDB(mongodb.Nodedb)
	//resl := mgdb.FindResult("nodecollection", nil)
	//err := resl.All(nodersl)
	//CommonLibrary.CheckError(err)
	//for _, nodeval := range nodersl {
	//	nodeval.Nodename
	//}
	//return nginxlog.QueryData("SELECT count(\"port\") FROM \"%s\" WHERE  time > now() - %s GROUP BY \"status\" fill(0)", nginxlogcon.Tablename, reqtime)
	//return nginxlog.QueryData("select cs as data,status as label from dg5telegraf.rp6h.countstatusarea where time >= now() - 10m")
	return nil, nil
}
func GetmongoDb(mongodb agentConf.MO) *mongo.MgoOp {
	if runglobal.GlobalMongdb != nil {
		log.Println("nill")
		return runglobal.GlobalMongdb
	} else {
		log.Println("else")
		return mongo.CreateMO(mongodb)
	}
}
func ConvertMoEntity() agentConf.MO {
	temp := agentConf.MO{}
	temp.Dbhost = runglobal.GlobalConf.Mongo.Dbhost
	temp.Dbname = runglobal.GlobalConf.Mongo.Dbname
	temp.Dbpassword = runglobal.GlobalConf.Mongo.Dbpassword
	temp.Dbport = runglobal.GlobalConf.Mongo.Dbport
	temp.Dbuser = runglobal.GlobalConf.Mongo.Dbuser
	temp.Nodedb = runglobal.GlobalConf.Mongo.Nodedb
	return temp
}

////统计大于某个阀值的时间 的所有请求 时间 now－1d  1m durations=refresh time
//func ListMaxRespTime(resptime float32, starttime string) ([]client.Result, error) {
//	nginxlogcon := &runglobal.GlobalConf.NLog
//	nginxlog := GetRepository(nginxlogcon)
//	defer DeferCloseconn(nginxlog)
//	return nginxlog.QueryData("select * from %s where request_time > %f and time > now() - %s ", nginxlogcon.Tablename, resptime, starttime)
//
//}
//
////统计大于某个阀值的时间 的所有请求 时间 now－1d  1m durations=refresh time
//func ListMaxBodySent(respbody float32, starttime string) ([]client.Result, error) {
//	nginxlogcon := &runglobal.GlobalConf.NLog
//	nginxlog := GetRepository(nginxlogcon)
//	defer DeferCloseconn(nginxlog)
//	return nginxlog.QueryData("select count(status) from %s where body_bytes_sent::float > %f and time > now() - %s ", nginxlogcon.Tablename, respbody, starttime)
//
//}
//
////统计大于某个阀值的时间 和响应大小的所有请求 时间 now－1d  1m
//func ListMaxRespTimeBodySent(resptime, respbody float32, starttime string) ([]client.Result, error) {
//	nginxlogcon := &runglobal.GlobalConf.NLog
//	nginxlog := GetRepository(nginxlogcon)
//	defer DeferCloseconn(nginxlog)
//	return nginxlog.QueryData("select count(status) from %s where body_bytes_sent::float > %f and request_time > %f and time > now() - %s ", nginxlogcon.Tablename, respbody, resptime, starttime)
//
//}
