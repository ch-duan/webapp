package model

import (
	"fmt"
)

//User 用户结构体
type User struct {
	ID       int
	UserName string //用户名
	PassWord string //密码
	Email    string //邮箱
	PhoneNum string //手机号
}

//Tostring 打印user中所有信息
func (user *User) Tostring() string {
	return fmt.Sprintf("%d,%s,%s,%s,%s", user.ID, user.UserName, user.PassWord, user.Email, user.PhoneNum)
}
