package main

import (
	"fmt"
	"net/http"
)

// server
func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Request URL Path:", r.URL.Path)
	_, _ = fmt.Fprintln(w, "Request URL Query:", r.URL.RawQuery)
	_, _ = fmt.Fprintln(w, "Request Header:", r.Header)
	_, _ = fmt.Fprintln(w, "Request Header Accept-Encoding:", r.Header["Accept-Encoding"], r.Header.Get("Accept-Encoding"))

}

func main() {
	http.HandleFunc("/request", handler)
	_ = http.ListenAndServe(":8000", nil)
}
