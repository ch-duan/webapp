package utils

import (
	"log"
	"reflect"
	"strings"
)

//StructFieldSlice 传递的必须是结构体指针,返回结构体字段切片，用于sql执行后的scan，不改变默认顺序，错误或没有值返回空切片
func StructFieldSlice(v interface{}) (values []interface{}) {
	el := reflect.TypeOf(v)
	ul := reflect.ValueOf(v)
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		log.Println("StructFieldSlice: Check type error not Ptr Struct")
		return nil
	}
	if el.Kind() == reflect.Ptr {
		el = el.Elem()
		ul = ul.Elem()
	}
	if el.Kind() != reflect.Struct {
		log.Println("StructFieldSlice: Check type error not Struct")
		return nil
	}
	length := el.NumField()
	for i := 0; i < length; i++ {
		values = append(values, ul.Field(i).Addr().Interface())
	}
	return values
}

//StructField 传递的必须是结构体指针,返回结构体字段切片，用于sql执行后的scan，不改变默认顺序，错误会返回""
func StructField(v interface{}) (fields []string) {
	arr := []string{}
	el := reflect.TypeOf(v)
	if el.Kind() == reflect.Ptr {
		el = el.Elem()
	}
	if el.Kind() != reflect.Struct {
		log.Println("StructField: Check type error not Struct")
		return fields
	}
	length := el.NumField()
	for i := 0; i < length; i++ {
		arr = append(arr, el.Field(i).Name)
	}
	return arr
}

//AllValues 传递的必须是指针,返回结构体字段名加逗号，加=?和每个字段切片及总共几个字段，用于sql select和update字段和给相应结构体赋值
func AllValues(v interface{}) (fields string, fields2 string, values []interface{}, length int) {
	arr := []string{}
	el := reflect.TypeOf(v)
	if el.Kind() == reflect.Ptr {
		el = el.Elem()
	} else {
		log.Println("AllValues:Check type error not Ptr")
		log.Println("AllValues:Check type error not Ptr")
		return
	}
	vl := reflect.ValueOf(v).Elem()
	length = el.NumField()
	for i := 0; i < length; i++ {
		arr = append(arr, el.Field(i).Name)
		values = append(values, vl.Field(i).Addr().Interface())
	}
	fields = "`" + strings.Join(arr, "`,`") + "`"
	arr = arr[1:]
	fields2 = strings.Join(arr, "=?, ") + "=? "
	return fields, fields2, values, length
}
