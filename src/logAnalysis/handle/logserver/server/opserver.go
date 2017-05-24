package server

import (
	"encoding/json"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"logAnalysis/CommonLibrary"
	"logAnalysis/handle/logserver/repository"
)

type CountStatusAreaData map[string]int

type CountStatusAreaDataNew struct {
	Label string `json:"label"`
	Data  int    `json:"data"`
}

//统计状态码  时间 now－1d  1m durations=refresh time
func CountStatusArea(reqtime string) ([]CountStatusAreaDataNew, error) {
	var acsad []CountStatusAreaDataNew
	var csad CountStatusAreaData = make(CountStatusAreaData)
	clresult, err := Repository.CountStatusArea(reqtime)
	if ok := CommonLibrary.CheckErrorPrintln("opserver.go", "CountStatusArea", err); ok {
		return nil, err
	}
	res, _ := json.Marshal(clresult)
	fmt.Println(string(res))
	for _, val := range clresult {
		for _, row := range val.Series {
			for _, vals := range row.Values {
				//f := row.Values[0][1].(json.Number)
				//rowval, _ := f.Int64()
				f := vals[1].(json.Number)
				rowval, _ := f.Int64()
				csad[vals[2].(string)] = int(rowval)
			}

		}
	}
	for key, val := range csad {
		data := CountStatusAreaDataNew{key, val}
		acsad = append(acsad, data)
	}
	return acsad, nil
}

//统计大于某个阀值的时间 的所有请求 时间 now－1d  1m durations=refresh time
func ListMaxRespTime(resptime float32, starttime string) ([]client.Result, error) {

	clresult, err := Repository.ListMaxRespTime(resptime, starttime)
	if ok := CommonLibrary.CheckErrorPrintln("opserver.go", "ListMaxRespTime", err); ok {
		return nil, err
	}
	res, _ := json.Marshal(clresult)
	fmt.Println(string(res))
	//for _, val := range clresult {
	//	for _, row := range val.Series {
	//		f := row.Values[0][1].(json.Number)
	//		rowval, _ := f.Int64()
	//		csad[row.Tags["status"]] = int(rowval)
	//	}
	//}
	return nil, nil
	//return Repository.ListMaxRespTime(resptime, starttime)
}

//统计大于某个阀值的时间 的所有请求 时间 now－1d  1m durations=refresh time
func ListMaxBodySent(respbody float32, starttime string) ([]client.Result, error) {
	return Repository.ListMaxBodySent(respbody, starttime)
}

//统计大于某个阀值的时间 和响应大小的所有请求 时间 now－1d  1m
func ListMaxRespTimeBodySent(resptime, respbody float32, starttime string) ([]client.Result, error) {
	return Repository.ListMaxRespTimeBodySent(resptime, respbody, starttime)

}
