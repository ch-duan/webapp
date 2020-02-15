package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"webapp/web_db/dao"
)

//Register 注册
func Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	if username == "" {
		fmt.Println("用户名为空")
		fmt.Println("注册失败")
		t := template.Must(template.ParseFiles("../bookstore/pages/user/regist.html"))
		t.Execute(w, "用户名不能为空!")
		return
	}
	user, err := dao.QueryUserName(username)
	if err != nil {
		fmt.Println("数据库检索失败", err)
	} else {
		if user.ID > 0 {
			fmt.Println("用户已经存在")
			fmt.Println("注册失败")
			t := template.Must(template.ParseFiles("../bookstore/pages/user/regist.html"))
			t.Execute(w, "用户名已存在!")
		} else {
			err = dao.AddUser(username, password, email)
			if err != nil {
				fmt.Println("注册失败")
				t := template.Must(template.ParseFiles("../bookstore/pages/user/regist.html"))
				t.Execute(w, "注册失败")
			} else {
				fmt.Println("注册成功")
				t := template.Must(template.ParseFiles("../bookstore/pages/user/regist_success.html"))
				user.Username = username
				t.Execute(w, user)
			}
		}
	}
}

//Login 登陆验证
func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	str, _ := os.Getwd()
	fmt.Println("login：", str)
	user, err := dao.CheckUser(username, password)
	if err != nil {
		fmt.Println("数据库检索失败", err)
	} else {
		if user != nil {
			if user.ID > 0 {
				//用户名和密码正确
				fmt.Println("登陆成功")
				log.Println(username, password, user)
				t := template.Must(template.ParseFiles("../bookstore/pages/user/login_success.html"))
				t.Execute(w, user)
			} else {
				fmt.Println("用户名或密码错误")
				t := template.Must(template.ParseFiles("../bookstore/pages/user/login.html"))
				t.Execute(w, "用户名或密码错误")
			}
		} else {
			fmt.Println("用户名或密码错误")
			t := template.Must(template.ParseFiles("../bookstore/pages/user/login.html"))
			t.Execute(w, "")
		}
	}
}

//QueryUserName 用户名可用验证
func QueryUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user, err := dao.QueryUserName(username)
	if err != nil {

	} else {
		if user.ID > 0 {
			w.Write([]byte("用户名已经存在!"))
		} else {
			w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
		}
	}

}
