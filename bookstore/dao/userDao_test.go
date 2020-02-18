package dao

import (
	"fmt"
	"testing"
	"webapp/bookstore/model"
)

func Test(t *testing.T) {
	fmt.Println("开始测试")
	// // t.Run("测试1", checkusernameTest)
	// user := &model.User{
	// 	ID:       1,
	// 	Username: "1111",
	// 	Password: "123",
	// 	Email:    "www@Gmail.com",
	// }
	// UpdateUser(user)
	book := &model.Books{
		ID:             42,
		Title:          "人间失格",
		Author:         "太宰治",
		Price:          6,
		Sales:          100,
		Stock:          100,
		Classification: "0",
		Publisher:      "0",
		ImgPath:        "static/img/default.jpg",
		Ebook:          true,
	}
	UpdateBook(book)
	QueryBookByID("42")
	fmt.Println("结束测试")
}
