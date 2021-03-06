package main

import (
	"net/http"
)

// set cookie
func setCookie(w http.ResponseWriter, r *http.Request)  {
	cookie := http.Cookie{
		Name:       "user",
		Value:      "admin",
		HttpOnly:   true,
	}
	// cookie_2 := http.Cookie{
	// 	Name:       "user",
	// 	Value:      "admin",
	// 	HttpOnly:   true,
	// }
	// w.Header().Set("Set-Cookie", cookie.String())
	// w.Header().Add("Set-Cookie", cookie_2.String())

	http.SetCookie(w, &cookie)
}

func main()  {
	http.HandleFunc("/setCookie", setCookie)
	_ = http.ListenAndServe(":8000", nil)
}