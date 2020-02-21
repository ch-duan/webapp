package controller

import (
	"log"
	"net/http"
	"text/template"
	"webapp/bookstore/dao"
)

//RootLogin 后台管理员验证
func RootLogin(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	root, err := dao.CheckRoot(username, password)
	if err != nil {
		log.Println("CheckUser:数据库检索失败", err)
	}

	if root.ID > 0 {
		log.Println("管理员登陆成功")
		t := template.Must(template.ParseFiles("../view/pages/manager/manager.html"))
		t.Execute(w, root)
	}
	log.Println("用户名或密码错误")
	t := template.Must(template.ParseFiles("../view/pages/user/rootLogin.html"))
	t.Execute(w, "用户名或密码错误")

}
