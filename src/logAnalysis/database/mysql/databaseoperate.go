package mysql

import (
	"database/sql"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	. "logAnalysis/CommonLibrary"
)

//OperateMysql  操作mysql 的数据操作 ，包括连接和关闭连接 以及常见的CURD
type OperateMysql struct {
	Sqluri string
	db     *sql.DB
}

func (om *OperateMysql) Init() {
	db, err := sql.Open("mysql", om.Sqluri)
	CheckError(err)
	om.db = db
}

// 插入数据 并返回id和error
func (om *OperateMysql) InsertData(insertstr string, args ...interface{}) (int64, error) {
	stmt, err := om.db.Prepare(insertstr)
	defer stmt.Close()
	CheckError(err)
	res, errres := stmt.Exec(args...)
	//fmt.Println(len(args), args)

	CheckError(errres)
	return res.LastInsertId()

}

// 更新数据 通过
func (om *OperateMysql) UpdateData(insertstr string, args ...interface{}) (int64, error) {
	stmt, err := om.db.Prepare(insertstr)
	defer stmt.Close()
	CheckError(err)
	res, errres := stmt.Exec(args)
	CheckError(errres)
	return res.RowsAffected()

}

// 查询数据数据 通过
func (om *OperateMysql) QueryData(insertstr string, v ...interface{}) (*sql.Rows, error) {
	return om.db.Query(insertstr)

}

// 删除数据数据 通过
func (om *OperateMysql) DeleteData(insertstr string, v ...interface{}) (int64, error) {
	stmt, err := om.db.Prepare(insertstr)
	defer stmt.Close()
	CheckError(err)
	res, errres := stmt.Exec(v)
	CheckError(errres)
	return res.RowsAffected()
}

// 关闭DB连接
func (om *OperateMysql) CloseDB() {
	om.db.Close()

}
