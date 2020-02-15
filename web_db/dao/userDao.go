package dao

import (
	"fmt"
	"log"
	"webapp/web_db/model"
	"webapp/web_db/utils"
)

//DbIDUpdate 更新自增ID重新排序
func DbIDUpdate() {
	sql := "ALTER TABLE users DROP id"
	_, err := utils.Db.Exec(sql)
	if err != nil {
		fmt.Println("删除自增主键失败", err)
	} else {
	}
	sql = "ALTER TABLE users ADD id int not null first"
	_, err1 := utils.Db.Exec(sql)
	if err1 != nil {
		fmt.Println("添加字段失败", err1)
	} else {
	}
	sql = "ALTER TABLE users MODIFY COLUMN id int not null AUTO_INCREMENT, ADD PRIMARY KEY(id)"
	_, err2 := utils.Db.Exec(sql)
	if err2 != nil {
		fmt.Println("更新自增主键失败", err2)
	} else {
	}
}

//AddUser 向数据库中添加用户并返回结果
func AddUser(username string, password string, email string) error {
	sql := "insert into users(username,password,email) values(?,?,?)"
	stmt, err := utils.Db.Prepare(sql)
	if err != nil {
		fmt.Println("添加用户SQL预处理失败", err)
		return err
	}
	defer stmt.Close()
	_, err1 := stmt.Exec(username, password, email)
	if err1 != nil {
		fmt.Println("添加用户失败", err1)
		return err1
	}
	fmt.Println("添加用户成功")
	DbIDUpdate()
	return nil
}

//DeleteUser 删除用户
func DeleteUser(user *model.User) error {
	sql := "delete from users where id=? and username=?"
	result, err := utils.Db.Exec(sql, user.ID, user.Username)
	if err != nil {
		fmt.Println("删除用户失败", err)
		return err
	}
	rowNum, _ := result.RowsAffected()
	fmt.Println("删除用户成功", rowNum, user.Tostring())
	DbIDUpdate()
	return nil
}

//UpdateUser 更新用户
func UpdateUser(user *model.User) error {
	sql := "update users set username=? ,password=? ,email=? where id=?"
	_, err := utils.Db.Exec(sql, user.Username, user.Password, user.Email, user.ID)
	if err != nil {
		fmt.Println("更新用户失败", err)
		return err
	}
	// rowNum,_:=result.LastInsertId()
	fmt.Println("更新用户成功", user.Tostring())
	DbIDUpdate()
	return nil
}

//QueryAll 查询user所有数据
func QueryAll() ([]*model.User, error) {
	sql := "select * from users"
	var users []*model.User
	row, err := utils.Db.Query(sql)
	if err != nil {
		fmt.Println("查询所有数据失败", err)
		return users, err
	}
	fmt.Println("查询所有数据成功：", err)
	defer row.Close()
	for row.Next() {
		user := &model.User{}
		row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
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
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	log.Println("result:", user)
	fmt.Println("result:", user)
	return user, nil

}

//QueryUserName 根据用户名查询
func QueryUserName(username string) (*model.User, error) {
	sql := "select * from users where username=?"
	row := utils.Db.QueryRow(sql, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	fmt.Println("result:", user.Tostring())
	return user, nil

}
