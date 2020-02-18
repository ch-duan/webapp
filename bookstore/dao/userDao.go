package dao

import (
	"fmt"
	"log"
	"webapp/bookstore/model"
	"webapp/bookstore/utils"
)

//AddUser 向数据库中添加用户并返回结果
func AddUser(username string, password string, email string, phonenum string) error {
	sql := "insert into users(username,password,email,phonenum) values(?,?,?,?)"
	stmt, err := utils.Db.Prepare(sql)
	if err != nil {
		fmt.Println("AddUser:添加用户SQL预处理失败", err)
		return err
	}
	defer stmt.Close()
	_, err1 := stmt.Exec(username, password, email, phonenum)
	if err1 != nil {
		fmt.Println("AddUser:添加用户失败", err1)
		return err1
	}
	fmt.Println("AddUser:添加用户成功")
	utils.DbIDUpdate("users")
	return nil
}

//DeleteUser 删除用户
func DeleteUser(user *model.User) error {
	sql := "delete from users where id=? and username=?"
	result, err := utils.Db.Exec(sql, user.ID, user.UserName)
	if err != nil {
		fmt.Println("DeleteUser:删除用户失败", err)
		return err
	}
	rowNum, _ := result.RowsAffected()
	fmt.Println("DeleteUser:删除用户成功", rowNum, user.Tostring())
	utils.DbIDUpdate("users")
	return nil
}

//UpdateUser 更新用户
func UpdateUser(user *model.User) error {
	sql := "update users set username=? ,password=? ,email=? ,phonenum=? where id=?"
	_, err := utils.Db.Exec(sql, user.UserName, user.PassWord, user.Email, user.PhoneNum, user.ID)
	if err != nil {
		fmt.Println("UpdateUser:更新用户失败", err)
		return err
	}
	fmt.Println("UpdateUser:更新用户成功", user.Tostring())
	utils.DbIDUpdate("users")
	return nil
}

//QueryAll 查询user所有数据
func QueryAll() ([]*model.User, error) {
	sql := "select * from users"
	var users []*model.User
	row, err := utils.Db.Query(sql)
	if err != nil {
		fmt.Println("QueryAll:查询所有用户数据失败", err)
		return nil, err
	}
	fmt.Println("QueryAll:查询所有用户数据成功：")
	defer row.Close()
	for row.Next() {
		user := &model.User{}
		_, _, values, _ := utils.AllValues(user)
		row.Scan(values...)
		users = append(users, user)
	}
	for c, v := range users {
		fmt.Println(c, v)
	}
	return users, nil
}

//CheckUser 检查用户名密码是否正确
func CheckUser(username string, password string) (*model.User, error) {
	sql := "select * from users where username=? and password=?"

	row := utils.Db.QueryRow(sql, username, password)
	user := &model.User{}
	_, _, values, _ := utils.AllValues(user)
	row.Scan(values...)
	log.Println("CheckUser:", user)
	fmt.Println("CheckUser:", user)
	return user, nil

}

//QueryUserName 根据用户名查询
func QueryUserName(username string) (*model.User, error) {
	sql := "select * from users where username=?"
	row := utils.Db.QueryRow(sql, username)
	user := &model.User{}
	_, _, values, _ := utils.AllValues(user)
	row.Scan(values...)
	fmt.Println("QueryUserName:", user.Tostring())
	return user, nil

}
