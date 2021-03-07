package dao

import (
	"fmt"
	"go_web/src/project/model"
	"go_web/src/project/utils"
)

func AddCart(cart *model.Cart) (err error) {
	fmt.Println(cart)
	sqlQuery := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	_, err = utils.Db.Exec(sqlQuery, cart.CartId, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserId)
	if err != nil {
		return err
	}
	for _, v := range cart.CartItems {
		err = AddCartItem(v)
	}
	return
}
