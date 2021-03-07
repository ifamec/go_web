package controller

import (
	"fmt"
	"go_web/src/project/dao"
	"go_web/src/project/model"
	"go_web/src/project/utils"
	"net/http"
	"strconv"
)

// AddBookToCart
func AddBookToCart(w http.ResponseWriter, r *http.Request)  {
	bookId, err := strconv.Atoi(r.FormValue("bookId"))
	if err != nil { fmt.Println(err) }

	book, err := dao.GetBookById(bookId)
	if err != nil { fmt.Println(err) }
	// has cart
	_, session := dao.IsLogin(r)
	userId := session.UserId
	cart, _ := dao.GetCartByUserId(userId)
	if cart != nil {

	} else {
		// create a cart
		CartId := utils.CreateUUID()
		var cartItems []*model.CartItem
		cartItem := &model.CartItem{
			Book: book,
			Count: 1,
			CartId: CartId,
		}
		cartItems = append(cartItems, cartItem)
		// to db
		cart := &model.Cart{
			CartId: CartId,
			UserId: userId,
			CartItems: cartItems,
		}
		_ = dao.AddCart(cart)
	}
}