package utils

import (
	"database/sql"
	"log"

	//mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

var (
	//Db is sql
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:1q2w3e4r@/bookstore?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err.Error())
	}

}

//DbIDUpdate 更新自增ID重新排序
func DbIDUpdate(table string) {
	sql := "ALTER TABLE " + table + " DROP id"
	_, err := Db.Exec(sql)
	if err != nil {
		log.Println("删除自增主键失败", err)
	} else {
	}
	sql = "ALTER TABLE " + table + " ADD id int not null first"
	_, err1 := Db.Exec(sql)
	if err1 != nil {
		log.Println("添加字段失败", err1)
	} else {
	}
	sql = "ALTER TABLE " + table + " MODIFY COLUMN id int not null AUTO_INCREMENT, ADD PRIMARY KEY(id)"
	_, err2 := Db.Exec(sql)
	if err2 != nil {
		log.Println("更新自增主键失败", err2)
	} else {
	}
	log.Println("更新自增ID成功")
}
