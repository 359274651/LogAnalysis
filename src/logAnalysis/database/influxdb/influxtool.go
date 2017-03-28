package influxdb

import (
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"logAnalysis/CommonLibrary"
	"time"
)

type InFluxDBTool struct {
	c      client.Client
	dbname string
}

// 创建客户端
func NewClient(addr string, username string, password string, dbname string) *InFluxDBTool {
	c, err := client.NewHTTPClient(client.HTTPConfig{Addr: addr, Username: username, Password: password})
	idb := InFluxDBTool{c, dbname}
	CommonLibrary.CheckError(err)
	return &idb
}

//close connection
func (i *InFluxDBTool) Close() (err error) {
	//client.NewBatchPoints(client.BatchPointsConfig{databas})
	//}
	err = i.c.Close()
	return
}

//插入数据 ,rp没有可以设置为空""
func (i *InFluxDBTool) InsertPoint(tablename, rp string, tags map[string]string, fields map[string]interface{}) (err error) {
	//client.NewBatchPoints(client.BatchPointsConfig{databas})
	//}
	bp, _ := NewBatchPoints(i.dbname, rp)
	p, _ := NewPoint(tablename, tags, fields)
	bp.AddPoint(p)
	err = i.c.Write(bp)
	return
}

// queryDB convenience function to query the database
func (i *InFluxDBTool) Query(cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: i.dbname,
	}
	if response, err := i.c.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

//  查询数据
func (i *InFluxDBTool) QueryData(cmd string, args ...interface{}) ([]client.Result, error) {
	q := fmt.Sprintf(cmd, args...)
	return i.Query(q)
}

//   创建rp  RETENTION POLICY
//dura  = 1d  1m  1s 1w 1h
//replication =1
//CREATE RETENTION POLICY <retention_policy_name> ON <database_name> DURATION <duration> REPLICATION <n> [SHARD DURATION <duration>] [DEFAULT]
func (i *InFluxDBTool) CreateRP(rpname, dbname, dura string, replication int) ([]client.Result, error) {
	q := fmt.Sprintf("CREATE RETENTION POLICY %s ON %s DURATION %s REPLICATION %d", rpname, dbname, dura, replication)
	return i.Query(q)
}

//   创建rp  RETENTION POLICY
//dura  = 1d  1m  1s 1w 1h
//replication =1
//CREATE RETENTION POLICY <retention_policy_name> ON <database_name> DURATION <duration> REPLICATION <n> [SHARD DURATION <duration>] [DEFAULT]
func (i *InFluxDBTool) ModifyRP(rpname, dbname, dura string, replication int) ([]client.Result, error) {
	q := fmt.Sprintf("ALTER RETENTION POLICY %s ON %s DURATION %s REPLICATION %d", rpname, dbname, dura, replication)
	return i.Query(q)
}

//   创建rp  RETENTION POLICY
//dura  = 1d  1m  1s 1w 1h
//replication =1
//CREATE RETENTION POLICY <retention_policy_name> ON <database_name> DURATION <duration> REPLICATION <n> [SHARD DURATION <duration>] [DEFAULT]
func (i *InFluxDBTool) DropRP(rpname, dbname string) ([]client.Result, error) {
	q := fmt.Sprintf("DROP RETENTION POLICY %s ON %s", rpname, dbname)
	return i.Query(q)
}

// 创建数据库
func (i *InFluxDBTool) CreateDB(dbname string) (err error) {
	_, err = i.Query(fmt.Sprintf("CREATE DATABASE %s", dbname))
	return
}

// 删除数据库
func (i *InFluxDBTool) DropDB(dbname string) (err error) {
	_, err = i.Query(fmt.Sprintf("drop DATABASE %s", dbname))
	return
}

func NewBatchPoints(db string, rp string) (bp client.BatchPoints, err error) {
	bp, err = client.NewBatchPoints(client.BatchPointsConfig{"", db, rp, ""})
	CommonLibrary.CheckError(err)
	return
}

func NewPoint(name string,
	tags map[string]string,
	fields map[string]interface{}) (point *client.Point, err error) {
	point, err = client.NewPoint(name, tags, fields, time.Now())
	CommonLibrary.CheckError(err)
	return
}
