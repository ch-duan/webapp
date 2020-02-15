package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
	"webapp/web_db/controller"
	"webapp/web_db/model"

	_ "github.com/go-sql-driver/mysql"
)

//MyHandler is a struct.
type MyHandler struct{}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!", r.URL.Path, r.Host)
	fmt.Fprintln(w, r.Header)
	fmt.Fprintln(w, r.URL.RawQuery)
	// len := r.ContentLength
	// body := make([]byte, len)
	// r.Body.Read(body)
	// fmt.Fprintln(w, string(body))
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func testhandler(w http.ResponseWriter, r *http.Request) {
	user := &model.User{
		ID:       1,
		Username: "abc",
		Password: "123",
		Email:    "abc@Gmail.com",
	}
	json, _ := json.Marshal(user)
	w.Write(json)
}

func myBookStore(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../bookstore/index.html"))
	t.Execute(w, "")
}

func main() {
	str, _ := os.Getwd()
	fmt.Println("main", str)
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/testJson", testhandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../bookstore/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("../bookstore/pages"))))
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/queryUserName", controller.QueryUserName)
	http.HandleFunc("/mybookstore", myBookStore)
	http.ListenAndServe(":8080", nil)
}
