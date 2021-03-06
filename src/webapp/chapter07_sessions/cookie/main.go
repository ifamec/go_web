package main

import (
	"fmt"
	"net/http"
)

// set cookie
func setCookie(w http.ResponseWriter, r *http.Request)  {
	cookie := http.Cookie{
		Name:     "user",
		Value:    "admin",
		HttpOnly: true,
		MaxAge:   60,
	}
	cookie2 := http.Cookie{
		Name:     "user1",
		Value:    "user01",
		HttpOnly: true,
	}
	// w.Header().Set("Set-Cookie", cookie.String())
	// w.Header().Add("Set-Cookie", cookie_2.String())

	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie2)
}

// get cookies
func getCookies(w http.ResponseWriter, r *http.Request) {
	// cookies := r.Header["Cookie"]
	cookies := r.Cookies()
	fmt.Println(cookies)
	c1, _ := r.Cookie("user")
	c2, _ := r.Cookie("user1")
	fmt.Println(c1, c2)
}

func main()  {
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookies", getCookies)
	_ = http.ListenAndServe(":8000", nil)
}