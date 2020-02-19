package dao

import (
	"fmt"
	"webapp/bookstore/model"
	"webapp/bookstore/utils"
)

//QueryPageBooks 分页获取图书
func QueryPageBooks(pageNum int, pageSize int) (*model.Pages, error) {
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
	fields, _, _, _ := utils.AllValues(&model.Books{})
	sql = "select " + fields + " from books limit ?,?"
	rows, err := utils.Db.Query(sql, (pageNum-1)*pageSize, pageSize)
	if err != nil {
		fmt.Println("QueryPageBooks:数据库操作失败", err)
		return nil, err
	}
	var books []*model.Books
	for rows.Next() {
		book := &model.Books{}
		_, _, values, _ := utils.AllValues(book)
		rows.Scan(values...)
		books = append(books, book)
	}
	page := &model.Pages{
		Books:        books,
		PageNum:      pageNum,
		PageSize:     pageSize,
		TotalPageNum: totalPageNum,
		TotalRecurd:  totalRecurd,
	}
	return page, nil
}

//QueryAllBooks 获取所有图书
func QueryAllBooks() ([]*model.Books, error) {
	var books []*model.Books
	sqlFields, _, _, _ := utils.AllValues(&model.Books{})
	sql := "select " + sqlFields + " from books"
	rows, err := utils.Db.Query(sql)
	if err != nil {
		fmt.Println("GetBooks:数据库读取失败", err)
		return nil, err
	}
	for rows.Next() {
		book := &model.Books{}
		_, _, values, _ := utils.AllValues(book)
		rows.Scan(values...)
		books = append(books, book)
	}
	return books, nil
}

//QueryBookByID 根据id查询书籍信息
func QueryBookByID(id string) (*model.Books, error) {
	book := &model.Books{}
	sqlFields, _, values, _ := utils.AllValues(book)
	sql := "select " + sqlFields + " from books where id=?"
	rows := utils.Db.QueryRow(sql, id)
	rows.Scan(values...)
	return book, nil
}

//QueryBookByTitle 根据书名查询书籍信息
func QueryBookByTitle(title string, pageNum int, pageSize int) (*model.Pages, error) {
	var books []*model.Books
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
	sqlFields, _, _, _ := utils.AllValues(&model.Books{})
	sql = "select " + sqlFields + " from books where title like '%" + title + "%' limit ?,?"
	rows, err := utils.Db.Query(sql, (pageNum-1)*pageSize, pageSize)
	if err != nil {
		fmt.Println("QueryBookByTitle:数据库操作失败", err)
		return nil, err
	}
	for rows.Next() {
		book := &model.Books{}
		_, _, values, _ := utils.AllValues(book)
		rows.Scan(values...)
		books = append(books, book)
	}
	page := &model.Pages{
		Books:        books,
		PageNum:      pageNum,
		PageSize:     pageSize,
		TotalPageNum: totalPageNum,
		TotalRecurd:  totalRecurd,
	}
	return page, nil
}

//AddBook 添加图书信息
func AddBook(book *model.Books) error {
	fields, _, values, length := utils.AllValues(book)
	sql := "insert into books (" + fields + ") values(?"
	length = length - 1
	for {
		if length == 0 {
			break
		}
		sql += ",?"
		length = length - 1
	}
	sql += ");"
	_, err := utils.Db.Exec(sql, values...)
	if err != nil {
		fmt.Println("AddBook:数据库操作错误", err)
		return err
	}
	return nil
}

//DeleteBook 删除图书
func DeleteBook(title string) error {
	sql := "delete from books where title=?"
	fmt.Println(sql, title)
	fmt.Println(title)
	_, err := utils.Db.Exec(sql, title)
	if err != nil {
		fmt.Println("DeleteBook:删除失败", err)
		return err
	}
	fmt.Println("DeleteBook:删除成功", title)
	return nil
}

//UpdateBook 修改图书
func UpdateBook(book *model.Books) error {
	_, fields2, values, _ := utils.AllValues(book)
	//将id放到最后
	s := values[0:1]
	values = append(values[1:], s...)
	sql := "update books set " + fields2 + "where id=?"
	result, err := utils.Db.Exec(sql, values...)
	if err != nil {
		fmt.Println("UpdateBook:失败", result, err)
		return err
	}
	return nil
}
