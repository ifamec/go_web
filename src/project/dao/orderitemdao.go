package dao

import (
	"go_web/src/project/model"
	"go_web/src/project/utils"
)

func AddOrderItem(oi *model.OrderItem) (err error) {
	sqlQuery := "insert into order_items(count,amount,title,author,price,img_path,order_id) values(?,?,?,?,?,?,?)"
	_, err = utils.Db.Exec(sqlQuery, oi.Count, oi.Amount, oi.Title, oi.Author, oi.Price, oi.ImgPath, oi.OrderId)
	return
}

func GetOrderItemsByOrderId(orderId string) (ois []*model.OrderItem, err error) {
	sqlQuery := "select id,count,amount,title,author,price,img_path,order_id from order_items where order_id = ?"
	rows, err := utils.Db.Query(sqlQuery, orderId)
	if err != nil {
		return nil, err
	}
	ois = make([]*model.OrderItem, 0)
	for rows.Next() {
		oi := &model.OrderItem{}
		_ = rows.Scan(&oi.OrderItemId, &oi.Count, &oi.Amount, &oi.Title, &oi.Author, &oi.Price, &oi.ImgPath, &oi.OrderId)
		ois = append(ois, oi)
	}
	return
}