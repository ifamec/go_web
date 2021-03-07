package dao

import (
	"go_web/src/project/model"
	"go_web/src/project/utils"
)

func AddCartItem(cartItem *model.CartItem) (err error) {
	sqlQuery := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"

	_, err = utils.Db.Exec(sqlQuery, cartItem.Count, cartItem.GetAmount(), cartItem.Book.Id, cartItem.CartId)
	return
}