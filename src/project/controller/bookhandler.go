package controller

import (
	"go_web/src/project/dao"
	"go_web/src/project/model"
	"html/template"
	"net/http"
	"strconv"
)

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