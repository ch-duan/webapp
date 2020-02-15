package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
	"webapp/bookstore/controller"

	_ "github.com/go-sql-driver/mysql"
)

func myBookStore(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../view/index.html"))
	t.Execute(w, "")
}

func main() {
	str, _ := os.Getwd()
	fmt.Println("mainPath:", str)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../view/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("../view/pages"))))
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/queryUserName", controller.QueryUserName)
	http.HandleFunc("/mybookstore", myBookStore)
	http.ListenAndServe(":8080", nil)
}