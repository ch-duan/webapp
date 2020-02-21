package dao

import (
	"log"
	"net/http"
	"webapp/bookstore/model"
	"webapp/bookstore/utils"
)

//AddSession 添加会话
func AddSession(session *model.Session) error {
	sql := "insert into sessions values(?,?,?)"
	_, err := utils.Db.Exec(sql, session.SessionID, session.Username, session.UserID)
	if err != nil {
		log.Println("AddSession:失败", err)
		return err
	}
	return nil
}

//DeleteSession 删除会话
func DeleteSession(sessionID string) error {
	sql := "delete from sessions where sessionid=?"
	_, err := utils.Db.Exec(sql, sessionID)
	if err != nil {
		log.Println("DeleteSession:失败", err)
		return err
	}
	return nil
}

//QuerySession 查询会话
func QuerySession(sessionID string) (*model.Session, error) {
	sql := "select sessionid,username,userid from sessions where sessionid=?"
	row := utils.Db.QueryRow(sql, sessionID)
	session := &model.Session{}
	row.Scan(&session.SessionID, &session.Username, &session.UserID)
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
