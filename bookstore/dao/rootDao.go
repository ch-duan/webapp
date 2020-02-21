package dao

import (
	"webapp/bookstore/model"
	"webapp/bookstore/utils"
)

//CheckRoot 检查用户名密码是否正确
func CheckRoot(username string, password string) (*model.Root, error) {
	sql := "select * from roots where username=? and password=?"
	row := utils.Db.QueryRow(sql, username, password)
	root := &model.Root{}
	values := utils.StructFieldSlice(root)
	row.Scan(values...)
	return root, nil
}
