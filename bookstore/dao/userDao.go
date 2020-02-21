package dao

import (
	"log"
	"strconv"
	"webapp/bookstore/model"
	"webapp/bookstore/utils"
)

//AddUser 向数据库中添加用户并返回结果
func AddUser(user *model.User) error {
	err := utils.SQLExecInsert(user, "users")
	if err != nil {
		log.Println("AddUser:添加用户失败", err)
		return err
	}
	return nil
}

//DeleteUser 删除用户
func DeleteUser(user *model.User) error {
	err := utils.SQLExecDelete("users", strconv.Itoa(user.ID), "id")
	if err != nil {
		return err
	}
	return nil
}

//UpdateUser 更新用户
func UpdateUser(user *model.User) error {
	err := utils.SQLExecUpdate(user, "users", strconv.Itoa(user.ID), "id")
	if err != nil {
		log.Println("UpdateUser:更新用户失败", err)
		return err
	}
	return nil
}

//QueryAll 查询user所有数据
func QueryAll() ([]*model.User, error) {
	sql := "select * from users"
	var users []*model.User
	row, err := utils.Db.Query(sql)
	if err != nil {
		log.Println("QueryAll:查询所有用户数据失败", err)
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		user := &model.User{}
		values := utils.StructFieldSlice(user)
		row.Scan(values...)
		users = append(users, user)
	}
	return users, nil
}

//CheckUser 检查用户名密码是否正确
func CheckUser(username string, password string) (*model.User, error) {
	sql := "select * from users where username=? and password=?"
	row := utils.Db.QueryRow(sql, username, password)
	user := &model.User{}
	values := utils.StructFieldSlice(user)
	row.Scan(values...)
	return user, nil

}

//QueryUserByUsername 根据用户名查询
func QueryUserByUsername(username string) (*model.User, error) {
	user, err := utils.SQLExecQuery(&model.User{}, "users", username, "username")
	if err != nil {
		return nil, err
	}
	return user.(*model.User), nil
}
