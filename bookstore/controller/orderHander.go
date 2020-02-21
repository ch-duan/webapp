package controller

import (
	"log"
	"net/http"
	"text/template"
	"time"
	"webapp/bookstore/dao"
	"webapp/bookstore/model"
	"webapp/bookstore/utils"
)

//Checkout 去结账
func Checkout(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.QueryCart(userID)
	orderID := utils.CreateUUID()
	time := time.Now()
	log.Println(time)
	order := &model.Order{
		ID:          orderID,
		CreateTime:  time,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      userID,
	}
	dao.AddOrder(order)
	cartItems := cart.CartItems
	for _, v := range cartItems {
		if v.ID > 0 {
			orderItem := &model.OrderItem{
				Count:          v.Count,
				Amount:         v.Amount,
				Title:          v.Book.Title,
				Author:         v.Book.Author,
				Price:          v.Book.Price,
				Classification: v.Book.Classification,
				Publisher:      v.Book.Publisher,
				ImgPath:        v.Book.ImgPath,
				OrderID:        orderID,
			}
			dao.AddOrderItem(orderItem)
			book := v.Book
			book.Sales = book.Sales + v.Count
			book.Stock = book.Stock - v.Count
			dao.UpdateBook(book)
		}
	}
	dao.DeleteOrderItemByCartID(cart.ID)
	dao.UpdateCart(&model.Cart{ID: cart.ID, CartItems: nil, TotalCount: 0, TotalAmount: 0, UserID: session.UserID})
	session.OrderID = orderID
	t := template.Must(template.ParseFiles("../view/pages/cart/checkout.html"))
	t.Execute(w, session)
}

//QueryMyOrderHandler 获取用户订单handler
func QueryMyOrderHandler(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	userID := session.UserID
	orders, err := dao.QueryOrderByUserID(userID)
	if err != nil {
		log.Println("QueryMyOrderHandler:QueryOrderByUserID失败", err)
	}
	session.Orders = orders
	t := template.Must(template.ParseFiles("../view/pages/order/order.html"))
	t.Execute(w, session)
}

//GetOrders 获取所有订单
func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := dao.QueryAllOrder()
	if err != nil {
		log.Println("GetOrders:QueryAllOrder失败", err)
	}
	t := template.Must(template.ParseFiles("../view/pages/order/order_manager.html"))
	t.Execute(w, orders)
}

//GetOrderInfo 获取订单信息
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderID")
	orderItems, err := dao.QueryAllOrderItemByOrderID(orderID)
	if err != nil {
		log.Println("GetOrderInfo:QueryAllOrderItemByOrderID失败", err)
	}
	t := template.Must(template.ParseFiles("../view/pages/order/order_info.html"))
	t.Execute(w, orderItems)

}

//SendOrder 发货
func SendOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderID")
	dao.UpdateOrderState(orderID, 1)
	//调用GetOrders函数再次查询一下所有的订单
	GetOrders(w, r)
}

//TakeOrder 收货
func TakeOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderID")
	dao.UpdateOrderState(orderID, 2)
	QueryMyOrderHandler(w, r)
}
