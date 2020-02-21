package controller

import (
	"html/template"
	"log"
	"net/http"
	"webapp/bookstore/dao"
	"webapp/bookstore/model"
	"webapp/bookstore/utils"
)

//Register 注册
func Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	phonenum := r.PostFormValue("phoneNum")
	if username == "" || password == "" {
		log.Println("注册失败，用户名为空")
		t := template.Must(template.ParseFiles("../view/pages/user/regist.html"))
		t.Execute(w, "用户名和密码不能为空!")
		return
	}
	user, err := dao.QueryUserByUsername(username)
	if err != nil {
		log.Println("QueryUserByUsername:数据库检索失败", err)
	} else {
		if user.ID > 0 {
			log.Println("QueryUserByUsername注册失败，用户已经存在")
			t := template.Must(template.ParseFiles("../view/pages/user/regist.html"))
			t.Execute(w, "用户名已存在!")
		} else {
			err = dao.AddUser(&model.User{Username: username, Password: password, Email: email, PhoneNum: phonenum})
			if err != nil {
				log.Println("AddUser:注册失败", err)
				t := template.Must(template.ParseFiles("../view/pages/user/regist.html"))
				t.Execute(w, "注册失败")
			} else {
				log.Println("注册成功")
				t := template.Must(template.ParseFiles("../view/pages/user/regist_success.html"))
				user.Username = username
				t.Execute(w, user)
			}
		}
	}
}

//Logout 注销
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookievalue := cookie.Value
		// session,err:=dao.QuerySession(cookievalue)
		dao.DeleteSession(cookievalue)
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	//首页
	GetPageBooks(w, r)
}

//Login 登陆
func Login(w http.ResponseWriter, r *http.Request) {

	flag, _ := dao.IsLogin(r)
	if flag {

	} else {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		user, err := dao.CheckUser(username, password)
		if err != nil {
			log.Println("CheckUser:数据库检索失败", err)
		} else {
			if user.ID > 0 {
				//用户名和密码正确
				uuid := utils.CreateUUID()
				session := &model.Session{
					SessionID: uuid,
					Username:  user.Username,
					UserID:    user.ID,
				}
				dao.AddSession(session)
				cookie := http.Cookie{
					Name:     "user",
					Value:    uuid,
					HttpOnly: true,
					MaxAge:   3600,
				}
				http.SetCookie(w, &cookie)
				t := template.Must(template.ParseFiles("../view/pages/user/login_success.html"))
				t.Execute(w, user)
				log.Println("登陆成功")
			} else {
				log.Println("用户名或密码错误")
				t := template.Must(template.ParseFiles("../view/pages/user/login.html"))
				t.Execute(w, "用户名或密码错误")
			}
		}
	}
}

//QueryUserByUsername 用户名可用验证
func QueryUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user, err := dao.QueryUserByUsername(username)
	if err != nil {
		log.Println("QueryUserByUsername:数据库检索失败,用户名:", username, err)
	} else {
		if user.ID > 0 {
			w.Write([]byte("用户名已经存在!"))
		} else {
			w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
		}
	}

}
