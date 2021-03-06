package dao

import (
	"go_web/src/project/model"
	"go_web/src/project/utils"
)

func AddCart(cart *model.Cart) (err error) {
	// fmt.Println(cart)
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

func GetCartByUserId(userId int) (cart *model.Cart, err error) {
	sqlQuery := "select id,total_count,total_amount,user_id from carts where user_id = ?"
	row := utils.Db.QueryRow(sqlQuery, userId)
	cart = &model.Cart{}
	err = row.Scan(&cart.CartId, &cart.TotalCount, &cart.TotalAmount, &cart.UserId)
	if err != nil {
		return nil, err
	}
	cart.CartItems, err = GetCartItemsByCartId(cart.CartId)

	return
}

func UpdateCart(cart *model.Cart) (err error) {
	sqlQuery := "update carts set total_count = ?, total_amount = ? where id = ?"
	_, err = utils.Db.Exec(sqlQuery, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartId)
	return
}

func DeleteCartByCartId(cartId string) (err error) {
	// delete cart items
	err = DeleteCartItemsByCartID(cartId)
	if err != nil {
		return err
	}
	sqlQuery := "delete from carts where id = ?"
	_, err = utils.Db.Exec(sqlQuery, cartId)
	return err
}