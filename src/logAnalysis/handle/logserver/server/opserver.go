package server

import (
	"github.com/influxdata/influxdb/client"
	"logAnalysis/handle/logserver/repository"
)

//统计状态码  时间 now－1d  1m durations=refresh time
func CountStatusArea(reqtime, starttime, durations string) ([]client.Result, error) {
	return Repository.CountStatusArea(reqtime, starttime, durations)

}

//统计大于某个阀值的时间 的所有请求 时间 now－1d  1m durations=refresh time
func ListMaxRespTime(resptime float32, starttime string) ([]client.Result, error) {
	return Repository.ListMaxRespTime(resptime, starttime)
}

//统计大于某个阀值的时间 的所有请求 时间 now－1d  1m durations=refresh time
func ListMaxBodySent(respbody float32, starttime string) ([]client.Result, error) {
	return Repository.ListMaxBodySent(respbody, starttime)
}

//统计大于某个阀值的时间 和响应大小的所有请求 时间 now－1d  1m
func ListMaxRespTimeBodySent(resptime, respbody float32, starttime string) ([]client.Result, error) {
	return Repository.ListMaxRespTimeBodySent(resptime, respbody, starttime)

}
