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
	var pageNum int64
	if r.FormValue("pageNum") == "" {
		pageNum = 1
	} else {
		pageNum, _ = strconv.ParseInt(r.FormValue("pageNum"), 10, 0)
	}
	page, err := dao.QueryBookByTitle(title, int(pageNum), 4)
	if err != nil {
		fmt.Println("QueryBookByTitle:数据库检索失败,书名:", title, err)
		return
	}

	flag, session := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.UserName = session.UserName
	}
	for _, k := range page.Books {
		if k.ID == 0 {
			fmt.Println("图书没有找到,书名:", title)
			t := template.Must(template.ParseFiles("../view/index.html"))
			t.Execute(w, page)
			return
		}
	}
	t := template.Must(template.ParseFiles("../view/index.html"))
	t.Execute(w, page)
}

//GetPageBooks 分区获取图书handler
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	var pageNum int64
	if r.FormValue("pageNum") == "" {
		pageNum = 1
	} else {
		pageNum, _ = strconv.ParseInt(r.FormValue("pageNum"), 10, 0)
	}
	page, err := dao.QueryPageBooks(int(pageNum), 4)
	flag, session := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.UserName = session.UserName
	}
	t := template.Must(template.ParseFiles("../view/index.html"))
	if err != nil {
		t.Execute(w, "")
		return
	}
	t.Execute(w, page)

}

/*图书后台管理功能*/

//GetBooks 图书后台管理查询所有图书
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := dao.QueryAllBooks()
	t := template.Must(template.ParseFiles("../view/pages/manager/book_manager.html"))
	if err != nil {
		fmt.Println("GetBooks:数据库操作失败", err)
		t.Execute(w, nil)
	}
	t.Execute(w, books)
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
	bookID, _ := strconv.ParseInt(r.PostFormValue("bookID"), 10, 0)
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 32)
	sales, _ := strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
	stock, _ := strconv.ParseInt(r.PostFormValue("stock"), 10, 0)
	imgFile, imgHead, imgErr := r.FormFile("newImgPath")
	if imgErr != nil {
		fmt.Println("UpdateOrAddBook:图片上传失败", imgErr)
	}
	imgPath := r.PostFormValue("oldImgPath")
	if imgFile != nil {
		dstf, err := os.OpenFile("E:\\Go_WorkSpaces\\src\\webapp\\bookstore\\view\\static\\img\\"+imgHead.Filename, os.O_CREATE, os.ModePerm)
		defer dstf.Close()
		if err != nil {
			fmt.Println("UpdateOrAddBook:图片创建失败", err)
		}
		_, err2 := io.Copy(dstf, imgFile)
		if err2 != nil {
			fmt.Println("UpdateOrAddBook:图片复制失败", err2)
		}
		imgPath = "/static/img/" + imgHead.Filename
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
			fmt.Println("UpdateOrAddBook:更新失败", err)
		}

	} else {
		err := dao.AddBook(book)
		if err != nil {
			fmt.Println("UpdateOrAddBook:添加失败", err)
		}
	}
	books, _ := dao.QueryAllBooks()
	t := template.Must(template.ParseFiles("../view/pages/manager/book_manager.html"))
	t.Execute(w, books)
}

//DeleteBook 删除图书handler
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	err := dao.DeleteBook(title)
	if err != nil {
		fmt.Println("DeleteBookhandler:", err)
	}
	GetBooks(w, r)
}
