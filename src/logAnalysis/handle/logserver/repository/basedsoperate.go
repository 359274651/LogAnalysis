package Repository

//基础的数据操作，后面的操作基于这个对象进行封装
import (
	"github.com/influxdata/influxdb/client/v2"

	"logAnalysis/config/serverConf"
	"logAnalysis/database/influxdb"
)

type BaseRepository struct {
	iidb *influxdb.InFluxDBTool
}

func (rps *BaseRepository) InsertPoint(tablename, rp string, tags map[string]string, fields map[string]interface{}) (err error) {

	return rps.iidb.InsertPoint(tablename, rp, tags, fields)
}

func (rps *BaseRepository) Query(cmd string) ([]client.Result, error) {

	return rps.iidb.Query(cmd)

}

func (rps *BaseRepository) QueryData(cmd string, args ...interface{}) ([]client.Result, error) {

	return rps.iidb.QueryData(cmd, args...)
}

func (rps *BaseRepository) CreateRP(rpname, dbname, dura string, replication int) ([]client.Result, error) {

	return rps.iidb.CreateRP(rpname, dbname, dura, replication)
}

func (rps *BaseRepository) ModifyRP(rpname, dbname, dura string, replication int) ([]client.Result, error) {

	return rps.iidb.ModifyRP(rpname, dbname, dura, replication)
}

func (rps *BaseRepository) DropRP(rpname, dbname string) ([]client.Result, error) {

	return rps.iidb.DropRP(rpname, dbname)
}

func (rps *BaseRepository) CreateDB(dbname string) error {

	return rps.iidb.CreateDB(dbname)
}

func (rps *BaseRepository) DropDB(dbname string) error {

	return rps.iidb.DropDB(dbname)
}

func GetRepository(config *serverConf.CommonConf) *BaseRepository {
	return &BaseRepository{influxdb.NewClient(config.Addr, config.Username, config.Password, config.Dbname)}
}

func DeferCloseconn(repository *BaseRepository) {
	return
	//err := repository.Close()
	//if err != nil {
	//	panic(err)
	//}
}
