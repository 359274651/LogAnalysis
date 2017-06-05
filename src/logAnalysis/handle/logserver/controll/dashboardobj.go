package controll

import (
	"gopkg.in/mgo.v2/bson"
	"logAnalysis/CommonLibrary"
	. "logAnalysis/handle/logserver"
	"time"
)

//ReqCountStatusArea 统计状态码分布
type ReqCountStatusArea struct {
	DB         string `json:"db"` //查询的数据库
	Collection string `json:"cc"` //mongo Collection  ,一般就是日志文件的名字,可以通过path.Split(path) (dir,name string) 获取
	Sd         string `json:"sd"`
	Ed         string `json:"ed"`
}

func (rs *ReqCountStatusArea) ToQueryData() *QueryData {
	layout := "2006-01-02 15:04:05"
	sd := Conver2DateLocation(layout, rs.Sd)
	ed := Conver2DateLocation(layout, rs.Ed)
	qd := &QueryData{}
	qd.DB = DB(rs.DB)
	qd.Collection = rs.Collection
	qd.Contidions = []KeyContidion{KeyContidion{"match": []Contidion{Contidion{"time_local", "大于等于", sd}, Contidion{"time_local", "小于等于", ed.UTC()}}, "group": []Contidion{Contidion{"_id", "", "$status"}, Contidion{"sum_status", "", bson.M{"$sum": 1}}}}}

	return qd

}

func Conver2Date(layout string, val string) time.Time {
	t, err := time.Parse(layout, val)
	CommonLibrary.CheckPrintlnError(err)
	return t
}
func Conver2DateLocation(layout string, val string) time.Time {
	t, err := time.ParseInLocation(layout, val, time.UTC)
	CommonLibrary.CheckPrintlnError(err)
	t = t.Add(-8 * time.Hour)
	return t
}
