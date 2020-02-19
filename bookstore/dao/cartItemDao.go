package dao

import (
	"fmt"
	"webapp/bookstore/model"
	"webapp/bookstore/utils"
)

//AddCartItem 添加购物项
func AddCartItem(cartItem *model.CartItem) error {
	sql := "insert into cartitems(count,amount,bookid,cartid) values(?,?,?,?)"
	fmt.Println(cartItem)
	_, err := utils.Db.Exec(sql, cartItem.Count, cartItem.Amount, cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		fmt.Println("AddCartItem:失败", err)
		return err
	}
	return nil
}

//QueryCartItem 查找购物项
func QueryCartItem(bookID string, cartID string) (*model.CartItem, error) {
	sql := "select id,count,amount,cartid from cartitems where bookid=? and cartid=?"
	fmt.Println(sql)
	row := utils.Db.QueryRow(sql, bookID, cartID)
	cartItem := &model.CartItem{}
	row.Scan(&cartItem.CartID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	book, err := QueryBookByID(string(bookID))
	if err != nil {
		fmt.Println("QueryCartIttem:失败", err)
		return nil, err
	}
	cartItem.Book = book
	return cartItem, nil
}

//QueryCartItems 根据cartID查找购物项切片
func QueryCartItems(cartID string) ([]*model.CartItem, error) {
	var cartItems []*model.CartItem
	sql := "select id,count,amount,bookid,cartid from cartitems where cartid=?"
	fmt.Println(cartID, sql)
	rows, err := utils.Db.Query(sql, cartID)
	if err != nil {
		fmt.Println("QueryCartItems:失败err", err)
		return nil, err
	}
	for rows.Next() {
		cartItem := &model.CartItem{}
		var bookID string
		rows.Scan(&cartItem.ID, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		book, err1 := QueryBookByID(bookID)
		if err1 != nil {
			fmt.Println("QueryCartItems:失败err1", err1)
			return nil, err1
		}
		cartItem.Book = book
		fmt.Println("测试", cartItem, book)
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

//UpdateCartItem 修改购物项
func UpdateCartItem(cartItem *model.CartItem) error {
	sql := "update cartitems set count=?,amount=?,bookid=?,cartid=? where id=?"
	fmt.Println(cartItem)
	_, err := utils.Db.Exec(sql, cartItem.Count, cartItem.Amount, cartItem.Book.ID, cartItem.CartID, cartItem.ID)
	if err != nil {
		fmt.Println("UpdateCartItem:失败", err)
		return err
	}
	return nil
}

//DeleteCartItem 删除购物项
func DeleteCartItem(ID int) error {
	sql := "delete from cartitems where id=?"
	_, err := utils.Db.Exec(sql, ID)
	if err != nil {
		fmt.Println("DeleteCartItem:失败", err)
		return err
	}
	return nil
}
