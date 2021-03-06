package controller

import (
	"go_web/src/project/dao"
	"go_web/src/project/model"
	"go_web/src/project/utils"
	"html/template"
	"net/http"
)

// Login
func Login(w http.ResponseWriter, r *http.Request)  {
	// get username password
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	// validate
	user, _ := dao.LoginValidate(username, password)
	if user.Id != 0 {
		// login success
		uuid := utils.CreateUUID()
		_ = dao.AddSession(&model.Session{
			SessionId: uuid,
			Username:  user.Username,
			UserId:    user.Id,
		})
		http.SetCookie(w, &http.Cookie{
			Name:     "user",
			Value:    uuid,
			HttpOnly: true,
		})
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		_ = t.Execute(w, user)
	} else {
		// error
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		_ = t.Execute(w, "username or password invalid")
	}
}
// Signup
func Signup(w http.ResponseWriter, r *http.Request)  {
	// get username password
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	// validate
	isAvailable, _ := dao.UsernameAvailability(username)
	if  isAvailable {
		// available
		err := dao.CreateUser(username, password, email)
		if err != nil {
			t := template.Must(template.ParseFiles("views/pages/user/signup.html"))
			_ = t.Execute(w, "error occurs during signup")
		} else {
			t := template.Must(template.ParseFiles("views/pages/user/signup_success.html"))
			_ = t.Execute(w, "")
		}
	} else {
		// not available
		t := template.Must(template.ParseFiles("views/pages/user/signup.html"))
		_ = t.Execute(w, "username exist")
	}
}
// CheckUsername use ajax
func CheckUsername(w http.ResponseWriter, r *http.Request)  {
	username := r.PostFormValue("username")
	isAvailable, _ := dao.UsernameAvailability(username)
	if isAvailable {
		_, _ = w.Write([]byte("<font style='color:green'>username valid</font>"))
	} else {
		_, _ = w.Write([]byte("username exist, retry"))
	}

}