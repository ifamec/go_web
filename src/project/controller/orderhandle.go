package controller

import (
	"go_web/src/project/dao"
	"go_web/src/project/model"
	"go_web/src/project/utils"
	"html/template"
	"net/http"
	"time"
)

func Checkout(w http.ResponseWriter, r *http.Request)  {
	// session
	_, session := dao.IsLogin(r)
	userId := session.UserId
	cart, _ := dao.GetCartByUserId(userId)
	orderId := utils.CreateUUID()
	order := &model.Order{
		OrderId:     orderId,
		Timestamp:   time.Now(),
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		Status:      0,
		UserId:      userId,
		Username:    session.Username,
	}
	_ = dao.AddOrder(order)
	// save order items
	for _, v := range cart.CartItems {
		oi := &model.OrderItem{
			Count: v.Count,
			Amount: v.GetAmount(),
			Title: v.Book.Title,
			Author: v.Book.Author,
			Price: v.Book.Price,
			ImgPath: v.Book.ImgPath,
			OrderId: orderId,
		}
		_ = dao.AddOrderItem(oi)
		// update book sales stock
		book := v.Book
		book.Sales += v.Count
		book.Stock -= v.Count
		_ = dao.UpdateBook(book)
	}
	// clear cart
	_ = dao.DeleteCartByCartId(cart.CartId)
	tp := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	_ = tp.Execute(w, order)
}

func GetOrders(w http.ResponseWriter, r *http.Request)  {
	orders, _ := dao.GetOrders()
	tp := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	_ = tp.Execute(w, orders)
}

func GetOrderDetails(w http.ResponseWriter, r *http.Request)  {
	orderId := r.FormValue("orderId")
	orderItems, _ := dao.GetOrderItemsByOrderId(orderId)
	tp := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	_ = tp.Execute(w, orderItems)
}