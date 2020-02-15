package dao

import (
	"fmt"
	"webapp/bookstore/model"
	"webapp/bookstore/utils"
)

//GetBooks 获取所有图书
func GetBooks() ([]*model.Books, error) {
	var books []*model.Books
	sqlFields, _ := utils.AllValues(&model.Books{})
	sql := "select " + sqlFields + " from books"
	fmt.Println(sql)
	rows, err := utils.Db.Query(sql)
	if err != nil {
		fmt.Println("获取所有图书失败，数据库读取失败")
		return nil, err
	}
	for rows.Next() {
		book := &model.Books{}
		_, values := utils.AllValues(book)
		rows.Scan(values...)
		books = append(books, book)
		fmt.Println(book.ImgPath)
	}
	for _, k := range books {
		fmt.Println(k, k.ImgPath)
	}
	return books, nil
}
