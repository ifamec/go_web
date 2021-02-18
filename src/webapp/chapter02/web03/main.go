package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "My ServeMux", r.URL.Path)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/myServeMux", handler)
	_ = http.ListenAndServe(":8000", mux)
}
