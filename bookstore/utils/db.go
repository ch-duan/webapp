package utils

import (
	"database/sql"

	//mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

var (
	//Db is sql
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:1q2w3e4r@/bookstore")
	if err != nil {
		panic(err.Error())
	}

}
