package dao

import (
	"log"
	"strconv"
	"strings"
	"webapp/bookstore/model"
	"webapp/bookstore/utils"
)

//QueryPageBooks 分页获取图书
func QueryPageBooks(pageNum int, pageSize int) (*model.Page, error) {
	sql := "select count(*) from books"
	var totalRecurd int
	var totalPageNum int
	row := utils.Db.QueryRow(sql)
	row.Scan(&totalRecurd)
	if totalRecurd%pageSize == 0 {
		totalPageNum = totalRecurd / pageSize
	} else {
		totalPageNum = totalRecurd/pageSize + 1
	}
	fields := utils.StructField(&model.Book{})
	sql = "select " + strings.Join(fields, ",") + " from books limit ?,?"
	rows, err := utils.Db.Query(sql, (pageNum-1)*pageSize, pageSize)
	if err != nil {
		log.Println("QueryPageBooks:数据库操作失败", err)
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		values := utils.StructFieldSlice(book)
		rows.Scan(values...)
		books = append(books, book)
	}
	page := &model.Page{
		Books:        books,
		PageNum:      pageNum,
		PageSize:     pageSize,
		TotalPageNum: totalPageNum,
		TotalRecurd:  totalRecurd,
	}
	return page, nil
}

//QueryAllBooks 获取所有图书
func QueryAllBooks() ([]*model.Book, error) {
	var books []*model.Book
	sql := utils.SQLSelectAll(&model.Book{}, "books")
	rows, err := utils.Db.Query(sql)
	if err != nil {
		log.Println("GetBooks:数据库读取失败", err)
		return nil, err
	}
	for rows.Next() {
		book := &model.Book{}
		values := utils.StructFieldSlice(book)
		rows.Scan(values...)
		books = append(books, book)
	}
	return books, nil
}

//QueryBookByID 根据id查询书籍信息
func QueryBookByID(id string) (*model.Book, error) {
	book, err := utils.SQLExecQuery(&model.Book{}, "books", id, "id")
	if err != nil {
		log.Println("QueryBookByID:失败", err)
		return nil, err
	}
	return book.(*model.Book), nil
}

//QueryBookByTitle 根据书名查询书籍信息
func QueryBookByTitle(title string, pageNum int, pageSize int) (*model.Page, error) {
	var books []*model.Book
	var totalRecurd int
	var totalPageNum int
	sql := "select count(*) from books where title like '%" + title + "%' "
	row := utils.Db.QueryRow(sql)
	row.Scan(&totalRecurd)
	if totalRecurd%pageSize == 0 {
		totalPageNum = totalRecurd / pageSize
	} else {
		totalPageNum = totalRecurd/pageSize + 1
	}
	sqlFields := utils.StructField(books)
	sql = "select " + strings.Join(sqlFields, ",") + " from books where title like '%" + title + "%' limit ?,?"
	rows, err := utils.Db.Query(sql, (pageNum-1)*pageSize, pageSize)
	if err != nil {
		log.Println("QueryBookByTitle:数据库操作失败", err)
		return nil, err
	}
	for rows.Next() {
		book := &model.Book{}
		values := utils.StructFieldSlice(book)
		rows.Scan(values...)
		books = append(books, book)
	}
	page := &model.Page{
		Books:        books,
		PageNum:      pageNum,
		PageSize:     pageSize,
		TotalPageNum: totalPageNum,
		TotalRecurd:  totalRecurd,
	}
	return page, nil
}

//AddBook 添加图书信息
func AddBook(book *model.Book) error {
	err := utils.SQLExecInsert(book, "books")
	if err != nil {
		log.Println("AddBook:添加图书失败", err)
		return err
	}
	return nil
}

//DeleteBook 删除图书
func DeleteBook(ID string) error {
	err := utils.SQLExecDelete("books", ID, "id")
	if err != nil {
		return err
	}
	return nil
}

//UpdateBook 修改图书
func UpdateBook(book *model.Book) error {
	err := utils.SQLExecUpdate(book, "books", strconv.Itoa(book.ID), "id")
	if err != nil {
		log.Println("UpdateBook:失败", err)
		return err
	}
	return nil
}
