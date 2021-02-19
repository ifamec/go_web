package main

import (
	"fmt"
	"net/http"
)

func handler (w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Test HTTP Protocol", r)
}

func main() {
	// handler
	http.HandleFunc("/http", handler)
	// routing
	_ = http.ListenAndServe(":8000", nil)

}
