package agentConf

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/BurntSushi/toml"
	//"database/sql"
	"fmt"
	
	"log"
)

type Config struct {
	Agentname string
	NLog      NginxLog `toml:"nginxLog"`
	AtsLog    AtsLog `toml:"atsLog"`
	MysqlConf Mysql `toml:"mysql"`
}

type NginxLog struct {
	Title             string
	Separator         string
	Filterconditions  string
	NginxAcessLogPath string
	NginxErrorLogPath string
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
