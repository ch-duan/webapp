package dao

import (
	"fmt"
	"testing"
	"webapp/web_db/model"
)

func Test(t *testing.T) {
	fmt.Println("开始测试")
	// t.Run("测试1", checkusernameTest)
	user := &model.User{
		ID:1,
		Username: "1111",
		Password: "123",
		Email:    "www@Gmail.com",
	}
	UpdateUser(user)
	fmt.Println("结束测试")
}
