package model

import (
	"fmt"
)

//User 用户结构体
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

//Tostring 打印user中所有信息
func (user *User) Tostring() string {
	return fmt.Sprintf("%d,%s,%s,%s", user.ID, user.Username, user.Password, user.Email)
}
