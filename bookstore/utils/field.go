package utils

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

//AllValues 传递的必须是指针,返回结构体字段名和每个字段切片，用于sql字段和给相应结构体赋值
func AllValues(v interface{}) (fields string, values []interface{}) {
	arr := []string{}
	el := reflect.TypeOf(v)
	if el.Kind() == reflect.Ptr {
		el = el.Elem()
	} else {
		fmt.Println("AllValues:Check type error not Ptr")
		log.Println("AllValues:Check type error not Ptr")
		return
	}
	vl := reflect.ValueOf(v).Elem()
	for i := 0; i < el.NumField(); i++ {
		arr = append(arr, el.Field(i).Name)
		values = append(values, vl.Field(i).Addr().Interface())
	}
	fields = "`" + strings.Join(arr, "`,`") + "`"
	// fmt.Println(fields, values)
	return fields, values
}
