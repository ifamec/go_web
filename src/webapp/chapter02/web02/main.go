package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct{}
func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Through My Handler & My Server", r)
}


func main() {
	myHandler := MyHandler{}
	server := http.Server{
		Addr: ":8000",
		Handler: &myHandler,
		ReadTimeout: time.Second * 2,
	}
	_ = server.ListenAndServe()
}
