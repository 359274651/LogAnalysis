package serverConf

import (
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	//"database/sql"

	"log"
)

type Config struct {
	ServerAddr string
	NLog       CommonConf `toml:"nginxLog"`
	AtsLog     CommonConf `toml:"atsLog"`
	NELog      CommonConf `toml:"nginxErrorLog"`
	AtsELog    CommonConf `toml:"atsErrorLog"`
	Mongo      MO         `toml:"mongo"`
}

type CommonConf struct {
	Addr      string
	Username  string
	Password  string
	Dbname    string
	Tablename string
}

type MO struct {
	Dbhost     string
	Dbport     string
	Dbuser     string
	Dbpassword string
	Dbname     string
	Nodedb     string //用来存储节点信息的数据库包括节点名称，节点对应的日志路径
}

func (cc *CommonConf) ToString() string {
	return ""
}

func ReadConfig(filepath string) *Config {
	var AllConfig Config
	if _, err := toml.DecodeFile(filepath, &AllConfig); err != nil {
		log.Fatalln(err)
	}
	return &AllConfig
}
