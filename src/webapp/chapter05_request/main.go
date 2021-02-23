package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Test struct {
	Username string
	Password string
}

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

func handlerResponseText(w http.ResponseWriter, r *http.Request)  {
	_, _ = w.Write([]byte("Request Respond\n"))
}

func handlerResponseHtml(w http.ResponseWriter, r *http.Request)  {
	_, _ = w.Write([]byte(`<html><b>Request Respond</b></html>`))
}

func handlerResponseJson(w http.ResponseWriter, r *http.Request)  {
	// json
	w.Header().Set("Content-Type", "application/json")
	jsonStr, _ := json.Marshal(Test{
		Username: "admin",
		Password: "admin",
	})
	_, _ = w.Write(jsonStr)
}

func handlerResponseRedirect(w http.ResponseWriter, r *http.Request)  {
	// redirect
	w.Header().Set("Location", "https://www.google.com")
	w.WriteHeader(302)}

func main() {
	http.HandleFunc("/request", handlerResponseRedirect)
	_ = http.ListenAndServe(":8000", nil)
}
