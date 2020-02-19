package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"webapp/bookstore/dao"
	"webapp/bookstore/model"
	"webapp/bookstore/utils"
)

//QueryCart 购物车信息handler
func QueryCart(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.QueryCart(userID)
	fmt.Println(cart)
	t := template.Must(template.ParseFiles("../view/pages/cart/cart.html"))
	if cart.UserID > 0 {
		fmt.Println("测试1")
		session.Cart = cart
		t.Execute(w, session)
	} else {
		fmt.Println("测试2")
		t.Execute(w, session)
	}
}

//UpdateCart 修改购物车
func UpdateCart(w http.ResponseWriter, r *http.Request) {
	cartItemID, _ := strconv.Atoi(r.FormValue("cartItemID"))
	bookCount, _ := strconv.Atoi(r.FormValue("bookCount"))
	_, session := dao.IsLogin(r)
	cart, err := dao.QueryCart(session.UserID)
	if err != nil {
		fmt.Println("UpdateCart：QueryCart失败", err)
	}
	fmt.Println("1", cart)
	cartItems := cart.CartItems
	for _, v := range cartItems {
		if v.ID == cartItemID {
			fmt.Println("v2", v)
			v.Count = bookCount
			v.Amount = float64(v.Count) * v.Book.Price
			fmt.Println("v2", v)
			err1 := dao.UpdateCartItem(v)
			if err1 != nil {
				fmt.Println("UpdateCart:UpdateCartItem失败", err1)
			}

		}
	}
	cart.TotalCount = cart.GetTotalCount()
	cart.TotalAmount = cart.GetTotalAmount()
	fmt.Println("2", cart)
	err2 := dao.UpdateCart(cart)
	if err2 != nil {
		fmt.Println("UpdateCart:UpdateCart失败", err2)
	}
	cart, _ = dao.QueryCart(session.UserID)
	fmt.Println("3", cart)

	totalCount := cart.TotalCount
	//获取购物车中图书的总金额
	totalAmount := cart.TotalAmount
	var amount float64
	//获取购物车中更新的购物项中的金额小计
	cIs := cart.CartItems
	for _, v := range cIs {
		if cartItemID == v.ID {
			//这个就是我们寻找的购物项，此时获取当前购物项中的金额小计
			amount = v.Amount
		}
	}
	//创建Data结构
	data := model.Data{
		Amount:      amount,
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
	}
	json, _ := json.Marshal(data)
	w.Write(json)
}

//AddCart 添加购物
func AddCart(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookID")
	flag, session := dao.IsLogin(r)
	if flag {
		cart, err := dao.QueryCart(session.UserID)
		if err != nil {
			fmt.Println("AddBookCart:失败，没找到cart", err)
		}
		book, err1 := dao.QueryBookByID(bookID)
		if err1 != nil {
			fmt.Println("AddBookCart:没有找到图书信息", err1, book)
		}
		//有购物车
		if cart.UserID > 0 {
			cartItem0, err := dao.QueryCartItem(bookID, cart.ID)
			if cartItem0.ID > 0 {
				fmt.Println("购物车中有该项", err, cartItem0)
				cts := cart.CartItems
				for _, v := range cts {
					if v.Book.ID == cartItem0.Book.ID {
						v.Count = v.Count + 1
						v.Amount = float64(v.Count) * v.Book.Price
						dao.UpdateCartItem(v)
					}
				}
			} else {
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					Amount: book.Price * 1,
					CartID: cart.ID,
				}
				dao.AddCartItem(cartItem)
				cart.CartItems = append(cart.CartItems, cartItem)
				fmt.Println(cartItem)
			}
		} else {
			cartID := utils.CreateUUID()
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				Amount: book.Price * 1,
				CartID: cartID,
			}
			cartItems = append(cartItems, cartItem)
			cart1 := &model.Cart{
				ID:        cartID,
				CartItems: cartItems,
				UserID:    session.UserID,
			}
			cart1.TotalCount = cart1.GetTotalCount()
			cart1.TotalAmount = cart1.GetTotalAmount()
			dao.AddCart(cart1)
			dao.AddCartItem(cartItem)
		}
		w.Write([]byte("您刚刚将" + book.Title + "添加到了购物车！"))
	} else {
		//没有登录
		w.Write([]byte("请先登录！"))
	}

}

//DeleteCartItem 删除购物车中的东西handler
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID, _ := strconv.Atoi(r.FormValue("cartItemID"))
	_, session := dao.IsLogin(r)
	cart, err := dao.QueryCart(session.UserID)
	if err != nil {
		fmt.Println("DeleteCartItem:QueryCart失败", err)
	}
	for k, v := range cart.CartItems {
		if v.ID == cartItemID {
			cart.CartItems = append(cart.CartItems[:k], cart.CartItems[k+1:]...)
			err1 := dao.DeleteCartItem(cartItemID)
			if err1 != nil {
				fmt.Println("DeleteCartItem:DeleteCartItem失败", err1)
			}
		}
	}
	cart.TotalCount = cart.GetTotalCount()
	cart.TotalAmount = cart.GetTotalAmount()
	err2 := dao.UpdateCart(cart)
	if err2 != nil {
		fmt.Println("DeleteCartItem:UpdateCart失败", err2)
	}
	QueryCart(w, r)
}
