package controller

import (
	"encoding/json"
	"fmt"
	"go_web/src/project/dao"
	"go_web/src/project/model"
	"go_web/src/project/utils"
	"html/template"
	"net/http"
	"strconv"
)

// AddBookToCart
func AddBookToCart(w http.ResponseWriter, r *http.Request)  {
	// has cart
	isLogin, session := dao.IsLogin(r)
	if !isLogin {
		_, _ = w.Write([]byte("Please Login and retry."))
		return
	}

	bookId, err := strconv.Atoi(r.FormValue("bookId"))
	if err != nil { fmt.Println(err) }

	book, err := dao.GetBookById(bookId)
	if err != nil { fmt.Println(err) }
	userId := session.UserId
	cart, _ := dao.GetCartByUserId(userId)
	if cart != nil {
		cartItem , _ := dao.GetCartItemByBookIdCartId(bookId, cart.CartId)
		if cartItem != nil { // book exist, count++
			for _, v := range cart.CartItems {
				if bookId == v.Book.Id {
					v.Count++
					_ = dao.UpdateBookCount(v, bookId, cart.CartId)
				}
			}
		} else { // new book
			newBook := &model.CartItem{
				Book: book,
				Count: 1,
				CartId: cart.CartId,
			}
			cart.CartItems = append(cart.CartItems, newBook)
			_ = dao.AddCartItem(newBook)
		}
		_ = dao.UpdateCart(cart)
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
	_, _ = w.Write([]byte("\"" + book.Title + "\" added to cart."))
}

func GetCartInfo(w http.ResponseWriter, r *http.Request)  {
	_, session := dao.IsLogin(r)
	userId := session.UserId

	// get cart
	cart, _ := dao.GetCartByUserId(userId)
	if cart != nil {
		cart.Username = session.Username
		tp := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		_ = tp.Execute(w, cart)
	} else {
		emptyCart := &model.Cart{
			Username: session.Username,
			UserId:   session.UserId,
		}
		tp := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		_ = tp.Execute(w, emptyCart)
	}
}

func ClearCart(w http.ResponseWriter, r *http.Request)  {
	cartId := r.FormValue("cartId")
	_ = dao.DeleteCartByCartId(cartId)
	GetCartInfo(w, r)
}

func DeleteCartItem(w http.ResponseWriter, r *http.Request)  {
	cartItemId, err := strconv.Atoi(r.FormValue("cartItemId"))
	if err != nil { return }
	_, session := dao.IsLogin(r)
	userId := session.UserId
	cart, _ := dao.GetCartByUserId(userId)
	cartItems := cart.CartItems
	for k, v := range cartItems {
		if cartItemId == v.CartItemId {
			 // remove from slice
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			cart.CartItems = cartItems // OPTIONAL
			_ = dao.DeleteCartItemByCartItemID(cartItemId)
		}
	}
	_ = dao.UpdateCart(cart)
	GetCartInfo(w, r)
}

func UpdateCartItem(w http.ResponseWriter, r *http.Request)  {
	cartItemId, err := strconv.Atoi(r.FormValue("cartItemId"))
	bookCount, err := strconv.Atoi(r.FormValue("bookCount"))
	if err != nil { return }
	_, session := dao.IsLogin(r)
	userId := session.UserId
	cart, _ := dao.GetCartByUserId(userId)
	cartItems := cart.CartItems
	for _, v := range cartItems {
		if cartItemId == v.CartItemId {
			// update from slice
			v.Count = bookCount
			_ = dao.UpdateBookCount(v, v.Book.Id, v.CartId)
		}
	}
	_ = dao.UpdateCart(cart)
	// GetCartInfo(w, r)
	cart, _ = dao.GetCartByUserId(userId)
	var updateItemAmount float64
	for _, v := range cart.CartItems {
		if cartItemId == v.CartItemId {
			updateItemAmount = v.Amount
		}
	}
	response := utils.UpdateCartResponse{
		TotalCount:       cart.GetTotalCount(),
		TotalAmount:      cart.GetTotalAmount(),
		UpdateItemAmount: updateItemAmount,
	}
	js, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Marshall Error", err)
	}
	_, _ = w.Write(js)
}