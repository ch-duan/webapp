package dao

import (
	"log"
	"webapp/bookstore/model"
	"webapp/bookstore/utils"
)

//AddCart 添加购物车
func AddCart(cart *model.Cart) error {
	sql := "insert into carts (id,totalcount,totalamount,userid) values(?,?,?,?)"
	_, err := utils.Db.Exec(sql, cart.ID, cart.TotalCount, cart.TotalAmount, cart.UserID)
	if err != nil {
		log.Println("AddCart:失败", err)
		return err
	}
	return nil
}

//QueryCart 查询用户购物车
func QueryCart(userID int) (*model.Cart, error) {
	cart := &model.Cart{}
	sql := "select id,totalcount,totalamount,userid from carts where userid=?"
	row := utils.Db.QueryRow(sql, userID)
	row.Scan(&cart.ID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	cartItems, _ := QueryCartItems(cart.ID)
	cart.CartItems = cartItems
	return cart, nil
}

//UpdateCart 修改购物车
func UpdateCart(cart *model.Cart) error {
	sql := "update carts set totalcount=?,totalamount=? where userid=?"
	_, err := utils.Db.Exec(sql, cart.TotalCount, cart.TotalAmount, cart.UserID)
	if err != nil {
		log.Println("UpdateCart:修改购物车失败", err)
		return err
	}
	return nil
}

//DeleteCart 修改购物车
func DeleteCart(ID string) error {
	sql := "Delette from carts where id=?"
	_, err := utils.Db.Exec(sql, ID)
	if err != nil {
		log.Println("DeleteCart:删除购物车失败", err)
		return err
	}
	return nil
}
