package db

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"

	"database/sql"
)

// sql 基本功能，执行sql语句，映射结果到结构体

var source = "root:vspnmysql123@tcp(192.168.0.12:3312)/battlesothebys?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
}

func Exec(sqlStr string, args ...interface{}) (int64, error) {
	res, err := db.Exec(sqlStr, args...)
	if err != nil {
		return 0, errors.New()
	}

	return 0, nil
}

func QueryRows(v interface{}, sqlStr string, args ...interface{}) error {

	return nil
}

func QueryRow(v interface{}, sqlStr string, args ...interface{}) error {

	return nil
}
