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
