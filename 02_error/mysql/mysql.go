package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var TestMysql *sql.DB

var (
	dbType       = "mysql"
	address      = ""
	username     = ""
	password     = ""
	charset      = "utf8"
	databaseName = ""
	maxOpenConns = 10
	maxIdleConns = 10
	maxLifetime  = 2 * time.Minute
)

func init() {
	err := InitMysql()
	if err != nil {
		log.Fatalf("InitMysql err:%v", err)
	}
}

func InitMysql() (err error) {
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		address,
		databaseName,
		charset)
	TestMysql, err = sql.Open(dbType, source)
	if err != nil {
		return
	}
	TestMysql.SetMaxOpenConns(maxOpenConns)
	TestMysql.SetMaxIdleConns(maxIdleConns)
	TestMysql.SetConnMaxLifetime(maxLifetime)
	err = TestMysql.Ping()
	return
}
