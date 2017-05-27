package runglobal

import (
	"sync"

	"logAnalysis/config/serverConf"
	"logAnalysis/database/mongo"
)

//全局的配置文件，为了自己的Sb 犯下的错，通用一个配置变量，比较sb过渡一下吧
var GlobalConf *serverConf.Config
var GlobalMongdb *mongo.MgoOp
var PublicLock sync.Mutex = sync.Mutex{}
var PulicRegex map[string]string = map[string]string{}
