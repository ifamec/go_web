package main

import (
	"go_web/src/project/controller"
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
	// login
	http.HandleFunc("/login", controller.Login)
	// signup
	http.HandleFunc("/signup", controller.Signup)
	// check username
	http.HandleFunc("/checkUsername", controller.CheckUsername)
	// get books
	http.HandleFunc("/getBooks", controller.GetBooks)
	// add book
	// http.HandleFunc("/addBook", controller.AddBook)
	// delete book
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	// modify book
	http.HandleFunc("/modifyBookPage", controller.ModifyBookPage)
	// update book
	http.HandleFunc("/addOrUpdateBook", controller.AddOrUpdateBook)



	_ = http.ListenAndServe(":8000", nil)
}
