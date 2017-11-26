package entities

import (
    _ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"database/sql"
)

var myEngine *xorm.Engine

func init() {
	engine, err := xorm.NewEngine("mysql","root:root@tcp(127.0.0.1:3308)/test?charset=utf8&parseTime=true")
	if err != nil {
        panic(err)
	}
	myEngine = engine
	myEngine.SetMapper(core.SameMapper{})
}

// SQLExecer interface for supporting xorm.Engine and sql.Tx to do sql statement
type SQLExecer interface {
	//Exec(query string, args ...interface{}) (sql.Result, error)
	//Prepare(query string) (*sql.Stmt, error)
    //Query(query string, args ...interface{}) (*sql.Rows, error)
	//QueryRow(query string, args ...interface{}) *sql.Row
	Exec(sql string, args ...interface{}) (sql.Result, error)
	Rows(bean interface{}) (*xorm.Rows, error)
	//Insert(beans ...interface{}) (int64, error)
	//Query(sqlorArgs ...interface{}) (resultsSlice []map[string][]byte, err error)
	Where(query interface{}, args ...interface{}) *xorm.Session
}

// DaoSource Data Access Object Source
type DaoSource struct {
    // if DB, each statement execute sql with random conn.
    // if Tx, all statements use the same conn as the Tx's connection
    SQLExecer
}


func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}