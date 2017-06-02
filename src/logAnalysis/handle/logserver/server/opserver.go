package server

import (
	"fmt"
	//"log"
	"gopkg.in/mgo.v2/bson"
	. "logAnalysis/CommonLibrary"
	"logAnalysis/handle/logserver"
	"logAnalysis/handle/logserver/repository"
)

type CountStatusAreaData map[string]int

type CountStatusAreaDataNew struct {
	Label string `json:"label"`
	Data  int    `json:"data"`
}

func InitMenu() (*logserver.MenuData, error) {
	var nodec []logserver.NodeCollection
	result := Repository.InitMenu()
	err := result.All(&nodec)
	CheckPrintlnError(err)
	var lmd logserver.MenuData = logserver.MenuData{}
	for _, nodeval := range nodec {
		if nodeval.Nodename == "" {
			continue
		}
		lmd[logserver.DB(nodeval.Nodename)] = []string{
			nodeval.AtsErrlog,
			nodeval.NLlog,
			nodeval.Atslog,
			nodeval.HttpsErrorNllog,
			nodeval.HttpsNLlog,
			nodeval.NLErrlog}
	}

	return &lmd, err
}

func InitDocumentKey(qk *logserver.QueryKey) []string {
	var nodec bson.M
	result := Repository.InitDocumentKey(qk)
	err := result.One(&nodec)
	CheckPrintlnError(err)
	title := []string{}
	for key, _ := range nodec {
		title = append(title, key)
	}
	return title
	//return &lmd, err
}

//统计状态码  时间 now－1d  1m durations=refresh time
func CountStatusArea(conditon *logserver.QueryData) ([]CountStatusAreaDataNew, error) {
	var acsad []CountStatusAreaDataNew
	clresult, err := Repository.CountStatusArea(conditon)
	if err != nil {
		return nil, err
	}
	bs := []bson.M{}
	clresult.Iter().All(&bs)
	fmt.Println(bs)

	for _, val := range bs {
		data := CountStatusAreaDataNew{val["_id"].(string), val["sum_status"].(int)}
		acsad = append(acsad, data)
	}
	return acsad, nil
}

////统计大于某个阀值的时间 的所有请求 时间 now－1d  1m durations=refresh time
//func ListMaxRespTime(resptime float32, starttime string) ([]client.Result, error) {
//
//	clresult, err := Repository.ListMaxRespTime(resptime, starttime)
//	if ok := CheckErrorPrintln("opserver.go", "ListMaxRespTime", err); ok {
//		return nil, err
//	}
//	res, _ := json.Marshal(clresult)
//	fmt.Println(string(res))
//	//for _, val := range clresult {
//	//	for _, row := range val.Series {
//	//		f := row.Values[0][1].(json.Number)
//	//		rowval, _ := f.Int64()
//	//		csad[row.Tags["status"]] = int(rowval)
//	//	}
//	//}
//	return nil, nil
//	//return Repository.ListMaxRespTime(resptime, starttime)
//}

////统计大于某个阀值的时间 的所有请求 时间 now－1d  1m durations=refresh time
//func ListMaxBodySent(respbody float32, starttime string) ([]client.Result, error) {
//	return Repository.ListMaxBodySent(respbody, starttime)
//}
//
////统计大于某个阀值的时间 和响应大小的所有请求 时间 now－1d  1m
//func ListMaxRespTimeBodySent(resptime, respbody float32, starttime string) ([]client.Result, error) {
//	return Repository.ListMaxRespTimeBodySent(resptime, respbody, starttime)
//
//}
