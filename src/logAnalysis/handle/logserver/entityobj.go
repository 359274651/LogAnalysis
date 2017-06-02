package logserver

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type QueryKey struct {
	DB         `json:"db"` //查询的数据库
	Collection string      `json:"collection"` //mongo Collection  ,一般就是日志文件的名字,可以通过path.Split(path) (dir,name string) 获取
}

func (qk *QueryKey) String() string {
	return fmt.Sprintf("%s %s", qk.DB, qk.Collection)
}

//存储查询条件用的结构体 ,默认只支持and方法不支持or
type QueryData struct {
	DB         `json:"db"`    //查询的数据库
	Collection string         `json:"collection"` //mongo Collection  ,一般就是日志文件的名字,可以通过path.Split(path) (dir,name string) 获取
	Contidions []KeyContidion `json:"contidion"`
}

type KeyContidion map[string][]Contidion

type Contidion struct {
	Key   string      `json:"key"`
	Oper  string      `json:"oper"` //运算符表示 $lt/$lte/$gt/$gte/$ne，依次等价于</<=/>/>=/!=。
	Value interface{} `json:"value"`
}

//生成对应的mongo查询语句
//以下方法对符合类型的查询操作存在不支持的情况，需要后面补充完善,思路混乱只能先写别的
func (q *QueryData) GenerateMql() (bson.M, error) {
	var res bson.M = bson.M{}
	var precon []bson.M = []bson.M{}
	for _, con := range q.Contidions {
		for key, val := range con {
			if op, err := MatchOper(key); err != nil || op == "" {
				return nil, err
			} else {

				bm, err := generateMql(val)
				if err != nil {
					return nil, err
				}
				precon = append(precon, bm)
			}
			op, _ := MatchOper(key)
			res[op] = precon
		}

	}
	return res, nil
}

//生成对应的mongo管道查询语句结构pipe
func (q *QueryData) GenerateMqlPipe() ([]bson.M, error) {
	var precon []bson.M = []bson.M{}
	for _, con := range q.Contidions {
		for key, val := range con {
			if op, err := MatchOper(key); err != nil || op == "" {
				return nil, err
			} else {
				op, _ := MatchOper(key)
				if key == "group" {
					bm, err := generateGroupMql2(val)
					if err != nil {
						return nil, err
					}
					precon = append(precon, bson.M{op: bm})
					continue
				}
				bm, err := generateMql(val)
				if err != nil {
					return nil, err
				}
				precon = append(precon, bson.M{op: bm})
			}
		}

	}
	return precon, nil
}

//针对group的格式，单独生成数据结构
func generateGroupMql2(data []Contidion) (bson.M, error) {
	var temp bson.M = bson.M{}
	for _, kc := range data {
		temp[kc.Key] = kc.Value
		//if _, err := MatchOper(kc.Oper); err != nil {
		//	return nil, err
		//}
		//if val, ok := temp[kc.Key]; ok {
		//	tv := val.(bson.M)
		//	tv[kc.Oper] = kc.Value
		//} else {
		//	temp[kc.Key] = bson.M{kc.Oper: kc.Value}
		//}

	}
	return temp, nil
}

func generateMql(data []Contidion) (bson.M, error) {
	var temp bson.M = bson.M{}
	for _, kc := range data {
		if _, err := MatchOper(kc.Oper); err != nil {
			return nil, err
		}
		op, _ := MatchOper(kc.Oper)
		if val, ok := temp[kc.Key]; ok {
			tv := val.(bson.M)

			tv[op] = kc.Value
		} else {
			temp[kc.Key] = bson.M{op: kc.Value}
		}

	}
	return temp, nil
}

//匹配管道符号，用于操作
func MatchOper(n string) (string, error) {
	switch n {
	case "match":
		return "$match", nil
	case "project":
		return "$project", nil
	case "group":
		return "$group", nil
	case "sort":
		return "$sort", nil
	case "skip":
		return "$skip", nil
	case "lt", "小于", "<":
		return "$lt", nil
	case "大于", "gt", ">":
		return "$gt", nil
	case "大于等于", "gte", ">=":
		return "$gte", nil
	case "lte", "<=", "小于等于":
		return "$lte", nil
	case "等于", "equal", "=", "eq":
		return "$eq", nil
	case "不等于", "ne", "!=":
		return "$ne", nil
	case "或", "|", "||", "or":
		return "$or", nil
	case "并且", "and", "&&", "&":
		return "$and", nil
	case "in", "include":
		return "$in", nil
	case "nin":
		return "$nin", nil
	case "regex":
		return "$regex", nil
	//case "":
	//case "":
	//case "":

	default:
		return "", errors.New("表达式没有匹配到，请检查")
	}
	return "", errors.New("表达式没有匹配到，请检查")
}

/*

$project：修改输入文档的结构。可以用来重命名、增加或删除域，也可以用于创建计算结果以及嵌套文档。
$match：用于过滤数据，只输出符合条件的文档。$match使用MongoDB的标准查询操作。
$limit：用来限制MongoDB聚合管道返回的文档数。
$skip：在聚合管道中跳过指定数量的文档，并返回余下的文档。
$unwind：将文档中的某一个数组类型字段拆分成多条，每条包含数组中的一个值。
$group：将集合中的文档分组，可用于统计结果。
$sort：将输入文档排序后输出。
$geoNear：输出接近某一地理位置的有序文档。

等于	{<key>:<value>}	db.col.find({"by":"菜鸟教程"}).pretty()	where by = '菜鸟教程'
小于	{<key>:{$lt:<value>}}	db.col.find({"likes":{$lt:50}}).pretty()	where likes < 50
小于或等于	{<key>:{$lte:<value>}}	db.col.find({"likes":{$lte:50}}).pretty()	where likes <= 50
大于	{<key>:{$gt:<value>}}	db.col.find({"likes":{$gt:50}}).pretty()	where likes > 50
大于或等于	{<key>:{$gte:<value>}}	db.col.find({"likes":{$gte:50}}).pretty()	where likes >= 50
不等于	{<key>:{$ne:<value>}}	db.col.find({"likes":{$ne:50}}).pretty()	where likes != 50
MongoDB AND 条件
MongoDB 的 find() 方法可以传入多个键(key)，每个键(key)以逗号隔开，及常规 SQL 的 AND 条件。
语法格式如下：
>db.col.find({key1:value1, key2:value2}).pretty()
实例
以下实例通过 by 和 title 键来查询 菜鸟教程 中 MongoDB 教程 的数据
> db.col.find({"by":"菜鸟教程", "title":"MongoDB 教程"}).pretty()
以上实例中类似于 WHERE 语句：WHERE by='菜鸟教程' AND title='MongoDB 教程'
MongoDB OR 条件
MongoDB OR 条件语句使用了关键字 $or,语法格式如下：
>db.col.find(
   {
      $or: [
	     {key1: value1}, {key2:value2}
      ]
   }
).pretty()
实例
以下实例中，我们演示了查询键 by 值为 菜鸟教程 或键 title 值为 MongoDB 教程 的文档。
>db.col.find({$or:[{"by":"菜鸟教程"},{"title": "MongoDB 教程"}]}).pretty()
*/

//存储下拉的列表用的
type MenuData map[DB]Collections

type NodeCollection struct {
	Nodename        string
	NLlog           string
	NLErrlog        string
	Atslog          string
	AtsErrlog       string
	HttpsNLlog      string
	HttpsErrorNllog string
}

type DB string
type Collections []string
