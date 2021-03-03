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

func handlerActionWith(w http.ResponseWriter, r *http.Request)  {
	tp := template.Must(template.ParseFiles("with.html"))
	_ = tp.Execute(w, "Banana")
}

func handlerActionTemp(w http.ResponseWriter, r *http.Request)  {
	tp := template.Must(template.ParseFiles("template.html", "template_sub.html"))
	_ = tp.Execute(w, "SHOW")
}

func handlerActionDefine(w http.ResponseWriter, r *http.Request)  {
	tp := template.Must(template.ParseFiles("define.html"))
	_ = tp.ExecuteTemplate(w, "model", "")
}

func handlerActionDefineS(w http.ResponseWriter, r *http.Request)  {
	var content string
	num := 12
	if num <= 18 {
		content = "define_content_1.html"
	} else {
		content = "define_content_2.html"
	}
	tp := template.Must(template.ParseFiles("define_.html", content))
	_ = tp.ExecuteTemplate(w, "model", "")
}

func handlerActionBlock(w http.ResponseWriter, r *http.Request)  {
	tp := template.Must(template.ParseFiles("block.html"))
	_ = tp.ExecuteTemplate(w, "model", "")
}

func main() {
	http.HandleFunc("/template", handlerTemplate02)
	http.HandleFunc("/action_if", handlerActionIf)
	http.HandleFunc("/action_range", handlerActionRange)
	http.HandleFunc("/action_with", handlerActionWith)
	http.HandleFunc("/action_temp", handlerActionTemp)
	http.HandleFunc("/action_define", handlerActionDefine)
	http.HandleFunc("/action_define_s", handlerActionDefineS)
	http.HandleFunc("/action_block", handlerActionBlock)
	_ = http.ListenAndServe(":8000", nil)
}
