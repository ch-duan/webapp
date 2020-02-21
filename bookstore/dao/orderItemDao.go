package dao

import (
	"log"
	"webapp/bookstore/model"
	"webapp/bookstore/utils"
)

//QueryAllOrderItemByOrderID 根据订单号查询所有订单项
func QueryAllOrderItemByOrderID(orderID string) ([]*model.OrderItem, error) {
	var orderItems []*model.OrderItem
	sql := utils.SQLSelectAll(&model.OrderItem{}, "orderitems") + " where orderid=?"
	rows, err := utils.Db.Query(sql, orderID)
	if err != nil {
		log.Println("QueryAllOrderItemByOrderID:失败", err)
		return nil, err
	}
	for rows.Next() {
		orderItem := &model.OrderItem{}
		values := utils.StructFieldSlice(orderItem)
		rows.Scan(values...)
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}

//AddOrderItem 添加订单项
func AddOrderItem(orderItem *model.OrderItem) error {
	values := utils.StructFieldSlice(orderItem)
	sql := utils.SQLInsertAll(orderItem, "orderitems")
	_, err := utils.Db.Exec(sql, values...)
	if err != nil {
		log.Println("AddOrderItem:数据库操作错误", err)
		return err
	}
	return nil
}

//UpdateOrderItem 修改订单项
func UpdateOrderItem(orderItem *model.OrderItem) error {
	values := utils.StructFieldSlice(orderItem)
	sql := utils.SQLUpdateAll(orderItem, "orderitems")
	_, err := utils.Db.Exec(sql, values...)
	if err != nil {
		log.Println("UpdateOrderItem:失败", err)
		return err
	}
	return nil
}

//DeleteOrderItemByCartID 删除购物车里面的购物项
func DeleteOrderItemByCartID(ID string) error {
	sql := "delete from cartitems where cartid=?"
	_, err := utils.Db.Exec(sql, ID)
	if err != nil {
		return err
	}
	return nil
}
