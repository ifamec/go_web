package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello World", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	_ = http.ListenAndServe(":8000", nil)
}
