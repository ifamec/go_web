package controller

import (
	"go_web/src/project/dao"
	"html/template"
	"net/http"
)

// GetBooks
func GetBooks(w http.ResponseWriter, r *http.Request)  {
	books, _ := dao.GetBooks()
	// template
	tp := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	_ = tp.Execute(w, books)
}