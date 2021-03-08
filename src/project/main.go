package main

import (
	"go_web/src/project/controller"
	_ "html/template"
	"net/http"
)

func main()  {
	// handle static content
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))
	http.HandleFunc("/main", controller.GetPageBooksByPrice)
	// login
	http.HandleFunc("/login", controller.Login)
	// logout
	http.HandleFunc("/logout", controller.Logout)
	// signup
	http.HandleFunc("/signup", controller.Signup)
	// check username
	http.HandleFunc("/checkUsername", controller.CheckUsername)
	// get books
	// http.HandleFunc("/getBooks", controller.GetBooks)
	// add book
	// http.HandleFunc("/addBook", controller.AddBook)
	// delete book
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	// modify book
	http.HandleFunc("/modifyBookPage", controller.ModifyBookPage)
	// update book
	http.HandleFunc("/addOrUpdateBook", controller.AddOrUpdateBook)
	// get page books
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	// get page books by price
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	// add book to cart
	http.HandleFunc("/addBookToCart", controller.AddBookToCart)
	// get cart info
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	// clear cart
	http.HandleFunc("/clearCart", controller.ClearCart)
	// clear cart item
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	// update cart item
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)



	_ = http.ListenAndServe(":8000", nil)
}
