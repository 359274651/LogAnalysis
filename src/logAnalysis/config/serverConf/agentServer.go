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
}

type CommonConf struct {
	Addr      string
	Username  string
	Password  string
	Dbname    string
	Tablename string
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
