package controller

import (
	"go_web/src/project/dao"
	"go_web/src/project/model"
	"html/template"
	"net/http"
	"strconv"
)

// GetBooks
func GetBooks(w http.ResponseWriter, r *http.Request)  {
	books, _ := dao.GetBooks()
	// template
	tp := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	_ = tp.Execute(w, books)
}

// AddBook
func AddBook(w http.ResponseWriter, r *http.Request)  {

	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	sales, _ := strconv.Atoi(r.PostFormValue("sales"))
	stock, _ := strconv.Atoi(r.PostFormValue("stock"))

	_ = dao.AddBook(&model.Book{
		Title: r.PostFormValue("title"),
		Author: r.PostFormValue("author"),
		Price: price,
		Sales: sales,
		Stock: stock,
		ImgPath: r.PostFormValue("img_path"),
	})

	GetBooks(w, r)
}