package utils

import (
	"log"
	"strings"
)

//SQLExecInsert SQL插入接口
func SQLExecInsert(v interface{}, table string) error {
	values := StructFieldSlice(v)
	sql := SQLInsertAll(v, table)
	log.Println(v, sql, table)
	_, err := Db.Exec(sql, values...)
	if err != nil {
		log.Println("SQLExecInsert:数据库操作错误", table, err)
		return err
	}
	return nil
}

//SQLExecQuery 根据where=value查询table,一次只能一个,返回v，v应该是结构体指针
func SQLExecQuery(v interface{}, table string, value string, where string) (interface{}, error) {
	values := StructFieldSlice(v)
	sql := SQLSelectAll(v, table)
	if where != "" {
		sql = sql + " where " + where + "=?"
		log.Println("SQLExecQuery", sql)
		rows := Db.QueryRow(sql, value)
		rows.Scan(values...)
	} else {
		log.Println("SQLExecQuery", sql)
		rows := Db.QueryRow(sql)
		rows.Scan(values...)
	}
	return v, nil
}

//SQLExecUpdate 根据where=value修改table,一次只能一个，v应该是结构体指针
func SQLExecUpdate(v interface{}, table string, value string, where string) error {
	values := StructFieldSlice(v)
	values = append(values, value)
	sql := SQLUpdateAll(v, table) + " where " + where + "=?"
	_, err := Db.Exec(sql, values...)
	if err != nil {
		log.Println("SQLExecUpdate:失败", err)
		return err
	}
	return nil
}

//SQLExecDelete 根据where=value删除table,一次只能一个,返回v，v应该是结构体指针
func SQLExecDelete(table string, value string, where string) error {
	sql := "delete from " + table + " where " + where + "=?"
	log.Println(sql)
	_, err := Db.Exec(sql, value)
	if err != nil {
		log.Println("SQLExecDelete:删除用户失败", err)
		return err
	}
	return nil
}

//SQLSelectAll 返回一个查询语句，查询结构体内所有字段，没有条件
func SQLSelectAll(v interface{}, table string) string {
	fields := StructField(v)
	sql := "select " + strings.Join(fields, ",") + " from " + table
	log.Println(sql)
	return sql
}

//SQLInsertAll 返回一个插入语句去除第一个字段，请把主键ID放第一个字段，插入结构体内所有字段，没有条件
func SQLInsertAll(v interface{}, table string) string {
	fields := StructField(v)
	// fields = append(fields[1:])
	length := len(fields)
	sql := "insert into " + table + " (" + strings.Join(fields, ",") + ") values(?"
	length = length - 1
	for {
		if length == 0 {
			break
		}
		sql += ",?"
		length = length - 1
	}
	sql += ")"
	log.Println(sql)
	return sql
}

//SQLUpdateAll 返回一个更新语句，更新结构体内所有字段，没有条件
func SQLUpdateAll(v interface{}, table string) string {
	fields := StructField(v)
	//将id放到最后
	sql := "update " + table + " set " + strings.Join(fields, "=?, ") + "=? "
	log.Println(sql)
	return sql
}
