package main

import (
	"html/template"
	"net/http"
)

func handlerTemplate01(w http.ResponseWriter, r *http.Request)  {
	tp, _ := template.ParseFiles("index.html")
	_ = tp.Execute(w, "Hello")
}
func handlerTemplate02(w http.ResponseWriter, r *http.Request)  {
	tp := template.Must(template.ParseFiles("index.html", "main.html"))
	_ = tp.ExecuteTemplate(w, "main.html", "Hello")
}

func main() {
	http.HandleFunc("/template", handlerTemplate02)
	_ = http.ListenAndServe(":8000", nil)
}
