package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	//"log"
	"logAnalysis/config/agentConf"
)

type MgoOp struct {
	mgodb   *mgo.Database
	session *mgo.Session
}

//切换数据库
func (m *MgoOp) SwitchDB(mdbname string) {
	m.mgodb = m.session.DB(mdbname)
}

//更新信息，如果不存在则插入数据
func (m *MgoOp) UpdateResultAndInsert(collectionname string, selector, updatedata interface{}) (*mgo.ChangeInfo, error) {
	colname := m.mgodb.C(collectionname)
	return colname.Upsert(selector, updatedata)
}

//查找一群结果，并返回
func (m *MgoOp) InsertResult(collectionname string, doc ...interface{}) error {
	colname := m.mgodb.C(collectionname)
	return colname.Insert(doc...)
}

//查找一群结果，并返回
func (m *MgoOp) CreateIndex(collectionname string, index mgo.Index) error {
	var match = false
	colname := m.mgodb.C(collectionname)
	i, err := colname.Indexes()
	if err != nil {
		return err
	} else {
		for _, val := range i {
			//如果已经存在索引，则不新建，直接返回
			if stringEq(index.Key, val.Key) {
				return nil
			}
		}

	}

	return colname.EnsureIndex(index)
}

//删除索引根据key
func (m *MgoOp) DropIndex(collectionname string, index ...string) error {
	colname := m.mgodb.C(collectionname)
	return colname.DropIndex(index...)
}

//查找一个结果，并将结果返回给一个结构体
func (m *MgoOp) FindOneResult(collectionname string, query interface{}, result interface{}) error {
	colname := m.mgodb.C(collectionname)
	return colname.Find(query).One(result)
}

//查找一群结果，并返回
func (m *MgoOp) FindResult(collectionname string, query interface{}) *mgo.Query {
	colname := m.mgodb.C(collectionname)
	return colname.Find(query)
}

//查找Aggreate
func (m *MgoOp) AggregatPipeResult(collectionname string, pipeline interface{}) *mgo.Pipe {
	colname := m.mgodb.C(collectionname)
	return colname.Pipe(pipeline)
}

//创建mongo数据库对象，可以操作collection
func CreateMO(moo agentConf.MO) *MgoOp {
	var surl string
	if moo.Dbuser != "" {
		//[mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
		surl = fmt.Sprintf("%s:%s@%s:%s/%s", moo.Dbuser, moo.Dbpassword, moo.Dbhost, moo.Dbport, "admin")
	} else {
		surl = fmt.Sprintf("%s:%s/%s", moo.Dbhost, moo.Dbport, moo.Dbname)
	}

	session, err := mgo.Dial(surl)
	if err != nil {
		panic(err)
	}

	return &MgoOp{session.DB(moo.Dbname), session}

}

func stringEq(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
