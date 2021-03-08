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

func GetCartItemByBookIdCartId(bookId int, cartId string) (cartItem *model.CartItem, err error) {
	sqlQuery := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id = ?"
	row := utils.Db.QueryRow(sqlQuery, bookId, cartId)
	cartItem = &model.CartItem{}
	err = row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.CartId)
	if err != nil {
		return nil, err
	}
	return
}
func UpdateBookCount(bCount int, bookId int, cartId string) (err error) {
	sqlQuery := "update cart_items set COUNT = ? where book_id = ? and cart_id = ?"
	_, err = utils.Db.Exec(sqlQuery, bCount, bookId, cartId)
	return
}

func GetCartItemsByCartId(cartId string) (cartItems []*model.CartItem, err error) {
	sqlQuery := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?"
	row, err := utils.Db.Query(sqlQuery, cartId)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var bookId int
		cartItem := &model.CartItem{}
		err = row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &bookId, &cartItem.CartId)
		if err != nil {
			return nil, err
		}
		cartItem.Book, _ = GetBookById(bookId)
		cartItems = append(cartItems, cartItem)
	}
	return
}