package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
	"webapp/bookstore/controller"
	"webapp/bookstore/dao"

	_ "github.com/go-sql-driver/mysql"
)

func myBookStore(w http.ResponseWriter, r *http.Request) {
	books, err := dao.QueryAllBooks()
	if err != nil {
		fmt.Println("myBookStore:失败，数据库操作失败", err)
	}
	for _, k := range books {
		fmt.Println(k)
	}
	t := template.Must(template.ParseFiles("../view/index.html"))
	t.Execute(w, books)
}

func main() {
	str, _ := os.Getwd()
	fmt.Println("mainPath:", str)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../view/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("../view/pages"))))
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/queryUserName", controller.QueryUserName)
	http.HandleFunc("/search", controller.SearchBooks)
	http.HandleFunc("/bookManager", controller.GetBooks)
	http.HandleFunc("/upDateBook", controller.UpdateBook)
	http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBook)
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	http.HandleFunc("/myBookStore", controller.GetPageBooks)
	http.HandleFunc("/cart", controller.QueryCart)
	http.HandleFunc("/addCart", controller.AddCart)
	http.HandleFunc("/updateCartItem", controller.UpdateCart)
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	fmt.Println("程序开始监听8080：")
	http.ListenAndServe(":8080", nil)
}
