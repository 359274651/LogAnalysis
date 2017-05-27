package logserver

type QueryKey struct {
	DB         `json:"db"` //查询的数据库
	Collection string      `json:"collection"` //mongo Collection  ,一般就是日志文件的名字,可以通过path.Split(path) (dir,name string) 获取
}

//存储查询条件用的结构体
type QueryData struct {
	DB         `json:"db"` //查询的数据库
	Collection string      `json:"collection"` //mongo Collection  ,一般就是日志文件的名字,可以通过path.Split(path) (dir,name string) 获取
	Contidions []Contidion `json:"contidion"`
}

type Contidion struct {
	Key   string `json:"key"`
	Oper  string `json:"oper"` //运算符表示 $lt/$lte/$gt/$gte/$ne，依次等价于</<=/>/>=/!=。
	Value string `json:"value"`
}

/*
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
