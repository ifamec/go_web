package main

import (
	"fmt"
	"net/http"
)

// server
func handlerHeader(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Request URL Path:", r.URL.Path)
	_, _ = fmt.Fprintln(w, "Request URL Query:", r.URL.RawQuery)
	_, _ = fmt.Fprintln(w, "Request Header:", r.Header)
	_, _ = fmt.Fprintln(w, "Request Header Accept-Encoding:", r.Header["Accept-Encoding"], r.Header.Get("Accept-Encoding"))

}

func handlerRequestBody(w http.ResponseWriter, r *http.Request)  {
	// get request body
	length := r.ContentLength
	// byte slice
	body := make([]byte, length)
	// read request to body
	_, _ = r.Body.Read(body)
	_, _ = fmt.Fprintln(w, "Request Body:", string(body))
}

func handlerRequestForm(w http.ResponseWriter, r *http.Request)  {
	// parse form
	_ = r.ParseForm()
	_, _ = fmt.Fprintln(w, "Request Body ParseForm:", r.Form)

	// post form
	_, _ = fmt.Fprintln(w, "Request PostForm:", r.PostForm)
}

func handlerRequestFormValue(w http.ResponseWriter, r *http.Request)  {
	_, _ = fmt.Fprintln(w, "Request FormValue:", r.FormValue("username"), r.FormValue("password"))

	_, _ = fmt.Fprintln(w, "Request PostFormValue:", r.PostFormValue("username"), r.PostFormValue("password"))
}

func main() {
	http.HandleFunc("/request", handlerRequestFormValue)
	_ = http.ListenAndServe(":8000", nil)
}
