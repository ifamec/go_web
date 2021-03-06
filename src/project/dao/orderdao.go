package dao

import (
	"go_web/src/project/model"
	"go_web/src/project/utils"
)

func AddOrder(o *model.Order) (err error) {
	sqlQuery := "insert into orders(id,timestamp,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	_, err = utils.Db.Exec(sqlQuery, o.OrderId, o.Timestamp, o.TotalCount, o.TotalAmount, o.Status, o.UserId)
	return
}

func GetOrders() (orders []*model.Order, err error) {
	sqlQuery := "select id,timestamp,total_count,total_amount,state,user_id from orders"
	rows, err := utils.Db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	orders = make([]*model.Order, 0)
	for rows.Next() {
		o := &model.Order{}
		_ = rows.Scan(&o.OrderId, &o.Timestamp, &o.TotalCount, &o.TotalAmount, &o.Status, &o.UserId)
		orders = append(orders, o)
	}
	return
}

func GetMyOrders(userId int, userName string) (orders []*model.Order, err error) {
	sqlQuery := "select id,timestamp,total_count,total_amount,state,user_id from orders where user_id = ?"
	rows, err := utils.Db.Query(sqlQuery, userId)
	if err != nil {
		return nil, err
	}
	orders = make([]*model.Order, 0)
	for rows.Next() {
		o := &model.Order{}
		_ = rows.Scan(&o.OrderId, &o.Timestamp, &o.TotalCount, &o.TotalAmount, &o.Status, &o.UserId)
		o.Username = userName
		orders = append(orders, o)
	}
	return
}

func UpdateOrderState(orderId string, state int) (err error) {
	sqlQuery := "update orders set state =? where id = ?"
	_, err = utils.Db.Exec(sqlQuery, state, orderId)
	return
}