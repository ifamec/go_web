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

func GetMyOrders(w http.ResponseWriter, r *http.Request)  {
	_, session := dao.IsLogin(r)
	orders, _ := dao.GetMyOrders(session.UserId, session.Username)
	tp := template.Must(template.ParseFiles("views/pages/order/order.html"))
	_ = tp.Execute(w, orders)
}

func GetOrderDetails(w http.ResponseWriter, r *http.Request)  {
	isLogin, session := dao.IsLogin(r)
	var userName string = ""
	if isLogin {
		userName = session.Username
	}
	orderId := r.FormValue("orderId")
	orderItems, _ := dao.GetOrderItemsByOrderId(orderId, userName)
	tp := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	_ = tp.Execute(w, orderItems)
}

func ShipOrder(w http.ResponseWriter, r *http.Request)  {
	orderId := r.FormValue("orderId")
	_ = dao.UpdateOrderState(orderId, 1)
	GetOrders(w, r)
}
func ConfirmOrder(w http.ResponseWriter, r *http.Request)  {
	orderId := r.FormValue("orderId")
	_ = dao.UpdateOrderState(orderId, 2)
	GetMyOrders(w, r)
}