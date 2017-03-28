package Repository

import (
	"fmt"

	"github.com/influxdata/influxdb/client"

	"logAnalysis/runglobal"
)

//统计状态码  时间 now－1d  1m durations=refresh time
func CountStatusArea(reqtime, starttime, durations string) ([]client.Result, error) {
	nginxlogcon := &runglobal.GlobalConf.NLog
	nginxlog := GetRepository(nginxlogcon)
	defer DeferCloseconn(nginxlog)

	nginxlog.QueryData("select count(status) from %s where time_local > %s and time > now - %s group by status time(%s) fill(0)", nginxlogcon.Tablename, reqtime, starttime, durations)

}

//统计大于某个阀值的时间 的所有请求 时间 now－1d  1m durations=refresh time
func ListMaxRespTime(resptime float32, starttime string) ([]client.Result, error) {
	nginxlogcon := &runglobal.GlobalConf.NLog
	nginxlog := GetRepository(nginxlogcon)
	defer DeferCloseconn(nginxlog)
	return nginxlog.QueryData("select * from %s where request_time > %f and time > now() - %s ", nginxlogcon.Tablename, resptime, starttime)

}

//统计大于某个阀值的时间 的所有请求 时间 now－1d  1m durations=refresh time
func ListMaxBodySent(respbody float32, starttime string) ([]client.Result, error) {
	nginxlogcon := &runglobal.GlobalConf.NLog
	nginxlog := GetRepository(nginxlogcon)
	defer DeferCloseconn(nginxlog)
	return nginxlog.QueryData("select count(status) from %s where body_bytes_sent::float > %f and time > now() - %s ", nginxlogcon.Tablename, respbody, starttime)

}

//统计大于某个阀值的时间 和响应大小的所有请求 时间 now－1d  1m
func ListMaxRespTimeBodySent(resptime, respbody float32, starttime string) ([]client.Result, error) {
	nginxlogcon := &runglobal.GlobalConf.NLog
	nginxlog := GetRepository(nginxlogcon)
	defer DeferCloseconn(nginxlog)
	return nginxlog.QueryData("select count(status) from %s where body_bytes_sent::float > %f and request_time > %f and time > now() - %s ", nginxlogcon.Tablename, respbody, resptime, starttime)

}
