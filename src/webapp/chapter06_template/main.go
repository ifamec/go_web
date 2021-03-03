package main

import (
	"go_web/src/webapp/chapter06_template/model"
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

func handlerActionIf(w http.ResponseWriter, r *http.Request)  {
	tp := template.Must(template.ParseFiles("if.html"))
	num := 17
	_ = tp.Execute(w, num > 18)
}

func handlerActionRange(w http.ResponseWriter, r *http.Request)  {
	tp := template.Must(template.ParseFiles("range.html"))
	var empSlice []*model.Employee
	emp1 := &model.Employee{Name: "a1", Id: 1, Email: "a1@panda.com"}
	emp2 := &model.Employee{Name: "a2", Id: 2, Email: "a2@panda.com"}
	emp3 := &model.Employee{Name: "a3", Id: 3, Email: "a3@panda.com"}
	empSlice = append(empSlice, emp1)
	empSlice = append(empSlice, emp2)
	empSlice = append(empSlice, emp3)
	// fmt.Println(empSlice)
	_ = tp.Execute(w, empSlice)
}

func main() {
	http.HandleFunc("/template", handlerTemplate02)
	http.HandleFunc("/action_if", handlerActionIf)
	http.HandleFunc("/action_range", handlerActionRange)
	_ = http.ListenAndServe(":8000", nil)
}
