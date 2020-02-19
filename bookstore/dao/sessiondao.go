package dao

import (
	"fmt"
	"net/http"
	"webapp/bookstore/model"
	"webapp/bookstore/utils"
)

//AddSession 添加会话
func AddSession(session *model.Session) error {
	sql := "insert into sessions values(?,?,?)"
	_, err := utils.Db.Exec(sql, session.SessionID, session.UserName, session.UserID)
	if err != nil {
		fmt.Println("AddSession:失败", err)
		return err
	}
	fmt.Println("AddSession:添加")
	return nil
}

//DeleteSession 删除会话
func DeleteSession(sessionID string) error {
	sql := "delete from sessions where sessionid=?"
	_, err := utils.Db.Exec(sql, sessionID)
	if err != nil {
		fmt.Println("DeleteSession:失败", err)
		return err
	}
	fmt.Println("DeleteSession:删除成功")
	return nil
}

//QuerySession 查询会话
func QuerySession(sessionID string) (*model.Session, error) {
	sql := "select sessionid,username,userid from sessions where sessionid=?"
	row := utils.Db.QueryRow(sql, sessionID)
	session := &model.Session{}
	row.Scan(&session.SessionID, &session.UserName, &session.UserID)
	return session, nil
}

//IsLogin 判断是否登陆
func IsLogin(r *http.Request) (bool, *model.Session) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookievalue := cookie.Value
		session, _ := QuerySession(cookievalue)
		if session.UserID > 0 {
			return true, session
		}
	}
	//未登陆
	return false, nil
}
