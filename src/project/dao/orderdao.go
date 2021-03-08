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
