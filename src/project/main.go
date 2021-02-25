package main

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
 	t := template.Must(template.ParseFiles("views/index.html"))
 	_ = t.Execute(w, "")
}

func main()  {
	// handle static content
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))
	http.HandleFunc("/main", IndexHandler)
	_ = http.ListenAndServe(":8000", nil)
}
