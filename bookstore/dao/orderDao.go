package dao

import (
	"log"
	"webapp/bookstore/model"
	"webapp/bookstore/utils"
)

//AddOrder 添加订单
func AddOrder(order *model.Order) error {
	values := utils.StructFieldSlice(order)
	sql := utils.SQLInsertAll(order, "orders")
	_, err := utils.Db.Exec(sql, values...)
	if err != nil {
		log.Println("AddOrder:失败", err)
		return err
	}
	return nil
}

//QueryOrderByUserID 查询用户订单
func QueryOrderByUserID(userID int) ([]*model.Order, error) {
	var orders []*model.Order
	sql := utils.SQLSelectAll(&model.Order{}, "orders") + " where userid=?"
	rows, err := utils.Db.Query(sql, userID)
	if err != nil {
		log.Println("QueryOrderByUserID:失败", err)
		return nil, err
	}
	for rows.Next() {
		order := &model.Order{}
		values := utils.StructFieldSlice(order)
		rows.Scan(values...)
		orders = append(orders, order)
	}
	return orders, nil
}

//QueryAllOrder 查询所有订单
func QueryAllOrder() ([]*model.Order, error) {
	var orders []*model.Order
	sql := utils.SQLSelectAll(&model.Order{}, "orders")
	rows, err := utils.Db.Query(sql)
	if err != nil {
		log.Println("QueryAllOrder:失败", err)
		return nil, err
	}
	for rows.Next() {
		order := &model.Order{}
		values := utils.StructFieldSlice(order)
		rows.Scan(values...)
		orders = append(orders, order)
	}
	return orders, nil
}

//UpdateOrderState 更新订单状态
func UpdateOrderState(orderID string, state int) error {
	sql := "update orders set state=? where id=?"
	_, err := utils.Db.Exec(sql, state, orderID)
	if err != nil {
		log.Println("UpdateOrderState:失败", err)
		return err
	}
	return nil
}
