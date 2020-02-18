package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"webapp/bookstore/dao"
	"webapp/bookstore/model"
)

//SearchBooks 搜索图书
func SearchBooks(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("search")
	books, err := dao.QueryBookByTitle(title)
	if err != nil {
		fmt.Println("QueryBookByTitle:数据库检索失败,书名:", title, err)
		return
	}
	for _, k := range books {
		if k.ID == 0 {
			fmt.Println("图书没有找到,书名:", title)
			t := template.Must(template.ParseFiles("../view/index.html"))
			t.Execute(w, books)
			return
		}
	}
	t := template.Must(template.ParseFiles("../view/index.html"))
	t.Execute(w, books)
}

/*图书后台管理功能*/

//GetBooks 图书后台管理查询所有图书
func GetBooks(w http.ResponseWriter, r *http.Request) {
	book, err := dao.QueryAllBooks()
	if err != nil {
		fmt.Println("GetBooks：数据库操作失败", err)
		t := template.Must(template.ParseFiles("../view/pages/manager/book_manager.html"))
		t.Execute(w, book)
	} else {
		t := template.Must(template.ParseFiles("../view/pages/manager/book_manager.html"))
		t.Execute(w, book)
	}
}

//UpdateBook 传递要修改图书的信息
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("bookID")
	book, err := dao.QueryBookByID(id)
	if err != nil {
		fmt.Println("失败:", err)
	} else {
		t := template.Must(template.ParseFiles("../view/pages/manager/book_edit.html"))
		if book.ID > 0 {
			t.Execute(w, book)
		} else {
			t.Execute(w, "")
		}
	}
}

//UpdateOrAddBook 修改或添加图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	bookID, _ := strconv.ParseInt(r.PostFormValue("bookId"), 10, 0)
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 32)
	sales, _ := strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
	stock, _ := strconv.ParseInt(r.PostFormValue("stock"), 10, 0)
	imgFile, imgHead, imgErr := r.FormFile("newimgpath")
	imgPath := r.PostFormValue("oldimgpath")
	if imgFile != nil {
		fmt.Println(imgFile, imgHead.Filename, imgErr)
		dstf, err := os.OpenFile("E:\\Go_WorkSpaces\\src\\webapp\\bookstore\\view\\static\\img\\"+imgHead.Filename, os.O_CREATE, os.ModePerm)
		defer dstf.Close()
		_, err2 := io.Copy(dstf, imgFile)
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println(err)
		imgPath = "/static/img/" + imgHead.Filename
		defer dstf.Close()
	}
	fmt.Println(imgPath)
	book := &model.Books{
		ID:             int(bookID),
		Title:          title,
		Author:         author,
		Price:          price,
		Sales:          int(sales),
		Stock:          int(stock),
		Classification: "",
		Publisher:      "",
		ImgPath:        imgPath,
		Ebook:          false,
	}
	if book.ID > 0 {
		err := dao.UpdateBook(book)
		if err != nil {
			fmt.Println("失败b ")
		}

	} else {
		err := dao.AddBook(book)
		if err != nil {
			fmt.Println("失败a ")
		}
	}
	books, _ := dao.QueryAllBooks()
	t := template.Must(template.ParseFiles("../view/pages/manager/book_manager.html"))
	t.Execute(w, books)
}
