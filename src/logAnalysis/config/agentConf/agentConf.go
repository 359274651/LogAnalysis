package agentConf

import (
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	//"database/sql"
	"fmt"

	"log"
)

type Config struct {
	Agentname string
	NLog      NginxLog `toml:"nginxLog"`
	AtsLog    AtsLog   `toml:"atsLog"`
	MysqlConf Mysql    `toml:"mysql"`
	MongoC    MO       `toml:"mongo"`
}

type NginxLog struct {
	Title             string
	Separator         string
	Filterconditions  string
	NginxAcessLogPath string
	NginxErrorLogPath string
	HttpsNLlog        string
	HttpsNLErrorlog   string
}

type AtsLog struct {
	Title            string
	Separator        string
	Filterconditions string
	AtsAcessLogPath  string
	AtsErrorLogPath  string
}

type Mysql struct {
	Dbhost     string
	Dbport     string
	Dbuser     string
	Dbpassword string
	Dbname     string
}

type MO struct {
	Dbhost     string
	Dbport     string
	Dbuser     string
	Dbpassword string
	Dbname     string
	Nodedb     string //用来存储节点信息的数据库包括节点名称，节点对应的日志路径
}

type InfluxDb struct {
	Addr     string
	Username string
	Password string
}

func (thisgo *Mysql) ToString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", thisgo.Dbuser, thisgo.Dbpassword, thisgo.Dbhost, thisgo.Dbport, thisgo.Dbname)
}

func ReadConfig(filepath string) *Config {
	var AllConfig Config
	if _, err := toml.DecodeFile(filepath, &AllConfig); err != nil {
		log.Fatalln(err)
	}
	return &AllConfig
}
