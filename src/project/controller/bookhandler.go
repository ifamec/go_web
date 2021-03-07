package controller

import (
	"go_web/src/project/dao"
	"go_web/src/project/model"
	"html/template"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// get pageNumber
	pageNumber, _ := strconv.Atoi(r.FormValue("pageNumber"))
	if pageNumber == 0 { pageNumber = 1 }
	// get page bools
	page, _ := dao.GetPageBooks(pageNumber)
	t := template.Must(template.ParseFiles("views/index.html"))
	_ = t.Execute(w, page)
}

// GetBooks
// func GetBooks(w http.ResponseWriter, r *http.Request)  {
// 	books, _ := dao.GetBooks()
// 	// template
// 	tp := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
// 	_ = tp.Execute(w, books)
// }

// AddBook
// func AddBook(w http.ResponseWriter, r *http.Request)  {
//
// 	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
// 	sales, _ := strconv.Atoi(r.PostFormValue("sales"))
// 	stock, _ := strconv.Atoi(r.PostFormValue("stock"))
//
// 	_ = dao.AddBook(&model.Book{
// 		Title: r.PostFormValue("title"),
// 		Author: r.PostFormValue("author"),
// 		Price: price,
// 		Sales: sales,
// 		Stock: stock,
// 		ImgPath: r.PostFormValue("img_path"),
// 	})
//
// 	GetBooks(w, r)
// }

// DeleteBook
func DeleteBook(w http.ResponseWriter, r *http.Request)  {
	id, _ := strconv.Atoi(r.FormValue("bookId"))
	_ = dao.DeleteBook(id)
	GetPageBooks(w, r)
}

// ModifyBookPage
func ModifyBookPage(w http.ResponseWriter, r *http.Request)  {
	id, _ := strconv.Atoi(r.FormValue("bookId"))
	book, _ := dao.GetBookById(id)
	if book.Id > 0 {
		// update book
		tp := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		_ = tp.Execute(w, book)
	} else {
		// add book
		tp := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		_ = tp.Execute(w, "")
	}
}

// UpdateBook
func AddOrUpdateBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PostFormValue("bookId"))
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	sales, _ := strconv.Atoi(r.PostFormValue("sales"))
	stock, _ := strconv.Atoi(r.PostFormValue("stock"))

	book := &model.Book{
		Id:      id,
		Title:   r.PostFormValue("title"),
		Author:  r.PostFormValue("author"),
		Price:   price,
		Sales:   sales,
		Stock:   stock,
		ImgPath: r.PostFormValue("img_path"),
	}
	if id > 0 {
		_ = dao.UpdateBook(book)
	} else {
		_ = dao.AddBook(book)
	}

	GetPageBooks(w, r)
}

// GetPageBooks
func GetPageBooks(w http.ResponseWriter, r *http.Request)  {
	// get pageNumber
	pageNumber, _ := strconv.Atoi(r.FormValue("pageNumber"))
	if pageNumber == 0 { pageNumber = 1 }
	// get page bools
	page, _ := dao.GetPageBooks(pageNumber)
	// template
	tp := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	_ = tp.Execute(w, page)
}
// GetPageBooksByPrice
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request)  {
	// get pageNumber
	pageNumber, _ := strconv.Atoi(r.FormValue("pageNumber"))
	priceMin, _   := strconv.ParseFloat(r.FormValue("min"), 64)
	priceMax, _   := strconv.ParseFloat(r.FormValue("max"), 64)
	if pageNumber == 0 { pageNumber = 1 }
	var page *model.Page
	if priceMax == 0.0 && priceMin == 0.0 {
		page, _ = dao.GetPageBooks(pageNumber)
	} else {
		// get page bools
		page, _ = dao.GetPageBooksByPrice(pageNumber, priceMin, priceMax)
		page.PriceMin = priceMin
		page.PriceMax = priceMax
	}
	// get cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		sessionId := cookie.Value
		session, _ := dao.GetSession(sessionId)
		if session.UserId != 0 {
			page.Username = session.Username
		}
	}
	// template
	tp := template.Must(template.ParseFiles("views/index.html"))
	_ = tp.Execute(w, page)
}