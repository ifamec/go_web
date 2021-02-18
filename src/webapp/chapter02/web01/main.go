package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}
func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Through My Handler", r)
}


func main() {
	myHandler := MyHandler{}
	http.Handle("/myHandler", &myHandler)
	_ = http.ListenAndServe(":8000", nil)
}
